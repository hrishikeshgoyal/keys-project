package in_memory

import (
	"fmt"
	"github.com/jwangsadinata/go-multimap/slicemultimap"
	"keys-project/entities"
)

type InMemoryIndexRepo struct {
	Index map[string]*Index
}

func NewInMemoryIndexRepo() *InMemoryIndexRepo {
	return &InMemoryIndexRepo{
		Index: make (map[string]*Index),
	}
}

type Index struct {
	mp *slicemultimap.MultiMap
	itype string
}

func(r *InMemoryIndexRepo) Create(key string, a entities.Attribute) {
	if _, ok := r.Index[key]; !ok {
		r.Index[key] = &Index{
			mp:    slicemultimap.New(),
			itype: a.Type,
		}
	}
}


func (r *InMemoryIndexRepo) AddTo(attrKey string, attrVal string, dataKey string) error  {
	index, ok := r.Index[attrKey]
	if !ok {
		return fmt.Errorf("index not found")
	}

	index.mp.Put(attrVal, dataKey)

	return nil

}

func(r *InMemoryIndexRepo) List(attrKey string) ([]string, error)  {
	index, found := r.Index[attrKey]
	if!found {
		return nil, fmt.Errorf("no index found")
	}
	val, found := index.mp.Get(attrKey)
	if found {
		var ret []string
		for _, item := range val{
			ret = append(ret, item.(string))
		}
		return ret, nil
	}
	return nil, fmt.Errorf("not found")
}