package models

type Category struct {
	Cid      int    `json:"cid" db:"cid"`
	Name     string `json:"name" db:"name"`
	CreateAt string `json:"create_at" db:"create_at"`
	UpdateAt string `json:"update_at" db:"update_at"`
}

type CategoryResponse struct {
	HomeResponse
	CategoryName string `json:"category_name" db:"category_name"`
}
