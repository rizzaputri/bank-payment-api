package initializers

import "test-mnc/models"

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to migrate User database: " + err.Error())
	}

	err = DB.AutoMigrate(&models.Customer{})
	if err != nil {
		panic("Failed to migrate Customer database: " + err.Error())
	}

	err = DB.AutoMigrate(&models.Payment{})
	if err != nil {
		panic("Failed to migrate Payment database: " + err.Error())
	}
}
