package fetcher

import (
	"fmt"
	"testing"
)

func TestFetcher(t *testing.T) {
	content, err := Fetcher("http://www.zhenai.com/zhenghun")
	if err != nil {
		t.Errorf("err:%v", err)
	}
	fmt.Println(string(content))
}
