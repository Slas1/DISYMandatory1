package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/courses", getCourses)
	router.POST("/courses", postCourse)
	router.GET("/courses/:courseId", getCourseByID)
	router.PUT("/courses/:courseId", updateCourse)
	router.DELETE("/courses/:courseId", deleteCourse)

	router.Run("localhost:8080")
}

// albums slice to seed record album data.
var courses = []course{
	{ID: 1, Name: "Danish", Ects: 1, CourseResponsible: new(teacher), NRatings: 100, AvgRatings: 1.1, ActiveStudents: make([]*student, 10)},
	{ID: 2, Name: "Math", Ects: 15, CourseResponsible: new(teacher), NRatings: 10, AvgRatings: 9.9, ActiveStudents: make([]*student, 2)},
	{ID: 3, Name: "History", Ects: 5, CourseResponsible: new(teacher), NRatings: 30, AvgRatings: 5.0, ActiveStudents: make([]*student, 7)},
	{ID: 4, Name: "SoftwareEngenering", Ects: 60, CourseResponsible: new(teacher), NRatings: 50, AvgRatings: 6.0, ActiveStudents: make([]*student, 12)},
}

/*
var courses = []course{
	{ID: 1, Name: "Danish", Ects: 1, NRatings: 100, AvgRatings: 1.1},
	{ID: 2, Name: "Math", Ects: 15, NRatings: 10, AvgRatings: 9.9},
	{ID: 3, Name: "History", Ects: 5, NRatings: 30, AvgRatings: 5.0},
	{ID: 4, Name: "SoftwareEngenering", Ects: 60, NRatings: 50, AvgRatings: 6.0},
}
*/
