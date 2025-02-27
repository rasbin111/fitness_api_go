package repositories

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"log"
	"time"
)

func ListUsers(users models.Users) (models.Users, error){
    db := storage.GetDB()

    if db == nil {
        log.Fatal("Database connection is nil")
    }

    sqlStatement := `select id, name, email from users`

    rows, err := db.Query(sqlStatement)

    if err != nil {
        log.Fatal(err)
    }

    for rows.Next() {
        user := models.User{}
        err := rows.Scan(&user.Id, &user.Name, &user.Email)
        if err != nil {
            log.Fatal(err)
        }
        users = append(users, user)
    }

    return users, nil

}

func CreateUser(user models.User) (models.User, error) {
	db := storage.GetDB()

	if db == nil {
		log.Fatal("Database connection is nil")
	}
	sqlStatement := `INSERT INTO users (name, email, password)
					VALUES ($1, $2, $3) RETURNING id`
	// err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password).Scan(&user.Id)
	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUser(user models.User, id int) (models.User, error) {
	db := storage.GetDB()

	if db == nil {
		log.Fatal("Database connection is nil")
	}

	sqlStatement := `
	update users 
	set name = $2, email = $3, password = $4, updated_at = $5
	where id = $1
	returning id
	`

	err := db.QueryRow(sqlStatement,
		id,
		user.Name,
		user.Email,
		user.Password,
		time.Now()).Scan(&id)
	if err != nil {
		return models.User{}, err
	}

	user.Id = id
	return user, nil
}
