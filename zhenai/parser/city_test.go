package parser

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

func TestParseCity(t *testing.T) {
	var nextPageUrlRe = regexp.MustCompile(
		`<li class="paging-item"><a href="(http://www.zhenai.com/zhenghun/[^"]+)">下一页</a>`)
	contents, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}
	findSubmatch := nextPageUrlRe.FindSubmatch(contents)

	fmt.Printf("%s, %s\n", string(findSubmatch[0]), string(findSubmatch[1]))
}
