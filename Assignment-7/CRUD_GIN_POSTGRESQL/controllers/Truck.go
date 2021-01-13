package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
	guuid "github.com/google/uuid"
)


type Truck struct {
	TruckID       string      `json:"truckId""`
	DriverName    string   `json:"driverName"`
	CleanerName   string   `json:"cleanerName"`
	TruckNo       string      `json:"truckNo"`
}


func CreateTruckTable(db *pg.DB) error {                    // Create User Table
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Truck{}, opts)
	if createError != nil {
		log.Printf("Error while creating truck table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Truck table created")
	return nil
}


var dbConnect *pg.DB                                  // INITIALIZE DB CONNECTION (to avoid to many connections)

func InitiateDB(db *pg.DB) {
	dbConnect = db
}

func GetAllTrucks(c *gin.Context) {
	var trucks []Truck
	err := dbConnect.Model(&trucks).Select()

	if err != nil {
		log.Printf("Error while getting all trucks, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Trucks",
		"data":    trucks,
	})
	return
}

func CreateTruck(c *gin.Context) {
	var truck Truck
	c.BindJSON(&truck)
	driverName := truck.DriverName
	cleanerName := truck.CleanerName
	truckNo := truck.TruckNo
	truckId := guid.New().String()

	insertError := dbConnect.Insert(&Truck{
		TruckID:        truckId,
		DriverName:     driverName,
		CleanerName:    cleanerName,
		TruckNo:        truckNo,
		// CreatedAt: time.Now(),
		// UpdatedAt: time.Now(),
	})
	if insertError != nil {
		log.Printf("Error while inserting new truck into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Truck created Successfully",
	})
	return
}

func GetSingleTruck(c *gin.Context) {
	truckid := c.Param("truckid")
	truck := &Truck{TruckID: truckid}
	err := dbConnect.Select(truck)

	if err != nil {
		log.Printf("Error while getting a single truck, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Truck",
		"data":    truck,
	})
	return
}

func EditTruck(c *gin.Context) {
	truckid := c.Param("truckid")
	var truck Truck
	c.BindJSON(&truck)
	truckNo := truck.TruckNo

	_, err := dbConnect.Model(&Truck{}).Set("truckNo = ?", truckNo).Where("truckId = ?", truckid).Update()
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Truck Edited Successfully",
	})
	return
}

func DeleteTruck(c *gin.Context) {
	truckid := c.Param("truckid")
	truck := &Truck{TruckID: truckid}

	err := dbConnect.Delete(truck)
	if err != nil {
		log.Printf("Error while deleting a single truck, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Truck deleted successfully",
	})
	return
}
