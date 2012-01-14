package anagrams

import (
    "io/ioutil"
    "sort"
    "strings"
)

type b []byte

func (b b) Len() int { return len(b) }
func (b b) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b b) Less(i, j int) bool { return b[i] < b[j] }

func ReadDictionary(file string) (words []string) {
    contents, _ := ioutil.ReadFile(file)
    words = strings.Split(string(contents), "\n")
    return
}

func Anagrams(file string) (anagrams map[string][]string) {
    anagrams = make(map[string][]string)
    words := UpperStrings(ReadDictionary(file))
    for _, word := range words {
        sortedWord := b([]byte(word))
        sort.Sort(sortedWord)
        stringWord := string([]byte(sortedWord))
        anagrams[stringWord]  = append(anagrams[stringWord], word)
    }
    
    return
}

func UpperStrings(words []string) (stringsUpper []string) {
    for _, word := range words {
        stringsUpper = append(stringsUpper, strings.ToUpper(word))
    }
    
    return
}
