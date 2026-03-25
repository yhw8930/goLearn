package main

import (
	"fmt"
	"math/rand"
	"time"
)

var bitMap = make(map[int32]int)

// test 哈希表暴力解法，对数器
func test(arr []int32, k, m int) int32 {
	countMap := make(map[int32]int)
	for _, num := range arr {
		countMap[num]++
	}
	for num, cnt := range countMap {
		if cnt == k {
			return num
		}
	}
	return -1
}

// 给定一个数组 arr，保证满足： 只有一种数出现了 K 次；其余所有数都出现了 M 次；K < M
// 要求： 不用哈希表、O (N) 时间、O (1) 空间，找出出现 K 次的数。
func onlyKTimes(arr []int32, k, m int) int32 {
	// 第一步：准备 32 位计数器
	if len(bitMap) == 0 {
		mapCreater()
	}

	// 	第二步：遍历所有数字，统计 1 的个数
	//	把每个数的二进制 1 挨个统计到 t 数组里。
	t := make([]int, 32)
	for _, num := range arr {
		n := num
		for n != 0 {
			rightOne := n & (-n)  // 提取最右侧的1
			t[bitMap[rightOne]]++ // 对应位计数+1
			n ^= rightOne         // 删掉这个1
		}
	}

	//	第三步：根据位数组，拼出答案
	//	其他数都出现 m 次
	// 	所以每一位的 1 总数 一定是 m 的倍数
	// 	只有目标数会破坏规律
	//	被破坏的位 → 就是答案的二进制位
	var ans int32 = 0
	for i := 0; i < 32; i++ {
		if t[i]%m != 0 {
			if t[i]%m == k {
				ans |= 1 << i
			} else {
				return -1
			}
		}
	}

	// 第四步：特殊情况 —— 答案是 0. 因为 0 的二进制全是 0，不会被统计，所以要单独判断。
	if ans == 0 {
		count := 0
		for _, num := range arr {
			if num == 0 {
				count++
			}
		}
		if count != k {
			return -1
		}
	}

	return ans
}

// mapCreater 建立 2^i → 第i位 的映射
func mapCreater() {
	var value int32 = 1
	for i := 0; i < 32; i++ {
		bitMap[value] = i
		value <<= 1
	}
}

// randomArray 生成 int32 数组
func randomArray(maxKinds, rangeLimit int, k, m int) []int32 {
	rand.Seed(time.Now().UnixNano())
	ktNum := randomNumber(rangeLimit)
	times := k
	if rand.Float64() < 0.5 {
		times = rand.Intn(m-1) + 1
	}

	numKinds := rand.Intn(maxKinds) + 2
	arr := make([]int32, times+(numKinds-1)*m)
	index := 0

	for ; index < times; index++ {
		arr[index] = ktNum
	}
	numKinds--

	set := make(map[int32]bool)
	set[ktNum] = true

	for numKinds > 0 {
		var curNum int32
		for {
			curNum = randomNumber(rangeLimit)
			if !set[curNum] {
				break
			}
		}
		set[curNum] = true
		numKinds--
		for i := 0; i < m; i++ {
			arr[index] = curNum
			index++
		}
	}

	// 打乱
	for i := range arr {
		j := rand.Intn(len(arr))
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

// randomNumber 返回 int32
func randomNumber(rangeLimit int) int32 {
	a := rand.Intn(rangeLimit) + 1
	b := rand.Intn(rangeLimit) + 1
	return int32(a - b)
}

func main() {
	kinds := 5
	rangeVal := 30
	testTime := 100000
	maxKM := 9

	fmt.Println("测试开始")
	for i := 0; i < testTime; i++ {
		a := rand.Intn(maxKM) + 1
		b := rand.Intn(maxKM) + 1
		k := min(a, b)
		m := max(a, b)

		if k == m {
			m++
		}

		arr := randomArray(kinds, rangeVal, k, m)
		ans1 := test(arr, k, m)
		ans2 := onlyKTimes(arr, k, m)

		if ans1 != ans2 {
			fmt.Println("ans1=", ans1)
			fmt.Println("ans2=", ans2)
			fmt.Println("出错了！")
			return
		}
	}
	fmt.Println("测试结束 ✅")
}
