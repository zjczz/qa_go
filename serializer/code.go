package serializer

import "strconv"

type errorCode uint16

// 通用错误码
const (
	CodeOk                errorCode = 0
	CodeUnknownError      errorCode = 1
	CodeParamError        errorCode = 1001
	CodeTokenNotExitError errorCode = 1002
	CodeTokenExpiredError errorCode = 1003
	CodeDatabaseError     errorCode = 1004
)

// 用户相关的错误
const (
	// 注册错误
	CodeUserExistError       errorCode = 2001
	CodePasswordConfirmError errorCode = 2002

	// 登录错误
	CodeUserNotExistError errorCode = 2101
	CodePasswordError     errorCode = 2102
)

// 问题相关的错误
const (
	CodeQuestionIdError errorCode = 3001
	CodeQuestionNotOwn  errorCode = 3002
)

// 回答相关的错误
const (
	CodeAnswerIdError    errorCode = 4001
	CodeQidMismatchError errorCode = 4002
	CodeAnswerNotOwn     errorCode = 4003
	CodeVoterTypeError   errorCode  = 4004
)

// 错误码与描述信息map
var msgMap = map[errorCode]string{
	CodeOk:                "ok",
	CodeUnknownError:      "未知错误",
	CodeParamError:        "请求参数错误",
	CodeTokenNotExitError: "需要权限",
	CodeTokenExpiredError: "token过期或不正确",
	CodeDatabaseError:     "数据库错误",

	CodeUserExistError:       "注册失败，用户已存在",
	CodePasswordConfirmError: "注册失败，两次输入密码不一致",

	CodeUserNotExistError: "登录失败，用户名不存在",
	CodePasswordError:     "登录失败，密码错误",

	CodeQuestionIdError: "无效的问题id",
	CodeQuestionNotOwn:  "用户无权处理该问题",

	CodeAnswerIdError:    "无效的回答id",
	CodeQidMismatchError: "回答所属的问题id错误",
	CodeAnswerNotOwn:     "用户无权处理该回答",
	CodeVoterTypeError:	  "点赞类型错误",
}

// 根据错误码得到对应说明
func GetErrorMsg(code errorCode) string {
	msg, ok := msgMap[code]
	if !ok {
		msg = "未知错误，错误码：" + strconv.Itoa(int(code))
	}
	return msg
}
