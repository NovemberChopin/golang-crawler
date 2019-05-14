package engine

// 请求结构
type Request struct {
	Url       string // 请求地址
	ParseFunc func([]byte) ParseResult
}

// 解析结果结构
type ParseResult struct {
	Requests []Request // 解析出的请求
	Items    []Item    // 解析出的内容
}

type Item struct {
	Url     string      // 个人信息Url地址
	Type    string      // table
	Id      string      // Id
	Payload interface{} // 详细信息
}

func NilParseFun([]byte) ParseResult {
	return ParseResult{}
}
