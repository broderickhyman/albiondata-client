package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

func SaveMarketItems(marketItems []string) {
	date := time.Now().Local().Format("2006-01-02")
	filename := fmt.Sprintf("marketorder-%v.txt", date)

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("Error opeing file for saving: %v", err)
		return
	}

	defer f.Close()

	for _, order := range marketItems {
		_, err := f.WriteString(fmt.Sprintf("%v\n", order))
		if err != nil {
			log.Printf("Error appending to file: %v", err)
			return
		}
	}

	log.Printf("Saved %v market orders to disk.", len(marketItems))
}
