package mysql

import (
	"blue/models"
	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	err = DB.Find(&communityList).Error
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		return
	}
	return
}
