package main

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"golang.org/x/example/hello/reverse"
)

type Message struct {
	Name string
	Body string
	Time time.Time
}

func main() {
	fmt.Println(reverse.String("Hello"))

	m := Message{Name: "Bod", Body: "The world is mine", Time: time.Now()}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Ошибка кодирования:", err)
		return
	}
	fmt.Println("Coded JSON", string(b))

	var decoded Message
	err = json.Unmarshal(b, &decoded)
	if err != nil {
		fmt.Println("Decode failed:", err)
		return
	}

	fmt.Println("Decoded Message:", decoded)

	// Generic JSON with any
	var varible any
	varible = 1
	varible = "world"
	varible = 2.5

	r := varible.(float64)
	fmt.Println("The circle's area", math.Pi*r*r)
}
