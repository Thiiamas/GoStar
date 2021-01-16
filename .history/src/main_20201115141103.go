package main

import (
	"fmt"
	"Thiiamas/Utils"
	"math/rand"
	"time"
)

func main() {
	board := createBoard(25)
	drawBoard(board)
	grid := utils.NewGrid(3, 3)
	grid
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

/* func createBoard(size int) []int {
	//construction de la liste des cases
	numbers := make([]int, size)
	cases := make([]int, size)
	for i :=0; i<size; i++ {
		numbers[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	for x :=0; x<size; x++{
		random := rand.Intn(len(numbers))
		cases[x] = numbers[random]
		numbers = removeIndex(numbers, random)
	}
	return cases
} */
func createBoard(size int) []int {
	//construction de la liste des cases
	numbers := make([]int, size)
	board := grid.New([]int, size)
	cases := make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	for x := 0; x < size; x++ {
		random := rand.Intn(len(numbers))
		cases[x] = numbers[random]
		numbers = removeIndex(numbers, random)
	}
	return cases
}

func drawBoard(board []int) {
	var dim int
	switch len(board) {
	case 9:
		dim = 3
	case 16:
		dim = 4
	case 25:
		dim = 5
	case 36:
		dim = 6
	}
	vBoard := board
	for i := 0; i < dim; i++ {
		fmt.Println(vBoard[:dim])
		vBoard = vBoard[dim:]
	}
}
