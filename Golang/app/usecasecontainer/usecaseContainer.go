package usecasecontainer

import (
	"record/app/domains"
	"record/app/repocontainer"
	"record/app/usecase"
)

type UsecaseContainer struct {
	DataUsecase domains.DataUsecase
}

func NewUsecaseContainer(rp *repocontainer.RepoContainer) *UsecaseContainer {
	return &UsecaseContainer{
		DataUsecase: usecase.NewDataUsecase(rp),
	}
}
