package app

type Modeler interface {
	Resource() interface{}
	SetResource(interface{})
}

type Model struct {
	resource interface{}
}

func (m *Model) SetResource(r interface{}) {
	m.resource = r
}

func (m *Model) Resource() interface{} {
	return m.resource
}
