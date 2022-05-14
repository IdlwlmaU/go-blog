package logic

import (
	"go-blog/dao/mysql"
	"go-blog/models"
	"go-blog/settings"
	"go.uber.org/zap"
	"html/template"
)

func GetIndexInfo(page, pageSize int) models.HomeResponse {
	category, err := mysql.GetAllCategory()
	if err != nil {
		zap.L().Error("mysql.GetAllCategory() failed, err:", zap.Error(err))
		panic(err)
	}
	posts, err := mysql.GetPostPage(page, pageSize)
	if err != nil {
		zap.L().Error("mysql.GetAllPosts() failed, err:", zap.Error(err))
		panic(err)
	}
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName, err := mysql.GetCategoryNameById(post.CategoryId)
		if err != nil {
			zap.L().Error("mysql.GetCategoryNameById() failed, err:", zap.Error(err))
			panic(err)
		}
		userName, err := mysql.GetUserNameById(post.UserId)
		if err != nil {
			zap.L().Error("mysql.GetUserNameById() failed, err:", zap.Error(err))
			panic(err)
		}
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(content),
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
		postMores = append(postMores, postMore)
	}
	total, err := mysql.GetAllPostCount()
	if err != nil {
		zap.L().Error("mysql.GetAllPostCount() failed, err:", zap.Error(err))
		panic(err)
	}
	pageCount := (total-1)/pageSize + 1
	pages := make([]int, 0)
	for i := 1; i <= pageCount; i++ {
		pages = append(pages, i)
	}
	var hr = models.HomeResponse{
		Viewer:    settings.Conf.Viewer,
		Categorys: category,
		Posts:     postMores,
		Total:     total,
		Page:      page,
		Pages:     pages,
		PageEnd:   page != pageCount,
	}
	return hr
}
