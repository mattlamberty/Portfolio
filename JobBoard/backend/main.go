package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Job represents a job posting
type Job struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Salary      int    `json:"salary"`
}

// In-memory job storage
var jobs = []Job{
	{ID: 1, Title: "Software Engineer", Description: "Develop and maintain software applications.", Company: "Tech Corp", Location: "New York, NY", Salary: 120000},
	{ID: 2, Title: "Data Scientist", Description: "Analyze and interpret complex data.", Company: "Data Inc", Location: "San Francisco, CA", Salary: 130000},
}

func main() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access")
		c.Next()
	})