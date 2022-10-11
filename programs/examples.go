package programs

import (
	"bufio"
	"log"
	"time"
	"timeplay/utils"
)

func RunExamples(reader *bufio.Reader) {
	locations := []string{
		"Europe/London",
		"Europe/Paris",
		"Europe/Vienna",
		"Europe/Brussels",
		"Europe/Prague",
		"Europe/Berlin",
		"Europe/Budapest",
		"Europe/Bucharest",
		"Europe/Luxembourg",
		"Europe/Amsterdam",
		"Europe/Warsaw",
		"Europe/Bratislava",
		"Europe/Sofia",
		"Europe/Madrid",
		"Europe/Lisbon",
		"Europe/Rome",
		"America/Chicago",
		"Asia/Kolkata",
	}
	for {
		log.Printf("Write date (YYYY-mm-dd):")
		input := utils.Read(reader)

		if input == "exit" {
			break
		}

		t, err := time.Parse("2006-01-02", input)
		if err != nil {
			log.Printf("Could not parse '%v'. Error: %v", input, err)
			continue
		}

		for _, l := range locations {

			location, err := time.LoadLocation(l)
			if err != nil {
				log.Printf("Could not load location with name %v", l)
				continue
			}

			timeInSelectedLocation := t.In(location)
			timeZone, offset := timeInSelectedLocation.Zone()
			offsetDuration := time.Duration(offset) * time.Second
			timeInSelectedLocation = t.Add(-offsetDuration).In(location)

			locationSpace := "\t"
			if len(location.String()) <= len("Europe/Budapest") {
				locationSpace = "\t\t"
			}

			log.Printf("Location: %v, %vTime: %v, TZ: %v", location, locationSpace, timeInSelectedLocation.Format("2006-01-02 15:04"), timeZone)
		}
	}
}
