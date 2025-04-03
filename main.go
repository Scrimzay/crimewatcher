package main

import (
	"os"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/Scrimzay/loglogger"
)

var log *logger.Logger

func init() {
	var err error
	log, err = logger.New("details.txt")
	if err != nil {
		log.Fatalf("Could not start logger: %v", err)
	}
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.html")
	r.Static("/static", "./static")
	r.Static("/statestats", "./templates/statestats")

	r.GET("/", indexHandler)
	r.GET("/crimemap", crimeMapHandler)
	r.GET("/crime-data/:state", crimeDataPerStateHandler)

	err := r.Run(":6677")
	if err != nil {
		log.Fatalf("Could not start gin: %v", err)
	}
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func crimeMapHandler(c *gin.Context) {
	c.HTML(200, "map.html", nil)
}

func crimeDataPerStateHandler(c *gin.Context) {
    state := c.Param("state")
    fileName := fmt.Sprintf("%sstats.html", strings.ToLower(state))

    // Check if the file exists
    if _, err := os.Stat(fmt.Sprintf("templates/statestats/%s", fileName)); os.IsNotExist(err) {
        c.String(404, "File not found for state: %s", state)
        return
    }

    c.HTML(200, fileName, gin.H{"state": state})
}