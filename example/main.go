package main

import (
	"fmt"

	"github.com/wirekang/prototag/pkg/prototag"
)

func main() {
	m, err := prototag.ParseFile("example.proto")
	if err != nil {
		return
	}

	tag := m.Messages[0].Tags[0]
	fmt.Println(tag.Key, tag.Value)

	// for Message(name) Field(name) Tag(name) instead of Message[n]
	m.Cache()

	fmt.Println(m.Message("Person").Field("email").Tag("key").Value)

}
