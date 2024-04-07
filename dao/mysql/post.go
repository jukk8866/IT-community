package mysql

import (
	"blue/models"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {

	err = DB.Create(&p).Error
	return
}
