package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// filepath variable
var filepath string

// define a struct for each proverb in the string
type proverb struct {
	line string
	chars map[rune]int
}
// add a method to loop through the string and count the number of characters
func (p *proverb) countChars() int {

	for i := 0; i < len(p.line); i++ {
		chr := rune(p.line[i])
		_, chrIdx := p.chars[chr]

		if chrIdx {
			p.chars[chr] = p.chars[chr] + 1
		} else {
			p.chars[chr] = 1
		}
	}
	
	return len(p.chars)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// function returns a pointer to a new proverb type
func newProverb(s string) *proverb {
	p := proverb{line: s, chars: map[rune]int{}}
	p.countChars()
	return &p
}

// function returns a slice of pointer to a proverb and an error value
func loadProverbs(filepath string) ([]*proverb, error) {
	dat, err := ioutil.ReadFile(filepath)

	check(err)

	proverbTexts := string(dat)
	proverbs := strings.Split(proverbTexts, "\n")

	proverbSlice := make([]*proverb, len(proverbs))

	for i := 0; i < len(proverbs); i++ {
		proverbSlice[i] = newProverb(proverbs[i])
	}

	return proverbSlice, err
}

func main() {
	fileFlag := flag.String("f", "", "command line file flag")
	flag.Parse()

	envFilePath := os.Getenv("FILE")

	if *fileFlag != "" && len(os.Args) > 0 {
		filepath = *fileFlag
	} else if envFilePath != "" {
		filepath = envFilePath
	} else {
		check(fmt.Errorf("no filepath assigned"))
	}

	quotes, err := loadProverbs(filepath)
	check(err)
	
	for i := 0; i < len(quotes); i++ {
		var currNum int = i + 1
		p := quotes[i]
		var numWords int = len(strings.Split(p.line, " "))

		fmt.Printf("%d. %s (WC: %d)\n", currNum, p.line, numWords)

		for k, v := range p.chars {
			fmt.Printf("'%s'=%d, ", string(k), v)
		}
		fmt.Printf("\n\n")
	}

}
