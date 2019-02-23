package algo

import "fmt"

// More general stack implementation is here.
// https://godoc.org/github.com/golang-collections/collections/stack

/* StackInt */

type StackInt struct {
	data   []int
	length int
}

func NewStackInt() *StackInt {
	return &StackInt{[]int{}, 0}
}

func (this *StackInt) Len() int {
	return this.length
}

func (this *StackInt) Peek() int {
	if this.length == 0 {
		panic("Stack is empty!")
	}
	return this.data[this.length-1]
}

func (this *StackInt) Pop() int {
	if this.length == 0 {
		panic("Stack is empty!")
	}
	v := this.data[this.length-1]
	this.data = this.data[0 : this.length-1]
	this.length--

	return v
}

func (this *StackInt) Push(v int) {
	this.data = append(this.data, v)
	this.length++
}

func (this *StackInt) Print() {
	fmt.Printf("[")
	for i := 0; i < this.length; i++ {
		fmt.Printf(" %d", this.data[i])
	}
	fmt.Printf(" ]")
}

/* QueueInt */

type QueueInt struct {
	data   []int
	length int
}

func NewQueueInt() *QueueInt {
	return &QueueInt{[]int{}, 0}
}

func (this *QueueInt) Len() int {
	return this.length
}

func (this *QueueInt) Peek() int {
	if this.length == 0 {
		panic("Queue is empty!")
	}
	return this.data[0]
}

func (this *QueueInt) Dequeue() int {
	if this.length == 0 {
		panic("Queue is empty!")
	}
	v := this.data[0]
	this.data = this.data[1:this.length]
	this.length--

	return v
}

func (this *QueueInt) Enqueue(v int) {
	this.data = append(this.data, v)
	this.length++
}

func (this *QueueInt) Print() {
	fmt.Printf("[")
	for i := 0; i < this.length; i++ {
		fmt.Printf(" %d", this.data[i])
	}
	fmt.Printf(" ]")
}
