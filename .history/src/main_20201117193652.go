package main

import (
	"fmt"
	"Thiiamas/Utils"
	"math"
	"math/rand"
	"time"
)

func main() {
	size := 3
	board := createBoard(size,"rand")
	goal := findGoal(size)

	board.Draw()
	fmt.Println("//////////////////")
	modified := make([]int, len(board.Values))
	ii := copy(modified, board.Values)
	fmt.Println(ii)
	fmt.Println(modified)
	index0 := board.GetIndex(0)
	modified[index0] = len(modified)
	sorted, count := quickSort(modified, 0, len(modified) - 1)
	fmt.Println(sorted)
	fmt.Println(count)
	board.Draw()
	/* way := aStar(board, goal)
	utils.Draw(way) */
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



func aStar(start , goal *utils.Grid) []*utils.Grid{
	/* openSet := make(map[*utils.Grid]bool)
	openSet[start] = true */
	openSet := make([]*utils.Grid,0)
	openSet = append(openSet, start)
	cameFrom := make(map[*utils.Grid]*utils.Grid)
	gScore := make(map[*utils.Grid]int)
	gScore[start] = 0
	fScore := make(map[*utils.Grid]int)
	fScore[start] = distance(start, goal)

	for len(openSet) != 0 {
		currentIndex := 0
		currentBoard := utils.NewGrid(1,1)
		currentScore := 500

		//Select the node with lowest value
		for i, b := range openSet {
			//calculate fscore, and lowest goes into current
			tempScore := distance(b, goal)
			if tempScore < currentScore  {
				currentIndex = i
				currentBoard = b
				currentScore = tempScore 
			}
		} 

		if currentBoard.Equals(goal) {
			return reconstrucPath(cameFrom, currentBoard)
		}
		openSet = remove(openSet, currentIndex)

		//calcul des voisin/fils
		neighbors := findSonStates(currentBoard)
		for _, neigh := range neighbors {
			tentative_score := gScore[currentBoard] + 1 // 1 bc each edge has a weight of 1
			countScore := 0
			for i := range gScore {
				if neigh.Equals(i){
					continue
				}else if countScore == len(gScore) - 1{
					gScore[neigh] = math.MaxInt32
				}
				countScore += 1
			}
			if tentative_score < gScore[neigh] {
				cameFrom[neigh] = currentBoard
				gScore[neigh] = tentative_score
				unsolved := distance(neigh, goal)
				fScore[neigh] = gScore[neigh] + unsolved
				fmt.Println(unsolved)
				if len(openSet) == 0 {
					openSet = append(openSet, neigh)
				}else {
					for i, b := range openSet{
						if neigh.Equals(b){
							continue
						}else if i == len(openSet) - 1 {
							openSet = append(openSet, neigh)
						}
					}
				}

			}
		}

	
	}
	fmt.Println("EPIC FAIL")
	return nil
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func remove(slice []*utils.Grid, s int) []*utils.Grid {
    return append(slice[:s], slice[s+1:]...)
}

func quickSort(slice []int, lo, hi int) ([]int, int) {
	count := 0
	if lo < hi {
		p, nCount := partition(slice, lo, hi)
		count += nCount
		quickSort(slice, lo, p - 1)
		quickSort(slice, p + 1, hi)
	}
	return slice, count
}

func partition(slice []int, lo, hi int) (int,int) {
	pivot := slice[hi]
	i := lo
	count := 0
	for j:= lo; j <= hi; j += 1 {
		if slice[j] <  pivot {
			v1 := slice[i]
			v2 := slice[j]
			slice[i] = v2
			slice[j] = v1
			i += 1
			count += 1
		}
	}
	v1 := slice[i]
	v2 := slice[hi]
	slice[i] = v2
	slice[hi] = v1
	count += 1
	return i, count
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
		case "test":
			board.Values = []int{8, 4, 2, 7, 0, 1, 6, 5, 3} 
	}
	
	return board
}

func reconstrucPath(cameFrom map[*utils.Grid]*utils.Grid, current *utils.Grid) []*utils.Grid{
	total_path := make([]*utils.Grid,0)
	grid, exist := cameFrom[current]
	for exist {
		total_path = append([]*utils.Grid{grid}, total_path...)
		grid, exist = cameFrom[grid]
	}
	return total_path
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

func duplicateSliceInt(slice []int) []int{
	duplicated := make([]int, len(slice))
	_ = copy(slice, duplicated)
	/* for i,k := range slice {
		duplicated[i] = k
	} */
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

func distance(board, goal *utils.Grid) int {
	score := 0
	for i,k := range board.Values {
		if k == i + 1 {
			continue
		}else{
			point := utils.IndexToPoint(i, board)
			final := utils.IndexToPoint(k - 1, board)
			xDist := final.X - point.X
			yDist := final.Y - point.Y
			distance := math.Abs(float64(xDist)) + math.Abs(float64(yDist))
			score += int(distance)
		}

	}
	return score
}

func count0ToFinal(board *utils.Grid) int {
	i := board.GetIndex(0)
	point := utils.IndexToPoint(i, board)
	final := utils.Point{X: board.Cols, Y:board.Rows}
	xDist := final.X - point.X
	yDist := final.Y - point.Y
	distance := math.Abs(float64(xDist)) + math.Abs(float64(yDist))
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

