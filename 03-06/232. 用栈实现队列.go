package main

// 两个栈即可模拟队列
type MyQueue struct {
	pushStack []int
	popStack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		pushStack: make([]int,0),
		popStack:  make([]int,0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.pushStack = append(this.pushStack,x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.popStack) == 0{
		this.Push2Pop()
	}
	v:=this.popStack[len(this.popStack)-1]
	this.popStack = this.popStack[:len(this.popStack)-1]
	return v
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.popStack) == 0{
		this.Push2Pop()
	}
	return this.popStack[len(this.popStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.pushStack) == 0 && len(this.popStack) == 0
}

// 将push栈中内容全部放入pop栈
func (this *MyQueue)Push2Pop(){
	for len(this.pushStack)!=0{
		this.popStack = append(this.popStack,this.pushStack[len(this.pushStack)-1])
		this.pushStack = this.pushStack[:len(this.pushStack)-1]
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
