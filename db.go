package common

import (
	"strings"

	"git.iptq.io/nso/common/models"
	"github.com/jinzhu/gorm"
)

type DB struct {
	*gorm.DB
}

func ConnectDB(provider string, connection string) (db *DB, err error) {
	h, err := gorm.Open(provider, connection)

	var score models.Score
	// var scores []models.Score
	var user models.User
	h.AutoMigrate(
		&models.Beatmap{},
		&models.BeatmapSet{},
		&score, &user,
	)

	if err != nil {
		return
	}
	return &DB{h}, nil
}

func (db *DB) GetUser(username string) (user models.User, err error) {
	err = db.Where("username = ?", strings.ToLower(username)).First(&user).Error
	return
}
