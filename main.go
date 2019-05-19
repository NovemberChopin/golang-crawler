package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	zhipin "crawler/zhipin/parser"
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
		WorkerCount: 5,
		ItemChan:    itemChan,
	}
	// 配置抓取任务信息
	e.Run(engine.Request{
		Url:       "https://www.zhipin.com/c101010100/a_%E4%B8%8A%E5%9C%B0-b_%E6%B5%B7%E6%B7%80%E5%8C%BA/?ka=sel-area-1",
		ParseFunc: zhipin.ParsePositionList,
	})

	// 爬取单个城市
	//e.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun/yantai",
	//	ParseFunc: parser.ParseCity,
	//})
}
