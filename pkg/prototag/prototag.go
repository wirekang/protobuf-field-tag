package prototag

import (
	"io"
	"os"

	"github.com/yoheimuta/go-protoparser/v4"
)

func ParseFile(name string) (m *Model, err error) {
	f, err := os.Open(name)
	if err != nil {
		return
	}

	return Parse(f)
}

func Parse(r io.Reader) (m *Model, err error) {
	p, err := protoparser.Parse(r, protoparser.WithBodyIncludingComments(true))
	if err != nil {
		panic(err)
	}

	rv := &rootVisitor{}
	p.Accept(rv)
	m = &Model{
		Messages: rv.messages,
		Enums:    rv.enums,
	}
	return
}
