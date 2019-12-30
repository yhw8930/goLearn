package main

import (
	"goLearn/crawler/engine"
	"goLearn/crawler/persist"
	"goLearn/crawler/scheduler"
	"goLearn/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
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
