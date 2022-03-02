package repository

import (
	"github.com/otaviolarrosa/golang-gin-gonic/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	sqlDb, err := db.connection.DB()
	sqlDb.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}
func (db *database) Update(video entity.Video) {
	db.connection.Save(&video)
}

func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Preload("Author", true).Find(&videos)
	return videos
}
