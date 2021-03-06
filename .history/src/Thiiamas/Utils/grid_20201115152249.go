package utils

import (
	"fmt"
)
//Grid s
type (
	Grid struct {
		values     []int
		cols, rows int
	}
)

func NewGrid(cols, rows int) *Grid {
	return &Grid{
		values: make([]int, cols*rows),
		cols:   cols,
		rows:   rows,
	}
}

func (this *Grid) Do(f func(p Point, value int)) {
	for x := 0; x < this.cols; x++ {
		for y := 0; y < this.rows; y++ {
			f(Point{x, y}, this.values[x*this.cols+y])
		}
	}
}

func (this *Grid) Get(p Point) interface{} {
	if p.X < 0 || p.Y < 0 || p.X >= this.cols || p.Y >= this.rows {
		return nil
	}
	v := this.values[p.X*this.cols+p.Y]
	return v
}

func (this *Grid) Draw(){
	pointCols := 0
	for x := 0; x < this.rows ; x++ {
		fmt.Println(this.values[pointCols:this.cols+pointCols])
		pointCols += this.cols
	}
	
}

func (this *Grid) Rows() int {
	return this.rows
}

func (this *Grid) Cols() int {
	return this.cols
}

func (this *Grid) Len() int {
	return this.rows * this.cols
}

func (this *Grid) Set(p Point, value int) {
	this.values[p.X*this.rows+p.Y] = value
}
