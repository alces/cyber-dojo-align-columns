package align

import (
    "strings"
)

func Align(text string) (result string) {
    data := split(text)
    sizes := maxWidth(data)
    
    for _, row := range data {
        result += alignedLine(row, sizes) + "\n"
    }
    
    return
}

func addSpaces(word string, width int) string {
    for i := len(word); i < width; i++ {
        word += " "
    }
    
    return word
}

func alignedLine(words []string, sizes []int) string {
    result := make([]string, len(words))
    lastWord := len(words) - 1

    for i := 0; i < lastWord; i++ {
        result[i] = addSpaces(words[i], sizes[i])
    }
    
    result[lastWord] = words[lastWord]
    
    return strings.Join(result, " ")
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
