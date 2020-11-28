package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username      string      `gorm:"unique;not null;"` // 用户名
	Password      string      `gorm:"not null;"`        // 密码
	Email         string      `gorm:"unique;not null;"` // 邮箱
	Nickname      string      // 昵称
	Avatar        string      `gorm:"type:text;not null;"` // 头像
	Status        int         `gorm:"not null;"`           // 状态
	UserProfileID int         `gorm:"unique;not null;"`    // 用户信息外键
	UserProfile   UserProfile // 用户信息
}

// UserProfile 用户信息模型
type UserProfile struct {
	gorm.Model
	Description string // 个人描述
}

const (
	// PasswordCost 密码加密难度
	PasswordCost = bcrypt.DefaultCost
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
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
