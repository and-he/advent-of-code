package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func get_calibration(line string) (first int, last int) {
	left, right := 0, len(line)-1
	for left < len(line) {
		if '0' <= line[left] && line[left] <= '9' {
			first_conv, err := strconv.Atoi(string(line[left]))
			if err != nil {
				log.Fatalf("error converting %s at index %d to int", line, left)
			}
			first = first_conv
			break
		} else {
			left += 1
		}
	}
	for right >= 0 {
		if '0' <= line[right] && line[right] <= '9' {
			last_conv, err := strconv.Atoi(string(line[right]))
			if err != nil {
				log.Fatalf("error converting %s at index %d to int", line, right)
			}
			last = last_conv
			break
		} else {
			right -= 1
		}
	}
	return
}

// use regex to get the index of both occurrences:
// word that contains either "one", "twp", "three", etc.
// digit
// whichever index is least, that will be returned as "first"
// whichever index is most, that will be returned as "last"
func get_calibration_2(line string) (first int, last int) {

}

func main() {
	file, err := os.Open("day1_input.txt")
	if err != nil {
		log.Fatalf("error reading input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		first, last := get_calibration(scanner.Text())
		calibration := (first * 10) + last
		sum += calibration
		fmt.Printf("%s ---> len=%d ---> first=%d, last=%d, calibration=%d\n", scanner.Text(), len(scanner.Text()), first, last, calibration)
	}
	fmt.Printf("sum is %d", sum)
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error in scanner: %v", err)
	}
}
