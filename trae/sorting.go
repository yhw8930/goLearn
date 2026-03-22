package main

// BubbleSort 冒泡排序
//
// 思路:
// 重复地遍历要排序的数列，一次比较两个元素，如果它们的顺序错误就把它们交换过来。
// 遍历数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。
//
// 复杂度分析:
// - 时间复杂度: O(n²)。在最好的情况下（数组已经排序），时间复杂度为 O(n)。
// - 空间复杂度: O(1)。这是一个原地排序算法。
// - 稳定性: 稳定。相等的元素不会改变它们的相对顺序。
func BubbleSort(arr []int) {
	n := len(arr)
	// 外层循环控制排序的轮数
	for i := 0; i < n-1; i++ {
		// 内层循环用于比较和交换
		// -i 是因为每轮排序后，末尾的元素已经是有序的了
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 如果前一个元素大于后一个元素，则交换
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// SelectionSort 选择排序
//
// 思路:
// 首先在未排序序列中找到最小（或最大）元素，存放到排序序列的起始位置，
// 然后再从剩余未排序元素中继续寻找最小（或最大）元素，然后放到已排序序列的末尾。
//
// 复杂度分析:
// - 时间复杂度: O(n²)。选择排序的性能与输入数据无关。
// - 空间复杂度: O(1)。这是一个原地排序算法。
// - 稳定性: 不稳定。在交换过程中，相等元素的相对位置可能会改变。
func SelectionSort(arr []int) {
	n := len(arr)
	// 外层循环遍历数组
	for i := 0; i < n-1; i++ {
		// 假设当前位置的元素是最小的
		minIndex := i
		// 内层循环从当前位置的下一个元素开始，寻找最小的元素
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				// 如果找到更小的元素，则更新最小元素的索引
				minIndex = j
			}
		}
		// 将找到的最小元素与当前位置的元素交换
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// InsertionSort 插入排序
//
// 思路:
// 构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
//
// 复杂度分析:
// - 时间复杂度: O(n²)。在最好的情况下（数组已经排序），时间复杂度为 O(n)。
// - 空间复杂度: O(1)。这是一个原地排序算法。
// - 稳定性: 稳定。相等的元素不会改变它们的相对顺序。
func InsertionSort(arr []int) {
	// 从第二个元素开始遍历
	for i := 1; i < len(arr); i++ {
		preIndex := i - 1
		current := arr[i] // 当前需要插入的元素
		// 在有序序列中从后向前扫描
		for preIndex >= 0 && arr[preIndex] > current {
			// 如果已排序的元素大于新元素，将该元素向后移动
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		// 将新元素插入到正确的位置
		arr[preIndex+1] = current
	}
}

// ShellSort 希尔排序
//
// 工作原理:
// 希尔排序是插入排序的一种更高效的改进版本。它通过比较相距一定间隔的元素来工作，
// 以此来弥补插入排序只交换相邻元素的弱点。
//  1. 选择一个增量序列 t1, t2, ..., tk，其中 ti > tj, tk = 1。
//  2. 按增量序列个数 k，对序列进行 k 趟排序。
//  3. 每趟排序，根据对应的增量 ti，将待排序列分割成若干长度为 m 的子序列，分别对各子表进行直接插入排序。
//     仅增量因子为 1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。
func ShellSort(arr []int) {
	n := len(arr)
	// 初始增量设置为数组长度的一半，然后逐步减半
	for gap := n / 2; gap > 0; gap /= 2 {
		// 从增量值开始，对每个子序列进行插入排序
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			// 在子序列内部进行插入排序
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
	}
}

// MergeSort 归并排序
//
// 思路 (分治法):
// 1. 分解：将 n 个元素的序列递归地分成两半。
// 2. 解决：当子序列无法再分割时，进行排序。
// 3. 合并：将已排序的子序列合并成一个最终的排序序列。
//
// 复杂度分析:
// - 时间复杂度: O(n log n)。无论最好、最坏还是平均情况，性能都非常稳定。
// - 空间复杂度: O(n)。需要额外的空间来存储合并过程中的临时数组。
// - 稳定性: 稳定。在合并过程中，可以保证相等元素的相对顺序不变。
func MergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	mid := len(arr) / 2
	left := make([]int, mid)
	right := make([]int, len(arr)-mid)
	copy(left, arr[:mid])
	copy(right, arr[mid:])

	MergeSort(left)
	MergeSort(right)
	merge(arr, left, right)
}

// merge 函数用于合并两个已排序的数组
func merge(arr, left, right []int) {
	i, j, k := 0, 0, 0
	// 比较左右两个子数组的元素，将较小的元素放入结果数组
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}
	// 如果左边数组还有剩余，直接复制到结果数组
	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}
	// 如果右边数组还有剩余，直接复制到结果数组
	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}

// QuickSort 快速排序
//
// 工作原理 (分治法):
//  1. 从数列中挑出一个元素，称为 “基准”（pivot）。
//  2. 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面。
//     这个称为分区（partition）操作。
//  3. 递归地把小于基准值元素的子数列和大于基准值元素的子数列排序。
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)  // 递归排序基准左边的部分
		quickSort(arr, pi+1, high) // 递归排序基准右边的部分
	}
}

// partition 函数用于分区操作，并返回基准的最终位置
func partition(arr []int, low, high int) int {
	pivot := arr[high] // 选择最后一个元素作为基准
	i := low - 1       // i 是小于基准的区域的边界
	for j := low; j < high; j++ {
		// 如果当前元素小于或等于基准
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // 将其移动到小于基准的区域
		}
	}
	// 将基准元素放到正确的位置
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// HeapSort 堆排序
//
// 工作原理:
//  1. 将初始待排序关键字序列(R1, R2, ..., Rn)构建成大顶堆，此堆为初始的无序区。
//  2. 将堆顶元素 R[1] 与最后一个元素 R[n] 交换，此时得到新的无序区(R1, R2, ..., Rn-1)和新的有序区(Rn),
//     且满足 R[1, 2, ..., n-1] <= R[n]。
//  3. 由于交换后新的堆顶 R[1] 可能违反堆的性质，因此需要对当前无序区(R1, R2, ..., Rn-1)调整为新堆，
//     然后再将 R[1] 与无序区最后一个元素 R[n-1] 交换，得到新的无序区和新的有序区。
//  4. 不断重复此过程直到有序区的元素个数为 n-1，则整个排序过程完成。
func HeapSort(arr []int) {
	n := len(arr)
	// 1. 构建大顶堆
	// 从最后一个非叶子节点开始，从下至上，从右至左调整结构
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 2. 一个个从堆顶取出元素
	for i := n - 1; i > 0; i-- {
		// 将当前堆顶元素（最大值）移动到数组末尾
		arr[0], arr[i] = arr[i], arr[0]
		// 对剩余的元素重新进行堆化
		heapify(arr, i, 0)
	}
}

// heapify 用于调整堆
// n 是堆的大小，i 是需要调整的节点的索引
func heapify(arr []int, n, i int) {
	largest := i     // 初始化最大值为根节点
	left := 2*i + 1  // 左子节点
	right := 2*i + 2 // 右子节点

	// 如果左子节点存在且大于根节点
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// 如果右子节点存在且大于目前的最大值
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大值不是根节点
	if largest != i {
		// 交换根节点和最大值节点
		arr[i], arr[largest] = arr[largest], arr[i]
		// 递归地调整受影响的子树
		heapify(arr, n, largest)
	}
}
