package usecases

import (
	"time"

	"github.com/YoshikiGit/poc-go-clean-architecture/adapters/repository"
	"github.com/YoshikiGit/poc-go-clean-architecture/entities"
)

// interfaceを返す
func NewCreateUserInteractor(
	repo repository.IUserRepository,
	presenter ICreateUserPresenter,
) ICreateUserUseCase {
	return createUserInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

type (
	// UseCase
	ICreateUserUseCase interface {
		CreateUserInteractor(CreateUserInput) (CreateUserOutput, error)
	}

	CreateUserInput struct {
		Name string `json:"name" validate:"required"`
		Age  int    `json:"age" validate:"required"`
	}

	ICreateUserPresenter interface {
		Output(entities.User) CreateUserOutput
	}

	CreateUserOutput struct {
		Id        string `json:"id" validate:"required"`
		Name      string `json:"name" validate:"required"`
		Age       int    `json:"age" validate:"required"`
		CreatedAt string `json:"created_at"`
	}

	createUserInteractor struct {
		repo      repository.IUserRepository
		presenter ICreateUserPresenter
	}
)

// UseCaseの実装部分
func (a createUserInteractor) CreateUserInteractor(input CreateUserInput) (CreateUserOutput, error) {
	var user = entities.NewUser(
		entities.UserID(entities.NewUUID()),
		input.Name,
		input.Age,
		time.Now(),
	)
	user, err := a.repo.Create(user)

	if err != nil {
		return a.presenter.Output(entities.User{}), err
	}

	return a.presenter.Output(user), nil
}
