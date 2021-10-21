package repository

import "keys-project/entities"

type DataRepo interface{
	Add(obj *entities.Object)
	//Update()
	//Delete()

}

