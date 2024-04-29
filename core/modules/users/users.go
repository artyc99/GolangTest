package users

import (
	"EchoAPI/core/models"
	"EchoAPI/core/repositories"
	"context"
)

type Users struct {
	usersRepo repositories.UsersI
}

func New(repoFactory repositories.RepoFactoryI) *Users {
	return &Users{
		usersRepo: repoFactory.Users(),
	}
}

func (m *Users) Create(ctx context.Context, req CreateReq) (_ CreateRes, err error) {
	user := models.User{
		ID:            req.ID,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Age:           req.Age,
		RecordingDate: req.RecordingDate,
	}

	err = m.usersRepo.Create(ctx, user)
	if err != nil {
		return CreateRes{}, err
	}

	return CreateRes{}, nil
}

func (m *Users) GetList(ctx context.Context, req GetListReq) (res GetListRes, err error) {
	props := repositories.FilterProps{
		FromAge: req.FromAge,
		ToAge:   req.ToAge,

		FromDate: req.FromDate,
		ToDate:   req.ToDate,
	}

	users, total, err := m.usersRepo.GetList(ctx, props)
	if err != nil {
		return GetListRes{}, err
	}

	usersRes := make([]User, 0)
	for _, user := range users {
		userRes := User{
			ID:            user.ID,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			Age:           user.Age,
			RecordingDate: user.RecordingDate,
		}
		usersRes = append(usersRes, userRes)
	}

	res.Users = usersRes
	res.Total = total

	return res, nil
}
