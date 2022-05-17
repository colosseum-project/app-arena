package utils

import (
	"math/rand"
	"time"
)

type Rand struct{}

func (Rand) InitSeed() {
	rand.Seed(time.Now().UnixNano())
}

func (Rand) Range(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func (Rand) OneinXChance(x int) bool {
	return rand.Intn(x) == 0
}

func (Rand) XPercentChance(x int) bool {
	return rand.Intn(101) <= x
}
