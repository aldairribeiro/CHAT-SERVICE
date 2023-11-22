package entity

type Model struct {
	Name string
	Maxtokens int
}

func NewModel (name string, maxTokes int) *Model {
	return &Model{
		Name: name,
		Maxtokens:maxTokes,
	}
}

func (m *Model) GetMaxTokens() int{
	return m.Maxtokens
}

func (m *Model) GetModelName() string  {
	return m.Name
}