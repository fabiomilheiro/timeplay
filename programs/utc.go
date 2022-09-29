package programs

import (
	"bufio"
	"log"
	"time"
	"timeplay/utils"
)

func RunUtc(reader *bufio.Reader) {
	for {
		log.Printf("Write time you want to calculate (nothing for today):")
		input := utils.Read(reader)

		if input == "exit" {
			break
		}

		if input == "" {
			input = time.Now().Format("2006-01-02")
		}

		t, err := time.Parse("2006-01-02", input)
		if err != nil {
			log.Println("Could not parse date", err)
			continue
		}

		utils.WriteTimeInLocation("Local", t)
		utils.WriteTimeInLocation("UTC", t)
		utils.WriteTimeInLocation("Europe/London", t)
		utils.WriteTimeInLocation("America/Chicago", t)
	}
}
