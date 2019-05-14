package persist

import (
	"context"
	"crawler/engine"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() (chan engine.Item, error) {
	// Must turn off sniff in docker
	//client, err := elastic.NewClient(elastic.SetSniff(false))
	//if err != nil {
	//	return nil, err
	//}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v\n",
				itemCount, item)
			itemCount++

			// Save Data
			//err := save(client, item)
			//if err != nil {
			//	// if have err, ignore it
			//	log.Printf("Item Saver: error, saving item %v: %v",
			//		item, err)
			//}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().
		Index("data_profile").
		Type(item.Type)
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
