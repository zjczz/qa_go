package v1

import (
    "likezh/serializer"
    "likezh/model"
)

func FindOneQuestion(id uint) *serializer.Response {
    if question, err := model.GetQuestion(id); err == nil {
        return serializer.OkResponse(serializer.BuildQuestionResponse(question))
    }
    return serializer.ErrorResponse(serializer.CodeQuestionIdError)
}
