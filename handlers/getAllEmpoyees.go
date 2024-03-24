package handlers

import (
	"employee-rest-api/config"
	"employee-rest-api/models"
	"encoding/json"
	"log"
	"net/http"
)

/*
GET /employees - Menampilkan seluruh data karyawan
- Response dari endpoint ini adalah array of object dari karyawan. Data karyawan yang 
ditampilkan pada endpoint ini hanya ID, nama, dan email saja.
*/
func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	var employees []models.Employees
	var response models.Response

	// Get All Employees Query
	rows, err := config.DB.Query("SELECT id, name, email, phone, created_at, updated_at FROM employees")
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = "Error " + err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		log.Println("Error on Query Get All Employees", err.Error()) 
		return
	}
	defer rows.Close()

	for rows.Next() {
		var employee models.Employees
		err := rows.Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Phone, &employee.CreatedAt, &employee.UpdatedAt)
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Message = "Error " + err.Error()
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
			log.Println("Error on rows.Scan Get All Employees", err.Error()) 
			return
		}
		employees = append(employees, employee)
	}
	// Success Message
	response.Status = 200
	response.Message = "Success Get All Employees Data"
	response.Data = employees

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
