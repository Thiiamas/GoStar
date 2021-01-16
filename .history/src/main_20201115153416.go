package main

import (
	"fmt"
	"Thiiamas/Utils"
	"math/rand"
	"time"
)

func main() {
	board := createBoard(5)
	board.Draw()
	fmt.println("fin")

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
func createBoard(size int) *utils.Grid {
	//construction de la liste des cases
	numbers := make([]int, size)
	board := utils.NewGrid(size,size)
	for i := 0; i < size; i++ {
		numbers[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	for x := 0; x < size; x++ {
		random := rand.Intn(len(numbers))
		board.SetIndex(x, numbers[random])
		numbers = removeIndex(numbers, random)
	}
	return board
}

