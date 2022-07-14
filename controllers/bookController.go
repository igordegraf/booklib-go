package controllers

import (
	"encoding/json"
	"igor/booklib/models"
	"igor/booklib/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AddAuthor2BookParams struct {
	AuthorId uint `json:"author_id"`
}

var CreateBook = func(w http.ResponseWriter, r *http.Request) {

	book := models.Book{}
	err := json.NewDecoder(r.Body).Decode(&book) //декодирует тело запроса в struct и завершается неудачно в случае ошибки

	if err != nil {
		utils.JsonResponse(w, utils.ErrorResultMessage(err.Error()), http.StatusBadRequest)
		return
	}

	id, err := book.Create()
	if err != nil {
		utils.JsonResponse(w, utils.ErrorResultMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, utils.CreateEntityResultMessage(id), http.StatusCreated)
}

var GetBook = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["book_id"])

	if id == 0 {
		utils.JsonResponse(w, utils.ErrorResultMessage("bad book ID"), http.StatusBadRequest)
		return
	}

	book, err := models.GetBook(uint(id))

	if err != nil {
		utils.JsonResponse(w, utils.ErrorResultMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	if book == nil {
		utils.JsonResponse(w, utils.ErrorResultMessage("Record not found"), http.StatusNotFound)
		return
	}

	utils.JsonResponse(w, book, http.StatusOK)
}

var DeleteBook = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["book_id"])

	if id == 0 {
		utils.JsonResponse(w, utils.ErrorResultMessage("bad book ID"), http.StatusBadRequest)
		return
	}

	book, err := models.GetBook(uint(id))

	if err != nil {
		utils.JsonResponse(w, utils.ErrorResultMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	if book == nil {
		utils.JsonResponse(w, utils.ErrorResultMessage("Record not found"), http.StatusNotFound)
		return
	}

	ok, err := book.Delete()

	if ok {
		utils.JsonResponse(w, nil, http.StatusNoContent)
		return
	}

	if err != nil {
		utils.JsonResponse(w, utils.ErrorResultMessage("Error deleting record: "+err.Error()), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, utils.ErrorResultMessage("Unknown error"), http.StatusInternalServerError)
}

var GetAllBooks = func(w http.ResponseWriter, r *http.Request) {

	var books []models.Book

	err := models.GetBooks(&books)

	if err != nil {
		utils.JsonResponse(w, utils.ErrorResultMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, books, http.StatusOK)
}

var AddAuthor2Book = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["book_id"])

	if id == 0 {
		utils.JsonResponse(w, utils.ErrorResultMessage("bad book ID"), http.StatusBadRequest)
		return
	}

	params := AddAuthor2BookParams{}
	err := json.NewDecoder(r.Body).Decode(&params) //декодирует тело запроса в struct и завершается неудачно в случае ошибки

	if err != nil {
		utils.JsonResponse(w, utils.ErrorResultMessage(err.Error()), http.StatusBadRequest)
		return
	}

	if params.AuthorId == 0 {
		utils.JsonResponse(w, utils.ErrorResultMessage("bad author ID"), http.StatusBadRequest)
		return
	}

	authorId := params.AuthorId

	//поиск экземпляра книги
	book, err := models.GetBook(uint(id))

	if err != nil {
		utils.JsonResponse(w, utils.ErrorResultMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	if book == nil {
		utils.JsonResponse(w, utils.ErrorResultMessage("Book not found"), http.StatusNotFound)
		return
	}

	//поиск автора в базе
	//поиск экземпляра книги
	author, err := models.GetAuthor(authorId)
	
	if err != nil {
		utils.JsonResponse(w, utils.ErrorResultMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	if author == nil {
		utils.JsonResponse(w, utils.ErrorResultMessage("Author not found"), http.StatusNotFound)
		return
	}

	err = book.AddAuthor(authorId)

	if err != nil {
		utils.JsonResponse(w, utils.ErrorResultMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, "Author added to book", http.StatusCreated)
}
