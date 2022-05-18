package main

import (
	"fmt"
	"strconv"
)

type employee struct {
	id     int
	name   string
	age    uint32
	salary uint64
}

func (e *employee) toString() string {
	return "Employee [ Id: " + strconv.Itoa(e.id) + " Name: " + e.name + " Age: " + strconv.FormatUint(uint64(e.age), 10) + " Salary: " + strconv.FormatUint(uint64(e.salary), 10) + " ]"
}

type geometry interface {
	area() float64
	perim() float64
	typeOfGeometry() string
}
type rectangle struct {
	width, height float64
}
type square struct {
	side float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}
func (r rectangle) typeOfGeometry() string {
	return "Rectangle width: " + fmt.Sprint(r.width) + " Height: " + fmt.Sprint(r.height)
}
func (sq square) area() float64 {
	return sq.side * sq.side
}
func (sq square) perim() float64 {
	return 4 * sq.side
}
func (sq square) typeOfGeometry() string {
	return "Square side: " + fmt.Sprint(sq.side)
}
func measure(g geometry) {
	fmt.Println(g.typeOfGeometry())
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func typeInterface(a interface{}) {

	// Using type switch
	switch a.(type) {

	case int:
		fmt.Println("Type: int, Value:", a.(int))
	case string:
		fmt.Println("\nType: string, Value: ", a.(string))
	case float64:
		fmt.Println("\nType: float64, Value: ", a.(float64))
	default:
		fmt.Println("\nType not found")
	}
}

func methodInterface() {
	emp := employee{id: 1, name: "Employee 1", age: 24, salary: 15000}
	fmt.Println(emp.toString())
	rect := rectangle{width: 10, height: 20}
	sq := square{side: 4}
	measure(rect)
	measure(sq)
	typeInterface("Hello")
	typeInterface(244)
	typeInterface(244.26)
	typeInterface(0x2b)

}
