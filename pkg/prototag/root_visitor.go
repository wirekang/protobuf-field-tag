package prototag

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

type rootVisitor struct {
	emptyVisitor
	messages []*Message
	enums    []*Enum
}

func newStruct(name string, comments []*parser.Comment, fields []*Field) Struct {
	var c *parser.Comment
	if len(comments) > 1 {
		c = comments[len(comments)-1]
	}
	return Struct{
		Name:   name,
		Fields: fields,
		Tags:   makeTags(c),
	}
}

func (v *rootVisitor) VisitMessage(msg *parser.Message) bool {
	mv := &messageVisitor{}
	msg.Accept(mv)
	v.messages = append(v.messages, &Message{Struct: newStruct(msg.MessageName, msg.Comments, mv.fields)})
	return true
}

func (v *rootVisitor) VisitEnum(enum *parser.Enum) bool {
	ev := &enumVisitor{}
	enum.Accept(ev)
	v.enums = append(v.enums, &Enum{Struct: newStruct(enum.EnumName, enum.Comments, ev.fields)})
	return true
}
