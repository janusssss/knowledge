package algorithm

/*
基础比较排序
- 冒泡排序 (Bubble Sort)
- 选择排序 (Selection Sort)
- 插入排序 (Insertion Sort)
- 希尔排序 (Shell Sort)

 高效比较排序
- 快速排序 (Quick Sort)
- 归并排序 (Merge Sort)
- 堆排序 (Heap Sort)
- TimSort (Python/Java 内置)

 非比较排序
- 计数排序 (Counting Sort)
- 桶排序 (Bucket Sort)
- 基数排序 (Radix Sort)

 特殊用途排序
- 鸡尾酒排序 (Cocktail Shaker Sort)
- 梳排序 (Comb Sort)
- 鸽巢排序 (Pigeonhole Sort)
- 循环排序 (Cycle Sort)

 现代混合排序
- 内省排序 (Introsort - C++ std::sort)
- 块排序 (Block Sort)
- 平滑排序 (Smooth Sort)
*/

/*
冒泡排序

时间复杂度: O(n²)
空间复杂度: O(1)
*/
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

// 快速排序
func quickSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	current, pivot := -1, 0
	for i := range n {
		if arr[i] < arr[pivot] {
			current++
			arr[i], arr[current] = arr[current], arr[i]
			if current == pivot {
				pivot = i
			}
		}
	}

	current++
	arr[pivot], arr[current] = arr[current], arr[pivot]

	quickSort(arr[:current])
	quickSort(arr[current+1:])
}


// 标准Lomuto快速排序
func quickSortLomuto(arr []int) {
    if len(arr) <= 1 {
        return
    }
    quickSortHelper(arr, 0, len(arr)-1)
}

// 递归辅助函数
func quickSortHelper(arr []int, low, high int) {
    if low < high {
        // 分区并获取基准位置
        pi := partitionLomuto(arr, low, high)
        
        // 递归排序左右部分
        quickSortHelper(arr, low, pi-1)
        quickSortHelper(arr, pi+1, high)
    }
}

// Lomuto分区方案
func partitionLomuto(arr []int, low, high int) int {
    // 选择最后一个元素作为基准
    pivot := arr[high]
    
    // i 指向小于基准的区域的末尾
    i := low - 1
    
    // j 遍历所有元素（除了最后一个基准元素）
    for j := low; j < high; j++ {
        // 如果当前元素小于等于基准
        if arr[j] <= pivot {
            i++  // 扩大小于基准的区域
            arr[i], arr[j] = arr[j], arr[i]  // 交换元素
        }
    }
    
    // 将基准放到正确位置（i+1 是基准的最终位置）
    i++
    arr[i], arr[high] = arr[high], arr[i]
    return i 
}