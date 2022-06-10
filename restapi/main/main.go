package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"restapi/controller"	
	"restapi/db"	
)


func handleReq(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",controller.HomePage)
    myRouter.HandleFunc("/students", controller.GetAll)
    myRouter.HandleFunc("/addstudent", controller.Addstudent).Methods("POST")
	myRouter.HandleFunc("/student/{id}",controller.DeleteStudent).Methods("DELETE")
	myRouter.HandleFunc("/student/{id}",controller.UpdateStudent).Methods("PUT")
	myRouter.HandleFunc("/student/{id}",controller.GetStudent)
	log.Println("listening on port 3000....")
    log.Fatal(http.ListenAndServe(":3000", myRouter))
}
func main() {
	db.CreateDb()
	db.CreateInstance()
	handleReq()
}


