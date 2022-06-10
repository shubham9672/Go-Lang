package controller
import (
	"io/ioutil"
	"strconv"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"restapi/db"	
)
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page") // sending text response
}
func GetAll(w http.ResponseWriter,r *http.Request){
	studentsJson,_:=json.MarshalIndent(db.GetAllStudents(),""," ")
	fmt.Fprintf(w,string(studentsJson))
}
func GetStudent(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	id,_:=strconv.Atoi(vars["id"])
	json.NewEncoder(w).Encode(db.GetStudent(id))
}
func DeleteStudent(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	id,_:=strconv.Atoi(vars["id"])
	if db.CheckIfStudentExists(id) == false {
		json.NewEncoder(w).Encode("Student Not Found!")
		return
	}
	fmt.Fprintf(w,db.DeleteStudent(id))
}
func Addstudent(w http.ResponseWriter, r *http.Request){
	 reqBody, _ := ioutil.ReadAll(r.Body)
	 var student db.Student
	 json.Unmarshal(reqBody,&student)
	 fmt.Fprintf(w,db.AddStudent(student))
}
func UpdateStudent(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	id,_:=strconv.Atoi(vars["id"])
	if db.CheckIfStudentExists(id) == false {
		json.NewEncoder(w).Encode("Student Not Found!")
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newstudent db.Student
	json.Unmarshal(reqBody,&newstudent)
	fmt.Fprintf(w,db.UpdateStudent(id,newstudent))
}