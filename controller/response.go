package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*

{
	"code": 10000, // 程序中的错误码
	"msg": xx,     // 提示信息
	"data": {},    // 数据
}

*/

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

/*
Msg() 方法是定义在 ResCode 类型上的，意味着可以通过ResCode 类型的实例来调用该方法。
如果调用方式是Msg(CodeSuccess)，那么编译器会提示错误，因为 Msg() 方法是在ResCode
类型上定义的方法，并且该方法没有参数，只能通过ResCode 类型的实例来调用。
因此，正确的调用方式应该是CodeSuccess.Msg()，其中CodeSuccess 是一个 ResCode 类型的实例,
调用该实例上的Msg()方法。
*/
