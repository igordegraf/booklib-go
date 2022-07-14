package controllers

import (
	"encoding/json"
	"igor/booklib/models"
	"igor/booklib/utils"
	"strconv"
	"github.com/gorilla/mux"
	"net/http"
)

var CreateAuthor = func(w http.ResponseWriter, r *http.Request) {

	author := models.Author{}
	err := json.NewDecoder(r.Body).Decode(&author) //декодирует тело запроса в struct и завершается неудачно в случае ошибки

	if err != nil {
		utils.JsonResponse(w, utils.ErrorResultMessage(err.Error()), http.StatusBadRequest)
		return
	}

	id, err := author.Create()
	if err != nil {
		utils.JsonResponse(w, utils.ErrorResultMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, utils.CreateEntityResultMessage(id), http.StatusCreated)
}

var GetAuthor = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["author_id"])

	if id == 0 {
		utils.JsonResponse(w, utils.ErrorResultMessage("bad author ID"), http.StatusBadRequest)
		return
	}

	author, err := models.GetAuthor(uint(id))

	if err != nil {
		utils.JsonResponse(w, utils.ErrorResultMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	if author == nil {
		utils.JsonResponse(w, utils.ErrorResultMessage("Record not found"), http.StatusNotFound)
		return
	}

	utils.JsonResponse(w, author, http.StatusOK)
}

var DeleteAuthor = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["author_id"])

	if id == 0 {
		utils.JsonResponse(w, utils.ErrorResultMessage("bad author ID"), http.StatusBadRequest)
		return
	}

	author, err := models.GetAuthor(uint(id))

	if err != nil {
		utils.JsonResponse(w, utils.ErrorResultMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	if author == nil {
		utils.JsonResponse(w, utils.ErrorResultMessage("Record not found"), http.StatusNotFound)
		return
	}

	ok, err := author.Delete()

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

var GetAllAuthors = func(w http.ResponseWriter, r *http.Request) {

	var authors []models.Author

	err := models.GetAuthors(&authors)

	if err != nil {
		utils.JsonResponse(w, utils.ErrorResultMessage(err.Error()), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, authors, http.StatusOK)
}
