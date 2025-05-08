package generator

import "time"

func GenerateId() uint {
	return uint(time.Now().Unix())
}
