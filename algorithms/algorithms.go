package algorithms

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

// Двоичный поиск

func BinarySearch(arr []int, target int) (int,error){ // BinarySearch проводит двоичный поиск числа в отсортиванном массиве
	if len(arr) == -1{
		return 0,errors.New("array is empty")
	}
	left,right := 0,len(arr)-1
	for left <= right{
		mid := left + (right-left)/2
		if arr[mid] == target{
			return mid,nil
		} else if arr[mid] < target{
			left = mid + 1
		} else{
			right = mid - 1
		}
	}
	return  -1,nil
}

// Префиксные суммы 

func NewPrefixSum(arr []int) []int{ // NewPrefixSum создание новой суммы
	n := len(arr)
	prefixSum := make([]int,n+1)
	for i := 1; i <= n; i++{
		prefixSum[i] = prefixSum[i-1] + arr[i-1]
	}
	return prefixSum
}

func RangeSum(prefixSum []int, i, j int) int{ // RangeSum находит сумму от i до j включительно
	if i == 0{
		return prefixSum[j]
	}
	return prefixSum[j] - prefixSum[i-1]
}

// Две суммы

func TwoSum(arr []int,target int) (int,int,bool){ // TwoSum ищет 2 числа, сумма которых будет равна target
	left, right := 0, len(arr) - 1
	for left < right{
		sum := arr[left] + arr[right]
		if sum == target{
			return arr[left],arr[right],true
		} else if sum < target{
			left++
		} else{
			right--
		}
	}
	return 0,0,false 
}

// Удаление дубликатов

func RemoveDuplicates(arr []int) int{ // RemoveDuplicates возвращает длину массива без дупликатов
	if len(arr) == 0{
		return 0
	}
	uniqueIndex := 0
	for i := 1; i < len(arr); i++{
		if arr[i] != arr[uniqueIndex]{
			uniqueIndex++
			arr[uniqueIndex] = arr[i]
		}
	}
	return uniqueIndex + 1
}

// Сортировка событий 

type Event struct{ // Event структура событий
	time int
	start bool
}

func MaxOverlappingIntervals(intervals [][2]int) int{ // MaxOverlappingIntervals возвращает максимальное количество одновременно активных событий
	events := []Event{}
	for _, interval := range intervals{
		events = append(events,Event{interval[0],true})
		events = append(events,Event{interval[1],false})
	}
	sort.Slice(events,func(i, j int) bool{
		if events[i].time == events[j].time{
			return !events[i].start && events[j].start
		}
		return events[i].time < events[j].time
	})
	maxActive, active := 0, 0
	for _, event := range events{
		if event.start{
			active++
			if active > maxActive{
				maxActive = active
			}
		} else{
			active--
		}
	}
	return maxActive
}

// Двоичное дерево

type Node struct{ // Node структура узла
	value int
	Left *Node
	Right *Node
}

func NewNode(value int) *Node{ // NewNode создание нового узла (корня)
	return &Node{
		value:value,
		Left : nil,
		Right : nil,
	}
}

func (n *Node) AddLeft(value int){ // AddLeft создание узла в левом поддереве
	n.Left = NewNode(value)
}

func (n *Node) AddRight(value int){ // AddRight создание узла в правом поддереве
	n.Right = NewNode(value)
}

func PreOrder(n *Node){  // PreOrder прямой обход и печать дерева
	if n == nil{
		return
	}
	fmt.Printf("%d ",n.value)
	PreOrder(n.Left)
	PreOrder(n.Right)
}

func InOrder(n *Node){ // InOrder симметричный обход и печать дерева
	if n == nil{
		return
	}
	InOrder(n.Left)
	fmt.Printf("%d ",n.value)
	InOrder(n.Right)
}

func PostOrder(n *Node){ // PostOrder обратный обход и печать дерева
	if n == nil{
		return
	}
	PostOrder(n.Left)
	PostOrder(n.Right)
	fmt.Printf("%d ",n.value)
}

// Одномерное динамическое программирование (зайчик)

func MinCostJump(cost []int) int{ // MinCostJump вычисляет минимальную цену на прохождение массива
	n := len(cost)
	if n == 0{
		return 0
	}
	if n == 1{
		return cost[0]
	}
	prev2 := cost[0]
	prev1 := cost[0] + cost[1]
	for i := 2; i < n; i++{
		current := int(math.Min(float64(prev1+cost[i]), float64(prev2+cost[i])))
		prev2 = prev1
		prev1 = current
		fmt.Println(prev1)
	}
	return prev1
}

// Двумерное динамическое программирование (черепашка)

func MinCostWalk(cost [][]int, n, m int) int{ // MinCostWalk вычисляет минимальную цену на прохождение двумерного массива
	dp := make([][]int,n)
	for i := range dp{
		dp[i] = make([]int,m)
	}
	dp[0][0] = cost[0][0]
	for i:=1; i < n; i++{
		dp[0][i] = cost[0][i] + dp[0][i - 1]
	}
	for i:=1; i < m; i++{
		dp[i][0] = cost[i][0] + dp[i -1][0]
	}
	for i := 1; i < n; i++{
		for j := 1; j < m; j++{
			dp[i][j] = cost[i][j] + int(math.Min(float64(dp[i-1][j]),float64(dp[i][j-1])))
		}
	}
	return dp[n-1][m-1]
}