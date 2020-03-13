package service

import (
	"gitlab.intelligentb.com/devops/bplus/stm"
	"gitlab.intelligentb.com/devops/__package_name__/model"
)


// a dummy repo implemented as a map
type __PackageName__Repo struct {
	entities map[string]*model.__PackageName__
}

func (or *__PackageName__Repo) Create(se stm.StateEntity) (stm.StateEntity, error) {
	ent := se.(*model.__PackageName__)
	ent.ID = "first"
	repo.entities[ent.ID] = ent
	return ent, nil
}

func (or *__PackageName__Repo) Retrieve(ID string) (stm.StateEntity, error) {
	return repo.entities[ID], nil
}

// repo - the repository used for persisting the data into the store
var repo = &__PackageName__Repo{
	entities: make(map[string]*model.__PackageName__),
}
