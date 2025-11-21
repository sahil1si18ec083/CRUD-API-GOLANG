package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func CreateBook(b *Book, db *gorm.DB) error {
	res := db.Create(b)
	return res.Error

}
func GetBooks(res *[]Book, db *gorm.DB) error {
	temp := db.Find(res)
	return temp.Error

}
func GetBookById(b *Book, db *gorm.DB, idnumd int) error {

	temp := db.First(b, "id = ?", idnumd)
	return temp.Error
}
func DeleteById(b *Book, db *gorm.DB, idnumd int) error {

	temp := db.Delete(b, idnumd)
	return temp.Error
}
func UpdateBookById(b *Book, db *gorm.DB) error {

	temp := db.Save(b)
	return temp.Error
}
