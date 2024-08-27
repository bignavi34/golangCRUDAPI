
package user

import (
	"encoding/json"
	"net/http"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)
var DB *gorm.DB
var err error

const DNS ="postgres://postgres:123@localhost:5432/go"

type User struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
}
func  InitialMigration(){
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	
	DB.AutoMigrate(&User{})
	
}
func GetUsers(w http.ResponseWriter, r *http.Request){

}
func CreateUsers(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)

}
func GetUser(w http.ResponseWriter,r *http.Request){

}
func UpdateUsers(w http.ResponseWriter,r *http.Request){

}

func DeleteUser(w http.ResponseWriter,r *http.Request){

}


