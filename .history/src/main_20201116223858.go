package main

import (
	"fmt"
	"Thiiamas/Utils"
	"math/rand"
	"time"
)

func main() {
	size := 5
	board := createBoard(size,"rand")
	goal := findGoal(size)

	board.Draw()
	fmt.Println("//////////////////")
	sons := findSonStates(board)
	utils.Draw(sons)
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

func reconstrucPath(cameFrom map[*utils.Grid]*utils.Grid, current *utils.Grid){
	step := 0
	total_path := make(map[*utils.Grid]int)
	for _, exist := cameFrom[current]{
		
	}
}

func aStar(start , goal *utils.Grid) {
	openSet := make(map[*utils.Grid]bool)
	openSet[start] = true
	cameFrom := make(map[*utils.Grid]*utils.Grid)
	gScore := make(map[*utils.Grid]int)
	gScore[start] = 0
	fScore := make(map[*utils.Grid]int)
	fScore[start] = countUnsolved(start, goal)

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
			return cameFrom
		}
		delete(openSet, currentBoard)

		//calcul des voisin/fils
		neighbors := findSonStates(currentBoard)
		for _, neigh := range neighbors {
			tentative_score := gScore[currentBoard] + 1 // 1 bc each edge has a weight of 1
			if tentative_score < gScore[neigh] {
				cameFrom[neigh] = currentBoard
				gScore[neigh] = tentative_score
				fScore[neigh] = gScore[neigh] + countUnsolved(neigh, goal)
				_, exist := openSet[neigh]
				if !exist {
					openSet[neigh] = true
				}
			}
		}

	
	}
	fmt.Println("EPIC FAIL")
}
func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}


// createBoard: Return a random generated board
func createBoard(size int, mode string) *utils.Grid {
	//construction de la liste des cases
	var numbers []int

	switch mode{
		case "rand":
			numbers = make([]int, size*size)
			for i := 0; i < size*size; i++ {
				numbers[i] = i
			}
		case "00":
			numbers = make([]int, size*size - 1)
			for i := 1; i < size*size; i++ {
				numbers[i - 1] = i
		}
	}
	board := utils.NewGrid(size,size)
	rand.Seed(time.Now().UnixNano())
	switch mode{
		case "rand":
			for x := 0; x < size*size; x++ {
			random := rand.Intn(len(numbers))
			board.SetIndex(x, numbers[random])
			numbers = removeIndex(numbers, random)
		}
		case "00":
			board.SetIndex(0,0)
			for x := 1; x < size*size; x++ {
				random := rand.Intn(len(numbers))
				board.SetIndex(x, numbers[random])
				numbers = removeIndex(numbers, random)
			}
	}
	
	return board
}



// duplicateBoard: Create a new Grid with the same values and independant
func duplicateBoard(board *utils.Grid) *utils.Grid {
	valuesTest := make([]int, board.Cols * board.Rows)
	for i,k := range board.Values {
		valuesTest[i] = k
	}
	duplicated := utils.NewGrid(board.Cols, board.Rows)
	duplicated.Values = valuesTest
	return duplicated
}

//findGoal: Return a Grid being the Taquin's game solution of a given size
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

//countUnsolved: Return the number of unsolved (= wrongly placed) squares
func countUnsolved(board, goal *utils.Grid) int{
	score := 0
	for x :=0; x < board.Len(); x++ {
		if board.GetValue(x) != goal.GetValue(x) {
			score += 1
		}
	}
	return score
}

// findSonStates: return an array of all possible board on the next step 
func findSonStates(board *utils.Grid) []*utils.Grid {
	var states []*utils.Grid
	var potentials []*utils.Grid
	i := board.GetIndex(0)
	p := utils.IndexToPoint(i, board)
	ichi := duplicateBoard(board)
	ichi.Swap(p, utils.Point{X:p.X, Y:p.Y-1})
	ini := duplicateBoard(board)
	ini.Swap(p, utils.Point{X:p.X+1, Y:p.Y})
	san := duplicateBoard(board)
	san.Swap(p, utils.Point{X:p.X, Y:p.Y+1})
	yon := duplicateBoard(board)
	yon.Swap(p, utils.Point{X:p.X-1, Y:p.Y})
	potentials = append(potentials, ichi, ini, san, yon)

	for _, b := range potentials {
		if !b.Equals(board){
			states = append(states, b)
		}
	}
	return states
}	

