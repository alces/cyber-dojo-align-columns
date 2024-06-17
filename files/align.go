package align

import (
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
    return
}

func split(text string) (result [][]string) {
    for _, l := range strings.Split(text, "\n") {
        result = append(result, strings.Split(l, "$"))
    }
    
    return
}
