package main

func run_1()  {

}


type Solution struct {
	initial []int
	changed []int
}


func Constructor1(nums []int) Solution {
	return Solution{
		nums,
		[]int{},
	}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.initial
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	var temp []int
	result := []int{}
	if len(this.changed) == 0 {
		temp = this.initial
	} else {
		temp = this.changed
	}
	for i, d := range temp  {
		if i == 0 {
			continue
		}
		result = append(result, d)
	}
	result = append(result, temp[0])
	this.changed = result
	return result
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */