package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:root@tcp(127.0.0.1:8889)/goLang?charset=utf8mb4&parseTime=True&loc=Local"

type Employee struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Employee{})
}

func GetEmployes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee []Employee
	DB.Find(&employee)
	json.NewEncoder(w).Encode(employee)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var employee Employee
	DB.First(&employee, params["id"])
	json.NewEncoder(w).Encode(employee)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	json.NewDecoder(r.Body).Decode(&employee)
	DB.Create(&employee)
	json.NewEncoder(w).Encode(employee)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var employee Employee
	DB.First(&employee, params["id"])
	json.NewDecoder(r.Body).Decode(&employee)
	DB.Save(&employee)
	json.NewEncoder(w).Encode(employee)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var employee Employee
	DB.Delete(&employee, params["id"])
	json.NewEncoder(w).Encode("The Employee Deatils are Deleted Successfully!")
}
