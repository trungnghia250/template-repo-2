package usecase

import (
	"SecondAssignment/service/model"
	"SecondAssignment/service/model/dto"
	"SecondAssignment/service/repo"
	"context"
	"github.com/google/martian/log"
	"time"
)

type UserUseCase interface {
	GetUserByID(ctx context.Context, id string) (dto.GetUserByIDResponse, error)
	GetAllUser(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, req dto.CreateUserRequest) error
	UpdateUser(ctx context.Context, req dto.UpdateUserRequest) error
	DeleteUser(ctx context.Context, id string) error
}

type userUseCase struct {
	repo repo.UserRepo
}

func NewUserUseCase(repo repo.UserRepo) UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (u *userUseCase) GetAllUser(ctx context.Context) ([]model.User, error) {
	var (
		err  error
		resp []model.User
	)
	resp, err = u.repo.GetAll(ctx)
	if err != nil {
		log.Errorf("GetAllUser fail, err", err)
		return resp, err
	}

	return resp, err
}

func (u *userUseCase) GetUserByID(ctx context.Context, id string) (dto.GetUserByIDResponse, error) {
	var (
		resp dto.GetUserByIDResponse
		err  error
	)

	user, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		log.Errorf("GetUserByID with ID %s fail, err", id, err)
		return resp, err
	}

	return dto.GetUserByIDResponse{
		ID:          user.ID,
		UserName:    user.UserName,
		Email:       user.Email,
		Phone:       user.Phone,
		DateOfBirth: user.DateOfBirth,
	}, nil
}

func (u *userUseCase) CreateUser(ctx context.Context, req dto.CreateUserRequest) error {
	user := model.User{
		ID:          req.ID,
		UserName:    req.UserName,
		Email:       req.Email,
		Phone:       req.Phone,
		DateOfBirth: req.DateOfBirth,
		CreatedAt:   time.Now(),
	}
	err := u.repo.CreateUser(ctx, &user)
	if err != nil {
		log.Errorf("CreateUser with userID %s fail, err", req.ID, err)
		return err
	}

	return nil
}

func (u *userUseCase) UpdateUser(ctx context.Context, req dto.UpdateUserRequest) error {
	var user model.User
	if len(req.ID) > 0 {
		user.ID = req.ID
	}
	if len(req.UserName) > 0 {
		user.UserName = req.UserName
	}
	if len(req.Phone) > 0 {
		user.Phone = req.Phone
	}
	if req.DateOfBirth.Unix() > 0 {
		user.DateOfBirth = req.DateOfBirth
	}
	err := u.repo.Update(ctx, &user)
	if err != nil {
		log.Errorf("UpdateUser with userID %s fail, err", req.ID, err)
		return err
	}

	return nil
}

func (u *userUseCase) DeleteUser(ctx context.Context, id string) error {
	user, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		log.Errorf("GetUserByID with ID %s fail, err", id, err)
		return err
	}

	err = u.repo.Delete(ctx, user)
	if err != nil {
		log.Errorf("DeleteUser with userID %s fail, err", id, err)
		return err
	}

	return nil
}
