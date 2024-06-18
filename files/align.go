package align

import (
    "strings"
)

func addSpaces(word string, width int) string {
    return ""
}

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
            
            if c < len(row) && len(row[c]) > max {
                max = len(row[c])
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
