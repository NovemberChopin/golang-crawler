package parser

import (
	"crawler/engine"
	"crawler/zhipin/util"
	"regexp"
)

// 解析 每个区中的 area 列表
func ParseAreaList(bytes []byte) engine.ParseResult {
	var areaRe = regexp.MustCompile(`<a href="(/c101[^"]+)" ka="(sel-area-[^"]+)">([^<]+)</a>`)
	// 出去字符串空格
	content := util.RemoveSpace(bytes)
	// 匹配 area 信息
	submatch := areaRe.FindAllSubmatch([]byte(content), -1)
	result := engine.ParseResult{}
	for _, item := range submatch {
		result.Requests = append(result.Requests, engine.Request{
			Url:       "https://www.zhipin.com" + string(item[1]) + "?ka=" + string(item[2]),
			ParseFunc: ParseAreaList,
		})
	}
	return result
}
