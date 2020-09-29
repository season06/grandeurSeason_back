package model

import (
	"fmt"
	"log"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Image struct {
	Id          int       `gorm: "column: id"`
	Name        string    `gorm: "column: name"`
	Image_url   string    `gorm: "column: image_url"`
	Type        string    `gorm: "column: type"`
	Location    string    `gorm: "column: location"`
	Favorite    int       `gorm: "column: favorite"`
	Tag         int       `gorm: "column: tag"`
	Create_time time.Time `gorm: "column: create_time"`
}

// ORM method to access Mysql
func (image *Image) GetAllImage() (img []Image) {
	db := connect()
	defer db.Close()

	db.Table("image").Order("create_time DESC").Find(&img)
	return img
}

func (image *Image) GetImageByTag(tag_id int) (img []Image) {
	db := connect()
	defer db.Close()

	db.Table("image").Where("tag = ?", tag_id).Order("create_time DESC").Find(&img)
	return img
}

func (image *Image) GetImageById(image_id int) (img Image) {
	db := connect()
	defer db.Close()

	db.Table("image").Where("id = ?", image_id).Find(&img)
	return img
}

func (image *Image) GetNewPost() (img []Image) {
	db := connect()
	defer db.Close()

	db.Table("image").Limit(5).Order("create_time DESC").Find(&img)
	return img
}

func (image *Image) InsertImage(data Image) (img Image) {
	db := connect()
	defer db.Close()

	db.Create(&data)
	return data
}

func (image *Image) DeleteImageById(image_id int) bool {
	db := connect()
	defer db.Close()
	fmt.Println(image_id)

	err := db.Where("id = ?", image_id).Delete(Image{}).Error
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (image *Image) CountImageNum(tag int) int {
	db := connect()
	defer db.Close()

	var count int
	db.Table("image").Where("tag = ?", tag).Count(&count)
	return count
}

func (image *Image) UpdateFavorite(image_id int) (img Image) {
	db := connect()
	defer db.Close()

	db.Table("image").Where("id = ?", image_id).Find(&img)
	if img.Favorite == 0 {
		img.Favorite = 1
	} else {
		img.Favorite = 0
	}
	db.Save(&img)
	return img
}

func (image *Image) UpdateLocation(image_id int, location string) (img Image) {
	db := connect()
	defer db.Close()

	db.Table("image").Where("id = ?", image_id).Find(&img)
	img.Location = location

	db.Save(&img)
	return img
}
