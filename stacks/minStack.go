package main

import "math"

/**
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.
You must implement a solution with O(1) time complexity for each function.

Example 1:

Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2

Constraints:

-231 <= val <= 231 - 1
Methods pop, top and getMin operations will always be called on non-empty stacks.
At most 3 * 104 calls will be made to push, pop, top, and getMin.
*/

/**
maintain two stacks, push to both stacks if required
pop only when values are same
can consider map approach as well with maintaining count
*/

type MinStack struct {
	data    []int
	minData []int
}

func Constructor() MinStack {
	return MinStack{
		data:    []int{},
		minData: []int{math.MaxInt},
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if this.minData[len(this.minData)-1] >= val {
		this.minData = append(this.minData, val)
	}
}

func (this *MinStack) Pop() {
	var lastData int
	if len(this.data) != 0 {
		lastData = this.data[len(this.data)-1]
		this.data = this.data[:len(this.data)-1]
	}
	if len(this.minData) != 0 {
		if lastData == this.minData[len(this.minData)-1] {
			this.minData = this.minData[:len(this.minData)-1]
		}
	}
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minData) != 0 {
		return this.minData[len(this.minData)-1]
	}
	return math.MaxInt
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
