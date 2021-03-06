package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* StackInt */

func TestStackInt(t *testing.T) {
	s := NewStackInt()

	if s.Len() != 0 {
		t.Errorf("Length of an empty stack should be 0")
	}

	s.Push(1)

	if s.Print() != "[ 1 ]" {
		t.Errorf("Print should be \"[ 1 ]\"")
	}

	if s.Len() != 1 {
		t.Errorf("Length should be 0")
	}

	if s.Peek() != 1 {
		t.Errorf("Top item on the stack should be 1")
	}

	if s.Pop() != 1 {
		t.Errorf("Top item should have been 1")
	}

	if s.Print() != "[ ]" {
		t.Errorf("Print should be \"[ ]\"")
	}

	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}

	s.Push(1)
	s.Push(2)

	if s.Len() != 2 {
		t.Errorf("Length should be 2")
	}

	if s.Peek() != 2 {
		t.Errorf("Top of the stack should be 2")
	}
}

/* StackString */

func TestStackString(t *testing.T) {
	s := NewStackString()

	if s.Len() != 0 {
		t.Errorf("Length of an empty stack should be 0")
	}

	s.Push("I")

	if s.Print() != "[ \"I\" ]" {
		t.Errorf("Print should be [ \"I\" ]")
	}

	if s.Len() != 1 {
		t.Errorf("Length should be 0")
	}

	if s.Peek() != "I" {
		t.Errorf("Top item on the stack should be 1")
	}

	if s.Pop() != "I" {
		t.Errorf("Top item should have been 1")
	}

	if s.Print() != "[ ]" {
		t.Errorf("Print should be \"[ ]\"")
	}

	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}

	s.Push("I")
	s.Push("am")
	s.Push("gopher.")

	if s.Len() != 3 {
		t.Errorf("Length should be 3")
	}

	if s.Peek() != "gopher." {
		t.Errorf("Top of the stack should be \"gopher.\"")
	}
}

/* QueueInt */

func TestQueueInt(t *testing.T) {
	q := NewQueueInt()

	if q.Len() != 0 {
		t.Errorf("Length should be 0")
	}

	q.Enqueue(1)

	if q.Print() != "[ 1 ]" {
		t.Errorf("Print should be \"[ 1 ]\"")
	}

	if q.Len() != 1 {
		t.Errorf("Length should be 1")
	}

	if q.Peek() != 1 {
		t.Errorf("Enqueued value should be 1")
	}

	v := q.Dequeue()

	if v != 1 {
		t.Errorf("Dequeued value should be 1")
	}

	if q.Len() != 0 {
		t.Errorf("Length should be 0")
	}

	q.Enqueue(1)
	q.Enqueue(2)

	if q.Len() != 2 {
		t.Errorf("Length should be 2")
	}

	if q.Peek() != 1 {
		t.Errorf("First value should be 1")
	}

	q.Dequeue()

	if q.Print() != "[ 2 ]" {
		t.Errorf("Print should be \"[ 1 ]\"")
	}

	if q.Peek() != 2 {
		t.Errorf("Next value should be 2")
	}
}

/* Union Find */

func TestUnionFind(t *testing.T) {
	n := 4
	uf := NewUnionFind(n)

	for i := 0; i < n; i++ {
		assert.Equal(t, i, uf.Root(i))
		assert.Equal(t, 1, uf.Rank(i))
	}

	// connect 0 and 1
	assert.False(t, uf.Same(0, 1))
	uf.Merge(0, 1)
	assert.True(t, uf.Same(0, 1))
	assert.Equal(t, 2, uf.Rank(0))
	assert.Equal(t, 2, uf.Rank(1))

	// connect 2 and 3, 1 and 3
	uf.Merge(2, 3)
	assert.Equal(t, 2, uf.Rank(2))
	uf.Merge(1, 3)
	assert.Equal(t, 4, uf.Rank(1))
	assert.Equal(t, 4, uf.Rank(0))
	assert.True(t, uf.Same(0, 2))
}
