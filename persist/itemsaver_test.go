package persist

import (
	"context"
	"crawler/engine"
	"crawler/zhenai/model"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"testing"
)

func TestSave(t *testing.T) {

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1946858930",
		Type: "zhenai",
		Id:   "1946858930",
		Payload: model.Profile{
			Name:     "為你垨候",
			Gender:   "女士",
			Age:      40,
			Height:   163,
			Weight:   54,
			Income:   "5-8千",
			Marriage: "未婚",
			Address:  "佛山顺德区",
		},
	}

	// TODO: Try to start up elastic search
	// here using docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "data_test"
	// Save expected
	err = es_save(client, index, expected)
	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)

	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	// Verify result
	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}

}

func TestMongoSave(t *testing.T) {
	// mongodb connect
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1946858930",
		Type: "zhenai",
		Id:   "1946858930",
		Payload: model.Profile{
			Name:     "為你垨候",
			Gender:   "女士",
			Age:      40,
			Height:   163,
			Weight:   54,
			Income:   "5-8千",
			Marriage: "未婚",
			Address:  "佛山顺德区",
		},
	}
	// 保存数据
	err = mongo_save(session, "crawler", expected)
	if err != nil {
		panic(err)
	}

	c := session.DB("crawler").C("zhenai")

	var result engine.Item
	err = c.Find(bson.M{"id": "1946858930"}).One(&result)
	// result 为 Json 类型
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s, %s, %v\n", result.Url, result.Id, result.Payload)
}
