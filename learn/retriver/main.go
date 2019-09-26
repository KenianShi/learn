package main

import (
	"fmt"
	"shikenian/learn/retriver/mock"
	"shikenian/learn/retriver/real"
)

const url = "http://www.imooc.com"

type Retriver interface {
	Get(url string) string
}

func download(r Retriver) string {
	return r.Get(url)
}

type Post interface {
	Post(url string, m map[string]string) string
}

type PostRetriver interface {
	Retriver
	Post
}

func session(s PostRetriver) string {
	s.Post(url, map[string]string{"content": "this is an another faked imooc.com"})
	return s.Get(url)
}

func main() {
	var r Retriver
	retriver := mock.Retriver{"this is a fake imooc"} //一个结构实现两个接口
	r = &retriver
	inspect(r)
	a := session(&retriver)
	fmt.Println("return session:", a)
}

func inspect(r Retriver) {
	fmt.Printf(" > %T : %v \n", r, r)
	switch v := r.(type) {
	case *mock.Retriver:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriver:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
