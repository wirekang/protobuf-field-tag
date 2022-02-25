package prototag

type Model struct {
	Messages    []*Message `json:"messages,omitempty"`
	Enums       []*Enum    `json:"enums,omitempty"`
	nameMessage map[string]*Message
	nameEnum    map[string]*Enum
	isCached    bool
}

func (m *Model) Cache() {
	if m.isCached {
		return
	}

	m.isCached = true
	m.nameMessage = make(map[string]*Message, len(m.Messages))
	m.nameEnum = make(map[string]*Enum, len(m.Enums))

	for _, msg := range m.Messages {
		msg.Cache()
		m.nameMessage[msg.Name] = msg
	}

	for _, enum := range m.Enums {
		enum.Cache()
		m.nameEnum[enum.Name] = enum
	}
}

func (m *Model) Message(name string) *Message {
	if !m.isCached {
		return nil
	}

	return m.nameMessage[name]
}

func (m *Model) Enum(name string) *Enum {
	if !m.isCached {
		return nil
	}

	return m.nameEnum[name]
}
