package main

// 两个数组即可
type MinStack struct {
	base []int
	cpy []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		base: make([]int,0),
		cpy:  make([]int,0),
	}
}


func (this *MinStack) Push(x int)  {
	if len(this.base) == 0{
		this.base = append(this.base,x)
		this.cpy = append(this.cpy,x)
		return
	}
	this.base = append(this.base,x)
	if x<this.cpy[len(this.cpy)-1]{
		this.cpy = append(this.cpy,x)
	}else{
		this.cpy = append(this.cpy,this.cpy[len(this.cpy)-1])
	}
}


func (this *MinStack) Pop()  {
	this.base = this.base[:len(this.base)-1]
	this.cpy = this.cpy[:len(this.cpy)-1]

}


func (this *MinStack) Top() int {
	return this.base[len(this.base)-1]
}


func (this *MinStack) GetMin() int {
	return this.cpy[len(this.cpy)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */