package in_memory

import (
	"keys-project/entities"
)

type InMemoryDataRepo struct {
	Data map[string]map[string]entities.Attribute
	Attrs map[string]entities.Attribute
}

func NewInMemoryDataRepo() *InMemoryDataRepo {
	return &InMemoryDataRepo{
		Data: make(map[string]map[string]entities.Attribute),
		Attrs: make (map[string]entities.Attribute),
	}
}

func (ds *InMemoryDataRepo) Add(obj *entities.Object)  {
	ds.Data[obj.Key] = obj.Value
}



