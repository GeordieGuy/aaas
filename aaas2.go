package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func timeError(c *gin.Context) {

	content := gin.H{"APIVersion": 2.0, "DayOfAdvent": "Error", "Description": "It is not advent anymore, or not yet"}
	c.IndentedJSON(451, content)
}

func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
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
		saintImage = "https://commons.wikimedia.org/wiki/File:Artus_Wolffort_-_St_Andrew_-_WGA25857.jpg"
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
		description = "Second Monday of Advent"
		saintImage = ""
		reading = "Isaiah 35: 1–10:	God himself is coming to save you.."
		psalm = "Ps 84:9–14 r. Is 35:4:	Look, our God is coming to save us."
		acclamation = "Luke 3:4. 6: Prepare a way for the Lord, make his paths straight. And all mankind shall see the salvation of God."
		gospel = "Luke 5:17–26: We have seen strange things today."
	case 10:
		description = "First Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 11:
		description = "Second Tuesday of Advent"
		saintImage = ""
		reading = "Isaiah 40:1–11: God consoles his people.."
		psalm = "Ps 95: 1-3. 10–13 r. cf. Is 40:9–10: Here is our God coming with power."
		acclamation = "Come Lord! Do not delay. Forgive the sins of your people."
		gospel = "Matthew 18:12–14: God does not wish the little ones to be lost."
	case 12:
		description = "Second Wednesday of Advent"
		saintImage = ""
		reading = "Isaiah 40:25–31: The Lord almighty gives strength to the wearied."
		psalm = "Ps 102: 1–4 8. 10 r.1: My soul, give thanks to the Lord."
		acclamation = "Behold, our Lord will come with power and will enlighten the eyes of his servants."
		gospel = "Matthew 11:28–30: Come to me, all you who labour."
	case 13:
		description = "Second Thursday of Advent"
		saintImage = ""
		reading = "Isaiah 41:13–20: I, the Holy One of Israel, am your redeemer."
		psalm = "Ps 144:1, 9–13 r. 8: The Lord is kind and full of compassion, slow to anger,abounding in love."
		acclamation = "	Come to us, Lord, with your peace that we may rejoice in your presence with sincerity of heart."
		gospel = "Matthew 11:11–15: A greater than John the Baptist has never been seen."
	case 14:
		description = "Second Friday of Advent"
		saintImage = ""
		reading = "Isaiah 48:17–19: If only you had been alert to my commandments."
		psalm = "Ps 1: 1–4. 6 r. cf. Jn 8:12: Anyone who follows you, Lord, will have the light of life."
		acclamation = "See, the king, the Lord of the world, will come. He will free us from the yoke of our bondage."
		gospel = "Matthew 11:16–19	They heed neither John nor the Son of Man."
	case 15:
		description = "Second Saturday of Advent"
		saintImage = ""
		reading = "Ecclesiasticus 48: 1–4.9–11: Elijah will come again."
		psalm = "Ps 79:2–3. 15–16. 18–19 r. 4: God of hosts, bring us back; let your face shine on us and we shall be saved."
		acclamation = "The day of the Lord is near; Look he comes to save us."
		gospel = "Matthew 17:10–13: Elijah has come already and they did not recognise him"
	case 16:
		description = "Second Sunday of Advent"
		saintImage = ""
		reading = "Isaiah 11:110: On him the spirit of the Lord rests."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:21-24: Jesus is filled with joy by the Holy Spirit."
	case 17:
		description = "Third Monday of Advent"
		saintImage = ""
		reading = "Numbers 24:2–7. 15–17: A star from Jacob takes the leadership."
		psalm = "Ps 71: 12. 78. 1213. 17 r. 7: In his days justice shall flourish and peace till the moon fails."
		acclamation = "Ps 84:8:	Let us see, O Lord, your mercy and give us your saving help."
		gospel = "Luke 10:2124: Jesus is filled with joy by the Holy Spirit."
	case 18:
		description = "Third Tuesday of Advent"
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

	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	Info.Println("Starting aaas v2 service")

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

	app.Run(":80")

}
