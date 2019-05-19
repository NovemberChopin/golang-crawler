package parser

import (
	"crawler/zhipin/util"
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

func TestParseAreaList(t *testing.T) {
	bytes, err := ioutil.ReadFile("business_test_data.html")
	if err != nil {
		panic(err)
	}
	content := util.RemoveSpace(bytes)
	parseResult := ParseAreaList(content)
	for _, item := range parseResult.Requests {
		fmt.Println(item.Url)
	}
}

func TestPageUrlRe(t *testing.T) {
	var pageUrlRe = regexp.MustCompile(
		`<a href="(/c101[^"]+)" class="selected" ka="(sel-business-[^"]+)">[^<]+</a>`)
	bytes, err := ioutil.ReadFile("business_test_data.html")
	if err != nil {
		panic(err)
	}
	content := util.RemoveSpace(bytes)
	submatch := pageUrlRe.FindSubmatch(content)
	fmt.Println(string(submatch[0]))
}
