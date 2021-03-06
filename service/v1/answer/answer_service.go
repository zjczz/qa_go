package v1

import (
	"qa_go/cache"
	"qa_go/model"
	"qa_go/serializer"
	"strconv"
)

// AddAnswerService 管理回答问题的服务
type AddAnswerService struct {
	Content string `form:"content" json:"content" binding:"required"`
}

// UpdateAnswerService 管理修改回答
type UpdateAnswerService struct {
	Content string `form:"content" json:"content" binding:"required"`
}

// VoterService 管理点赞
type VoterService struct {
	Type string `form:"type" json:"type" binding:"required"`
}

// 回答问题
func (service *AddAnswerService) AddAnswer(user *model.User, qid uint) *serializer.Response {
	answer := &model.Answer{
		UserID:     user.ID,
		QuestionID: qid,
		Content:    service.Content,
	}

	if err := model.DB.Create(answer).Error; err != nil {
		return serializer.ErrorResponse(serializer.CodeDatabaseError)
	}
	// 对应问题回答数+1
	question, _ := model.GetQuestion(qid)
	if _, err := model.UpdateQuestion(qid, map[string]interface{}{
		"answer_count": question.AnswerCount + 1,
	}); err != nil {
		return serializer.ErrorResponse(serializer.CodeDatabaseError)
	}
	// 如果对应问题为热门问题但没有热门回答，则新回答成为热门回答
	qidStr := strconv.Itoa(int(qid))
	if _, err := cache.RedisClient.ZScore(cache.KeyHotQuestions, qidStr).Result(); err == nil {
		if !cache.RedisClient.HExists(cache.KeyHotAnswer, qidStr).Val() {
			cache.RedisClient.HSet(cache.KeyHotAnswer, qidStr, answer.ID)
		}
	}
	return serializer.OkResponse(serializer.BuildAnswerResponse(answer, user.ID))
}

// 查看单个回答
func FindOneAnswer(qid uint, aid uint, uid uint) *serializer.Response {
	if answer, err := model.GetAnswer(aid); err == nil {
		if answer.QuestionID != qid {
			return serializer.ErrorResponse(serializer.CodeQidMismatchError)
		}
		return serializer.OkResponse(serializer.BuildAnswerResponse(answer, uid))
	}
	return serializer.ErrorResponse(serializer.CodeAnswerIdError)
}

// 修改回答
func (service *UpdateAnswerService) UpdateAnswer(user *model.User, qid uint, aid uint) *serializer.Response {
	_, err := model.GetQuestion(qid)
	if err != nil {
		return serializer.ErrorResponse(serializer.CodeQuestionIdError)
	}
	answer, err := model.GetAnswer(aid)
	if err != nil {
		return serializer.ErrorResponse(serializer.CodeAnswerIdError)
	}

	if answer.QuestionID != qid {
		return serializer.ErrorResponse(serializer.CodeQidMismatchError)
	}
	if answer.UserID != user.ID {
		return serializer.ErrorResponse(serializer.CodeAnswerNotOwn)
	}

	answer.Content = service.Content
	if model.DB.Save(answer).Error != nil {
		return serializer.ErrorResponse(serializer.CodeDatabaseError)
	}
	return serializer.OkResponse(serializer.BuildAnswerResponse(answer, user.ID))
}

// 删除回答
func DeleteAnswer(user *model.User, qid uint, aid uint) *serializer.Response {
	_, err := model.GetQuestion(qid)
	if err != nil {
		return serializer.ErrorResponse(serializer.CodeQuestionIdError)
	}
	answer, err := model.GetAnswer(aid)
	if err != nil {
		return serializer.ErrorResponse(serializer.CodeAnswerIdError)
	}

	if answer.QuestionID != qid {
		return serializer.ErrorResponse(serializer.CodeQidMismatchError)
	}
	if answer.UserID != user.ID {
		return serializer.ErrorResponse(serializer.CodeAnswerNotOwn)
	}

	err = model.DeleteAnswer(aid)
	if err != nil {
		return serializer.ErrorResponse(serializer.CodeDatabaseError)
	}

	// 回答数减一
	question, _ := model.GetQuestion(qid)
	if _, err := model.UpdateQuestion(qid, map[string]interface{}{
		"answer_count": question.AnswerCount - 1,
	}); err != nil {
		return serializer.ErrorResponse(serializer.CodeDatabaseError)
	}
	// 如果redis中有热门问题与本条回答记录，删除
	if aidStr, err := cache.RedisClient.HGet(cache.KeyHotAnswer, strconv.Itoa(int(qid))).Result(); err == nil {
		ai, _ := strconv.Atoi(aidStr)
		if uint(ai) == aid {
			cache.RedisClient.HDel(cache.KeyHotAnswer, strconv.Itoa(int(qid)))
		}
	}
	return serializer.OkResponse(nil)
}

//Voter 点赞
func Voter(uid uint, aid uint, status string) *serializer.Response {
	var code uint
	if status == "up" {
		code = 1
	} else if status == "down" {
		code = 2
	} else if status == "neutral" {
		code = 0
	} else {
		return serializer.ErrorResponse(serializer.CodeVoterTypeError)
	}
	err := model.AddUserLike(uid, aid, code)
	if err != nil {
		return serializer.ErrorResponse(serializer.CodeDatabaseError)
	}
	return serializer.OkResponse(nil)
}

//获取赞的内容
func GetAwesomes(uid uint) *serializer.Response {
	ids, err := model.GetUserLikes(uid)
	if err != nil {
		return serializer.ErrorResponse(serializer.CodeAnswerIdError)
	}
	ans, _ := model.GetAnswers(ids)
	return serializer.OkResponse(serializer.BuildAwesomesResponse(ans, uid))
}
