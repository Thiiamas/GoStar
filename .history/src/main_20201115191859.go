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
	fmt.Println("similaire ?")
	fmt.Println(board.Equals(board))
	if false {
		fmt.Println("Board :")
		board.Draw()
		fmt.Println("Goal :")
		goal.Draw()
/* 		aStar(board, goal) 
 */		fmt.Print("FS: ")
		fmt.Println(computeFs(board, goal))
	}
	
}

func aStar(start , goal *utils.Grid) {
	openSet := make(map[*utils.Grid]bool)
	openSet[start] = true
	cameFrom := make(map[*utils.Grid]bool)
	gScore := make(map[*utils.Grid]int)
	fScore := make(map[*utils.Grid]int)

	for len(openSet) != 0 {
		currentBoard := utils.NewGrid(1,1)
		currentScore := 500
		//Select the node with lowest value
		for b := range openSet {
			//calculate fscore, and lowest goes into current
			tempScore := countUnsolved(b, goal)
			if tempScore < currentScore  {
				currentBoard = b
				currentScore = tempScore 
			}
		} 
		if board.Equals(goal) {
			// fin
		}
	}
	
}
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

func countUnsolved(board, goal *utils.Grid) int{
	score := 0
	for x :=0; x < board.Len(); x++ {
		if board.GetIndex(x) != goal.GetIndex(x) {
			score += 1
		}
	}
	return score
}

