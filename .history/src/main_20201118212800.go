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
	fmt.Print("permutation :")
	countPerm := countPerm(board)
	pCountPerm := countPerm % 2
	countZero := count0ToFinal(board)
	pCountZero := countZero % 2
	fmt.Println(countPerm)
	fmt.Print("0 to final :")
	fmt.Println(countZero)
	if pCountPerm != pCountZero {
		fmt.Println("Impossible")
		main()
		return
	}
	fmt.Println("//////////////////")
	/* fmt.Println(countLinearConflict(board, goal)) */
	way := aStarVariation(board, goal)
	utils.Draw(way)
	/* modified := make([]int, len(board.Values))
	copy(modified, board.Values)
	index0 := board.GetIndex(0)
	modified[index0] = len(modified)
	fmt.Print("permutation :")
	countPerm := countPerm(board)
	pCountPerm := countPerm % 2
	countZero := count0ToFinal(board)
	pCountZero := countZero % 2
	fmt.Println(countPerm)
	fmt.Print("0 to final :")
	fmt.Println(countZero)
	if pCountPerm != pCountZero {
		fmt.Println("Impossible")
		return
	}
	way := aStar(board, goal)
	utils.Draw(way)
	fmt.Print("permutation :")
	fmt.Println(countPerm)
	fmt.Print("Zero to final :")
	fmt.Println(countZero) */

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


func aStarVariation(start , goal *utils.Grid) []*utils.Grid{
	/* openSet := make(map[*utils.Grid]bool)
	openSet[start] = true */
	openSet := make([]*utils.Grid,0)
	closed := make([]*utils.Grid,0)
	openSet = append(openSet, start)
	cameFrom := make(map[*utils.Grid]*utils.Grid)
	gScore := make(map[*utils.Grid]int)
	gScore[start] = 0
	fScore := make(map[*utils.Grid]int)
	fScore[start] = distance(start, goal)

	nbr := 0
	for len(openSet) != 0 {
		currentIndex := 0
		currentBoard := utils.NewGrid(1,1)
		currentScore := 1000
		//Select the node with lowest value
		for i, b := range openSet {
			//calculate fscore, and lowest goes into current
			tempScore := distance(b, goal) + gScore[b]
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
		closed = append(closed, currentBoard)
		//calcul des voisin/fils
		neighbors := findSonStates(currentBoard)
		nbr += 1
		for _, neigh := range neighbors {
			//initialise gScore if needed (still usefull ?)
			contains, _ := utils.Contains(utils.GridKeys(gScore), neigh)
 			if !contains {
				gScore[neigh] = math.MaxInt32
			}
			in,_ := utils.Contains(openSet, neigh)
			in2,iClosed := utils.Contains(closed, neigh);
			if !(in && in2) {
				openSet = append(openSet, neigh)
				cameFrom[neigh] = currentBoard
				// g(s) = g(e) + C(e-->)
				gScore[neigh] = gScore[currentBoard] + 1
			} else {
				Gneigh := gScore[neigh] + distance(neigh, goal)
				Gcurrent := gScore[currentBoard] + 1 + distance(neigh, goal)
				if Gneigh > Gcurrent {
					gScore[neigh] = gScore[currentBoard] + 1
					cameFrom[neigh] = currentBoard
					fmt.Println("fin")
					if in,_ := utils.Contains(closed, neigh); in {
						fmt.Println("ici")
						closed = remove(closed, iClosed)
						openSet = append(openSet, neigh)
					}
				}
			}
		}

	
	}
	fmt.Println("Fail")
	return nil
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
			//initialise gScore if needed
			contains, _ := utils.Contains(utils.GridKeys(gScore), neigh)
 			if !contains {
				gScore[neigh] = math.MaxInt32
			}
			if tentative_score < gScore[neigh] {
				cameFrom[neigh] = currentBoard
				gScore[neigh] = tentative_score
				unsolved := distance(neigh, goal)
				fScore[neigh] = gScore[neigh] + unsolved
 				/* fmt.Println(unsolved) */
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
		if k ==0 {
			continue
			/* point := utils.IndexToPoint(i, board)
			final := utils.Point{X: board.Cols - 1, Y:board.Rows - 1}
			xDist := final.X - point.X
			yDist := final.Y - point.Y
			distance := math.Abs(float64(xDist)) + math.Abs(float64(yDist))
			score += int(distance) */
		}else if k == i + 1 {
			continue
		}else {
			point := utils.IndexToPoint(i, board)
			final := utils.IndexToPoint(k - 1, board)
			xDist := final.X - point.X
			yDist := final.Y - point.Y
			distance := math.Abs(float64(xDist)) + math.Abs(float64(yDist))
			score += int(distance)
		}

	}
	countL := countLinearConflict(board,goal)
	score = score + 2*countL
	return score
}

func countLinearConflict(board, goal *utils.Grid) int {
	conflict := make([][]int, 0)
	for i, k := range board.Values {
		if k == 0 {
			continue
		}
		point := utils.IndexToPoint(i, board)
		for i2, k2 := range board.Values {
			if k2 == 0 {
				continue
			}
			if k == k2 {
				continue
			}
			point2 := utils.IndexToPoint(i2, board)
			if point.X == point2.X || point.Y == point2.Y {
				goal1 := utils.IndexToPoint(goal.GetIndex(k), goal)
				goal2 := utils.IndexToPoint(goal.GetIndex(k2), goal)
				if goal1.X == goal2.X || goal1.Y == goal2.Y {
					if i == goal.GetIndex(k2) || i2 == goal.GetIndex(k){
						sol := []int{k, k2}
						for _, val := range conflict{
							if utils.Equal(val, sol) {
								continue
							}
						}
						conflict = append(conflict, sol)
					}
				}
			}

		}
	}
/* 	fmt.Println(conflict)
 */	return len(conflict)/2
}

func countPerm(board *utils.Grid) int{
	/* goal := make([]int,len(board.Values))
	copy(goal, board.Values) */
	count := 0
	goal := duplicateBoard(board)
	for i,k := range goal.Values {
		if k == 0 {
			if i == len(goal.Values) - 1{
				
			}else {
				iSol := goal.GetIndex(i+1)
				goal.SwapIndex(i, iSol)
				count += 1
			}
		}else if k != i + 1 {
			iSol := goal.GetIndex(i+1)
			goal.SwapIndex(i, iSol)
			count += 1
		} 
	}
	return count
}
func count0ToFinal(board *utils.Grid) int {
	i := board.GetIndex(0)
	point := utils.IndexToPoint(i, board)
	final := utils.Point{X: board.Cols - 1, Y:board.Rows - 1}
	xDist := final.X - point.X
	yDist := final.Y - point.Y
	distance := math.Abs(float64(xDist)) + math.Abs(float64(yDist))
	return int(distance)
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

