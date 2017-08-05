package asserts

import (
	"log"
)

func NotNull(v interface{}, message string) {
	if v == nil {
		log.Fatal(message)
	}
}

func NoError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
