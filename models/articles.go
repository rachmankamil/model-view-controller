package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title         string `gorm:"unique;column:name;comment:this is title;size:50;"`
	Author        string `gorm:"-"`
	ThumbnailPath string
	RawHTML       string `gorm:"type:text;"`
}
