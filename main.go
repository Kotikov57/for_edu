package main

import(
	"for_edu/algorithms"
	"fmt"
)

func main(){
	cost := []int{1,1,1,3,4,1,6}
	fmt.Println(algorithms.MinCostJump(cost))
}