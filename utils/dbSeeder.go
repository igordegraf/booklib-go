package utils

import (
	"fmt"
	"igor/booklib/models"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	
	////// AUTHORS

	author1 := models.Author{Fio: "Фамилия1 Имя1 Отчество1", BirthDate: "1990-01-01", DeathDate: ""}
	if err := db.Create(&author1).Error; err != nil {
		return err;
	}

	fmt.Println("SEED. Author "+author1.Fio+" ID="+fmt.Sprint(author1.ID))

	author2 := models.Author{Fio: "Фамилия2 Имя2 Отчество2", BirthDate: "1991-03-01", DeathDate: "2020-05-01"}
	if err := db.Create(&author2).Error; err != nil {
		return err;
	}

	fmt.Println("SEED. Author "+author2.Fio+" ID="+fmt.Sprint(author2.ID))

	author3 := models.Author{Fio: "Фамилия3 Имя3 Отчество3", BirthDate: "1980-08-08", DeathDate: "2000-10-10"}
	if err := db.Create(&author3).Error; err != nil {
		return err;
	}

	fmt.Println("SEED. Author "+author3.Fio+" ID="+fmt.Sprint(author3.ID))

	////// BOOKS

	book1 := models.Book{Name: "Книга 1", Annotaion: "Описание книги 1", PublishDate: "2020-02-01", Authors: []models.Author{author1, author2}}
	if err := db.Create(&book1).Error; err != nil {
		return err;
	}

	fmt.Println("SEED. Book "+book1.Name+" ID="+fmt.Sprint(book1.ID))

	book2 := models.Book{Name: "Книга 2", Annotaion: "Описание книги 2", PublishDate: "2020-02-02", Authors: []models.Author{author1, author2, author3}}
	if err := db.Create(&book2).Error; err != nil {
		return err;
	}

	fmt.Println("SEED. Book "+book2.Name+" ID="+fmt.Sprint(book2.ID))

	book3 := models.Book{Name: "Книга 3", Annotaion: "Описание книги 3", PublishDate: "2020-03-02", Authors: []models.Author{author2, author3}}
	if err := db.Create(&book3).Error; err != nil {
		return err;
	}

	fmt.Println("SEED. Book "+book3.Name+" ID="+fmt.Sprint(book3.ID))

	return nil

}