package model

import (
	"gorm.io/gorm"
	"qa_go/cache"
)

// Answer 回答模型
type Answer struct {
	gorm.Model
	UserID     uint   `gorm:"not null;"`           // 回答所属用户Id
	QuestionID uint   `gorm:"not null;"`           // 回答所属问题Id
	Content    string `gorm:"type:text;not null;"` // 内容
	LikeCount  uint   `gorm:"not null;"`           // 点赞数
}

const (
	DeletedAnswers = "deleted_answers"
)

// GetAnswer 用ID获取回答
func GetAnswer(id uint) (*Answer, error) {
	var answer Answer
	result := DB.First(&answer, id)
	return &answer, result.Error
}

//GetAnswers 用ID获取回答列表
func GetAnswers(ids []uint) ([]Answer, error) {
	var ans []Answer
	for _, id := range ids {
		a, _ := GetAnswer(id)
		ans = append(ans, *a)
	}
	return ans, nil
}

// 删除回答
func DeleteAnswer(id uint) error {
	result := DB.Delete(&Answer{}, id).Error
	result = DeleteLikeByAnswer(id)
	return result
}

// 删除回答关联的点赞
func DeleteLikeByAnswer(id uint) error {
	result := cache.RedisClient.SAdd(DeletedAnswers, id).Err()
	result = DB.Where("answer_id = ?", id).Delete(&UserLike{}).Error
	return result
}

// 获取回答列表，按创建时间降序排列
func GetAnswersByTime(questionID uint, limit int, offset int) ([]Answer, error) {
	var answers []Answer
	result := DB.Where("question_id = ?", questionID).Order("created_at desc").Limit(limit).Offset(offset).Find(&answers)
	return answers, result.Error
}

// 获取回答列表，按热度（暂时只有点赞数）排序
func GetAnswersByScore(questionID uint, limit int, offset int) ([]Answer, error) {
	var answers []Answer
	result := DB.Where("question_id = ?", questionID).Order("like_count desc").Limit(limit).Offset(offset).Find(&answers)
	return answers, result.Error
}

//获取某回答的点赞总数
func GetAnswerLikedCount(aid uint) (uint, error) {
	exist, cnt, err := GetLikeCountInCache(aid)
	if err != nil || exist {
		return cnt, err
	}
	ans, err := GetAnswer(aid)
	cnt = ans.LikeCount
	return cnt, err
}

//获取某用户对某问题的点赞状态
func GetUserLikeStatus(uid uint, aid uint) (uint, error) {
	return GetUserLike(uid, aid)
}

// 获取指定用户ID发布的回答（时间倒序）
func GetUserAnswers(userID uint) ([]Answer, error) {
	var answers []Answer
	result := DB.Where("user_id=?", userID).Order("created_at desc").Find(&answers)
	return answers, result.Error
}

// 获取指定问题在当前数据库中的最高赞回答
func GetHotAnswer(questionID uint) *Answer {
	var answer Answer
	result := DB.Where("question_id = ?", questionID).Order("like_count desc").Limit(1).Find(&answer)
	if result.RowsAffected > 0 {
		return &answer
	}
	return nil
}
