package mysql

import (
	"database/sql"
	"go-blog/models"
	"go.uber.org/zap"
)

// GetPostPage 获取一个页面的帖子信息
func GetPostPage(page, pageSize int) (posts []*models.Post, err error) {
	sqlStr := `select * from blog_post limit ?,?`
	pageStart := (page - 1) * pageSize
	err = db.Select(&posts, sqlStr, pageStart, pageSize)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no post in db")
		err = nil
	}
	return
}

// GetAllPostCount 获取页面总数
func GetAllPostCount() (counts int, err error) {
	sqlStr := `select count(1) from blog_post`
	err = db.Get(&counts, sqlStr)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no value in db")
		err = nil
	}
	return
}
