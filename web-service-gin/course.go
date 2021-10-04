package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type course struct {
	ID                int32   `json:"courseId"`
	Name              string  `json:"name"`
	Ects              int32   `json:"ects"`
	CourseResponsible int32   `json:"courseResponsibleId"`
	NRatings          int32   `json:"nRatings"`
	AvgRatings        float32 `json:"avgRatings"`
	ActiveStudents    []int32 `json:"activeStudentsId"`
}

// getCourses responds with the list of all courses as JSON.
func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func postCourse(c *gin.Context) {
	var newCourse course

	// Call BindJSON to bind the received JSON to
	// newCourse
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}

	// Add the new album to the slice.
	courses = append(courses, newCourse)
	c.IndentedJSON(http.StatusCreated, newCourse)
}

// getCouseByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getCourseByID(c *gin.Context) {
	var id int
	id, _ = strconv.Atoi(c.Param("courseId"))
	var idd = int32(id)

	// Loop over the list of courses, looking for
	// a course whose ID value matches the parameter.
	for _, a := range courses {
		if a.ID == idd {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
}

func updateCourse(c *gin.Context) {
	var id int
	id, _ = strconv.Atoi(c.Param("courseId"))
	var idd = int32(id)

	var newCourse course
	for _, a := range courses {
		if a.ID == idd {
			newCourse = a
		}
	}
	name := c.DefaultQuery("name", "")
	ects := c.DefaultQuery("ects", "")
	courseResponsible := c.DefaultQuery("courseResponsible", "")
	nRatings := c.DefaultQuery("nRatings", "")
	avgRatings := c.DefaultQuery("avgRatings", "")
	activeStudents := c.DefaultQuery("activeStudents", "")

	if name != "" {
		newCourse.Name = name
	}
	if ects != "" {
		var standIn int
		standIn, _ = strconv.Atoi(ects)
		newCourse.Ects = int32(standIn)
	}
	if courseResponsible != "" {
		json.Unmarshal([]byte(courseResponsible), &newCourse.CourseResponsible)
	}
	if nRatings != "" {
		var standIn int
		standIn, _ = strconv.Atoi(nRatings)
		newCourse.NRatings = int32(standIn)
	}
	if avgRatings != "" {
		var standIn float64
		standIn, _ = strconv.ParseFloat(avgRatings, 64)
		newCourse.AvgRatings = float32(standIn)
	}
	if activeStudents != "" {
		json.Unmarshal([]byte(activeStudents), &newCourse.ActiveStudents)
	}

}

func deleteCourse(c *gin.Context) {
	var id int
	id, _ = strconv.Atoi(c.Param("courseId"))
	var idd = int32(id)

	var newCourses = make([]course, 0)
	for _, a := range courses {
		if a.ID != idd {
			newCourses = append(newCourses, a)
		}
	}
	for _, a := range newCourses {
		fmt.Println(a)
	}
	courses = newCourses
}
