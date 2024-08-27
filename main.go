package main

import (
	"Goprogram/internal"
	u"Goprogram/user"

	"github.com/gorilla/mux"
	"log"
	"net/http"

)
func InitializeRouter(){
    r:=mux.NewRouter()
    r.HandleFunc("/users",u.GetUsers).Methods("GET")
    r.HandleFunc("/users{id}",u.GetUser).Methods("GET")
    r.HandleFunc("/createusers",u.CreateUsers).Methods("POST")
    r.HandleFunc("/users{id}",u.UpdateUsers).Methods("PUT")
    r.HandleFunc("/users{id}",u.DeleteUser).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8080", r))
    
}

func main() {
	internal.Alu()
	u.InitialMigration()
	

	InitializeRouter()


}



