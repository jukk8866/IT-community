package mysql

import (
	"blue/models"
	"crypto/md5"
	"encoding/hex"
)

const secret = "happyboy"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	/*user := models.User{Username: username}
	msg := DB.Find(&user)
	fmt.Println(user)
	if msg.RowsAffected > 1 {
		fmt.Println(msg.RowsAffected)
		return true
	}*/

	user := models.User{}
	msg := DB.Where("username = ?", username).Find(&user)

	if msg.RowsAffected > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 插入用户数据
func InsertUser(user *models.User) (err error) {
	// 对密码加密
	user.Password = encryptPassword(user.Password)
	//插入数据
	DB.Create(&user)
	return
}

// 加密
func encryptPassword(oPassword string) string {
	h := md5.New()          //创建一个MD5对象
	h.Write([]byte(secret)) //加盐
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
	//h.Sum([]byte(oPassword)) 将 oPassword 转换为字节数组，
	//然后将其传递给哈希对象 h 的 Sum() 函数，生成原始密码的哈希摘要（即密码的加密结果）
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	msg := DB.Where("username = ?", user.Username).Find(&user)
	if msg.RowsAffected == 0 {
		return ErrorUserNotExist
	}
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}
