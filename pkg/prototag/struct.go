package prototag

type Enum struct {
	Struct
}

type Message struct {
	Struct
}

type Struct struct {
	Name      string   `json:"name"`
	Fields    []*Field `json:"fields"`
	isCached  bool
	nameField map[string]*Field
}

type Field struct {
	Name     string `json:"name"`
	Number   int    `json:"number"`
	Tags     []*Tag `json:"tags"`
	isCached bool
	keyTag   map[string]*Tag
}

func (m *Struct) Cache() {
	if m.isCached {
		return
	}

	m.isCached = true
	m.nameField = make(map[string]*Field, len(m.Fields))
	for _, field := range m.Fields {
		m.nameField[field.Name] = field
	}
}

func (m *Struct) Field(name string) *Field {
	if !m.isCached {
		return nil
	}

	return m.nameField[name]
}

func (f *Field) Cache() {
	if f.isCached {
		return
	}

	f.isCached = true
	f.keyTag = make(map[string]*Tag, len(f.Tags))
	for _, tag := range f.Tags {
		f.keyTag[tag.Key] = tag
	}
}

func (f *Field) Tag(key string) *Tag {
	if !f.isCached {
		return nil
	}

	return f.keyTag[key]
}
