package prototag

import (
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

type emptyVisitor struct {
}

func (v emptyVisitor) VisitEnum(enum *parser.Enum) (next bool) {
	return true
}

func (v emptyVisitor) VisitMessage(message *parser.Message) (next bool) {
	return true
}

func (v emptyVisitor) VisitComment(comment *parser.Comment) {}

func (v emptyVisitor) VisitEmptyStatement(statement *parser.EmptyStatement) (next bool) {
	return true
}

func (v emptyVisitor) VisitEnumField(field *parser.EnumField) (next bool) {
	return true
}

func (v emptyVisitor) VisitExtend(extend *parser.Extend) (next bool) {
	return true
}

func (v emptyVisitor) VisitExtensions(extensions *parser.Extensions) (next bool) {
	return true
}

func (v emptyVisitor) VisitField(field *parser.Field) (next bool) {
	return true
}

func (v emptyVisitor) VisitGroupField(field *parser.GroupField) (next bool) {
	return true
}

func (v emptyVisitor) VisitImport(i *parser.Import) (next bool) {
	return true
}

func (v emptyVisitor) VisitMapField(field *parser.MapField) (next bool) {
	return true
}

func (v emptyVisitor) VisitOneof(oneof *parser.Oneof) (next bool) {
	return true
}

func (v emptyVisitor) VisitOneofField(field *parser.OneofField) (next bool) {
	return true
}

func (v emptyVisitor) VisitOption(option *parser.Option) (next bool) {
	return true
}

func (v emptyVisitor) VisitPackage(p *parser.Package) (next bool) {
	return true
}

func (v emptyVisitor) VisitReserved(reserved *parser.Reserved) (next bool) {
	return true
}

func (v emptyVisitor) VisitRPC(rpc *parser.RPC) (next bool) {
	return true
}

func (v emptyVisitor) VisitService(service *parser.Service) (next bool) {
	return true
}

func (v emptyVisitor) VisitSyntax(syntax *parser.Syntax) (next bool) {
	return true
}
