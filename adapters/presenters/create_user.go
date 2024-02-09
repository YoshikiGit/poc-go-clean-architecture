package presenters

import (
	"time"

	"github.com/YoshikiGit/poc-go-clean-architecture/entities"
	"github.com/YoshikiGit/poc-go-clean-architecture/usecases"
)

func NewCreateUserPresenter() usecases.ICreateUserPresenter {
	return createUserPresenter{}
}

type createUserPresenter struct{}

func (a createUserPresenter) Output(user entities.User) usecases.CreateUserOutput {
	return usecases.CreateUserOutput{
		Id:        user.ID().String(),
		Name:      user.Name(),
		Age:       user.Age(),
		CreatedAt: user.CreatedAt().Format(time.RFC3339),
	}
}
