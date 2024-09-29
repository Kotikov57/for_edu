package structs

import (
//	"container/heap"
	"errors"
	"fmt"
)

// Стэк

type Node struct { // Node структура узла
	value int
	next  *Node
}

type Stack struct { // Stack структура стэка
	top *Node
}

func (s *Stack) StackPush(data int) { // StackPush добавляет элемент в конец стэка
	newNode := &Node{value: data, next: s.top}
	s.top = newNode
}

func (s *Stack) StackPop() (int, error) { // StackPop удаляет элемент из конца стэка
	if s.top == nil {
		return 0, errors.New("stack is empty")
	}
	v := s.top.value
	s.top = s.top.next
	return v, nil
}

func (s Stack) StackFind(data int) (int, error) { // StackFind проверяет наличие элемента в стэке
	if s.top == nil {
		return 0, errors.New("stack is empty")
	}
	current := s.top
	depth := 0
	for current != nil {
		if current.value == data {
			return depth, nil
		}
		depth++
		current = current.next
	}
	return 0, errors.New("not found")
}

// Очередь

type Queue struct { // Queue структура очереди
	head *Node
	tail *Node
}

func (q *Queue) Enqueue(data int) { // Enqueue добавляет элемент в конец очереди
	newNode := &Node{value: data, next: nil}
	if q.tail == nil {
		q.tail = newNode
		q.head = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

func (q *Queue) Dequeue() (int, error) { // Dequeue возвращает и удаляет элемент из начала очереди
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	v := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return v, nil
}

func (q Queue) Peak() (int, error) { // Peak возвращает первый элемент очереди, не удаляя его
	if q.head == nil {
		return 0, errors.New("queue is empty")
	}
	return q.head.value, nil
}

func (q Queue) IsEmpty() bool { // IsEmpty проверяет, пустая ли очередь
	return q.head == nil
}

func (q Queue) Size() (int, error) { // Size возвращает длину очереди
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	out := 0
	current := q.head
	for current != nil {
		out++
		current = current.next
	}
	return out, nil
}

// Бинарное дерево поиска

type TreeNode struct { // TreeNode структура узла дерева
	data  int
	left  *TreeNode
	right *TreeNode
}

type BinarySearchTree struct { // BinarySearchTree структура двоичного дерева
	root *TreeNode
}

func (bst *BinarySearchTree) Insert(data int) { // Insert добавляет элемент в дерево
	NewNode := &TreeNode{data: data, left: nil, right: nil}
	if bst.root == nil {
		bst.root = NewNode
		return
	}
	current := bst.root
	for {
		if data < current.data {
			if current.left == nil {
				current.left = NewNode
				return
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = NewNode
				return
			}
			current = current.right
		}
	}

}

func (bst BinarySearchTree) Search(data int) (string, error) { // Search проверяет наличие элемента в дереве
	if bst.root == nil {
		return "", errors.New("tree is empty")
	}
	current := bst.root
	for current != nil {
		if current.data == data {
			return "found", nil
		}
		if data > current.data {
			current = current.right
		} else {
			current = current.left
		}
	}
	return "", errors.New("not found")
}

func (bst BinarySearchTree) PrintInOrder() { // PrintInOrder печатает дерево по порядку
	inOrder(bst.root)
	fmt.Println()
}

func inOrder(node *TreeNode) { // InOrder определяет порядок печати дерева
	if node != nil {
		inOrder(node.left)
		fmt.Print(node.data, " ")
		inOrder(node.right)
	}
}

// Очередь с приоритетом на основе кучи

type Item struct { // Item структура элемента кучи
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item // PriorityQueue структура очереди с приоритетом

func (pq PriorityQueue) Len() int { // Len возвращает длину очереди
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool { // Less сравнивает приоритеты двух элементов
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) { // Swap меняет местами элементы
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) { // Push добавляет элемент в очередь
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} { // Pop возвращает и удаляет элемент из очереди
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}
