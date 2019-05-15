package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	// 启动数据存储
	itemChan, err := persist.ItemSaver("crawler")
	if err != nil {
		panic(err)
	}
	// 配置爬虫引擎
	e := engine.ConcurrendEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		//Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 30,
		ItemChan:    itemChan,
	}
	// 配置抓取任务信息
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})

	// 爬取单个城市
	//e.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun/yantai",
	//	ParseFunc: parser.ParseCity,
	//})
}
