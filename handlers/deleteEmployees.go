package handlers

import (
	"employee-rest-api/config"
	"employee-rest-api/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/*
DELETE /employees/:id - Menghapus data karyawan
- Response dari endpoint ini adalah data karyawan sesuai dengan ID pada parameter endpoint 
dan message yang yang menjelaskan bahwa proses penghapusan data karyawan berhasil, 
jika karyawan dengan ID tidak ditemukan, maka response harus terdiri dari message yang 
menjelaskan bahwa data karyawan tidak ditemukan
*/
func DeleteEmployees(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Get request params 
	idEmp := params.ByName("id")
	var response models.Response
  var employee models.Employees

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
	// if employee ID avaliable on database, continue to delete that employee ID data on the database
	} else {
		stmt, err := config.DB.Prepare("DELETE FROM employees WHERE id = $1")
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Message = "Error " + err.Error()
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
			log.Println("Error Prepare Statement on Delete Employees", err.Error()) 
			return
		}
		_, err = stmt.Exec(idEmp)
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Message = "Error " + err.Error()
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
			log.Println("Error Execution Statement on Delete Employees", err.Error()) 
			return
		}
	
		// Success Message 
		response.Status = http.StatusOK
		response.Message = "Success Deleted Employees Data by ID"
		response.Data = employee
	
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

}
