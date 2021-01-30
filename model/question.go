package model

import "gorm.io/gorm"

// Question 问题模型
type Question struct {
	gorm.Model
	UserID  uint     `gorm:"not null;"`                                     // 问题所属用户Id
	Title   string   `gorm:"not null;"`                                     // 标题
	Content string   `gorm:"type:text"`                                     // 内容
	Answers []Answer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 关联回答信息
}

// GetQuestion 用ID获取问题
func GetQuestion(id uint) (*Question, error) {
	var question Question
	result := DB.First(&question, id)
	return &question, result.Error
}

// UpdateQuestion 根据ID修改问题
func UpdateQuestion(id uint, columns map[string]interface{}) (*Question, error) {
	var question Question
	result := DB.Model(&question).Where("id = ?", id).Updates(columns).Find(&question)
	return &question, result.Error
}

// DeleteQuestion 根据ID删除问题
func DeleteQuestion(id uint) error {
	result := DB.Delete(&Question{}, id).Error
	return result
}

// GetQuestions 用获取首页推荐列表，按创建时间降序排列，如果问题下有回答加载一个最新回答
func GetQuestions(limit int, offset int) ([]Question, error) {
	var questions []Question
	result := DB.Order("created_at desc").Limit(limit).Offset(offset).Find(&questions)
	for i := 0; i < len(questions); i++ {
		var answer Answer
		DB.Where("question_id = ?", questions[i].ID).Limit(1).Find(&answer)
		if answer.ID != 0 {
			questions[i].Answers = append(questions[i].Answers, answer)
		}
	}
	return questions, result.Error
}

// GetHotQuestions 用于获取问题热榜，暂返回最近发表的50条问题
func GetHotQuestions() ([]Question, error) {
	var questions []Question
	result := DB.Order("created_at desc").Limit(50).Find(&questions)
	return questions, result.Error
}

// 获取指定用户ID发布的问题（时间倒序）
func GetUserQuestions(userID uint) ([]Question, error) {
	var questions []Question
	result := DB.Where("user_id=?", userID).Order("created_at desc").Find(&questions)
	return questions, result.Error
}
