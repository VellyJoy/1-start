package main

import (
    "errors"
    "fmt"
    "strings"
)

func main() {
    rate, amount := InputCurrency()
    fmt.Println("Ваш результат:", rate*amount)
}

func InputCurrency() (float64, float64) {
    var choice, target string
    var amount float64

    for {
        fmt.Println("Выберите исходную валюту (RUB/EUR/USD):")
        fmt.Scan(&choice)
        choice = strings.ToUpper(choice)

        if err := validateCurrency(choice); err != nil {
            fmt.Println("Ошибка:", err)
            continue
        }
        break
    }

    for {
        fmt.Println("Введите количество:")
        fmt.Scan(&amount)

        if err := validateAmount(amount); err != nil {
            fmt.Println("Ошибка:", err)
            continue
        }
        break
    }

    for {
        fmt.Println("Выберите целевую валюту:")
        switch choice {
        case "RUB":
            fmt.Println("(EUR/USD):")
        case "EUR":
            fmt.Println("(RUB/USD):")
        case "USD":
            fmt.Println("(RUB/EUR):")
        }

        fmt.Scan(&target)
        target = strings.ToUpper(target)

        if err := validateTarget(choice, target); err != nil {
            fmt.Println("Ошибка:", err)
            continue
        }
        break
    }

    switch choice {
    case "RUB":
        if target == "EUR" {
            return 0.011, amount
        }
        return 0.012, amount
    case "EUR":
        if target == "RUB" {
            return 93.83, amount
        }
        return 1.16, amount
    case "USD":
        if target == "RUB" {
            return 81.16, amount
        }
        return 0.87, amount
    }
    return 0, 0
}

func validateCurrency(c string) error {
    if c != "RUB" && c != "USD" && c != "EUR" {
        return errors.New("неизвестная валюта, используйте RUB, USD или EUR")
    }
    return nil
}

func validateAmount(a float64) error {
    if a <= 0 {
        return errors.New("количество должно быть положительным числом")
    }
    return nil
}

func validateTarget(base, target string) error {
    switch base {
    case "RUB":
        if target != "EUR" && target != "USD" {
            return errors.New("из RUB можно только в EUR или USD")
        }
    case "EUR":
        if target != "RUB" && target != "USD" {
            return errors.New("из EUR можно только в RUB или USD")
        }
    case "USD":
        if target != "RUB" && target != "EUR" {
            return errors.New("из USD можно только в RUB или EUR")
        }
    }
    return nil
}




