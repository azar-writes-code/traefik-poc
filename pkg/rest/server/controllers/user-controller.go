package controllers

import (
	"errors"
	"net/http"

	"github.com/azar-writes-code/traefik-poc/user-service/pkg/rest/server/daos/clients/nosqls"
	"github.com/azar-writes-code/traefik-poc/user-service/pkg/rest/server/models"
	"github.com/azar-writes-code/traefik-poc/user-service/pkg/rest/server/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() (*UserController, error) {
	userService, err := services.NewUserService()
	if err != nil {
		return nil, err
	}
	return &UserController{
		userService: userService,
	}, nil
}

func (userController *UserController) CreateUser(context *gin.Context) {
	// validate input
	var input models.User
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// trigger user creation
	userCreated, err := userController.userService.CreateUser(&input)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, userCreated)
}

func (userController *UserController) FetchUser(context *gin.Context) {
	// trigger user fetching
	user, err := userController.userService.GetUser(context.Param("id"))
	if err != nil {
		log.Error(err)
		if errors.Is(err, nosqls.ErrNotExists) {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, nosqls.ErrInvalidObjectID) {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, user)
}

func (userController *UserController) UpdateUser(context *gin.Context) {
	// validate input
	var input models.User
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// trigger user update
	if _, err := userController.userService.UpdateUser(context.Param("id"), &input); err != nil {
		log.Error(err)
		if errors.Is(err, nosqls.ErrNotExists) {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, nosqls.ErrInvalidObjectID) {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

func (userController *UserController) DeleteUser(context *gin.Context) {
	// trigger user deletion
	if err := userController.userService.DeleteUser(context.Param("id")); err != nil {
		log.Error(err)
		if errors.Is(err, nosqls.ErrNotExists) {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, nosqls.ErrInvalidObjectID) {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}
