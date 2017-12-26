package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"regexp"
	"strconv"
	"time"
)

type Remind struct {
	gorm.Model
	UserID  string
	Timer   int64
	Message string
}

var db *gorm.DB

func parseMessageFromUser(message string) (int64, string, string) {

	remindTimeRegex := regexp.MustCompile("^[Ч|ч]ерез ([[:digit:]]) (часов|минут)")
	result := remindTimeRegex.FindStringSubmatch(message)
	timeCount, _ := strconv.ParseInt(result[1], 10, 64)
	timeSize := result[2]

	remindMessageRegex := regexp.MustCompile("напомни (.+)$")
	result = remindMessageRegex.FindStringSubmatch(message)
	remindMessage := result[1]

	return timeCount, timeSize, remindMessage

}
	// Seconds in Hour
	var multiplyCounter int64
	multiplyCounter = 3600

	var until int64

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Remind{})

	id := "Some ID"
	message := "Через 5 часов напомни принять визин"
	c, _, m := parseMessageFromUser(message)
	timer := c*multiplyCounter + time.Now().Unix()

	db.Create(&Remind{UserID: id, Timer: timer, Message: m})

	until = 10000000000000000

	var reminds []Remind
	db.Where("Timer < ?", until).Find(&reminds)
	for _, rem := range reminds {
		fmt.Println(rem)
	}
}
