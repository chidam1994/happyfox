package gorp

import (
	"database/sql"

	"github.com/chidam1994/happyfox/models"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v2"
)

var db *sql.DB

func InitDB() *gorp.DbMap {
	dbConnString := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	//dbConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.GetString(config.DB_HOST), config.GetString(config.DB_PORT), config.GetString(config.DB_USERNAME), config.GetString(config.DB_PASSWORD), config.GetString(config.DB_DBNAME))
	var err error
	db, err = sql.Open("postgres", dbConnString)
	if err != nil {
		panic(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	dbmap.AddTableWithName(models.Contact{}, "contacts")
	dbmap.AddTableWithName(models.Email{}, "emails")
	dbmap.AddTableWithName(models.PhNum{}, "phnumbers")
	dbmap.AddTableWithName(models.Group{}, "groups")
	dbmap.AddTableWithName(models.Member{}, "members")
	return dbmap
}

func CloseDBConn() {
	db.Close()
}
