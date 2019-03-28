package common

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/nsogame/common/models"
	"golang.org/x/crypto/bcrypt"
)

type DB struct {
	*gorm.DB
}

func ConnectDB(provider string, connection string) (db *DB, err error) {
	h, err := gorm.Open(provider, connection)

	h.AutoMigrate(
		models.Beatmap{},
		models.BeatmapSet{},
		models.Score{},
		models.User{},
	)

	if err != nil {
		return
	}
	return &DB{h}, nil
}

func (db *DB) Authenticate(username, password string) (err error) {
	user, err := db.GetUserByName(username)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return
}

func (db *DB) GetUserByName(username string) (user models.User, err error) {
	err = db.Where("username = ?", strings.ToLower(username)).First(&user).Error
	return
}
