package main

import (
	"math/rand"
	"time"

	"github.com/marprin/assessment/fetchapp/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cmd.Execute()
}
