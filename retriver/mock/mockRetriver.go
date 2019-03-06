package mock

import "fmt"

type Retriver struct {
	Contents string
}

// 修改string方法,打印的时候将会出现这个
// 如果没有实现string方法的话，将会打印{"contents":"???"}
func (r *Retriver) String() string {
	return `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
}

//an POST method created by shikenian
func (r *Retriver) Post(url string, m map[string]string) string {
	r.Contents = m["content"]
	return "ok"
}

func (r *Retriver) Get(url string) string {
	fmt.Println("get URL:", url)
	return r.Contents
}
