package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username    string      `gorm:"unique;not null;"`        // 用户名
	Password    string      `gorm:"not null;"`               // 密码
	Email       string      `gorm:"default:null;unique;"`    // 邮箱
	Nickname    string      `gorm:"default:null"`            // 昵称
	Avatar      string      `gorm:"default:null;type:text;"` // 头像
	Status      int         `gorm:"default:0;not null;"`     // 状态
	UserProfile UserProfile // 关联用户信息
}

// UserProfile 用户信息模型
type UserProfile struct {
	gorm.Model
	UserID      uint
	Description string // 个人描述
}

const (
	// PasswordCost 密码加密难度
	PasswordCost = bcrypt.DefaultCost
	// Inactive 未激活用户
	Inactive int = 0
	// Active 激活用户
	Active int = 1
	// Suspend 被封禁用户
	Suspend int = 2
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
