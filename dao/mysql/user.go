package mysql

import (
	"database/sql"
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
