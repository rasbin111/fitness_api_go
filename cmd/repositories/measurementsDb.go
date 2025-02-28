package repositories

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"fmt"
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
    oldMeasurements := models.Measurements{} 

    if db == nil {
        log.Fatal("Database connection is nil")
    }

    fmt.Println("My data:", measurement)
    getMeasurement := `select id, user_id, weight, height, body_fat
                        from measurements
                        where id=$1`

    getErr := db.QueryRow(getMeasurement, id).Scan(
        &oldMeasurements.Id, 
        &oldMeasurements.UserId, 
        &oldMeasurements.Weight, 
        &oldMeasurements.Height, 
        &oldMeasurements.BodyFat)

    if getErr!= nil{
        log.Fatal(getErr)
    }

    sqlStatement := `update measurements
    set weight=$2, height=$3, body_fat=$4, user_id=$5
    where id=$1
    returning id
    `

    if measurement.UserId == 0{
        measurement.UserId = oldMeasurements.UserId
    }
    if measurement.Height == 0{
        measurement.Height = oldMeasurements.Height
    }
    if measurement.Weight == 0 {
        measurement.Weight = oldMeasurements.Weight
    }
    if measurement.BodyFat == 0 {
        measurement.BodyFat = oldMeasurements.BodyFat
    }

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
