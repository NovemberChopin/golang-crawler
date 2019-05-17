package parser

import (
	"crawler/engine"
	"crawler/zhipin/util"
	"regexp"
)

// 解析城市包含的所有 区 信息
func ParseBusinessList(bytes []byte) engine.ParseResult {
	var businessRe = regexp.MustCompile(
		`<a href="(/c101[^"]+)" ka="(sel-business-[^"]+)">([^<]+)</a>`)

	content := util.RemoveSpace(bytes)
	submatch := businessRe.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, item := range submatch {
		result.Requests = append(result.Requests, engine.Request{
			Url:       "https://www.zhipin.com" + string(item[1]) + "?ka=" + string(item[2]),
			ParseFunc: ParseAreaList,
		})
	}
	return result
}
