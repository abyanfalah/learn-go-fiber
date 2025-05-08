package helper

import (
	"log"
	"strconv"
)

func ToInt(data string) int {
	r, e := strconv.Atoi(data)
	if e != nil {
		log.Fatalf("Unable to convert to Int: %v", e)
		panic("what the fuck")
	}

	return r
}

func ToString(data int) string {
	return strconv.Itoa(data)
}
