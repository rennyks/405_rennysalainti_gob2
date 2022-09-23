package main

import (
	"fmt"
	"math"
)

func main()  {
	var diameter float64 = 15
	var  area, circumferense = calculate(diameter)

	fmt.Println("area :", area)
	fmt.Println("Circumferense :", circumferense)
}
// returning multiple value
// func  calculate(d float64) (float64, float64)  {
// 	//menghitung luas
// 	var area float64 = math.Pi * math.Pow(d/2, 2)
// 	// menghitung keliling

// 	var circumferense = math.Pi * d
// 	return area, circumferense
// }

//predefined return value
func calculate(d float64) (area float64, circumferense float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumferense = math.Pi * d
	return
	
}