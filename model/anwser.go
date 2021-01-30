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

//根据questionID获取回答总数
func GetAnswerCount(id uint) int64 {
	var cnt int64
	DB.Model(&Answer{}).Where("question_id = ?", id).Count(&cnt)
	return cnt
}

// 删除回答
func DeleteAnswer(id uint) error {
	result := DB.Delete(&Answer{}, id).Error
	return result
}

// 获取回答列表，按创建时间降序排列
func GetAnswersByTime(questionID uint, limit int, offset int) ([]Answer, error) {
	var answers []Answer
	result := DB.Where("question_id = ?", questionID).Order("created_at desc").Limit(limit).Offset(offset).Find(&answers)
	return answers, result.Error
}

// 获取回答列表，按热度排序，暂按时间升序
func GetAnswersByScore(questionID uint, limit int, offset int) ([]Answer, error) {
	var answers []Answer
	result := DB.Where("question_id = ?", questionID).Order("created_at").Limit(limit).Offset(offset).Find(&answers)
	return answers, result.Error
}

// 获取指定用户ID发布的回答（时间倒序）
func GetUserAnswers(userID uint) ([]Answer, error) {
	var answers []Answer
	result := DB.Where("user_id=?", userID).Order("created_at desc").Find(&answers)
	return answers, result.Error
}
