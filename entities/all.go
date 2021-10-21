package entities

type Attribute struct {
	Value string
	Type string
}

type Object struct {
	Key string
	Value map[string]Attribute
}


