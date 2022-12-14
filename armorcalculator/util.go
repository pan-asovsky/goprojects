package main

import (
	"fmt"
	"strconv"
)

func IntToString(value int) string {
	return strconv.Itoa(value)
}

func fts32(value float32) string {

	result := fmt.Sprintf("%f", value)
	return result
}

func stf32(value string) float32 {

	temp, _ := strconv.ParseFloat(value, 32)
	return float32(temp)
}

func stf64(value string) float64 {

	temp, _ := strconv.ParseFloat(value, 64)
	return temp
}

func InputFloat32() float32 {

	var float float32
	_, _ = fmt.Scan(&float)
	return float
}

func InputInt() int {

	var result int
	_, _ = fmt.Scan(&result)
	return result
}

func ios() string {

	var str string
	_, _ = fmt.Scanf("%s", &str)
	return str
}
