package lib

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBHandler struct {
	DB  *gorm.DB
	Err error
}

var dbConn *DBHandler

func DBOpen() {
	dbConn = NewDBHandler()
}

func DBClose() {
	sqlDB, _ := dbConn.DB.DB()
	sqlDB.Close()
}

func NewDBHandler() *DBHandler {
	// user := os.Getenv("DB_USERNAME")
	// password := os.Getenv("DB_PASSWORD")
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_DATABASE")
	// fmt.Println(user, password, host, port)

	// format := "%s:%s@%s:%s/%s"

	// dsn := fmt.Sprintf(format, user, password, host, port, dbName)
	// db, err = gorm.Open(postgres.Open(dsn))
	dsn := "test.sqlite3"
	db, err := gorm.Open(sqlite.Open(dsn))

	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(50 * time.Second)

	dbHandler := new(DBHandler)
	db.Logger.LogMode(4)
	dbHandler.DB = db
	return dbHandler
}

func GetDBConn() *DBHandler {
	return dbConn
}

func BeginTransaction() *gorm.DB {
	dbConn.DB = dbConn.DB.Begin()
	return dbConn.DB
}

func RollBack() {
	dbConn.DB.Rollback()
}
