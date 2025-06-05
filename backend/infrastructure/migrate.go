package database

import "Kuraz-Tech-Challenge/domain/entity"

func MigrateSchema() {
	DB.AutoMigrate(&entity.Task{})
}
