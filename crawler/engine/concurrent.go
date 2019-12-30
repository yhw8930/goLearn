package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}
type ReadyNotifier interface {
	WorkerReady(w chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	m := make(map[string]int)

	for {
		//chan传输的值 out必须有值 循环才可以正常运作 out的值从createWorker中获取
		result := <-out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}
		for _, request := range result.Requests {
			if isDuplicate(request.Url, m) {
				continue
			}
			//将新的地址传入函数 在函数内的结构体进行in值的改变
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			//注意函数内部接收方的类型 每一个并发要传入一个独特的in 并且从这个独特的in中获得返回值
			ready.WorkerReady(in)
			//这个独特的in被处理后执行这行代码 获得返回值
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			//将返回值放入out chan类型 out类似于无限数组 可以放入无限的数据 被Run吸收
			out <- result
		}
	}()
}

func isDuplicate(url string, m map[string]int) bool {
	if m[url] > 0 {
		return true
	} else {
		m[url]++

	}
	return false
}
