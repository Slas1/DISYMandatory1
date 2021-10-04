package main

import (
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
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}
	var oriCour *course
	for _, a := range courses {
		if a.ID == idd {
			oriCour = &a
		}
	}
	if newCourse.Name != "" {
		oriCour.Name = newCourse.Name
	}
	if newCourse.Ects != 0 {
		oriCour.Ects = newCourse.Ects
	}
	if newCourse.CourseResponsible != 0 {
		oriCour.CourseResponsible = newCourse.CourseResponsible
	}
	if newCourse.NRatings != 0 {
		oriCour.NRatings = newCourse.NRatings
	}
	if newCourse.AvgRatings != 0 {
		oriCour.AvgRatings = newCourse.AvgRatings
	}
	if newCourse.ActiveStudents != nil {
		oriCour.ActiveStudents = newCourse.ActiveStudents
	}

	var newCourses = make([]course, 0)
	for _, a := range courses {
		if a.ID != idd {
			newCourses = append(newCourses, a)
		} else {
			newCourses = append(newCourses, *oriCour)
		}
	}
	courses = newCourses

	c.IndentedJSON(http.StatusOK, newCourse)
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
	courses = newCourses
	c.IndentedJSON(http.StatusOK, courses)
}
