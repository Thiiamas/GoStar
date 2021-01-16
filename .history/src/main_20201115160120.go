package main

import (
	"fmt"
	"Thiiamas/Utils"
	"math/rand"
	"time"
)

func main() {
	size := 5
	board := createBoard(size)
	goal := findGoal(size)
	aStar(board, goal)
}

/* func aStar(start , goal *utils.Grid) {
	openSet := make(map[*utils.Grid]bool)
	openSet[start] = true
	cameFrom := make(map[*utils.Grid]bool)
	gScore := make(map[*utils.Grid]int)
	fScore := make(map[*utils.Grid]int)

	for k := range openSet {
		//calculate fscore, and lowest goes into current
		current := 
	}
} */
func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func createBoard(size int) *utils.Grid {
	//construction de la liste des cases

	numbers := make([]int, size*size)
	board := utils.NewGrid(size,size)
	for i := 0; i < size*size; i++ {
		numbers[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	for x := 0; x < size*size; x++ {
		random := rand.Intn(len(numbers))
		board.SetIndex(x, numbers[random])
		numbers = removeIndex(numbers, random)
	}
	return board
}

func findGoal(size int) *utils.Grid {
	numbers := make([]int, size*size)
	board := utils.NewGrid(size,size)
	for i := 0; i < size*size; i++ {
		numbers[i] = i
	}
	for x := 0; x < (size*size) - 1; x++ {
		board.SetIndex(x,x+1)
	}
	board.SetIndex(size*size - 1, 0)
	return board
}

func computeFs(board, goal *utils.Grid) int{
	bValues := board.Values
	gValues := goal.Values
	score := 0
	for x :=0; x < board.Len(); x++ {
		if bValues[x] == gValues[x] {
			score += 1
		}
	}
	return score
}
