package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
    content, err := ioutil.ReadFile("../input/input.txt")
	text := string(content)
	split := strings.Split(text, "\n")
    if err != nil {
        log.Fatal(err)
    }
	
	max := 0;
	sum:= 0;
	for i:=0; i<=len(split) -1; i++ {
		if split[i] == "" {
			if sum > max {
				max = sum
			}
			sum = 0;
			continue
		}
		parseNum, err := strconv.Atoi(split[i])
		if err != nil {
			log.Fatal(err)
		}
		sum += parseNum
	}

	fmt.Println("\nResult 1: ", max)
    
}