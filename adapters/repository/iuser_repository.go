package repository

import (
	"github.com/YoshikiGit/poc-go-clean-architecture/entities"
)

type (
	IUserRepository interface {
		Create(entities.User) (entities.User, error)
	}
)
