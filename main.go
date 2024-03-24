package main

import (
	"employee-rest-api/config"
	"employee-rest-api/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)
func main() {
	// Call the Connection to Database Function
	config.ConnectDB()
	defer config.DB.Close()

	router := httprouter.New()

	// Retrieve All Employees Data
	router.GET("/employees", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		handlers.GetAllEmployees(w, r)
	})

	// Retrieve Employees Data by ID
	router.GET("/employees/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handlers.GetEmployeesByID(w, r, p)
	})

	// Create New Employees Data
	router.POST("/employees", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		handlers.CreateEmployees(w, r)
	})

	// Update Employees Data By ID
	router.PUT("/employees/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handlers.UpdateEmployees(w, r, p)
	})

	// Delete Employees Data By ID
	router.DELETE("/employees/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handlers.DeleteEmployees(w, r, p)
	})
	
  fmt.Println("Running Server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}