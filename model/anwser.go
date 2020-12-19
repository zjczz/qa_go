package model

import "gorm.io/gorm"

// Answer 回答模型
type Answer struct {
	gorm.Model
	UserID     uint   `gorm:"not null;"`           // 回答所属用户Id
	QuestionID uint   `gorm:"not null;"`           // 回答所属问题Id
	Content    string `gorm:"type:text;not null;"` // 内容
}

// GetAnswer 用ID获取回答
func GetAnswer(id uint) (*Answer, error) {
	var answer Answer
	result := DB.First(&answer, id)
	return &answer, result.Error
}
