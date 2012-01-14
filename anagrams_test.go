package main

import (
    "./anagrams"
    "time"
    "fmt"
)


const dict = "/usr/share/dict/words"

func main() {
    start := time.Nanoseconds()
    words := anagrams.Anagrams(dict)
    finish := time.Nanoseconds()
    diff := finish - start
    fmt.Println("Anagrams runtime: ", diff, " ns (", diff / int64(1000000), " ms)")
    for _, ana := range words {
        fmt.Println(ana)
    }
}
