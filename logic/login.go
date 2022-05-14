package logic

import (
	"errors"
	"go-blog/dao/mysql"
	"go-blog/models"
	"go.uber.org/zap"
)

func Login(userName, password string) (*models.LoginRes, error) {
	user, err := mysql.GetUser(userName, password)
	if err != nil {
		zap.L().Error("mysql.GetUser() failed, err:", zap.Error(err))
		panic(err)
	}
	if user == nil {
		return nil, errors.New("账号或密码错误！")
	}
	var lr = &models.LoginRes{
		Token: "",
		UserInfo: models.UserInfo{
			Uid:      user.Uid,
			UserName: userName,
			Avatar:   user.Avatar,
		},
	}
	return lr, nil
}
