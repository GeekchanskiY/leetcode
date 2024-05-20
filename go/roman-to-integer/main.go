package main

import(
    "fmt"
)
var symbolValue = map[string]int{
    "I": 1,
    "V": 5,
    "X": 10,
    "L": 50,
    "C": 100,
    "D": 500,
    "M": 1000,
}

func romanToInt(s string) int {
    
    var l int = len(s)
    var r string
    recent_val := 0
    var current_val int
    res := 0
    for i := 0; i < l; i++ {
        r = string(rune(s[i]))
        current_val = symbolValue[r]
        if current_val > recent_val && recent_val != 0 {
            res += (current_val - recent_val)
            recent_val = 0
        } else {
            res += recent_val
            recent_val = current_val
        }
        
        
    }
    if recent_val != 0{
        res += recent_val
    }
    return res
}

func main(){
    fmt.Println(romanToInt("III"))
    fmt.Println(romanToInt("XIV"))
    fmt.Println(romanToInt("XXXIII"))


}
