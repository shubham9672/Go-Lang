package db


type Student struct{
	Id int `json:"id"`
	Name string  `json:"name"`
	Rollno string  `json:"rollno"`
	Branch string  `json:"branch"`
}