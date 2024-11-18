package usecase

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/entity"
	"github.com/auster-kaki/auster-mono/pkg/util/austerid"
	"github.com/auster-kaki/auster-mono/pkg/util/austerstorage"
)

type UserUseCase struct {
	repository repository.Repository
}

func NewUserUseCase(r repository.Repository) *UserUseCase {
	return &UserUseCase{repository: r}
}

func (u *UserUseCase) GetHobbies(ctx context.Context) (entity.Hobbies, error) {
	return u.repository.Hobby().GetAll(ctx)
}

type UsersGetOutput struct {
	Users entity.Users
}

func (u *UserUseCase) GetUsers(ctx context.Context) (*UsersGetOutput, error) {
	users, err := u.repository.User().GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return &UsersGetOutput{Users: users}, nil
}

type UserGetOutput struct {
	User    *entity.User
	Hobbies entity.Hobbies
	Photo   []byte
}

func (u *UserUseCase) GetUser(ctx context.Context, id entity.UserID) (*UserGetOutput, error) {
	user, err := u.repository.User().FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	photo, err := austerstorage.Get(user.ProfilePath)
	if err != nil {
		return nil, err
	}

	userHobbies, err := u.repository.UserHobby().GetByUserID(ctx, id)
	if err != nil {
		return nil, err
	}
	hobbies, err := u.repository.Hobby().GetByIDs(ctx, userHobbies.HobbyIDs())
	if err != nil {
		return nil, err
	}
	return &UserGetOutput{
		User:    user,
		Hobbies: hobbies,
		Photo:   photo,
	}, nil
}

type UserInput struct {
	ID     entity.UserID
	Name   string
	Gender string
	Age    int
	Photo  []byte
}

func (u *UserUseCase) CreateUser(ctx context.Context, input *UserInput) error {
	path, err := austerstorage.Save(austerstorage.JPEG, input.Photo)
	if err != nil {
		return err
	}

	return u.repository.User().Create(ctx, &entity.User{
		ID:          austerid.Generate[entity.UserID](),
		Name:        input.Name,
		Gender:      input.Gender,
		Age:         input.Age,
		ProfilePath: path,
	})
}

func (u *UserUseCase) UpdateUser(ctx context.Context, input *UserInput) error {
	path, err := austerstorage.Save(austerstorage.JPEG, input.Photo)
	if err != nil {
		return err
	}
	return u.repository.User().Update(ctx, &entity.User{
		ID:          input.ID,
		Name:        input.Name,
		Gender:      input.Gender,
		Age:         input.Age,
		ProfilePath: path,
	})
}
