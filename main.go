package main

import (
	"keys-project/entities"
	in_memory "keys-project/repository/in-memory"
	service2 "keys-project/service"
)

func main()  {
	dr := in_memory.NewInMemoryDataRepo()
	ir := in_memory.NewInMemoryIndexRepo()

	service := service2.NewKeyProjectService(dr, ir)

	obj := &entities.Object{
		Key:   "delhi",
		Value: map[string]entities.Attribute{
			"pl": entities.Attribute{
				Value: "v high",
				Type:  "str",
			},
			"po": entities.Attribute{
				Value: "10M",
				Type:  "str",
			},
		},
	}

	obj2 := &entities.Object{
		Key:   "jakarta",
		Value: map[string]entities.Attribute{
			"pl": entities.Attribute{
				Value: "high",
				Type:  "",
			},
			"lo": entities.Attribute{
				Value: "106",
				Type:  "str",
			},
		},
	}

	service.Add(obj)
	service.Add(obj2)
}
