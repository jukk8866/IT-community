package controller

import (
	"blue/dao/mysql"
	"blue/logic"
	"blue/models"
	"blue/pkg/jwt"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	//1.获取参数和参数校验

	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		//c.JSON(http.StatusOK, gin.H{
		//	"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		//})
	}

	/*
		  手动对请求参数进行详细的业务规则校验
		if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 ||
			p.Password != p.RePassword {
			zap.L().Error("SignUp with invalid param")
			c.JSON(http.StatusOK, gin.H{
				"msg": "请求参数有误",
			})
			return
		}
	*/

	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "注册失败",
			"result": err.Error(),
		})
		return
	}
	// 3. 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
	return
}

func LoginHandler(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			/*
				c.JSON(http.StatusOK, gin.H{
					"msg": "请求参数有误",
					"err": err.Error(),
				})
				return
			*/
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		/*
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
			})
		*/
	}

	//2.业务处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("usename", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserId),
		"user_name": user.Username,
		"atoken":    user.AToken,
		"rtoken":    user.RToken,
	})

	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "登陆成功",
	//})
	//return

}

// RefreshToken 刷新token
func RefreshToken(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	rToken := c.Request.Header.Get("rToken")
	if authHeader == "" && rToken == "" {
		ResponseError(c, CodeNeedLogin)
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		//if !(len(parts) == 2) {
		ResponseError(c, CodeInvalidToken)
		c.Abort()
		return
	}

	aToken, err := jwt.RefreshToken(parts[1], rToken)
	if err != nil {
		return
	}
	value, _ := c.Get(CtxUserIDKey)
	ResponseSuccess(c, gin.H{
		"user_id": fmt.Sprintf("%d", value),
		"atoken":  aToken,
	})
}
