package models

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbConnection *gorm.DB //база данных

func init() {


	godotenv.Load()
	dbName := os.Getenv("db_name")

	if dbName == "" {
		log.Fatal("В файле .env не заполнено значение db_name - путь к файлу БД")
	}

	//создаём пустой файл для sqlite СУБД если его нет
	if _, err := os.Stat(dbName); err != nil {
		file, err := os.Create(dbName); 
		if err != nil {
			log.Fatal("Не удалось создать файл БД: "+dbName+". "+err.Error())
		}

		file.Close()
	
	  }
	

	//dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //строка подключения
	fmt.Println(dbName)

	connection, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
	if err != nil {
		log.Fatal(err.Error())
	}

	dbConnection = connection

	//Миграция базы данных
	err = dbConnection.Debug().AutoMigrate( &Book{}, &Author{}) 
	//dbConnection.AutoMigrate(&Book{}, &Author{})
	if err != nil {
		log.Fatal("DB migration error. "+err.Error())
	}

}

// возвращает дескриптор объекта DB
func GetDB() *gorm.DB {
	return dbConnection
}
