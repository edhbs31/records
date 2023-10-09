package repocontainer

import (
	"record/app/domains"
	"record/app/repository"

	"gorm.io/gorm"
)

type RepoContainer struct {
	DataRepository domains.DataRepository
}

func NewRepoContainer(db *gorm.DB) *RepoContainer {
	containerRepo := &RepoContainer{}
	containerRepo.DataRepository = repository.NewDataRepository(db)
	return containerRepo
}
