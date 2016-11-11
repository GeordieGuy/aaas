package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func timeError(c *gin.Context) {

	content := gin.H{"APIVersion": 2.0, "DayOfAdvent": "Error", "Description": "It is not advent anymore, or not yet", "SaintImage": nil}
	c.IndentedJSON(451, content)
}

func index(c *gin.Context) {
	t := time.Now()
	adventDay := t.Day()
	description := ""
	saintImage := ""
	var readings [4]string

	switch adventDay {
	case 1:
		description = "1st Sunday of Advent"
		saintImage = ""
		readings = append("Is 63:16b-17, 19b,  64:2-7/1", "Cor 1:3-9", "Mk 13:33-37")

	case 2:
		description = "Feast of Saint Andrew, the apostle"
		saintImage = "https://en.wikipedia.org/wiki/Andrew_the_Apostle#/media/File:Artus_Wolffort_-_St_Andrew_-_WGA25857.jpg"
		readings = append("Is 2:1-5", "Mt 8:5-11")
	case 3:
		description = "Tuesday of the 1st Week"
		saintImage = ""
		readings = append("Is 11:1-10", "Lk 10:21-24")
	case 4:
		description = "Wednesday of the 1st Week"
		saintImage = ""
	case 5:
		description = "Saint Francis Xavier"
		saintImage = "https://upload.wikimedia.org/wikipedia/commons/thumb/d/db/Franciscus_de_Xabier.jpg/470px-Franciscus_de_Xabier.jpg"
	case 6:
		description = "Friday of the 1st Week, Immaculate Conception"
		saintImage = "https://en.wikipedia.org/wiki/John_of_Damascus#/media/File:John_Damascus_(arabic_icon).gif"
	case 7:
		description = "Saturday of the 1st Week, Juan Diego"
		saintImage = ""
	case 8:
		description = "Second Sunday in Advent"
		saintImage = ""
	case 9:
		description = "Saint Ambrose, consular prefect of Liguria and Emilia before becoming bishop of Milan by popular acclamation in 374"
		saintImage = "https://en.wikipedia.org/wiki/Ambrose#/media/File:AmbroseOfMilan.jpg"
	case 10:
		description = "Immaculate Conception"
		saintImage = "https://en.wikipedia.org/wiki/Mary_(mother_of_Jesus)#/media/File:Vladimirskaya.jpg"
	case 11:
		description = "Wednesday of the 2nd Week, Saint Juan Diego Cuauhtlatoatzin"
		saintImage = "https://en.wikipedia.org/wiki/Juan_Diego#/media/File:MIguel_Cabrera_-_Fiel_retrato_do_vener%C3%A1vel_Juan_Diego.jpg"
	case 12:
		description = "Thursday of the 2nd Week"
		saintImage = ""
	case 13:
		description = "Friday of the 2nd Week, Saint Damasus I, Pope"
		saintImage = "https://upload.wikimedia.org/wikipedia/commons/f/fa/Saintdamasus.jpg"
	case 14:
		description = "Saturday of the 2nd Week, Our Lady of Guadalupe"
		saintImage = "https://upload.wikimedia.org/wikipedia/commons/thumb/e/ee/1531_Nuestra_Se%C3%B1ora_de_Guadalupe_anagoria.jpg/275px-1531_Nuestra_Se%C3%B1ora_de_Guadalupe_anagoria.jpg"
	case 15:
		description = "Third Sunday of Advent"
		saintImage = ""
	case 16:
		description = "Saint John of the Cross, Priest, Doctor"
		saintImage = "https://en.wikipedia.org/wiki/John_of_the_Cross#/media/File:Zurbar%C3%A1n_St._John_of_the_Cross.jpg"
	case 17:
		description = "Tuesday of the 3rd Week"
		saintImage = ""
	case 18:
		description = "Wednesday of the 3rd Week"
		saintImage = ""
	case 19:
		description = "17th of December"
		saintImage = ""
	case 20:
		description = "18th of December"
		saintImage = ""
	case 21:
		description = "19th of December"
		saintImage = ""
	case 22:
		description = "4th Sunday of Advent"
		saintImage = ""
	case 23:
		description = "21st of Decemeber, Saint Peter Canisius"
		saintImage = "https://upload.wikimedia.org/wikipedia/commons/thumb/2/28/Saint_Petrus_Canisius.jpg/200px-Saint_Petrus_Canisius.jpg"
	case 24:
		description = "22nd of December"
		saintImage = ""
	case 25:
		description = "23rd of December, Saint John of Kęty"
		saintImage = "https://upload.wikimedia.org/wikipedia/commons/6/6d/Jan_Kanty.jpg"
	case 26:
		description = "24th of December, Christmas Eve"
		saintImage = ""
	case 27:
		description = "Christmas Day - Solenmity"
		saintImage = ""
	case 28:
		description = "Saint Stephen, the First Martyr"
		saintImage = "https://upload.wikimedia.org/wikipedia/commons/d/d3/St-stephen.jpg"
	case 29:
		description = "The Holy Family"
		saintImage = ""
	default:
		description = "It's not advent yet"
	}

	content := gin.H{"APIVersion": 2.0, "DayOfAdvent": adventDay, "Description": description, "SaintImage": saintImage}
	c.IndentedJSON(200, content)
}

func main() {

	start, _ := time.Parse(time.RFC822, "06 Dec 17 00:00 UTC")
	end, _ := time.Parse(time.RFC822, "24 Dec 17 23:59 UTC")
	today, _ := time.Parse(time.RFC822, "10 Nov 17 13:46 UTC")
	fmt.Println(start, end, today)
	app := gin.Default()

	if inTimeSpan(start, end, today) {

		app.GET("/api/aaas/v2/", index)

	} else {

		app.GET("/api/aaas/v2/", timeError)
	}

	app.Run(":8000")

}
