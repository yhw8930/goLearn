package engine

//请求数据结构
type Request struct {
	Url        string                   //请求url
	ParserFunc func([]byte) ParseResult //内容解析函数
}

//经过内容解析函数解析后的返回数据结构体
type ParseResult struct {
	Requests []Request     //请求数据结构切片
	Items    []interface{} //抓取到的有用信息项
}
