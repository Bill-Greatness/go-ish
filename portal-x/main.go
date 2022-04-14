package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// schema definitions for student Model
type Student struct {
	ID          string `json:"studentID"`
	Name        string `json:"name"`
	Class       string `json:"class"`
	Gender      string `json:"gender"`
	DOB         string `json:"dateOfBirth"`
	Phone       int    `json:"phone"`
	Image       string `json:"profileImage"`
	Status      string `json:"status"`
	Nationality string `json:"countryOfOrigin"`
	Group       int    `json:"yearGroup"`
}

// schema definition for staff Model
type Staff struct {
	ID            string `json:"studentID"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Gender        string `json:"gender"`
	Phone         int    `json:"phone"`
	Image         string `json:"profileImage"`
	Status        string `json:"status"`
	Qualification string `json:"qualification"`
	Group         int    `json:"groupAssigned"`
	WorkType      string `json:"employmentType"`
}

// create a slice of students Type : Temporal Data
var students = []Student{
	{
		Name:        "Bill Greatness",
		ID:          "BG-001-00A",
		Class:       "Basic 7",
		Gender:      "Male",
		Phone:       509343841,
		Image:       "https://some-random-image.com/hey.png",
		Status:      "Active",
		Nationality: "Ghanaian",
		Group:       2022,
		DOB:         "03/11/1998",
	},
}

// create slice of Type Staff: Temporal Data;
// data will be queried from a MongoDB source.
var staff = []Staff{
	{
		Name:          "",
		Gender:        "",
		Qualification: "",
		Phone:         2,
		Image:         "",
		Email:         "",
		Status:        "",
		Group:         2,
		WorkType:      "",
		ID:            "",
	},
}

// Student API Functions
func getStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, students)
}

func addStudent(c *gin.Context) {
	// create a new student variable with type Student
	var newStudent Student

	// I didn't really get the explanation for this.
	if err := c.BindJSON(&newStudent); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	students = append(students, newStudent)
	c.IndentedJSON(http.StatusCreated, students)
}

func queryStudent(c *gin.Context) {
	query := c.Request.URL.Query()
	fmt.Println(query)
}

func getStaff(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, staff)
}

func getStudentByID(c *gin.Context) {
	id := c.Param("id")

	for _, std := range students {
		if std.ID == id {
			c.IndentedJSON(http.StatusOK, std)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Data not Available"})
}

func main() {
	// initialize gin Router.
	router := gin.Default()

	// Defined API Endpoints: Students
	router.GET("/students", getStudents)
	router.POST("/students", addStudent)
	router.GET("/student/:id", getStudentByID)
	router.GET("/query-student", queryStudent)

	// Define API EndPoints: Staff
	router.GET("/staff", getStaff)

	// run Gin Router.
	router.Run("localhost:5001")

}
