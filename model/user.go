package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	Username    string      `gorm:"unique;not null;"`                              // 用户名
	Password    string      `gorm:"not null;"`                                     // 密码
	UserProfile UserProfile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 关联用户信息
	Questions   []Question  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 关联问题信息
	Answers     []Answer    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 关联回答信息
}

// UserProfile 用户信息模型
type UserProfile struct {
	gorm.Model
	UserID      uint
	Nickname    string `gorm:"default:null"`         // 昵称
	Email       string `gorm:"unique;default:null;"` // 邮箱
	Avatar      string `gorm:"default:null;"`        // 头像
	Status      int    `gorm:"not null;default:0;"`  // 状态
	Description string `gorm:"default:null"`         // 个人描述
}

const (
	// PasswordCost 密码加密难度
	PasswordCost = bcrypt.DefaultCost
	// Inactive 未激活用户
	Inactive int = 0
	// Active 激活用户
	Active int = 1
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// GetUserProfile 用ID获取用户详细信息
func GetUserProfile(ID interface{}) (UserProfile, error) {
	var profile UserProfile
	result := DB.Where("user_id = ?", ID).First(&profile)
	return profile, result.Error
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
