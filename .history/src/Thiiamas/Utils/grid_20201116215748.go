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

func (this *Grid) Swap(p1, p2 Point) {
	p1V := this.Get(p1)
	p2V := this.Get(p2)
	if p2V == nil || p1V == nil {
		return
	}
	this.Set(p1, p2V.(int))
	this.Set(p2, p1V.(int))
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
		val := this.Values[pointCols:this.Cols+pointCols]
		for i, v := range val {
			switch v/10{
			case 0: 
				fmt.Print(" ")
				fmt.Print(v)
				fmt.Print(" ")
			case 1:
				fmt.Print(v)
				fmt.Print(" ")
			case 2:
				fmt.Print(v)
			}
			if i % this.Cols == this.Cols - 1 {
				fmt.Println("")
			}
		}
		pointCols += this.Cols
	}
	
}

func Draw(array []*Grid){
	for i, grid := range array {
		grid.Draw()
		if (i != len(array) - 1){
			fmt.Println(",")
		}
	}
	
}

func IndexToPoint(index int, grid *Grid) (point Point) {
	return Point {
		X: index % grid.Cols,
		Y: index / grid.Cols,
	}
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
