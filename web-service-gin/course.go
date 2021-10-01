package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type course struct {
	ID                int32      `json:"courseId"`
	Name              string     `json:"name"`
	Ects              int32      `json:"ects"`
	CourseResponsible *teacher   `json:"courseResponsible"`
	NRatings          int32      `json:"nRatings"`
	AvgRatings        float32    `json:"avgRatings"`
	ActiveStudents    []*student `json:"activeStudents"`
}

// getCourses responds with the list of all courses as JSON.
func getCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
