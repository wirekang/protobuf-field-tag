package prototag

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

type enumVisitor struct {
	emptyVisitor
	fields []*Field
}

func (v *enumVisitor) VisitEnumField(field *parser.EnumField) bool {
	v.fields = append(
		v.fields, &Field{
			Name:   field.Ident,
			Number: intX(field.Number),
			Tags:   makeTags(field.InlineComment),
		},
	)
	return true
}
