package web

import (
	"github.com/gin-gonic/gin"
	"github.com/otumian-empire/go-to-do-list-restful-api/model"
)

type UserController struct {
	model model.UserModel
}

func (controller *UserController) SignUp() gin.HandlerFunc

func (controller *UserController) Login() gin.HandlerFunc

func (controller *UserController) UpdateUserUsername() gin.HandlerFunc

func (controller *UserController) UpdateUserPassword() gin.HandlerFunc

func (controller *UserController) ReadUser() gin.HandlerFunc

func (controller *UserController) DeleteUser() gin.HandlerFunc

func (controller *UserController) Logout() gin.HandlerFunc
