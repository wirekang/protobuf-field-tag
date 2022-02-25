package prototag_test

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirekang/prototag/pkg/prototag"
)

//go:embed data.proto
var data []byte

func TestReader(t *testing.T) {
	m, err := prototag.Parse(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	{
		assert.Len(t, m.Messages, 1)
		msg := m.Messages[0]
		msg.Cache()
		assert.Equal(t, msg.Name, "Message1")
		assert.Len(t, msg.Fields, 4)
		assert.Len(t, msg.Tags, 0)

		assertFields(t, &msg.Struct, []string{"first", "second", "t", "asdf"}, []int{1, 2, 4, 8})
		assert.Len(t, msg.Fields[0].Tags, 2)

		field := msg.Fields[0]
		{
			keys := []string{"json", "int"}
			values := []string{"omitempty,foo,string", "123"}
			ints := []int{0, 123}
			for i, tag := range field.Tags {
				assert.Equal(t, keys[i], tag.Key)
				assert.Equal(t, values[i], string(tag.Value))
				assert.Equal(t, ints[i], tag.Value.IntX())
			}
		}
	}

	{
		assert.Len(t, m.Enums, 1)
		enum := m.Enums[0]
		enum.Cache()
		assert.Len(t, enum.Tags, 2)
		assert.Equal(t, enum.Tags[0].Key, "key")
		assert.Equal(t, enum.Tags[1].Key, "key2")

		assert.Equal(t, string(enum.Tags[0].Value), "value")
		assert.Equal(t, string(enum.Tags[1].Value), "value 2")

		assert.Equal(t, enum.Name, "Enum1")
		assert.Len(t, enum.Fields, 3)
		assertFields(t, &enum.Struct, []string{"Asdf", "Qwer", "Zsdf"}, []int{0, 1, 2})
	}

}

func assertFields(t *testing.T, s *prototag.Struct, names []string, numbers []int) {
	{
		for i, field := range s.Fields {
			assert.Equal(t, names[i], field.Name)
			assert.Equal(t, numbers[i], field.Number)
			assert.Equal(t, field, s.Field(field.Name))
		}
	}
}
