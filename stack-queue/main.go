package main

import (
	"fmt"
)

func main() {
	max := 10

	// Stack
	top := 0
	st := make([]int, 3, max)

	initStack(&top)

	pushStack(3, &top, st, max)
	pushStack(5, &top, st, max)
	pushStack(7, &top, st, max)

	fmt.Printf("st => %d\n", st)

	fmt.Printf("pop() => %d\n", popStack(&top, st))
	fmt.Printf("pop() => %d\n", popStack(&top, st))

	pushStack(9, &top, st, max)
	pushStack(10, &top, st, max)

	fmt.Printf("st => %d\n", st)

	// queue
	tail := 0
	head := 0
	qu := make([]int, 5, max)

	initQueue(&head, &tail)

	enqueue(3, qu, head, &tail, max)
	enqueue(5, qu, head, &tail, max)
	enqueue(7, qu, head, &tail, max)

	fmt.Printf("qu => %d\n", qu)

	fmt.Printf("dequeue() => %d\n", dequeue(qu, &head, tail, max))
	fmt.Printf("dequeue() => %d\n", dequeue(qu, &head, tail, max))

	enqueue(9, qu, head, &tail, max)
	enqueue(10, qu, head, &tail, max)

	fmt.Printf("qu => %d\n", qu)
}

// Stack
func initStack(top *int) {
	*top = 0
}

func isFullStack(top int, max int) bool {
	return top == max
}

func isEmptyStack(top int) bool {
	return top == 0
}

func pushStack(x int, top *int, st []int, max int) {
	if isFullStack(*top, max) {
		fmt.Println("stack is fulll.")
		return
	}

	st[*top] = x
	*top++
}

func popStack(top *int, st []int) int {
	if isEmptyStack(*top) {
		fmt.Println("stack is empty.")
		return -1
	}

	*top--
	return st[*top]
}

// Queue
func initQueue(head *int, tail *int) {
	*head = 0
	*tail = 0
}

func isFullQueue(head int, tail int, max int) bool {
	return head == (tail+1)%max
}

func isEmptyQueue(head int, tail int) bool {
	return head == tail
}

func enqueue(x int, qu []int, head int, tail *int, max int) {
	if isFullQueue(head, *tail, max) {
		fmt.Println("error: queue is full.")
		return
	}

	qu[*tail] = x
	*tail++
	if *tail == max {
		*tail = 0
	}
}

func dequeue(qu []int, head *int, tail int, max int) int {
	if isEmptyQueue(*head, tail) {
		fmt.Println("error: queue is empty.")
		return -1
	}

	res := qu[*head]
	*head++
	if *head == max {
		*head = 0
	}
	return res
}
