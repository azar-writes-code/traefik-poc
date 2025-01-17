package services

import (
	"github.com/azar-writes-code/traefik-poc/user-service/pkg/rest/server/daos"
	"github.com/azar-writes-code/traefik-poc/user-service/pkg/rest/server/models"
)

type UserService struct {
	userDao *daos.UserDao
}

func NewUserService() (*UserService, error) {
	userDao, err := daos.NewUserDao()
	if err != nil {
		return nil, err
	}
	return &UserService{
		userDao: userDao,
	}, nil
}

func (userService *UserService) CreateUser(user *models.User) (*models.User, error) {
	return userService.userDao.CreateUser(user)
}

func (userService *UserService) GetUser(id string) (*models.User, error) {
	return userService.userDao.GetUser(id)
}

func (userService *UserService) UpdateUser(id string, user *models.User) (*models.User, error) {
	return userService.userDao.UpdateUser(id, user)
}

func (userService *UserService) DeleteUser(id string) error {
	return userService.userDao.DeleteUser(id)
}
