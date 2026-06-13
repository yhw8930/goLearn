package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
)

// Right 返回所有人喝完咖啡并让杯子变干净的最早结束时间。
// 这是验证用的彻底暴力：每个人枚举所有咖啡机，得到喝完时间后再枚举每个杯子洗或自然挥发。
//
// 时间复杂度：O(M^N * 2^N)，M 是咖啡机数量，N 是人数。
// 空间复杂度：O(M+N)。
func Right(arr []int, n, a, b int) int {
	times := make([]int, len(arr))
	drink := make([]int, n)
	return coffeeForceMake(arr, times, 0, drink, n, a, b)
}

func coffeeForceMake(arr, times []int, kth int, drink []int, n, a, b int) int {
	if kth == n {
		drinkSorted := append([]int(nil), drink[:kth]...)
		sort.Ints(drinkSorted)
		return coffeeForceWash(drinkSorted, a, b, 0, 0, 0)
	}
	ans := int(^uint(0) >> 1)
	for i := 0; i < len(arr); i++ {
		work := arr[i]
		pre := times[i]
		drink[kth] = pre + work
		times[i] = pre + work
		ans = min(ans, coffeeForceMake(arr, times, kth+1, drink, n, a, b))
		drink[kth] = 0
		times[i] = pre
	}
	return ans
}

func coffeeForceWash(drinks []int, a, b, index, washLine, time int) int {
	if index == len(drinks) {
		return time
	}
	wash := max(drinks[index], washLine) + a
	ans1 := coffeeForceWash(drinks, a, b, index+1, wash, max(wash, time))

	dry := drinks[index] + b
	ans2 := coffeeForceWash(drinks, a, b, index+1, washLine, max(dry, time))
	return min(ans1, ans2)
}

type coffeeMachine struct {
	timePoint int
	workTime  int
}

type coffeeMachineHeap []coffeeMachine

func (h coffeeMachineHeap) Len() int { return len(h) }

func (h coffeeMachineHeap) Less(i, j int) bool {
	return h[i].timePoint+h[i].workTime < h[j].timePoint+h[j].workTime
}

func (h coffeeMachineHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *coffeeMachineHeap) Push(x any) {
	*h = append(*h, x.(coffeeMachine))
}

func (h *coffeeMachineHeap) Pop() any {
	old := *h
	n := len(old)
	ans := old[n-1]
	*h = old[:n-1]
	return ans
}

// MinTime1 返回所有人喝完咖啡并让杯子变干净的最早结束时间。
// 先用小根堆贪心安排每个人最早喝完咖啡的时间，再递归枚举每个杯子洗或自然挥发。
//
// 时间复杂度：O(N*logM + 2^N)，M 是咖啡机数量，N 是人数。
// 空间复杂度：O(M+N)，递归调用栈深度和堆空间。
func MinTime1(arr []int, n, a, b int) int {
	drinks := coffeeGetDrinks(arr, n)
	return BestTime(drinks, a, b, 0, 0)
}

func coffeeGetDrinks(arr []int, n int) []int {
	machines := &coffeeMachineHeap{}
	heap.Init(machines)
	for _, workTime := range arr {
		heap.Push(machines, coffeeMachine{workTime: workTime})
	}
	drinks := make([]int, n)
	for i := 0; i < n; i++ {
		cur := heap.Pop(machines).(coffeeMachine)
		cur.timePoint += cur.workTime
		drinks[i] = cur.timePoint
		heap.Push(machines, cur)
	}
	return drinks
}

// BestTime 返回 drinks[index..] 所有杯子变干净的最早结束时间。
// free 表示洗杯机何时可用；每个杯子在“串行机洗”和“并行自然挥发”之间二选一。
//
// 时间复杂度：O(2^N)。
// 空间复杂度：O(N)，递归调用栈深度。
func BestTime(drinks []int, wash, air, index, free int) int {
	if index == len(drinks) {
		return 0
	}
	selfClean1 := max(drinks[index], free) + wash
	restClean1 := BestTime(drinks, wash, air, index+1, selfClean1)
	p1 := max(selfClean1, restClean1)

	selfClean2 := drinks[index] + air
	restClean2 := BestTime(drinks, wash, air, index+1, free)
	p2 := max(selfClean2, restClean2)
	return min(p1, p2)
}

// MinTime2 返回所有人喝完咖啡并让杯子变干净的最早结束时间。
// 咖啡机安排仍用小根堆贪心，洗杯过程把 BestTime 的递归尝试改成动态规划。
//
// 时间复杂度：O(N*logM + N*maxFree)，maxFree 为洗杯机可能的最晚空闲时间。
// 空间复杂度：O(N*maxFree)。
func MinTime2(arr []int, n, a, b int) int {
	drinks := coffeeGetDrinks(arr, n)
	return BestTimeDP(drinks, a, b)
}

// BestTimeDP 返回 drinks 中所有杯子变干净的最早结束时间，是 BestTime 的动态规划版本。
// dp[index][free] 表示从 index 号杯子开始处理、洗杯机在 free 时刻可用时的最优结束时间。
//
// 时间复杂度：O(N*maxFree)。
// 空间复杂度：O(N*maxFree)。
func BestTimeDP(drinks []int, wash, air int) int {
	n := len(drinks)
	maxFree := 0
	for i := 0; i < n; i++ {
		maxFree = max(maxFree, drinks[i]) + wash
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, maxFree+1)
	}
	for index := n - 1; index >= 0; index-- {
		for free := 0; free <= maxFree; free++ {
			selfClean1 := max(drinks[index], free) + wash
			if selfClean1 > maxFree {
				break
			}
			restClean1 := dp[index+1][selfClean1]
			p1 := max(selfClean1, restClean1)

			selfClean2 := drinks[index] + air
			restClean2 := dp[index+1][free]
			p2 := max(selfClean2, restClean2)
			dp[index][free] = min(p1, p2)
		}
	}
	return dp[0][0]
}

func coffeeRandomArray(length, maxValue int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(maxValue) + 1
	}
	return arr
}

func coffeePrintArray(arr []int) {
	fmt.Print("arr : ")
	for _, num := range arr {
		fmt.Print(num, ", ")
	}
	fmt.Println()
}
