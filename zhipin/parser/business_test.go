package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseAreaList(t *testing.T) {
	bytes, err := ioutil.ReadFile("business_test_data.html")
	if err != nil {
		panic(err)
	}

	parseResult := ParseAreaList(bytes)
	for _, item := range parseResult.Requests {
		fmt.Println(item.Url)
	}
}
