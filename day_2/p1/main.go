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
    file, err := os.Open("test")
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
    re := regexp.MustCompile(`(do\(\))|(mul\(\d{1,3}\,\d{1,3}\))|(don\'t\(\))`)
    contentArr := re.FindAllString(string(content), -1)
    fmt.Println(contentArr)

    sum := 0
    for i := 0; i < len(contentArr) ; i++ {
        fmt.Println(i)
        value := contentArr[i]
        if value == "don't()" {
            // Start a loop that will look for the next mull
            for index, value := range contentArr[i + 1:] {
                if value == "do()" {
                    i += index + 1
                    break
                }
            }
            continue
        } else if value == "do()" {
            for index, value := range contentArr[i + 1:] {
                if value == "don't()" {
                    for i, v := range contentArr[i + 1:] {
                        if value == "do()" {
                            i += index + 1
                            break
                        }
            }
                    break
                }
                if strings.Contains(value, "mul") {
                    fmt.Println("Contains mul")
                    i += index + 1
                    fmt.Println(i)
                    fmt.Println(value)
                    break
                }
            }
        }
        value = contentArr[i]
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
        fmt.Println(sum)
        //fmt.Println(sum)
     }
}

func main () {
    get_muls()
}
