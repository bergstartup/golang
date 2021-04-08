//using maps, func and for_loop

package main

import (
  "fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	asWords:=strings.Split(s," ")
	wordCount:=make(map[string]int)
	for _,j:= range asWords{
		wordCount[j]+=1
	}
	return wordCount
}

func main() {
  fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
}
