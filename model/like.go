package model

import (
	"fmt"
	"gorm.io/gorm"
	"qa_go/cache"
	"strconv"
)

// UserLike 用户点赞表
type UserLike struct {
	gorm.Model
	UserID   uint     `gorm:"not null;"`                  // 点赞用户Id
	AnswerID uint     `gorm:"not null;"`                  // 被操作的回答Id
	Status   likeType `gorm:"type:enum(0,1,2);not null;"` // 点赞状态，0：无操作，1：已点赞，2：已点踩
}

type likeType uint

const (
	NONE likeType = 0
	UP   likeType = 1
	DOWN likeType = 2
)

const (
	AnswerLikeCount = "answer_like_count"
	UserLikeRecord  = "user_like_record"
)

// GetUserLike 获取用户uid对回答aid的点赞情况
func GetUserLike(uid uint, aid uint) (likeType, error) {
	// 首先从redis中获取，获取到直接返回，否则从数据库获取
	key := fmt.Sprintf("%d:%d", uid, aid)
	// 在redis中找到了
	if res, err := cache.RedisClient.HGet(UserLikeRecord, key).Int(); err == nil {
		return likeType(res), err
	}
	// 在redis中没有找到，从数据库获取
	var userLike UserLike
	result := DB.Where("user_id = ? and answer_id = ?", uid, aid).First(&userLike)
	// 如果数据库中没有该记录，返回未点赞
	if result.RowsAffected == 0 {
		return NONE, nil
	}
	return userLike.Status, result.Error
}

// AddUserLike status=0：取消点赞，status=1：点赞，status=2：点踩
func AddUserLike(uid uint, aid uint, status likeType) error {
	// 获取之前的点赞状态
	pre, err := GetUserLike(uid, aid)
	if err != nil {
		return err
	}
	var incr int64 = 0
	// 根据前后的状态，判断点赞数的增减与否
	if (pre == NONE || pre == DOWN) && status == UP {
		incr = 1
	} else if pre == UP && (status == NONE || status == DOWN) {
		incr = -1
	}
	pipe := cache.RedisClient.TxPipeline()
	pipe.HIncrBy(AnswerLikeCount, strconv.Itoa(int(aid)), incr)
	key := fmt.Sprintf("%d:%d", uid, aid)
	pipe.HSet(UserLikeRecord, key, uint(status))
	_, err = pipe.Exec()
	return err
}

// GetLikeCountIdInCache 根据AnswerID获取缓存中的点赞数
func GetLikeCountInCache(aid uint) (uint, error) {
	res, err := cache.RedisClient.HGet(AnswerLikeCount, strconv.Itoa(int(aid))).Int()
	return uint(res), err
}
