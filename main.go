package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"timeplay/programs"
	"timeplay/utils"
)

func main() {
	log.Println("Starting...")

	for {
		log.Println("\n\nProgram:")
		reader := bufio.NewReader(os.Stdin)
		program := utils.Read(reader)

		if program == "clear" {
			command := exec.Command("cmd", "/c", "cls")
			command.Stdout = os.Stdout
			if err := command.Run(); err != nil {
				log.Fatal("Could not clear", err)
			}
			continue
		}

		if program == "exit" {
			break
		}

		if program == "utc" {
			programs.RunUtc(reader)
		} else if program == "fromlocation" {
			programs.RunFromLocation(reader)
		} else if program == "examples" {
			programs.RunExamples(reader)
		} else {
			log.Println("could not find program")
		}
	}

	log.Println("Ending...")
}
