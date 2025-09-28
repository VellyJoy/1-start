package main 

import "fmt"

func main () {
	const UsdToEur float64 = 0.85
	const UsdToRub float64 = 83.74
	var EurToRUb float64 
	EurToRUb = UsdToRub / UsdToEur

	fmt.Println("Курс евро в рубли:",EurToRUb)


	
}