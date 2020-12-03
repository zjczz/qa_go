package model

import "gorm.io/gorm"

// Question 问题模型
type Question struct {
    gorm.Model
    UserID  uint        `gorm:"not null;"`           // 问题所属用户Id
    Title   string      `gorm:"not null;"`           // 标题
    Content string      `gorm:"type:text;not null;"` // 内容
    Answers []Answer    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 关联回答信息
}

// GetQuestion 用ID获取问题
func GetQuestion(id uint) (Question, error) {
    var question Question
    result := DB.First(&question, id)
    return question, result.Error
}

// GetQuestions 用获取问题列表，按创建时间降序排列
func GetQuestions(limit int, offset int) ([]Question, error) {
    var questions []Question
    result := DB.Order("created_at desc").Limit(limit).Offset(offset).Find(&questions)
    return questions, result.Error
}
