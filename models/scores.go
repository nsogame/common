package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Score struct {
	gorm.Model

	UserID    int64
	BeatmapID int64

	BeatmapMD5 string
	ReplayMD5  string

	Count300, Count100, Count50     int
	CountGeki, CountKatu, CountMiss int
	Score, MaxCombo                 uint32
	FullCombo                       bool

	Mods int
	Pass bool
	Mode int
	Date time.Time
}
