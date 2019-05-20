package persist

import (
	"context"
	"crawler/engine"
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	// Must turn off sniff in docker
	//client, err := elastic.NewClient(elastic.SetSniff(false))
	//if err != nil {
	//	return nil, err
	//}

	// mongodb connect
	//session, err := mgo.Dial("localhost:27017")
	//if err != nil {
	//	panic(err)
	//}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			// 接收到发送的 item
			item := <-out
			log.Printf("Item Saver: got item #%d: %v\n",
				itemCount, item)
			itemCount++

			// Save Data in elasticsearch
			//err := es_save(client, index, item)

			// Save data in mongodb
			//err := mongo_save(session, index, item)
			//
			//if err != nil {
			//	// if have err, ignore it
			//	log.Printf("Item Saver: error, saving item %v: %v",
			//		item, err)
			//}
		}
	}()
	return out, nil
}

// 使用 elasticsearch 保存数据
func es_save(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().
		Index(index).   // 数据库名称
		Type(item.Type) // 表名
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.
		BodyJson(item).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// 使用 MongoDB 保存数据
func mongo_save(session *mgo.Session, dbName string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	c := session.DB(dbName).C(item.Type)
	err := c.Insert(item.Payload)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
