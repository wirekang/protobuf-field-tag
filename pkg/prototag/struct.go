package prototag

type Enum struct {
	Struct
}

type Message struct {
	Struct
}

type Struct struct {
	Name      string   `json:"name,omitempty"`
	Fields    []*Field `json:"fields,omitempty"`
	Tags      []*Tag   `json:"tags,omitempty"`
	isCached  bool
	nameField map[string]*Field
	keyTag    map[string]*Tag
}

type Field struct {
	Name     string `json:"name,omitempty"`
	Number   int    `json:"number,omitempty"`
	Tags     []*Tag `json:"tags,omitempty"`
	isCached bool
	keyTag   map[string]*Tag
}

func (s *Struct) Cache() {
	if s.isCached {
		return
	}

	s.isCached = true
	s.keyTag = make(map[string]*Tag, len(s.Tags))
	for _, tag := range s.Tags {
		s.keyTag[tag.Key] = tag
	}

	s.nameField = make(map[string]*Field, len(s.Fields))
	for _, field := range s.Fields {
		field.Cache()
		s.nameField[field.Name] = field
	}
}

func (s *Struct) Field(name string) *Field {
	if !s.isCached {
		return nil
	}

	return s.nameField[name]
}

func (s *Struct) Tag(key string) *Tag {
	if !s.isCached {
		return nil
	}

	return s.keyTag[key]
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
