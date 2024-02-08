package database

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

func DatabaseConnection() {
	err := godotenv.Load()
	if err != nil {
		panic(".env file is missing")
	}

	dbConnection := os.Getenv("DB_CONNECTION")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	switch strings.ToLower(dbConnection) {
	case "mysql":
		db = mysqlConnection(dbDatabase, dbHost, dbPort, dbUsername, dbPassword)
	case "postsql":
		db = postsqlConnection(dbDatabase, dbHost, dbPort, dbUsername, dbPassword)
	case "sqlserver":
		db = sqlserverConnection(dbDatabase, dbHost, dbPort, dbUsername, dbPassword)
	case "sqlite":
		db = sqliteConnection()
	default:
		panic("unsupported connection")
	}
	if sqlDB, err = db.DB(); err == nil {
	} else {
		fmt.Println("Error sqlDB")
	}
}

func mysqlConnection(dbDatabase, dbHost, dbPort, dbUsername, dbPassword string) *gorm.DB {
	fmt.Println("mysql preparations")
	dns := dbUsername + ":" + dbPassword + "@/"+ dbDatabase+ "?charset=utf8mb4&parseTime=True&loc=Africa%2fDar_es_Salaam"
	// dns := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbDatabase + "?charset=utf8mb4&parseTime=True&loc=Africa%2fDar_es_Salaam"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	fmt.Println("about to finish mysql preparations")
	if err != nil {
		panic(err)
	}
	return db
}

func postsqlConnection(dbDatabase, dbHost, dbPort, dbUsername, dbPassword string) *gorm.DB {
	dsn := "host=" + dbHost + " user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbDatabase + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func sqlserverConnection(dbDatabase, dbHost, dbPort, dbUsername, dbPassword string) *gorm.DB {
	dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func sqliteConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func DB() *gorm.DB {
	return db
}

func CloseDB() {
	if sqlDB == nil {
		return
	}
	if err := sqlDB.Close(); nil != err {
		fmt.Println("Disconnect from database failed: ", err.Error())
	}
}
