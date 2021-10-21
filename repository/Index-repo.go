package repository

import "keys-project/entities"

type IndexRepo interface {
	Create(key string, a entities.Attribute)
	AddTo(attrKey string, attrVal string, dataKey string) error
	List(attrKey string) ([]string, error)
}