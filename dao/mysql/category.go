package mysql

import (
	"database/sql"
	"go-blog/models"
	"go.uber.org/zap"
)

// GetAllCategory 获取所有的分类信息
func GetAllCategory() (category []models.Category, err error) {
	sqlStr := `select * from blog_category`
	err = db.Select(&category, sqlStr)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no category in db")
		err = nil
	}
	return
}

// GetCategoryNameById 根据分类的id获取分类名
func GetCategoryNameById(id int) (name string, err error) {
	sqlStr := `select name from blog_category where cid = ?`
	err = db.Get(&name, sqlStr, id)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no category_name in db")
		err = nil
	}
	return
}
