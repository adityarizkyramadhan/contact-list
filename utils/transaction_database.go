package utils

import "gorm.io/gorm"

func BeginTransaction(db *gorm.DB) *gorm.DB {
	return db.Begin()
}

func Rollback(tx *gorm.DB, rollback bool) {
	if rollback {
		tx.Rollback()
	}
}

func Commit(tx *gorm.DB) {
	tx.Commit()
}
