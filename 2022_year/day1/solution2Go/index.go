package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
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
	var arrayTotalCalories []int;
	sum:= 0;
	for i:=0; i<=len(split) -1; i++ {
		if split[i] == "" {
			arrayTotalCalories = append(arrayTotalCalories, sum)
			sum = 0;
			continue
		}
		parseNum, err := strconv.Atoi(split[i])
		if err != nil {
			log.Fatal(err)
		}
		sum += parseNum
	}

	sort.Slice(arrayTotalCalories, func(i, j int) bool {
		return arrayTotalCalories[i] > arrayTotalCalories[j]
	})
	 
	firstThreeElements := arrayTotalCalories[:3]
	res := 0

	for i:=0; i<len(firstThreeElements); i++ {
		res += firstThreeElements[i]
	 }

	fmt.Println("\nResult 1: ", res)
    
}