package model

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Tag struct {
	Id        int    `gorm: "column: id"`
	Name      string `gorm: "column: name"`
	Blog_url  string `gorm: "column: blog_url"`
	Image_num int    `gorm: "column: image_numm"`
}

// ORM method to access Mysql
func (t *Tag) GetTagName(tag_id int) string {
	db := connect()
	defer db.Close()

	var tag Tag
	db.Table("tag").Where("id = ?", tag_id).Find(&tag)
	return tag.Name
}

func (t *Tag) GetTagId(tag_name string) int {
	db := connect()
	defer db.Close()

	var tag Tag
	db.Table("tag").Where("name = ?", tag_name).Find(&tag)
	return tag.Id
}

func (t *Tag) GetTagNum(tag_name string) int {
	db := connect()
	defer db.Close()

	var tag Tag
	db.Table("tag").Where("name = ?", tag_name).Find(&tag)
	return tag.Image_num
}

///// 改成+1
func (t *Tag) UpdateTagImageNum(tag_id int, count int) {
	db := connect()
	defer db.Close()

	err := db.Table("tag").Where("id = ?", tag_id).Update("image_num", count).Error
	if err != nil {
		log.Fatal(err)
	}
}
