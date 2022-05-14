package models

import (
	"go-blog/settings"
	"html/template"
	"time"
)

type Post struct {
	Pid        int       `json:"pid"`         // 文章ID
	Title      string    `json:"title"`       // 文章标题
	Slug       string    `json:"slug"`        // 自定义页面 path
	Content    string    `json:"content"`     // 文章的html
	Markdown   string    `json:"markdown"`    // 文章的markdown
	CategoryId int       `json:"category_id"` // 分类的ID
	UserId     int       `json:"user_id"`     // 用户ID
	ViewCount  int       `json:"view_count"`  // 查看次数
	Type       int       `json:"type"`        // 文章类型 0 普通， 1 自定义文章
	CreateAt   time.Time `json:"create_at"`   // 创建时间
	UpdateAt   time.Time `json:"update_at"`   // 更新时间
}

type PostMore struct {
	Pid          int           `json:"pid"`           // 文章ID
	Title        string        `json:"title"`         // 文章标题
	Slug         string        `json:"slug"`          // 自定义页面 path
	Content      template.HTML `json:"content"`       // 文章的html
	Markdown     string        `json:"markdown"`      // 文章的markdown
	CategoryId   int           `json:"category_id"`   // 分类的ID
	CategoryName string        `json:"category_name"` // 分类名
	UserId       int           `json:"user_id"`       // 用户ID
	UserName     string        `json:"user_name"`     // 用户名
	ViewCount    int           `json:"view_count"`    // 查看次数
	Type         int           `json:"type"`          // 文章类型 0 普通， 1 自定义文章
	CreateAt     string        `json:"create_at"`     // 创建时间
	UpdateAt     string        `json:"update_at"`     // 更新时间
}

type PostReq struct {
	Pid        int    `json:"pid"`         // 文章ID
	Title      string `json:"title"`       // 文章标题
	Slug       string `json:"slug"`        // 自定义页面 path
	Content    string `json:"content"`     // 文章的html
	Markdown   string `json:"markdown"`    // 文章的markdown
	CategoryId int    `json:"category_id"` // 分类的ID
	UserId     int    `json:"user_id"`     // 用户ID
	Type       int    `json:"type"`        // 文章类型 0 普通， 1 自定义文章
}

type SearchResp struct {
	Pid   int    `orm:"pid" json:"pid"`     // 文章ID
	Title string `orm:"title" json:"title"` // 文章标题
}

type PostRes struct {
	settings.Viewer
	settings.SystemConfig
	Article PostMore
}
