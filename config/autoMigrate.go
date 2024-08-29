package config

import "tuxedo/database"

func RunMigrate(dataModel interface{}) {
	database.DB.AutoMigrate(dataModel)
}
