package controller

import (
	"assignment3/models"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatusData(ctx *gin.Context) {
	var data models.Status

	data.Water = rand.Intn(30)
	data.Wind = rand.Intn(30)

	if data.Water < 5 {
		data.WaterStatus = "aman"
	} else if data.Water >= 6 && data.Water <= 8 {
		data.WaterStatus = "siaga"
	} else if data.Water > 8 {
		data.WaterStatus = "bahaya"
	}

	if data.Wind < 6 {
		data.WindStatus = "aman"
	} else if data.Wind >= 7 && data.Wind <= 15 {
		data.WindStatus = "siaga"
	} else if data.Wind > 15 {
		data.WindStatus = "bahaya"
	}

	html, err := template.ParseFiles("./status.html")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	html.Execute(ctx.Writer, map[string]interface{}{
		"Title":   "Assignment 3",
		"Payload": data,
	})
	fmt.Println("inas", data)
}
