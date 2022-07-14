package models

import (
	"errors"
	//"fmt"
	/*"gorm.io/datatypes"*/
	"gorm.io/gorm"
)


type Book struct {
  //gorm.Model
  ID    uint `gorm:"primarykey;autoIncrement:true" json:"id"`
  Name string `json:"name"`
  PublishDate /*datatypes.Date*/ string `json:"publish_date"`
  Annotaion string `json:"annotation"`
  Authors []Author `gorm:"many2many:book_author" json:"authors"`
}

//Проверить входящие данные 
func (book *Book) Validate() (bool, error) {


  if book.Name == "" {
	return false, errors.New("не задано название книги")
  }

  if book.PublishDate == "" {
	return false, errors.New("не задана дата публикации книги")
  }

  if book.Annotaion == "" {
	return false, errors.New("не задано описание книги")
  }

  return true, nil
}

//Проверить существования книги пот названию
func (book *Book) ExistsByName() (bool, error) {

  existsBook := &Book{}

  err := GetDB().Model(book).Where("name = ?", book.Name).First(existsBook).Error


  if err != nil && err != gorm.ErrRecordNotFound {
    return false, errors.New("DB query error")
  }

  if existsBook.ID > 0 {
    return true, nil
  }

  return false, nil
}

// создание записи книги
func (book *Book) Create() (uint, error) {

  
  if ok, err := book.Validate(); !ok {
    return 0, err
  }
  
  ok, err := book.ExistsByName();

  if ok {
    return 0, errors.New("athor already exists")
  }

  if err != nil {
    return 0, err
  }

  err = GetDB().Create(book).Error

  if err != nil || book.ID <= 0 {
    return 0, err
  }

  return book.ID, nil
}

func (book *Book) Delete() (bool, error) {

  err := GetDB().Delete(book).Error
  
  if err != nil {
    return false, err
  }

  return true, nil
}

func GetBook(id uint) (*Book, error) {

  book := &Book{}
  //err := GetDB().Model(&Book{}).Where("id = ?", id).First(acc).Error
  err := GetDB().Preload("Authors").First(book, id).Error

  if err != nil && err != gorm.ErrRecordNotFound  {
	  return nil, err
  }

  if book.ID > 0 {
    return book, nil
  }

  return nil, nil
}

func GetBooks(books *[]Book) error {
  
  result := GetDB().Preload("Authors").Find(books)

  if result.Error != nil {
    return result.Error
  }

  return nil
}

func (book *Book) AddAuthor(authorId uint) (error)  {
	
	return GetDB().Model(book).Association("Authors").Append(&Author{ID: authorId})

}