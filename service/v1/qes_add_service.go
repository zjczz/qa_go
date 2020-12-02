package v1

import (
    "likezh/model"
    "likezh/serializer"
)

// 
type QesAddService struct {
    Title	string `form:"title" json:"title" binding:"required"`
    Content string  `form:"content" json:"content" binding:"required"`

}

// 
func(qesAddService * QesAddService) QuestionAdd(user *model.User) *serializer.Response {
    qes := model.Question{
        UserId: user.ID,
        Title: qesAddService.Title ,
        Content: qesAddService.Content,
    }

    //
    if err := model.DB.Create(&qes).Error; err != nil {
        return serializer.ErrorResponse(serializer.CodeDatabaseError)
    }
    return serializer.OkResponse(serializer.BuildQuestionResponse(&qes))
}
