package main

import (
    "fmt"
    "math"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func get_safe() {
    var reports = make([][]int, 0)
    file, err := os.Open("input")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        values := strings.Split(line, " ")
        num_values := make([]int, len(values))
        for i, str := range values {
            num, err := strconv.Atoi(str)
            if err != nil {
                fmt.Printf("Error")
                return
            }
            num_values[i] = num
        }
        reports = append(reports, num_values)

    }

    safe_count := 0
    for _, arr := range reports {
        isSafe := true
        increasing := false
        decreasing := false

        for i := 0; i < len(arr) - 1; i++ {
            difference := (float64(arr[i]) - float64(arr[i + 1]))
            absDiff := math.Abs(difference)
            if absDiff < 1 || absDiff > 3 {
                isSafe = false
                break
            }
            if difference > 0 {
                decreasing = true
            } else {
                increasing = true
            }
            if decreasing && increasing {
                isSafe = false
                break
            }
        }

        if isSafe {
            safe_count ++
        }
    }
    fmt.Println(safe_count)
}

func main () {
    get_safe()
}

arr[:len(arr) -1 ]
