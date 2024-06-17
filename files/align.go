package align

import (
    "fmt"
    "strings"
)

func maxFields(data [][]string) (result int) {
    for i := 0; i < len(data); i++ {
        size := len(data[i])
        
        if size > result {
            result = size
        }
    }
    
    return
}

func maxWidth(data [][]string) (result []int) {
    for c := 0; c < maxFields(data); c++ {
        max := 0
        
        for r := 0; r < len(data); r++ {
            row := data[r]
            fmt.Printf("row = %v, word = %v, max = %v\n", row, row[c], max)
            
            if c < len(row) && len(row[c]) > max {
                max = len(row)
            }
        }
        
        result = append(result, max)
    }
    
    return
}

func split(text string) (result [][]string) {
    for _, l := range strings.Split(text, "\n") {
        result = append(result, strings.Split(l, "$"))
    }
    
    return
}
