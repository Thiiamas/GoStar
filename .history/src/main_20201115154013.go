package main

import (
	"fmt"
	"Thiiamas/Utils"
	"math/rand"
	"time"
)

func main() {
	board := createBoard(3)
	board.Draw()
	fmt.Println("fin")

}

func aStar(){

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
	board.setIndex(size*size - 1, 0)
	return board
}

