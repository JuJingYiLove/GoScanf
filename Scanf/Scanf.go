package Scanf

import "fmt"

func ScanInt() int {
	var integer int
	_, err := fmt.Scan(&integer)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return integer
}

func ScanFloat32() float32 {
	var f32 float32
	_, err := fmt.Scan(&f32)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return f32
}

func ScanFloat64() float64 {
	var f64 float64
	_, err := fmt.Scan(&f64)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return f64
}

func ScanString() string {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return str
}
