package controllers

import (
	"crud-api/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var b models.Book
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			http.Error(w, "Error while parsing Request", http.StatusBadRequest)
			return
		}
		if b.Name == "" || b.Author == "" || b.Publication == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		err = models.CreateBook(&b, db)
		if err != nil {
			http.Error(w, "Error while creating book", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(b)
	}

}
func GetBooks(db *gorm.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var res []models.Book

		err := models.GetBooks(&res, db)
		if err != nil {
			http.Error(w, "Error while fetching books", http.StatusInternalServerError)
			return
		}

		_ = json.NewEncoder(w).Encode(res)

	}

}
func GetBookById(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val := mux.Vars(r)

		id := val["id"]
		fmt.Println(id)
		idnum, _ := strconv.Atoi(id)
		fmt.Println(idnum)

		var b models.Book

		err := models.GetBookById(&b, db, idnum)
		if err != nil {
			http.Error(w, "Unable to process request", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(b)

	}
}
func DeleteBookById(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mp := mux.Vars(r)
		idstr := mp["id"]
		id, err := strconv.Atoi(idstr)
		if err != nil {
			http.Error(w, "Unable to parse request", http.StatusBadRequest)
			return
		}
		var b models.Book
		err = models.DeleteById(&b, db, id)
		if err != nil {
			http.Error(w, "Unable to delete ", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"id":      id,
			"deleted": true,
		})

	}

}
func UpdateBookById(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mp := mux.Vars(r)
		idstr := mp["id"]
		idnum, err := strconv.Atoi(idstr)
		if err != nil {
			http.Error(w, "Unable to parse id", http.StatusBadRequest)
			return
		}
		var existingDetails models.Book

		err = models.GetBookById(&existingDetails, db, idnum)
		if err != nil {
			http.Error(w, "No details found", http.StatusNotFound)
			return
		}

		var payload models.Book

		err = json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, "Error while parsing request body", http.StatusBadRequest)
			return
		}
		if payload.Name != "" {
			existingDetails.Name = payload.Name
		}
		if payload.Author != "" {
			existingDetails.Author = payload.Author
		}
		if payload.Publication != "" {
			existingDetails.Publication = payload.Publication
		}
		err = models.UpdateBookById(&existingDetails, db)
		if err != nil {
			http.Error(w, "Unable to update ", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(existingDetails)

	}
}
