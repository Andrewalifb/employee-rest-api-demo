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
GET /employees/:id - Menampilkan data karyawan berdasarkan ID
- Response dari endpoint ini adalah data karyawan sesuai dengan ID pada parameter endpoint, 
jika karyawan dengan ID tidak ditemukan, maka response harus terdiri dari message yang 
menjelaskan bahwa data karyawan tidak ditemukan
*/
func GetEmployeesByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idEmp := params.ByName("id")

	var employee models.Employees
	var response models.Response
  // Get Employee By ID Query
	rows, err := config.DB.Query("SELECT id, name, email, phone, created_at, updated_at FROM employees WHERE id = $1", idEmp)
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = "Error " + err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		log.Println("Error Query on Get Employees by id", err.Error()) 
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Phone, &employee.CreatedAt, &employee.UpdatedAt)
		if err != nil {
			response.Status = http.StatusInternalServerError
			response.Message = "Error " + err.Error()
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
			log.Println("Error on rows.Scan on Get Employee by id", err.Error()) 
			return
		}
	}
	// Check if there is no employee ID found on database response error
	if employee.ID == 0 {
		response.Status = http.StatusNotFound
		response.Message = "Error Employee ID :" + idEmp + " Not Found"
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		log.Println("Error Employees ID " + idEmp + " Not Found") 
		return
	// Success Message 
	} else {
		response.Status = http.StatusOK
		response.Message = "Success Get Employees Data By ID"
		response.Data = employee
	
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

}
