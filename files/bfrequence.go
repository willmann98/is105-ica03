package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

const path = "C:/Users/willi/go/src/github.com/Willmann98/is105-ica03/files/pg100.txt"

func main() {

	fmt.Println(os.Args)

	lines := splitFile(bufio.ScanLines)
	fmt.Println(len(lines))

	runes := splitFile(bufio.ScanRunes)
	fmt.Println(len(runes))

	printMostFrequentRunes(runes)

}

func splitFile(split func(data []byte, atEOF bool) (advance int, token []byte, err error)) []string {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(split)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func printMostFrequentRunes(runes []string) {

	runeCount := make(map[string]int)

	for _, rune := range runes {

		runeCounter := runeCount[rune]
		runeCount[rune] = runeCounter + 1

	}

	type kv struct {
		Key   string
		Value int
	}

	var runeTeller []kv
	for k, v := range runeCount {
		runeTeller = append(runeTeller, kv{k, v})
	}

	sort.Slice(runeTeller, func(i, j int) bool {
		return runeTeller[i].Value > runeTeller[j].Value
	})

	for i := 0; i <=4; i++ {
		fmt.Printf("%s, %d\n", runeTeller[i].Key, runeTeller[i].Value )	
	}
}
