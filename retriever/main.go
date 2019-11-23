package main

import (
	"fmt"
	"goLearn/retriever/mock"
	"goLearn/retriever/real"
	"time"
)

type Retriever interface {
	Get(Url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "goLearn",
		"course": "goLang",
	})
}

func download(r Retriever) string {
	return r.Get(url)
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > Type:%T Value:%v\n", r, r)
	fmt.Print(" > Type switch: ")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	var r Retriever
	mockRetriever := mock.Retriever{Contents: "this is a fake imooc.com"}
	r = &mockRetriever

	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozelle/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	fmt.Println(
		"Try a session with mockRetriever")
	fmt.Println(session(&mockRetriever))

	//fmt.Println(download(r))
}
