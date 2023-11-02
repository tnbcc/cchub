package seeders

import (
	"cchub/database/factories"
	"cchub/pkg/console"
	"cchub/pkg/log"
	"cchub/pkg/seed"
	"fmt"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedTopicsTable", func(db *gorm.DB) {

		topics := factories.MakeTopics(10)

		result := db.Table("topics").Create(&topics)

		if err := result.Error; err != nil {
			log.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
