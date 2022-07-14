package main

import (
	"igor/booklib/controllers"
	"igor/booklib/models"
	"igor/booklib/utils"
	"log"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"fmt"
	"net/http"
	"os"
)

func main() {

	//Загрузить файл .env
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file. "+e.Error())
	}

	if os.Getenv("db_seed") == "true" {
		if err := utils.Seed(models.GetDB()); err != nil {
			log.Fatal("Error seeding DB. "+e.Error())
		}
	}

	router := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	//router.Use(app.Authentication) // middleware аутентификации

	router.HandleFunc("/book/raw", controllers.GetAllBooks).Methods("GET")                                     //get_all_books
	router.HandleFunc("/book", utils.UnderConstructionResponse).Methods("GET")                                 //get_all_books_paginate
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")                                         //create_book
	router.HandleFunc("/book/by_author/{book_id}", utils.UnderConstructionResponse).Methods("GET")             //find_books_by_author
	router.HandleFunc("/book/{book_id}", controllers.GetBook).Methods("GET")                                   //get_one_book
	router.HandleFunc("/book/{book_id}", utils.UnderConstructionResponse).Methods("PUT")                       //update_book
	router.HandleFunc("/book/{book_id}", controllers.DeleteBook).Methods("DELETE")                             //delete_one_book
	router.HandleFunc("/book/{book_id}/author", controllers.AddAuthor2Book).Methods("POST")                    //add_author
	router.HandleFunc("/book/{book_id}/author/{author_id}", utils.UnderConstructionResponse).Methods("DELETE") //delete_author_from_book
	router.HandleFunc("/book/{book_id}/authors", utils.UnderConstructionResponse).Methods("GET")               //get_book_authors

	router.HandleFunc("/author/raw", controllers.GetAllAuthors).Methods("GET")                     //get_all_authors
	router.HandleFunc("/author", utils.UnderConstructionResponse).Methods("GET")                   //get_all_authors_paginate
	router.HandleFunc("/author", controllers.CreateAuthor).Methods("POST")                         //create_author
	router.HandleFunc("/author/{author_id}", controllers.GetAuthor).Methods("GET")                 //get_one_author
	router.HandleFunc("/author/{author_id}", utils.UnderConstructionResponse).Methods("PUT")       //update_author
	router.HandleFunc("/author/{author_id}", controllers.DeleteAuthor).Methods("DELETE")           //delete_one_author
	router.HandleFunc("/author/{author_id}/books", utils.UnderConstructionResponse).Methods("GET") //authors_books

	router.NotFoundHandler = http.HandlerFunc(utils.UnknownApiCallResponse)

	port := os.Getenv("port") //Получить порт из файла .env
	if port == "" {
		port = "8081"
	}

	fmt.Println("Listening on port " + port)

	log.Fatal(http.ListenAndServe(":"+port, router))

}
