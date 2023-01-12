package database

import (
	"github.com/SirusDoma/btc-wallet/topup"
	"gorm.io/gorm"
	"log"
)

func CreateTopupTableUp(db *gorm.DB) {
	if err := db.AutoMigrate(&topup.Topup{}); err != nil {
		log.Fatalf("Failed to create topup table: %s", err.Error())
	}
}
