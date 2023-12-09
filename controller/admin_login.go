package controller

import (
	"github.com/e421083458/gin_scaffold/dto"
	"github.com/e421083458/gin_scaffold/middleware"
	"github.com/gin-gonic/gin"
)

type AdminLoginController struct{}

func RegisterAdminRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
	group.POST("/logout", adminLogin.AdminLoginOut)

}

func (adminLogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	err := params.BindValidParam(c)
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	middleware.ResponseSuccess(c, nil)
}

// AdminLoginOut /godoc
//
func (adminLogin *AdminLoginController) AdminLoginOut(c *gin.Context) {
	middleware.ResponseSuccess(c, nil)
}
