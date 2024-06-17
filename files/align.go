package align

import (
    "strings"
)

func maxWidth(data [][]string) (result []int) {
    return
}

func split(text string) (result [][]string) {
    for _, l := range strings.Split(text, "\n") {
        result = append(result, strings.Split(l, "$"))
    }
    
    return
}
