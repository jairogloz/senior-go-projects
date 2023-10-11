package main

import (
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/chatgpt/codereview/math"
	"log"
	"time"
)

func main() {

	now := time.Now()
	result, err := math.Divide(100000000, 3)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("time: ", time.Since(now).Nanoseconds())
	log.Println("result: ", result)

}
