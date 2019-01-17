package main

import "fmt"

func run_2()  {
	x := 0
	obj := Constructor();
	obj.Push(x);
	obj.Pop();
	param_3 := obj.Top();
	param_4 := obj.GetMin();
	fmt.Println(param_3, param_4)
}

type MinStack struct {
	stack []int
	length1 int
	min []int
	length2 int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		[]int{},
		0,
		[]int{},
		0,
	}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	this.length1 ++
	if x <= this.min[this.length2-1] {
		this.min = append(this.min, x)
		this.length2 ++
	}
}


func (this *MinStack) Pop()  {
	if  this.stack[this.length1-1] == this.min[this.length2-1]{
		this.length2--
	}
	this.length1 --
}


func (this *MinStack) Top() int {
	return this.stack[this.length1-1]
}


func (this *MinStack) GetMin() int {
	return this.min[this.length2-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
