package main

import (
	"fmt"
	"log"

	"git.iptq.io/nso/common/models"
	"git.iptq.io/nso/drift"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	emptyState := drift.EmptyState()
	state, err := drift.ConstructDatabaseState(
		&models.Beatmap{},
		&models.BeatmapSet{},
		&models.Score{},
		&models.User{},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("state:", state)

	diff := drift.CalculateDatabaseDiff(emptyState, state)
	fmt.Println("diff:", diff)
}
