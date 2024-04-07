package logic

import (
	"blue/dao/mysql"
	"blue/models"
	"blue/pkg/jwt"
	"blue/pkg/snowflake"
)

// 存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2.生成UID
	userId := snowflake.GenID()

	// 3.保存进数据库
	//构建User实例
	user := &models.User{
		UserId:   userId,
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	// 1.用username找password
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//传递的是在指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	//生成JWT accessToken and refreshToken
	aToken, rToken, err := jwt.GenToken(user.UserId)
	if err != nil {
		return
	}

	user.AToken, user.RToken = aToken, rToken
	return
}
