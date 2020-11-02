package main

import "fmt"

func main() {
    var num float64

    fmt.Println("Please, enter a floating point number:")

    n, err := fmt.Scan(&num)
    trunc := int(num)

    if err != nil {
        fmt.Println("Imput error")
    } else if n == 1 {
        fmt.Printf("Resulting integer is: %d\n", trunc)
    }
}
