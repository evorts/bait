package main

import (
	"banking-academy/testong"
	"log"
	"time"
)

func main() {
	// execute calculate function and print it out to screen
	// parameter indicate how deep it will iterate
	//section1.Fibonacci(10)
	//fmt.Println(section1.Fibonacci2(10))
	//fmt.Println(section2.TupleSquareCube(5))
	//section2.Array()
	//section2.Maps()
	//fmt.Printf("Binary search find: %v\n", section3.BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7))
	//section3.SelectionSort([]int{1,5,2,4,4,5,6,8,7,9,6,0,5})
	//section3.BubbleSort([]int{1,5,2,4,4,5,6,8,7,9,6,0,5}, func(i int, i2 int) bool {
	//	return i > i2
	//})
	testing.Hello()
	// Echo instance
	//e := echo.New()
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func bench(name string, f func(args ...interface{}), args ...interface{}) {
	start := time.Now()
	f(args)
	timeTrack(start, name)
}
