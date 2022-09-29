package programs

import (
	"bufio"
	"log"
	"strings"
	"time"
	"timeplay/utils"
)

func RunFromLocation(reader *bufio.Reader) {
	for {
		log.Printf("Write date and location (YYYY-mm-dd,<location>):")
		input := utils.Read(reader)

		if input == "exit" {
			break
		}

		values := strings.Split(input, ",")

		if len(values) != 2 {
			log.Printf("'%v' must have a sinlge comma", input)
			continue
		}

		timeToParse := values[0]
		locationName := values[1]

		t, err := time.Parse("2006-01-02", timeToParse)
		if err != nil {
			log.Printf("Could not parse '%v'. Error: %v", timeToParse, err)
			continue
		}

		location, err := time.LoadLocation(locationName)
		if err != nil {
			log.Printf("Could not load location with name %v", locationName)
			continue
		}

		timeInSelectedLocation := t.In(location)
		timeZone, offset := timeInSelectedLocation.Zone()
		offsetDuration := time.Duration(offset) * time.Second
		timeInSelectedLocation = t.Add(-offsetDuration).In(location)

		log.Printf("Location: %v, Time: %v, TZ: %v, Offset: %v", location, timeInSelectedLocation, timeZone, offset)
		utils.WriteTimeInLocation("UTC", timeInSelectedLocation)
		utils.WriteTimeInLocation("Local", timeInSelectedLocation)
	}
}
