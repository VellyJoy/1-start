package main 

import "fmt"

func main () {
	eur, rub := InputCurrency()
	var EurToRUb float64
	EurToRUb = eur / rub

	fmt.Println("Курс евро в рубли:",EurToRUb)
}

func InputCurrency () (float64,float64) {
var UsdToEur float64 
fmt.Println("Введите курс доллара к евро:")
fmt.Scan(&UsdToEur)
var UsdToRub float64 
fmt.Println("Введите курс доллара к рублю:")
fmt.Scan(&UsdToRub)
return UsdToEur,UsdToRub
}

func calculation (number int,OwnCurrency string,TartetCurrency string) ()