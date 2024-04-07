package logic

import (
	"blue/dao/mysql"
	"blue/dao/redis"
	"blue/models"
	"blue/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// 1.生成post_id
	p.ID = snowflake.GenID()
	// 2.保存到数据库
	if err = mysql.CreatePost(p); err != nil {
		zap.L().Error("CreatePost failed", zap.Error(err), zap.String("title", p.Title))
		return
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	return
}
