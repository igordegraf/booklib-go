package models

import (
	"errors"
	//"fmt"
	/*"gorm.io/datatypes"*/
	"gorm.io/gorm"
)


type Author struct {
	//gorm.Model
	ID                                 uint   `gorm:"primarykey;autoIncrement:true" json:"id"`
	Fio                                string `gorm:"not null" json:"fio"`
	BirthDate/*datatypes.Date*/ string        `gorm:"not null" json:"birth_date"`
	DeathDate/*datatypes.Date*/ string        `json:"death_date"`
	BooksCount                         int    `gorm:"-:migration;<-:false" json:"books_count"`
	Books                              []Book `gorm:"many2many:book_author" json:"books"`
}

//Проверить входящие данные
func (author *Author) Validate() (bool, error) {

	if author.Fio == "" {
		return false, errors.New("не задано ФИО автора")
	}

	if author.BirthDate == "" {
		return false, errors.New("не задана дата рождения автора")
	}

	return true, nil
}

//Проверить существование автора по ФИО
func (author *Author) ExistsByFio() (bool, error) {

	existsAuthor := &Author{}

	err := GetDB().Model(author).Where("fio = ?", author.Fio).First(existsAuthor).Error

	// fmt.Println(err)
	// fmt.Println(existsAuthor.ID)
	// fmt.Println(existsAuthor.Fio)

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, errors.New("DB query error")
	}

	if existsAuthor.ID > 0 {
		return true, nil
	}

	return false, nil
}

// создание записи автора
func (author *Author) Create() (uint, error) {

	if ok, err := author.Validate(); !ok {
		return 0, err
	}

	ok, err := author.ExistsByFio()

	if ok {
		return 0, errors.New("athor already exists")
	}

	if err != nil {
		return 0, err
	}

	err = GetDB().Create(author).Error

	if err != nil || author.ID <= 0 {
		return 0, err
	}

	return author.ID, nil
}

func (author *Author) Delete() (bool, error) {

	err := GetDB().Delete(author).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func GetAuthor(id uint) (*Author, error) {

	author := &Author{}
	
	err := GetDB().
		Select("*, (select count(1) from book_author where author_id=authors.id) as books_count").
		//Preload("Books").
		First(author, id).
		Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if author.ID > 0 {
		return author, nil
	}

	return nil, nil
}

func GetAuthors(authors *[]Author) error {

	result := GetDB().
		Select("*, (select count(1) from book_author where author_id=authors.id) as books_count").
		//Preload("Books").
		Find(authors)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
