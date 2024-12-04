package main
import (
    "fmt"
)


func get_similarity_score(location_ids[][]int) int{
    var appearance_map = make(map[int]int)
    var sum int = 0

    length := len(location_ids)

    arr1 := make([]int, length)
    arr2 := make([]int, length)

    var arr1_map = make(map[int]int)
    var arr2_map = make(map[int]int)

    for i := 0; i < length; i++ {
        arr1[i] = location_ids[i][0]
        arr2[i] = location_ids[i][1]
    }

    //Lets do this naively first // I'll find a better solution later
    for i := 0; i < len(arr1); i++ {
        val := arr1[i]
        sum = 
        for j := 0; j < len(arr2) ; j++ {
            if arr2[j] == val {
                sum 
            }
        }
    }

    // Get number count in the first array
    // Get number count in the second array

    return sum

}

func main() {
    try := [][]int{
        {3, 4},
        {4, 3},
        {2, 5},
        {1, 3},
        {3, 9},
        {3, 3},
      }
    fmt.Println(get_similarity_score(try))

}
