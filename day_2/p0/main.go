package main

import (
    "fmt"
    "os"
    "io"
    "regexp"
    "strings"
    "strconv"
)

func get_muls() {
    // read from the file
    file, err := os.Open("input")
    if err != nil {
        fmt.Println("Error: ", err)
        os.Exit(1)
    }
    defer file.Close()
    content, err := io.ReadAll(file)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    re := regexp.MustCompile(`mul\(\d{1,3}\,\d{1,3}\)`)
    contentArr := re.FindAllString(string(content), -1)
    sum := 0
    for _, value := range contentArr {
        start := strings.Index(value, "(")
        end := strings.Index(value, ")")

        numbers := value[start + 1: end]
        values := strings.Split(numbers, ",")
        firstNumber, err := strconv.Atoi(values[0])
        if err != nil {
            fmt.Println("Error converting the first number")
            return
        }
        secondNumber, err := strconv.Atoi(values[1])
        if err != nil {
            fmt.Println("Error converting the second number")
            return
        }

        sum += secondNumber * firstNumber
    }
    fmt.Println(sum)
}

func main () {
    get_muls()
}
