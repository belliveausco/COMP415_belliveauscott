package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

//Scott Belliveau
func main() {
	var i string

	fmt.Println("What is the name of the filename you would like to open?")
	fmt.Scan(&i)
	fileAsBytes, err := ioutil.ReadFile(i)
	if err != nil {
		log.Fatalln("Error reading file", err)
	}
	fileContents := string(fileAsBytes)
	reg, err := regexp.Compile("[^a-zA-Z0-9]")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(fileContents, "\n")
	countWords(processedString)
}
func countWords(str1 string) {
	words := strings.Split(str1, "\n")
	countWords := make(map[string]int)
	for _, word := range words {
		if word != word {
			countWords[word] = 1
		} else {
			countWords[word] += 1
		}
	}
	reportResults(countWords)
}
func reportResults(map1 map[string]int) {
	m := map1

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, "has a count of", m[k])
	}

}
