package main

import (
	"goLearn/crawler/engine"
	"goLearn/crawler/persist"
	"goLearn/crawler/scheduler"
	"goLearn/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	/*e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})*/
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/changsha",
		ParserFunc: parser.ParseCity,
	})

}
