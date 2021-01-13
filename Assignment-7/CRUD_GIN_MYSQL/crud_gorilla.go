package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Truck struct {
	TruckID       int      `json:"truckId" gorm:"primary_key"`
	DriverName    string   `json:"driverName"`
	CleanerName   string   `json:"cleanerName"`
	TruckNo       int      `json:"truckNo"`
}

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:password@/golang?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// db.Exec("CREATE DATABASE order_db")
	db.Exec("USE golang")
	db.AutoMigrate(&Truck{})
}

func createTruck(w http.ResponseWriter, r *http.Request) {
	var truck Truck
	json.NewDecoder(r.Body).Decode(&truck)
	db.Create(&truck)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(truck)
}

func getTrucks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var trucks []Truck
	db.Preload("Items").Find(&trucks)
	json.NewEncoder(w).Encode(trucks)
}

func getTruck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["truckId"]

	var turck Truck
	db.Preload("Items").First(&truck, id)
	json.NewEncoder(w).Encode(truck)
}

func updateTruck(w http.ResponseWriter, r *http.Request) {
	var updatedTruck Truck
	json.NewDecoder(r.Body).Decode(&updatedTruck)
	db.Save(&updatedTruck)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTruck)
}

func deleteTruck(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["truckId"]

	// db.Where("truck_id = ?", idToDelete).Delete(&Item{})
	db.Where("truck_id = ?", id).Delete(&Truck{})
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/trucks", getTrucks).Methods("GET")
	router.HandleFunc("/trucks", createTruck).Methods("POST")
	router.HandleFunc("/trucks/{truckId}", getTruck).Methods("GET")
	router.HandleFunc("/trucks/{truckId}", updateTruck).Methods("PUT")
	router.HandleFunc("/trucks/{truckId}", deleteTruck).Methods("DELETE")

	initDB()
	log.Fatal(http.ListenAndServe(":8080", router))
}
