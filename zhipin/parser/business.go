package parser

import (
	"crawler/engine"
	"crawler/zhipin/util"
	"regexp"
)

// 解析 每个区中的 area 列表
func ParseAreaList(bytes []byte) engine.ParseResult {
	// 匹配 area 正则
	var areaRe = regexp.MustCompile(`<a href="(/c101[^"]+)" ka="(sel-area-[^"]+)">([^<]+)</a>`)
	// 匹配当前页面的Url
	var pageUrlRe = regexp.MustCompile(
		`<a href="(/c101[^"]+)" class="selected" ka="(sel-business-[^"]+)">[^<]+</a>`)

	// 除去字符串中空格
	content := util.RemoveSpace(bytes)
	// 匹配 area 信息
	submatch := areaRe.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}

	if len(submatch) == 0 {
		// 如果在在当前 区 内 没有匹配到 area 数据，则直接在当前页面匹配 职位 信息
		urlMatch := pageUrlRe.FindSubmatch(content)
		//fmt.Printf("%s %s\n", string(urlMatch[1]), string(urlMatch[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       "http://www.zhipin.com" + string(urlMatch[1]) + "?ka=" + string(urlMatch[2]),
			ParseFunc: ParsePositionList,
		})
	} else {
		// 否则匹配下一级 area 信息
		for _, item := range submatch {
			result.Requests = append(result.Requests, engine.Request{
				Url:       "https://www.zhipin.com" + string(item[1]) + "?ka=" + string(item[2]),
				ParseFunc: ParsePositionList,
			})
		}
	}
	return result
}
