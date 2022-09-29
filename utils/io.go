package utils

import (
	"bufio"
	"log"
	"strings"
	"time"
)

func Read(reader *bufio.Reader) string {
	value, _ := reader.ReadString('\n')
	value = strings.ReplaceAll(value, "\r\n", "")
	return value
}

func WriteTimeInLocation(locationName string, t time.Time) {
	defer func() {
		recover()
	}()

	location, err := time.LoadLocation(locationName)
	if err != nil {
		log.Println("err: ", err.Error())
	}

	timeInLocation := t.In(location)

	timeZoneName, offset := timeInLocation.Zone()
	log.Printf("Location: %v, Time: %v, TZ: %v, Offset: %v", location, timeInLocation, timeZoneName, offset)
}
