package utils

import (
	"os"
	"time"
	"fmt"
	"log"
)

func SaveMarketItems(marketItems []string) {
	date := time.Now().Local().Format("2006-01-02")
	filename := fmt.Sprintf("marketorder-%v.txt", date)

	f, err := os.OpenFile(filename, os.O_APPEND, 0666)
	if err != nil{
		log.Printf("Error opeing file for saving: %v", err)
	}

	defer f.Close()

	for _, order := range marketItems {
		_, err := f.WriteString(order)
		if err != nil {
			log.Printf("Error appending to file: %v", err)
		}
	}

	log.Printf("Saved %v market orders to disk.", len(marketItems))
}
