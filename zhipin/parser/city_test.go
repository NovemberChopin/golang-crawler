package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseBusinessList(t *testing.T) {
	bytes, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}
	parseResult := ParseBusinessList(bytes)
	for _, item := range parseResult.Requests {
		fmt.Println(item.Url)
	}
}
