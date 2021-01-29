package model

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"qa_go/cache"
	"strconv"
	"strings"
)

// UserLike 用户点赞表
type UserLike struct {
	gorm.Model
	UserID   uint `gorm:"not null;"` // 点赞用户Id
	AnswerID uint `gorm:"not null;"` // 被操作的回答Id
	Status   uint `gorm:"not null;"` // 点赞状态，0：无操作，1：已点赞，2：已点踩
}

const (
	NONE uint = 0
	UP   uint = 1
	DOWN uint = 2
)

const (
	AnswerLikeCount = "answer_like_count"
	UserLikeRecord  = "user_like_record"
)

// GetUserLike 获取用户uid对回答aid的点赞情况
func GetUserLike(uid uint, aid uint) (uint, error) {
	// 首先从redis中获取，获取到直接返回，否则从数据库获取
	key := fmt.Sprintf("%d:%d", uid, aid)
	// 在redis中找到了
	if res, err := cache.RedisClient.HGet(UserLikeRecord, key).Int(); err == nil {
		return uint(res), err
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
func AddUserLike(uid uint, aid uint, status uint) error {
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
	pipe.HSet(UserLikeRecord, key, status)
	_, err = pipe.Exec()
	return err
}

// GetLikeCountIdInCache 根据AnswerID获取缓存中的点赞数
func GetLikeCountInCache(aid uint) (uint, error) {
	res, err := cache.RedisClient.HGet(AnswerLikeCount, strconv.Itoa(int(aid))).Int()
	return uint(res), err
}

// SyncUserLikeRecord 将redis中的用户点赞记录同步到数据库
func SyncUserLikeRecord() {
	fmt.Println("Start SyncUserLikeRecord...")
	defer fmt.Println("End SyncUserLikeRecord...")

	// 从redis中获得数据
	data := cache.RedisClient.HGetAll(UserLikeRecord).Val()

	for key, val := range data {
		//fmt.Printf("%s\t%s\n", key, val)

		split := strings.Split(key, ":")
		uid, _ := strconv.Atoi(split[0])
		aid, _ := strconv.Atoi(split[1])
		status, _ := strconv.Atoi(val)

		var userLike UserLike
		userLike.UserID = uint(uid)
		userLike.AnswerID = uint(aid)

		row := DB.Where(&userLike).Find(&userLike).RowsAffected

		var err error
		// 存在则更新，不存在则创建
		if row > 0 {
			userLike.Status = uint(status)
			err = DB.Save(&userLike).Error
		} else {
			userLike.Status = uint(status)
			err = DB.Create(&userLike).Error
		}
		if err != nil {
			panic(err)
		}
	}

	// 删除redis中的数据
	cache.RedisClient.Del(UserLikeRecord)
}

// SyncUserLikeRecord 将redis中的回答点赞数量同步到数据库
func SyncAnswerLikeCount() {
	fmt.Println("Start SyncLikeCount...")
	defer fmt.Println("End SyncLikeCount...")

	// 从redis中获得数据
	data := cache.RedisClient.HGetAll(AnswerLikeCount).Val()

	for key, val := range data {
		//fmt.Printf("%s\t%s\n", key, val)
		id, _ := strconv.Atoi(key)
		count, _ := strconv.Atoi(val)

		var expr clause.Expr
		if count > 0 {
			expr = gorm.Expr("like_count + ?", count)
		} else if count < 0 {
			expr = gorm.Expr("like_count - ?", -count)
		} else {
			continue
		}

		var answer Answer
		answer.ID = uint(id)

		err := DB.Model(&answer).UpdateColumn("like_count", expr).Error
		if err != nil {
			panic(err)
		}
	}

	// 删除redis中的数据
	cache.RedisClient.Del(AnswerLikeCount)
}
