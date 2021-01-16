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
	board2 := duplicateBoard(size, board.Values)
	goal := findGoal(size)
/* 	board2 := utils.Grid {
		Values: board.Values,
		Cols: board.Cols,
		Rows: board.Rows,
	} */
	/* a := *board
	x := board2 */
	/* fmt.Println(&board)
	fmt.Println(&board2) */
	/* fmt.Println(a.Values)
	fmt.Println(x.Values)
	fmt.Println(board3.Values) */
	
	fmt.Println("//////////////////")
	fmt.Println(board.Values)
	fmt.Println(board2.Values)
	fmt.Println("----------------------")
	board2.SwapIndex(0,1)
	fmt.Println(board.Values)
	fmt.Println(board2.Values)
	
	if false {
		fmt.Println("Board :")
		board.Draw()
		fmt.Println("Goal :")
		goal.Draw()
/* 		aStar(board, goal) 
 */		fmt.Print("FS: ")
		fmt.Println(countUnsolved(board, goal))
	}
	
}

/* func aStar(start , goal *utils.Grid) {
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

		if currentBoard.Equals(goal) {
			// retoune constructed path
	}
	delete(openSet, currentBoard)

	//calcul des voisin/fils
	// for each voisin

	
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

func duplicateBoard(size int, pValues []int) *utils.Grid {
	board := utils.NewGrid(size,size)
	board.Values = pValues
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
		if board.GetValue(x) != goal.GetValue(x) {
			score += 1
		}
	}
	return score
}

/* func findSonStates(board *utils.Grid) []*utils.Grid {
	var states []*utils.Grid
	i := board.GetIndex(0)
	if i == 0 {
		son := utils.Grid{
			Values: board.Values,	
			Cols: board.Cols,
			Rows : board.Rows,
		}

	} else if i % this.cols != this.cols -1 {

	}else if i != 0 || i != len(this.values) || i % this.cols != this.cols -1 || i % this.rows != 1 {
		// 4 possibilité
	} else {

	}

	return states
}
 */
