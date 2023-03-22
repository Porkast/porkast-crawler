package jobs

import (
	"math/rand"
	"time"
)

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

func getRandomStartTime() (startTime time.Duration) {
	var (
		randomInt int
	)

	randomInt = randInt(5, 20)

	startTime = time.Second * time.Duration(randomInt)
	
	return
}