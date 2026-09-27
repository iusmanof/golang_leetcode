package main

import "fmt"

func main() {
	intMap := map[string]int64{"first": 10, "second": 23, "third": 45}
	floatsMap := map[string]float64{"first": 10.2323, "second": 23.234, "third": 45.5434}
	fmt.Println(sumInts(intMap))

	// generic
	fmt.Println(sumIntsOrFloats(intMap))
	fmt.Println(sumIntsOrFloats(floatsMap))

	// generic + interface
	fmt.Println(sumInterface(intMap))
	fmt.Println(sumInterface(floatsMap))
}

func sumInts(numbers map[string]int64) int64 {
	var sum int64
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func sumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

type Number interface {
	int64 | float64
}

func sumInterface[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
