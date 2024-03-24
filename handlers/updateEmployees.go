package handlers

import (
	"employee-rest-api/config"
	"employee-rest-api/models"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

/*
PUT /employees/:id - Memperbaharui data karyawan
- Request dari endpoint ini harus meliputi nama, email, no handphone, dari karyawan baru tersebut
- Response dari endpoint ini harus berupa sebuah message sukses, dan data karyawan yang
berhasil diperbaharui, jika terdapat kesalahan pada request maka response harus terdiri
dari message yang menjelaskan, jika karyawan dengan ID tidak ditemukan, maka response harus
terdiri dari message yang menjelaskan bahwa data karyawan tidak ditemukan
*/
func UpdateEmployees(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idEmp := params.ByName("id")
	var employee models.Employees
	var response models.Response

	json.NewDecoder(r.Body).Decode(&employee)

 // First query on the database if the id params is avaliable on the database
 config.DB.QueryRow("SELECT id, name, email, phone FROM employees WHERE id = $1", idEmp).Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Phone)
 log.Println("Employee ID :", employee.ID)
 // If the employee ID not found it will write error as response
 if employee.ID == 0 {
 		response.Status = http.StatusNotFound
 		response.Message = "Error Employee ID Not Found"
 		w.WriteHeader(http.StatusInternalServerError)
 		json.NewEncoder(w).Encode(response)
 		log.Println("Error Employee ID Not Found") 
 		return
 // if employee ID avaliable on database, continue to update that employee ID data on the database
 } else {
  // Statement to update employees data
	stmt, err := config.DB.Prepare("UPDATE employees SET name = $1, email = $2, phone = $3, updated_at = $4 WHERE id = $5")
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = "Error " + err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		log.Println("Error Prepare Statement on Update Employees", err.Error()) 
		return
	}

	_, err = stmt.Exec(employee.Name, employee.Email, employee.Phone, time.Now(), idEmp)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = "Error " + err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		log.Println("Error Execution Statement on Update Employees", err.Error()) 
		return
	}

	// Success Message
	response.Status = 200
	response.Message = "Success Updated Employees Data by ID"
	response.Data = employee

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
 }

}
