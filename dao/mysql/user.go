package mysql

import (
	"database/sql"
	"go-blog/models"
	"go.uber.org/zap"
)

func GetUserNameById(id int) (name string, err error) {
	sqlStr := `select user_name from blog_user where uid = ?`
	err = db.Get(&name, sqlStr, id)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no user_name in db")
		err = nil
	}
	return
}

func GetUser(userName, password string) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select uid, user_name, passwd, avatar, create_at, update_at 
				from blog_user
				where user_name = ? and passwd = ? 
				limit 1`
	err = db.Get(user, sqlStr, userName, password)
	return
}
