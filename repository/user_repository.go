package repository

import (
	"context"
	"simple-bank/ent"
)

type CreateUserDto struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=8"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type GetUserDto struct {
	Username string `json:"username" binding:"required,alphanum"`
}

type IUserRepo interface {
	GetUser(ctx context.Context, dto GetUserDto) (user *ent.User, err error)
	GetUsers(ctx context.Context) (users []*ent.User, err error)
	CreateUser(ctx context.Context, dto CreateUserDto) (user *ent.User, err error)
}

type UserRepo struct {
	dbClient *ent.Client
}

func NewUserRepo(dbClient *ent.Client) IUserRepo {
	return &UserRepo{dbClient: dbClient}
}

func (u *UserRepo) GetUser(ctx context.Context, dto GetUserDto) (user *ent.User, err error) {
	user, err = u.dbClient.User.Get(ctx, dto.Username)
	return
}

func (u *UserRepo) GetUsers(ctx context.Context) (users []*ent.User, err error) {
	users, err = u.dbClient.User.Query().All(ctx)
	return
}

func (u *UserRepo) CreateUser(ctx context.Context, dto CreateUserDto) (user *ent.User, err error) {
	user, err = u.dbClient.User.Create().SetPassword(dto.Password).
		SetEmail(dto.Email).SetFullName(dto.FullName).
		SetID(dto.FullName).
		Save(ctx)
	return
}
