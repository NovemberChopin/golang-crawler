package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPraseCityList(t *testing.T) {
	bytes, err := ioutil.ReadFile("home_test_data.html")
	if err != nil {
		panic(err)
	}
	parseResult := PraseCityList(bytes)
	for _, item := range parseResult.Requests {
		fmt.Println(item.Url)
	}
}
