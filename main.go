package main 

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)
	
func main() {
var operation string 
var input string 

fmt.Println ("Введите тип операции (SUM,AVG,MED):")
fmt.Scanln(&operation)

fmt.Println ("Введите числа через запятую:")
fmt.Scanln (&input)

parts := strings.Split (input, ",")

var numbers []float64
for _, p := range parts {
	num , err := strconv.ParseFloat(strings.TrimSpace(p), 64)
	if err != nil {
		fmt.Println ("Введено неверное число:", p)
		return
	}
	numbers = append(numbers, num)
}

switch strings.ToUpper(operation) {
case "SUM":
	sum := 0.0 
	for _, n := range numbers {
		sum += n 
	}
	fmt.Println ("Сумма:", sum)

case "AVG":
	sum := 0.0 
	for _, n := range numbers {
		sum += n 
	}
	fmt.Println ("Среднее число:", sum / float64(len(numbers)))

case "MED": 
sort.Float64s(numbers)
n := len(numbers)
if n%2 == 0 {
	fmt.Println ("Медиана равна:", (numbers[n/2-1]+numbers[n/2])/2)
} else {
	fmt.Println("Медиана равна:", numbers[n/2])
}
default:
	fmt.Println ("Неизвестная операция Используйте операции из списка.")
}
}



