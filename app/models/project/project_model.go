// Package project 模型
package project

import (
	"cchub/app/models"
	"cchub/pkg/database"
)

type Project struct {
	models.BaseModel

	// Put fields in here

	models.CommonTimestampsField
}

func (project *Project) Create() {
	database.DB.Create(&project)
}

func (project *Project) Save() (rowsAffected int64) {
	result := database.DB.Save(&project)
	return result.RowsAffected
}

func (project *Project) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&project)
	return result.RowsAffected
}
