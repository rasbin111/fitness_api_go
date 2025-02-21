package repositories

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"log"
	"time"
)

func CreateMeasurement(measurement models.Measurements) (models.Measurements, error) {
	db := storage.GetDB()
	if db == nil {
		log.Fatal("Database connection is nil")
	}
	sqlStatement := `insert into measurements (user_id, weight, height, body_fat, created_at) 
	values ($1, $2, $3, $4, $5) 
	returning id`

	err := db.QueryRow(
		sqlStatement,
		measurement.UserId,
		measurement.Weight,
		measurement.Height,
		measurement.BodyFat,
		time.Now()).Scan(&measurement.Id)
	if err != nil {
		return measurement, err
	}
	return measurement, nil
}
