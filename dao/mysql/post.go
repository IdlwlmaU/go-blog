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

// GetPostPageByCategoryId 根据帖子分类id获取帖子信息
func GetPostPageByCategoryId(cId, page, pageSize int) (posts []*models.Post, err error) {
	sqlStr := `select * from blog_post where category_id = ? limit ?,?`
	pageStart := (page - 1) * pageSize
	err = db.Select(&posts, sqlStr, cId, pageStart, pageSize)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no post in db")
		err = nil
	}
	return
}

// GetAllPostCountById 根据cId获取该分类帖子数量
func GetAllPostCountById(cId int) (counts int, err error) {
	sqlStr := `select count(1) from blog_post where category_id = ?`
	err = db.Get(&counts, sqlStr, cId)
	if err == sql.ErrNoRows {
		zap.L().Warn("there is no value in db")
		err = nil
	}
	return
}
