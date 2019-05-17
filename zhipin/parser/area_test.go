package parser

import (
	"crawler/zhipin/util"
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

func TestParsePositionList(t *testing.T) {
	bytes, err := ioutil.ReadFile("area_test_data.html")
	if err != nil {
		panic(err)
	}
	parseResult := ParsePositionList(bytes)

	for _, item := range parseResult.Items {
		fmt.Println(item.Payload)
	}
	for _, item := range parseResult.Requests {
		fmt.Println("下一页：" + item.Url)
	}
	fmt.Printf("一共 %d 条数据\n", len(parseResult.Items))
}

// 测试职位相关要求正则
func TestPositionMsgRe(t *testing.T) {
	var postionRe = regexp.MustCompile(
		`</div> </a> </h3> <p>([^<]+)<em class="vline"></em>([^<]+)<em class="vline"></em>([^<]+)</p>`)
	bytes, err := ioutil.ReadFile("area_test_data.html")
	if err != nil {
		panic(err)
	}
	content := util.RemoveSpace(bytes)
	//fmt.Println(string(content))
	submatch := postionRe.FindAllSubmatch(content, -1)
	for _, item := range submatch {
		fmt.Printf("地址：%s 工作经验：%s 学历要求：%s\n", string(item[1]), string(item[2]), string(item[3]))
	}
	fmt.Printf("一共 %d 条数据\n", len(submatch))
}

// 测试公司状况正则
func TestFimeMsgRe(t *testing.T) {
	// 匹配公司要求整段代码
	var postionRe = regexp.MustCompile(
		`target="_blank">[^<]+</a></h3> <p>([^p]+)</p>`)
	// 判断数据个数，
	var isIntegrityRe = regexp.MustCompile(`<em class="vline"></em>`)
	// 数据完整时的匹配
	var integrityRe = regexp.MustCompile(
		`<p>([^<]+)<em class="vline"></em>([^<]+)<em class="vline"></em>([^<]+)</p>`)
	// 数据不完整时的匹配（缺少融资信息）
	var unIntegrityRe = regexp.MustCompile(
		`<p>([^<]+)<em class="vline"></em>([^<]+)</p>`)

	bytes, err := ioutil.ReadFile("area_test_data.html")
	if err != nil {
		panic(err)
	}
	content := util.RemoveSpace(bytes)

	submatch := postionRe.FindAllSubmatch(content, -1)
	for _, item := range submatch {
		dataNum := isIntegrityRe.FindAllSubmatch(item[0], -1)
		if len(dataNum) == 1 { // 当缺少融资信息时
			submatch := unIntegrityRe.FindSubmatch(item[0])
			fmt.Printf("公司类型：%s 公司规模：%s\n",
				string(submatch[1]), string(submatch[2]))
		} else { // 当不缺少融资信息时
			submatch := integrityRe.FindSubmatch(item[0])
			fmt.Printf("公司类型：%s 融资阶段：%s 公司规模：%s\n",
				string(submatch[1]), string(submatch[2]), string(submatch[3]))
		}
	}
	fmt.Printf("一共 %d 条数据\n", len(submatch))
}

func TestNextPageRe(t *testing.T) {
	var nextPageRe = regexp.MustCompile(`<a href="(/c101[^"]+)" ka="page-next" class="next"></a>`)
	bytes, err := ioutil.ReadFile("area_test_data.html")
	if err != nil {
		panic(err)
	}
	content := util.RemoveSpace(bytes)
	submatch := nextPageRe.FindSubmatch(content)
	fmt.Println("http://www.zhipin.com" + string(submatch[1]))
}
