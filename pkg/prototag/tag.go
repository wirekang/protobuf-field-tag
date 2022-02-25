package prototag

import (
	"strconv"
	"strings"

	"github.com/fatih/structtag"
	"github.com/yoheimuta/go-protoparser/v4/parser"
)

type Tag struct {
	Key   string   `json:"key,omitempty"`
	Value TagValue `json:"value,omitempty"`
}

type TagValue string

func (t TagValue) Int() (int, error) {
	return strconv.Atoi(string(t))
}

func (t TagValue) IntX() (i int) {
	return intX(string(t))
}

func intX(v string) (i int) {
	i, _ = strconv.Atoi(v)
	return
}

func makeTags(c *parser.Comment) (r []*Tag) {
	if c == nil {
		r = []*Tag{}
		return
	}

	lines := c.Lines()
	if len(lines) == 0 {
		r = []*Tag{}
		return
	}

	line := lines[0]
	line = insideBacktick(line)
	if line == "" {
		r = []*Tag{}
		return
	}

	tgs, err := structtag.Parse(line)
	if err != nil {
		r = []*Tag{}
		return
	}

	r = make([]*Tag, tgs.Len())

	for i, t := range tgs.Tags() {
		r[i] = &Tag{
			Key:   t.Key,
			Value: TagValue(t.Value()),
		}
	}
	return
}

func insideBacktick(v string) string {
	first := strings.IndexByte(v, '`')
	last := strings.LastIndexByte(v, '`')
	if first == last {
		return ""
	}

	return v[first+1 : last]
}
