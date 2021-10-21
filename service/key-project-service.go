package service

import (
	"keys-project/entities"
	"keys-project/repository"
)

type KeyProjectService struct {
	dataRepo repository.DataRepo
	indexRepo repository.IndexRepo
}

func NewKeyProjectService(dataRepo repository.DataRepo, indexRepo repository.IndexRepo) *KeyProjectService {
	return &KeyProjectService{
		dataRepo:  dataRepo,
		indexRepo: indexRepo,
	}
}

func (r *KeyProjectService)Add(obj *entities.Object) error {
	r.dataRepo.Add(obj)
	for k, v := range obj.Value{
		r.indexRepo.Create(obj.Key, v)
		r.indexRepo.AddTo(k, v.Value, obj.Key)
	}
	return nil
}


