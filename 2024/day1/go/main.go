package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func parseIds(line string) (int, int, error) {
	split_line_contents := strings.Split(line, "   ")
	id1, err := strconv.Atoi(split_line_contents[0])
	if err != nil {
		return -1, -1, err
	}
	id2, err := strconv.Atoi(split_line_contents[1])
	if err != nil {
		return -1, -1, err
	}
	return id1, id2, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error reading input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	left_ids := &IntHeap{}
	right_ids := &IntHeap{}
	heap.Init(left_ids)
	heap.Init(right_ids)
	for scanner.Scan() {
		id_left, id_right, err := parseIds(scanner.Text())
		if err != nil {
			log.Fatalf("error parsing line [%s]: %v", scanner.Text(), err)
		}
		heap.Push(left_ids, id_left)
		heap.Push(right_ids, id_right)
		// fmt.Printf("parsed->%d -- %d\n", id_left, id_right)
	}
	// for left_ids.Len() > 0 {
	// 	fmt.Printf("%d ", heap.Pop(left_ids))
	// }
	// fmt.Println("\n------")
	// for right_ids.Len() > 0 {
	// 	fmt.Printf("%d ", heap.Pop(right_ids))
	// }
	sum := 0
	for (left_ids.Len() > 0) && (right_ids.Len() > 0) {
		current_left := heap.Pop(left_ids).(int)
		current_right := heap.Pop(right_ids).(int)
		diff := current_left - current_right
		if diff < 0 {
			diff = -diff
		}
		sum += diff
		fmt.Printf("%d <-> %d --> %d\n", current_left, current_right, diff)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error in scanner: %v", err)
	}
	fmt.Printf("sum = %d\n", sum)
}
