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

func UpdateMeasurment(measurement models.Measurements, id int) (models.Measurements, error){
    db := storage.GetDB()

    if db == nil {
        log.Fatal("Database connection is nil")
    }

    sqlStatement := `update measurements
    set weight=$2, height=$3, body_fat=$3, user_id: $4
    where id=$1
    returning id
    `
    err := db.QueryRow(
        sqlStatement, 
        id,
        measurement.Weight,
        measurement.Height,
        measurement.BodyFat,
        measurement.UserId,
    ).Scan(&id)

    if err != nil {
        return models.Measurements{}, err
    }
    measurement.Id = id
    return measurement, nil
    

}
