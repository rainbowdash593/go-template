package repo

import (
	"context"
	"errors"
	"example/template/internal/adapters/repo/models"
	"example/template/internal/domain/dto"
	"example/template/internal/domain/entity"
	"example/template/internal/domain/exceptions"
	"example/template/pkg/database"
	"example/template/pkg/utils"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *database.Database
}

func NewUserRepo(db *database.Database) *UserRepo {
	return &UserRepo{db}
}

func (r UserRepo) Find(ctx context.Context, filter dto.UserFilter) (*entity.User, error) {
	var (
		user models.User
		err  error
	)
	result := r.db.Where(filter).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err = exceptions.ErrUserNotFound
		} else {
			err = utils.WrapErrors(exceptions.ErrUnhandled, result.Error)
		}
		return nil, err
	}

	return &entity.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (r UserRepo) Create(ctx context.Context, d dto.CreateUser) (*entity.User, error) {
	var (
		user models.User
		err  error
	)
	existingUser, err := r.Find(ctx, dto.UserFilter{Email: d.Email})

	if err != nil && !errors.Is(err, exceptions.ErrUserNotFound) {
		return nil, err
	}

	if existingUser != nil {
		return nil, exceptions.ErrUserAlreadyExists
	}

	user = models.User{Email: d.Email, Name: d.Name}
	result := r.db.Create(&user)

	if result.Error != nil {
		return nil, utils.WrapErrors(exceptions.ErrUnhandled, err)
	}

	return &entity.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
