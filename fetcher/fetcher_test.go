package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	content, err := Fetch("http://zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", content)
}
