package db

import (
	"fmt"
	"database/sql"
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

var sqliteDb *sql.DB

func fileNotExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return true
	}
	return info.IsDir()
}
func setupDb(){
		CreateInstance()
		CreateTable()
		AddStudent(Student{Rollno:"0001", Name:"Liana Kim",Branch : "EC"})
		AddStudent( Student{Rollno:"0002",Name: "Glen Rangel", Branch:"CS"})
		AddStudent( Student{Rollno:"0003",Name: "Martin Martins", Branch:"ME"})
		AddStudent( Student{Rollno:"0004",Name: "Alayna Armitage", Branch:"IT"})
		sqliteDb.Close()
}
func CreateDb() {
	if fileNotExists("sqlite-database.db") {
		log.Println("Creating sqlite-database.db...")
		file, err := os.Create("sqlite-database.db")
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		setupDb()
		log.Println("sqlite-database.db created")
	}

}
func CreateInstance() {
	sqliteDb, _ = sql.Open("sqlite3", "./sqlite-database.db")
}
func CreateTable() {
	createStudentTableSQL := `CREATE TABLE IF NOT EXISTS student (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"rollno" TEXT,
		"name" TEXT,
		"branch" TEXT		
	  );`

	statement, err := sqliteDb.Prepare(createStudentTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}
func AddStudent(student Student) string {
	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO student(rollno, name, branch) VALUES (?, ?, ?)`
	statement, err := sqliteDb.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
		return "Some Error Occured"
	}
	_, err = statement.Exec(student.Rollno,student.Name,student.Branch)
	if err != nil {
		log.Fatalln(err.Error())
		return "Some Error Occured"
	}
	return "Successfully Added"
}
func GetAllStudents() []Student{
	row, err := sqliteDb.Query("SELECT * FROM student ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var students []Student
	for row.Next() { 
		student:=Student{}
		row.Scan(&student.Id, &student.Rollno, &student.Name, &student.Branch)
		students=append(students,student)
		student=Student{}
	}
	return students
}
func GetStudent(id int) Student{
	row, err:=sqliteDb.Query("SELECT * FROM student WHERE id=?",id)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	student:=Student{}
	if row.Next(){
	row.Scan(&student.Id, &student.Rollno, &student.Name, &student.Branch)
	}
	return student
}
func CheckIfStudentExists(id int) bool {
	var student Student
	student=GetStudent(id)
	if student.Id == 0 {
		return false
	}
	return true
}
func UpdateStudent(id int, student Student) string{
	stmt, err := sqliteDb.Prepare("UPDATE student set name = ?, rollno = ?, branch = ? where id = ?")
	if err != nil {
		log.Fatalln(err.Error())
		return "Some Error Occured"
	}
	defer stmt.Close()

	_, err1 := stmt.Exec(student.Name,student.Rollno,student.Branch,student.Id)
	if err1 != nil {
		log.Fatalln(err.Error())
		return "Some Error Occured"
	}
	res:=fmt.Sprint("Student ID: ",id," got updated")
	return res

}
func DeleteStudent(id int) string{
stmt,err:=sqliteDb.Prepare("DELETE FROM student where id = ?")
if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatalln(err.Error())
		return "Some Error Occured"
	}
	return "Successfully Deleted"
}