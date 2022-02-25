package prototag

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

type messageVisitor struct {
	emptyVisitor
	fields []*Field
}

func (v *messageVisitor) VisitField(field *parser.Field) bool {
	v.fields = append(
		v.fields, &Field{
			Name:   field.FieldName,
			Number: intX(field.FieldNumber),
			Tags:   makeTags(field.InlineComment),
		},
	)
	return true
}
