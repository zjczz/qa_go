package model

import "gorm.io/gorm"

// Answer 回答模型
type Answer struct {
	gorm.Model
	UserId     int    `gorm:"not null;"`           // 回答所属用户Id
	QuestionId int    `gorm:"not null;"`           // 回答所属问题Id
	Content    string `gorm:"type:text;not null;"` // 内容
}
