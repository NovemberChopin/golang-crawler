package parser

import (
	"crawler/engine"
	"crawler/zhipin/model"
	"crawler/zhipin/util"
	"regexp"
)

var nameRe = regexp.MustCompile(`<div class="job-title">([^<]+)</div>`)
var paymentRe = regexp.MustCompile(`<span class="red">([^<]+)</span>`)

// 职位要求正则
var requireRe = regexp.MustCompile(
	`</div> </a> </h3> <p>([^<]+)<em class="vline"></em>([^<]+)<em class="vline"></em>([^<]+)</p>`)
var positionUrl = regexp.MustCompile(`<a href="(/job_detail/[^"]+)"`)
var fimeNameRe = regexp.MustCompile(
	`<h3 class="name"><a href="(/gongsi/[^"]+)" ka="(search_list_company_[^"]+)" target="_blank">([^<]+)</a></h3>`)

// 匹配公司要求整段代码
var positionRe = regexp.MustCompile(
	`target="_blank">[^<]+</a></h3> <p>([^p]+)</p>`)

// 判断数据个数，
var isIntegrityRe = regexp.MustCompile(`<em class="vline"></em>`)

// 数据完整时的匹配
var integrityRe = regexp.MustCompile(
	`<p>([^<]+)<em class="vline"></em>([^<]+)<em class="vline"></em>([^<]+)</p>`)

// 数据不完整时的匹配（缺少融资信息）
var unIntegrityRe = regexp.MustCompile(
	`<p>([^<]+)<em class="vline"></em>([^<]+)</p>`)

// 下一页正则
var nextPageRe = regexp.MustCompile(
	`<a href="(/c101[^"]+)" ka="page-next" class="next"></a>`)

// 解析每个 area 的职位
func ParsePositionList(bytes []byte) engine.ParseResult {

	//defer func() {
	//	err := recover()
	//	if err, ok := err.(error); ok {
	//		log.Println("Error occurred:", err)
	//	} else {
	//		panic(fmt.Sprintf("I do't know what to do %v", err))
	//	}
	//}()

	content := util.RemoveSpace(bytes)
	position := model.Position{}

	nameMatch := nameRe.FindAllSubmatch(content, -1)
	paymentMatch := paymentRe.FindAllSubmatch(content, -1)
	requireMatch := requireRe.FindAllSubmatch(content, -1)
	positionUrl := positionUrl.FindAllSubmatch(content, -1)
	fimeNameMatch := fimeNameRe.FindAllSubmatch(content, -1)

	submatch := positionRe.FindAllSubmatch(content, -1)
	nextPageUrl := nextPageRe.FindSubmatch(content) // 解析下一页Url

	items := []engine.Item{}
	requests := []engine.Request{}

	if len(nameMatch) == len(requireMatch) {
		for k, item := range nameMatch {
			position.Name = string(item[1])
			position.Payment = string(paymentMatch[k][1])
			position.Address = string(requireMatch[k][1])
			position.Experience = string(requireMatch[k][2])
			position.Education = string(requireMatch[k][3])
			position.PosiUrl = "http://www.zhipin.com" + string(positionUrl[k][1])

			position.FimeUrl = "http://www.zhipin.com" + string(fimeNameMatch[k][1]) + "?ka=" + string(fimeNameMatch[k][2])
			position.FirmName = string(fimeNameMatch[k][3])

			position.FirmType, position.FirmFinancing, position.FirmSize = extractFirmMsg(submatch[k])

			items = append(items, engine.Item{
				Url:     "",
				Type:    "zhipin_jinan",
				Id:      "",
				Payload: position,
			})
		}
	}

	// 请求下一页
	if len(nextPageUrl) != 0 {
		requests = append(requests, engine.Request{
			Url:       "http://www.zhipin.com" + string(nextPageUrl[1]),
			ParseFunc: ParsePositionList,
		})
	}
	// 本页解析数据
	result := engine.ParseResult{
		Items:    items,
		Requests: requests,
	}

	return result
}

// 传入公司信息整段代码，提取公司类型、融资阶段、公司规模
func extractFirmMsg(item [][]byte) (string, string, string) {
	dataNum := isIntegrityRe.FindAllSubmatch(item[0], -1)
	if len(dataNum) == 1 { // 当缺少融资信息时
		submatch := unIntegrityRe.FindSubmatch(item[0])
		return string(submatch[1]), "", string(submatch[2])
	} else { // 当不缺少融资信息时
		submatch := integrityRe.FindSubmatch(item[0])
		return string(submatch[1]), string(submatch[2]), string(submatch[3])
	}
}
