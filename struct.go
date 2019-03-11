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

func (this *StackInt) Print() string {
	str := "["
	for i := 0; i < this.length; i++ {
		str += fmt.Sprintf(" %d", this.data[i])
	}
	str += " ]"

	return str
}

/* StackString */

type StackString struct {
	data   []string
	length int
}

func NewStackString() *StackString {
	return &StackString{[]string{}, 0}
}

func (this *StackString) Len() int {
	return this.length
}

func (this *StackString) Peek() string {
	if this.length == 0 {
		panic("Stack is empty!")
	}
	return this.data[this.length-1]
}

func (this *StackString) Pop() string {
	if this.length == 0 {
		panic("Stack is empty!")
	}
	v := this.data[this.length-1]
	this.data = this.data[0 : this.length-1]
	this.length--

	return v
}

func (this *StackString) Push(v string) {
	this.data = append(this.data, v)
	this.length++
}

func (this *StackString) Print() string {
	str := "["
	for i := 0; i < this.Len(); i++ {
		str += fmt.Sprintf(" \"%s\"", this.data[i])
	}
	str += " ]"

	return str
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

func (this *QueueInt) Print() string {
	str := "["
	for i := 0; i < this.length; i++ {
		str += fmt.Sprintf(" %d", this.data[i])
	}
	str += " ]"

	return str
}

/* Union Find */

type UnionFind struct {
	Size  int
	Nodes []int
	Ranks []int
}

func NewUnionFind(n int) *UnionFind {
	nodes := make([]int, n)
	counts := make([]int, n)
	for i := 0; i < n; i++ {
		nodes[i] = i
		counts[i] = 1
	}
	return &UnionFind{n, nodes, counts}
}

func (p *UnionFind) Root(a int) int {
	nodes := []int{}

	par := p.Nodes[a]
	var root int
	for {
		if par == p.Nodes[par] {
			root = par
			break
		}
		nodes = append(nodes, par)
		par = p.Nodes[par]
	}

	for _, n := range nodes {
		p.Nodes[n] = root
	}

	return root
}

func (p *UnionFind) Merge(a, b int) {
	if a == b {
		return
	}

	if p.Rank(a) > p.Rank(b) {
		a, b = b, a
	}
	ra, rb := p.Root(a), p.Root(b)
	p.Nodes[rb] = ra

	p.Ranks[ra] += p.Ranks[rb]
	p.Ranks[rb] = 0
}

func (p *UnionFind) Same(a, b int) bool {
	return p.Root(a) == p.Root(b)
}

func (p *UnionFind) Rank(a int) int {
	return p.Ranks[p.Root(a)]
}
