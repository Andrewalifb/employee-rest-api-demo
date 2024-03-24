package handlers

import (
	"employee-rest-api/config"
	"employee-rest-api/models"
	"encoding/json"
	"log"
	"net/http"
)

/*
POST /employees - Menyimpan data karyawan baru,
- Request dari endpoint ini harus meliputi nama, email, no handphone, dari karyawan baru tersebut
- Response dari endpoint ini harus berupa sebuah message sukses, dan data karyawan yang berhasil 
disimpan, jika terdapat kesalahan pada request maka response harus terdiri dari message yang menjelaskan kesalahan pada input request nya.
*/
func CreateEmployees(w http.ResponseWriter, r *http.Request) {
	var employee models.Employees
	var response models.Response

	json.NewDecoder(r.Body).Decode(&employee)

	// Statement insert name, email and phone into employees
	stmt, err := config.DB.Prepare("INSERT INTO employees(name, email, phone) VALUES ($1, $2, $3)")
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = "Error " + err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		log.Println("Error Prepare Statement on Create Employees", err.Error()) 
		return
	}
   
  // Error message if there is error on exec and if the user input not unique email 
	_, err = stmt.Exec(employee.Name, employee.Email, employee.Phone)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = "Error " + err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		log.Println("Error Statement Execution on Create Employees", err.Error()) 
		return
	}
  // Success Message
	response.Status = http.StatusCreated
	response.Message = "Success Created New Employees Data"
	response.Data = employee

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
