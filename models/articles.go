package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	ID      string `json:"ID"`
	Title   string `json:"title" gorm:"unique;column:name;comment:this is title;size:50;"`
	Author  string `json:"author" gorm:"-"`
	RawHTML string `json:"raw_html" gorm:"type:text;"`
}
