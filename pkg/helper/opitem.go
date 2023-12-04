package helper

type opitem struct {
	ID     string    `json:"id"`
	Fields []opfield `json:"fields"`
}

func (i *opitem) matches(a *Attributes) bool {
	// required
	if f := i.field("hostname"); f == nil || f.Value == "" || f.Value != a.Host {
		return false
	}
	// optional
	if f := i.field("path"); f != nil && f.Value != "" && a.Path != "" && f.Value != a.Path {
		return false
	}
	if f := i.field("username"); f != nil && f.Value != "" && a.Username != "" && f.Value != a.Username {
		return false
	}
	return true
}

func (i *opitem) field(name string) *opfield {
	var f *opfield
	for _, field := range i.Fields {
		if field.ID == name || field.Label == name {
			f = &field
			break
		}
	}
	return f
}

type opfield struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Value string `json:"value"`
}
