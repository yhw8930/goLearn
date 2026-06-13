package main

// Jump 返回中国象棋棋盘上马从 (0,0) 出发，正好跳 k 步到达 (a,b) 的方法数。
// 棋盘大小固定为 10*9，每一步按马走日的 8 个方向递归尝试。
//
// 时间复杂度：O(8^K)。
// 空间复杂度：O(K)，递归调用栈深度。
func Jump(a, b, k int) int {
	return horseJumpProcess(0, 0, k, a, b)
}

func horseJumpProcess(x, y, rest, a, b int) int {
	if x < 0 || x > 9 || y < 0 || y > 8 {
		return 0
	}
	if rest == 0 {
		if x == a && y == b {
			return 1
		}
		return 0
	}
	ways := horseJumpProcess(x+2, y+1, rest-1, a, b)
	ways += horseJumpProcess(x+1, y+2, rest-1, a, b)
	ways += horseJumpProcess(x-1, y+2, rest-1, a, b)
	ways += horseJumpProcess(x-2, y+1, rest-1, a, b)
	ways += horseJumpProcess(x-2, y-1, rest-1, a, b)
	ways += horseJumpProcess(x-1, y-2, rest-1, a, b)
	ways += horseJumpProcess(x+1, y-2, rest-1, a, b)
	ways += horseJumpProcess(x+2, y-1, rest-1, a, b)
	return ways
}

// HorseJumpDP 返回马从 (0,0) 出发，正好跳 k 步到达 (a,b) 的方法数。
// dp[x][y][rest] 表示从 (x,y) 出发还剩 rest 步时到目标点的方法数，按 rest 从小到大填表。
//
// 时间复杂度：O(10*9*K)。
// 空间复杂度：O(10*9*K)。
func HorseJumpDP(a, b, k int) int {
	dp := make([][][]int, 10)
	for x := range dp {
		dp[x] = make([][]int, 9)
		for y := range dp[x] {
			dp[x][y] = make([]int, k+1)
		}
	}
	dp[a][b][0] = 1
	for rest := 1; rest <= k; rest++ {
		for x := 0; x < 10; x++ {
			for y := 0; y < 9; y++ {
				ways := horseJumpPick(dp, x+2, y+1, rest-1)
				ways += horseJumpPick(dp, x+1, y+2, rest-1)
				ways += horseJumpPick(dp, x-1, y+2, rest-1)
				ways += horseJumpPick(dp, x-2, y+1, rest-1)
				ways += horseJumpPick(dp, x-2, y-1, rest-1)
				ways += horseJumpPick(dp, x-1, y-2, rest-1)
				ways += horseJumpPick(dp, x+1, y-2, rest-1)
				ways += horseJumpPick(dp, x+2, y-1, rest-1)
				dp[x][y][rest] = ways
			}
		}
	}
	return dp[0][0][k]
}

func horseJumpPick(dp [][][]int, x, y, rest int) int {
	if x < 0 || x > 9 || y < 0 || y > 8 {
		return 0
	}
	return dp[x][y][rest]
}

// Ways 与 Jump 等价，保留 Java 课程中的第二组递归入口。
// 从 (0,0) 出发正好跳 step 步到 (a,b)，每层递归枚举 8 个马跳方向。
//
// 时间复杂度：O(8^step)。
// 空间复杂度：O(step)。
func Ways(a, b, step int) int {
	return horseJumpF(0, 0, step, a, b)
}

func horseJumpF(i, j, step, a, b int) int {
	if i < 0 || i > 9 || j < 0 || j > 8 {
		return 0
	}
	if step == 0 {
		if i == a && j == b {
			return 1
		}
		return 0
	}
	return horseJumpF(i-2, j+1, step-1, a, b) +
		horseJumpF(i-1, j+2, step-1, a, b) +
		horseJumpF(i+1, j+2, step-1, a, b) +
		horseJumpF(i+2, j+1, step-1, a, b) +
		horseJumpF(i+2, j-1, step-1, a, b) +
		horseJumpF(i+1, j-2, step-1, a, b) +
		horseJumpF(i-1, j-2, step-1, a, b) +
		horseJumpF(i-2, j-1, step-1, a, b)
}

// WaysDP 与 HorseJumpDP 等价，保留 Java 课程中的第二组动态规划入口。
// dp[i][j][step] 表示从 (i,j) 出发还剩 step 步时到目标点的方法数。
//
// 时间复杂度：O(10*9*step)。
// 空间复杂度：O(10*9*step)。
func WaysDP(a, b, step int) int {
	dp := make([][][]int, 10)
	for i := range dp {
		dp[i] = make([][]int, 9)
		for j := range dp[i] {
			dp[i][j] = make([]int, step+1)
		}
	}
	dp[a][b][0] = 1
	for curStep := 1; curStep <= step; curStep++ {
		for i := 0; i < 10; i++ {
			for j := 0; j < 9; j++ {
				dp[i][j][curStep] = horseJumpGetValue(dp, i-2, j+1, curStep-1) +
					horseJumpGetValue(dp, i-1, j+2, curStep-1) +
					horseJumpGetValue(dp, i+1, j+2, curStep-1) +
					horseJumpGetValue(dp, i+2, j+1, curStep-1) +
					horseJumpGetValue(dp, i+2, j-1, curStep-1) +
					horseJumpGetValue(dp, i+1, j-2, curStep-1) +
					horseJumpGetValue(dp, i-1, j-2, curStep-1) +
					horseJumpGetValue(dp, i-2, j-1, curStep-1)
			}
		}
	}
	return dp[0][0][step]
}

func horseJumpGetValue(dp [][][]int, i, j, step int) int {
	if i < 0 || i > 9 || j < 0 || j > 8 {
		return 0
	}
	return dp[i][j][step]
}
