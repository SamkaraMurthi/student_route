package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
}

// Inisiasi Student Database
var students []Student

// Inisiasi Web Server routes
func initRoutes() {
	e := echo.New()

	// Routing
	e.GET("/students", getStudents)
	e.GET("/students/:id", getStudent)
	e.POST("/students", createStudent)
	e.PUT("/students/:id", updateStudent)
	e.DELETE("/students/:id", deleteStudent)

	e.Start(":8080")
}

// TODO Problem and Testing 1: Buat Fungsi yang memberikan seluruh informasi student
func getStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, students)
}

// TODO Problem and Testing 2: Buat Fungsi yang memberikan informasi student berdasarkan ID
func getStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, student := range students {
		if student.ID == id {
			return c.JSON(http.StatusOK, student)
		}
	}
	return c.JSON(http.StatusOK, students)
}

// TODO Problem and Testing 3: Buat Fungsi yang membuat informasi student baru
func createStudent(c echo.Context) error {
	student := new(Student)
	if err := c.Bind(student); err != nil {
		return err
	}
	students = append(students, *student)
	return c.JSON(http.StatusCreated, student)
}

// TODO Problem and Testing 4: Buat Fungsi yang dapat mengupdate informasi student berdasarkan ID
func updateStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i := range students {
		if students[i].ID == id {
			updatedStudent := new(Student)
			if err := c.Bind(updatedStudent); err != nil {
				return err
			}

			students[i].Name = updatedStudent.Name
			students[i].Age = updatedStudent.Age
			students[i].Grade = updatedStudent.Grade

			return c.JSON(http.StatusOK, students[i])
		}
	}

	return c.JSON(http.StatusNotFound, "Student not found")
}

// TODO Problem and Testing 5: Buat Fungsi yang dapat menghapus informasi student berdasarkan ID
func deleteStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i := range students {
		if students[i].ID == id {
			students = append(students[:i], students[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, "Student not found")
}
