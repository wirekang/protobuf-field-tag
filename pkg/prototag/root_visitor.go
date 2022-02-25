package prototag

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

type rootVisitor struct {
	emptyVisitor
	messages []*Message
	enums    []*Enum
}

func (v *rootVisitor) VisitMessage(msg *parser.Message) bool {
	m := &Message{}
	m.Name = msg.MessageName
	mv := &messageVisitor{}
	msg.Accept(mv)
	m.Fields = mv.fields
	v.messages = append(v.messages, m)
	return true
}

func (v *rootVisitor) VisitEnum(enum *parser.Enum) bool {
	e := &Enum{}
	e.Name = enum.EnumName
	ev := &enumVisitor{}
	enum.Accept(ev)
	e.Fields = ev.fields
	v.enums = append(v.enums, e)
	return true
}
