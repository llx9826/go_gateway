package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/llx9826/gin_gateway/dto"
	"github.com/llx9826/gin_gateway/public"
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	Id        int       `json:"id" gorm:"primary_key" description:"自增主键"`
	UserName  string    `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Salt      string    `json:"salt" gorm:"column:salt" description:"盐"`
	Password  string    `json:"password" gorm:"column:password" description:"密码"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDeleted int       `json:"is_deleted" gorm:"column:is_deleted" description:"是否删除"`
}

func (t *Admin) TableName() string {
	return ""
}

func (t *Admin) Find(c *gin.Context, tx *gorm.DB, search *Admin) (*Admin, error) {
	out := &Admin{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (t *Admin) LoginCheck(c *gin.Context, tx *gorm.DB, param *dto.AdminLoginInput) (*Admin, error) {
	//1.判断用户是否存在
	adminInfo, err := t.Find(c, tx, &Admin{UserName: param.UserName, IsDeleted: 0})
	if err != nil {
		return nil, errors.New("用户信息不存在")
	}
	//2.判断密码是否正确
	saltPassword := public.GenSaltPassword(adminInfo.Salt, adminInfo.Password)
	if adminInfo.Password != saltPassword {
		return nil, errors.New("密码错误，请重新输入")
	}
	return adminInfo, nil
}
