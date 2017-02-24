package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"
const transformsFile = "sprinkle/transforms/transformations.txt"

var transforms []string

func main() {
	file, err := os.Open(transformsFile)
	if err != nil {
		fmt.Printf("Twang: %s\n", err.Error())
		os.Exit(0)
	}

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		t := fileScanner.Text()
		transforms = append(transforms, t)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	stdScan := bufio.NewScanner(os.Stdin)
	for stdScan.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, stdScan.Text(), -1))
	}
}
