package logic

import (
	"go-blog/dao/mysql"
	"go-blog/models"
	"go-blog/settings"
	"go.uber.org/zap"
	"html/template"
)

func GetPostDetail(pId int) *models.PostRes {
	post, err := mysql.GetPostDetailById(pId)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailById() failed, err:", zap.Error(err))
		return nil
	}
	categoryName, err := mysql.GetCategoryNameById(post.CategoryId)
	if err != nil {
		zap.L().Error("mysql.GetCategoryNameById() failed, err:", zap.Error(err))
		return nil
	}
	userName, err := mysql.GetUserNameById(post.UserId)
	if err != nil {
		zap.L().Error("mysql.GetUserNameById() failed, err:", zap.Error(err))
		return nil
	}
	postMore := models.PostMore{
		Pid:          post.Pid,
		Title:        post.Title,
		Slug:         post.Slug,
		Content:      template.HTML(post.Content),
		Markdown:     post.Markdown,
		CategoryId:   post.CategoryId,
		CategoryName: categoryName,
		UserId:       post.UserId,
		UserName:     userName,
		ViewCount:    post.ViewCount,
		Type:         post.Type,
		CreateAt:     post.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt:     post.UpdateAt.Format("2006-01-02 15:04:05"),
	}
	var pr = &models.PostRes{
		Viewer:       settings.Conf.Viewer,
		SystemConfig: settings.Conf.SystemConfig,
		Article:      postMore,
	}
	return pr
}
