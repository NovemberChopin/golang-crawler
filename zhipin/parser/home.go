package parser

import (
	"crawler/engine"
	"regexp"
)

// 解析首页 的 城市 列表
func PraseCityList(bytes []byte) engine.ParseResult {
	cityListRe := regexp.MustCompile(`<a href="(/c101[^"]+)" ka="(sel-city-101[^"]+)">([^<]+)</a>`)
	submatch := cityListRe.FindAllSubmatch(bytes, -1)
	result := engine.ParseResult{}
	for _, item := range submatch {
		result.Requests = append(result.Requests, engine.Request{
			Url:       "https://www.zhipin.com" + string(item[1]) + "?ka=" + string(item[2]),
			ParseFunc: ParseBusinessList,
		})
	}
	return result
}
