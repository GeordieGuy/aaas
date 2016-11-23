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
	reading := ""
	psalm := ""
	acclamation := ""
	gospel := ""

	adventDay = 4 // Debug

	switch adventDay {
	case 1:
		description = ""
		saintImage = ""
		reading = ""
		psalm = ""
		acclamation = ""
		gospel = ""
	case 2:
		description = "First Monday of Advent"
		saintImage = ""
		reading = "Isaiah 2:15: The Lord draws all the nations together into the eternal peace of God's kingdom."
		psalm = "Ps 121: 12. 45. 69 r. 1: I rejoiced when I heard them say:'Let us go to God's house.'"
		acclamation = "Ps 79:4: God of hosts, bring us back: let your face shine on us and we shall be saved."
		gospel = "Matthew 8:511: Many will come from east and west to take their places in the kingdom of heaven."
	case 3:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:1-10: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8: Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 4:
		description = "First Wednesday of Advent, St Andrew, Apostle"
		saintImage = "https://commons.wikimedia.org/wiki/File:Artus_Wolffort_-_St_Andrew_-_WGA25857.jpg#"
		reading = "Isaiah 25:6-10: The Lord invites us to his banquet and wipes away the tears from every cheek."
		psalm = "Ps 22 r. 6: In the Lord's own house shall I dwell for ever and ever."
		acclamation = "Is 33:22: The Lord is our judge, the Lord is our lawgiver, the Lord our king and our saviour."
		gospel = "Matthew 15:29-37: Jesus cures many and multiplies the loaves."
	case 5:
		description = "First Thursday of Advent"
		saintImage = ""
		reading = "Isaiah 26:16: Let the upright nation come in, she, the faithful one."
		psalm = "Ps 117: 1. 89. 1921. 2527 r. 26: Blessed in the name of the Lord is he who comes. or Alleluia!."
		acclamation = "Is 40: 910:	Shout with a loud voice, joyful messenger to Jerusalem. Here is the Lord God coming with power."
		gospel = "Matthew 7:21. 2427: The person who does the will of my Father will enter the kingdom of heaven."
	case 6:
		description = "First Friday of Advent"
		saintImage = ""
		reading = "Isaiah 29:17-24: That days the eyes of the blind will see.."
		psalm = "Ps 26: 1. 4. 13-14 r. 1: The Lord is my light and my help.."
		acclamation = "cf. Is 45:8: Send victory like a dew, you heavens, and let the clouds rain it down. Let the earth open and bring forth the saviour."
		gospel = "Matthew 9:27-31 Two blind men who believe in Jesus are cured."
	case 7:
		description = "First Saturday of Advent"
		saintImage = ""
		reading = "Isaiah 11:1-10: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8 Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 8:
		description = "First Sunday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 9:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 10:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 11:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 12:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 13:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 14:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
		description = "Third Sunday of Advent"
		saintImage = ""
	case 16:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 17:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 18:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 19:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 20:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 21:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 22:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 23:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 24:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 25:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 26:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 27:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 28:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 29:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	default:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	}

	content := gin.H{"APIVersion": 2.0, "DayOfAdvent": adventDay, "Reading": reading, "Description": description, "Responsorial": psalm, "SaintImage": saintImage, "Collect": acclamation, "Gospel": gospel}
	c.IndentedJSON(200, content)
}

func main() {

	start, _ := time.Parse(time.RFC822, "06 Dec 16 00:00 UTC")
	end, _ := time.Parse(time.RFC822, "24 Dec 16 23:59 UTC")
	today, _ := time.Parse(time.RFC822, "06 Dec 16 13:46 UTC")
	fmt.Println(start, end, today)
	app := gin.Default()

	if inTimeSpan(start, end, today) {

		app.GET("/api/aaas/v2/", index)

	} else {

		app.GET("/api/aaas/v2/", timeError)
	}

	app.Run(":8000")

}
