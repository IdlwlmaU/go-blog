package logic

import (
	"database/sql"
	"errors"
	"go-blog/dao/mysql"
	"go-blog/models"
	jwt2 "go-blog/pkg/jwt"
	"go.uber.org/zap"
)

func Login(userName, password string) (*models.LoginRes, error) {
	//password = md5.EnCryptPassword(password)
	user, err := mysql.GetUser(userName, password)
	if user == nil || err == sql.ErrNoRows {
		return nil, errors.New("账号或密码错误！")
	}
	uid := user.Uid
	// 生成JWT
	token, err := jwt2.GetToken(uid, userName)
	if err != nil {
		zap.L().Error("jwt2.GetToken() failed, err:", zap.Error(err))
		return nil, err
	}
	var lr = &models.LoginRes{
		Token: token,
		UserInfo: models.UserInfo{
			Uid:      user.Uid,
			UserName: userName,
			Avatar:   user.Avatar,
		},
	}
	return lr, nil
}
