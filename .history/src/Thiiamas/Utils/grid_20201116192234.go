package utils

import (
	"fmt"
)
//Grid s
type (
	Grid struct {
		Values     []int
		Cols, Rows int
	}
)

func NewGrid(Cols, Rows int) *Grid {
	return &Grid{
		Values: make([]int, Cols*Rows),
		Cols:   Cols,
		Rows:   Rows,
	}
}

func (this *Grid) Do(f func(p Point, value int)) {
	for x := 0; x < this.Cols; x++ {
		for y := 0; y < this.Rows; y++ {
			f(Point{x, y}, this.Values[x*this.Cols+y])
		}
	}
}

func (this *Grid) Get(p Point) interface{} {
	if p.X < 0 || p.Y < 0 || p.X >= this.Cols || p.Y >= this.Rows {
		return nil
	}
	v := this.Values[p.X+this.Cols*p.Y]
	return v
}
func (this *Grid) GetValue(i int) int {
	if i < 0 || i >= this.Rows * this.Cols {
		return -1
	}
	v := this.Values[i]
	return v
}

func (this *Grid) GetIndex(v int) int{
	for i, k := range this.Values {
		if k == v {
			return i
		}
	}
	return -1
}

func (this Grid) SwapIndex(i1, i2 int) {
	i1V := this.GetValue(i1)
	i2V := this.GetValue(i2)

	this.SetIndex(i1, i2V)
	this.SetIndex(i2, i1V)
}

func (this Grid) Draw(){
	pointCols := 0
	for x := 0; x < this.Rows ; x++ {
		fmt.Println(this.Values[pointCols:this.Cols+pointCols])
		pointCols += this.Cols
	}
	
}

func (this []*Grid) Draw(){
	
}

func (this Grid) GetObject() Grid {
	return this
}

func (this *Grid) GetRows() int {
	return this.Rows
}

func (this *Grid) GetCols() int {
	return this.Cols
}

func (this *Grid) Len() int {
	return this.Rows * this.Cols
}

func (this *Grid) GetValues() []int {
	return this.Values
}

func (this *Grid) Set(p Point, value int) {
	this.Values[p.X+this.Cols*p.Y] = value
}

func (this *Grid) SetIndex(index int, value int) {
	this.Values[index] = value
}

func (this *Grid) Equals(grid *Grid) bool {
	if this.Rows != grid.Rows || this.Cols != grid.Cols {
		return false
	}
	for x :=0; x < this.Len(); x++ {
		if this.GetValue(x) == grid.GetValue(x) {
			continue
		} else {
			return false
		}
	}
	return true
}
