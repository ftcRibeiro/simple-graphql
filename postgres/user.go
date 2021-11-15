package postgres

import (
	"github.com/ftcRibeiro/simple-graphql/graph/model"
	"github.com/go-pg/pg/v10"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) GetByID(id string) (*model.User, error) {
	var user model.User
	err := u.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) GetUsers() ([]*model.User, error) {
	var users []*model.User
	err := u.DB.Model(&users).Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepo) CreateUser(user *model.User) (*model.User, error) {
	_, err := u.DB.Model(user).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return user, err
}
