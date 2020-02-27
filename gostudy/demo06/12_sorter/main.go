package main

import "fmt"

var arr = []int{3, 6, 4, 2, 11, 10, 5}

func main() {
	fmt.Println(arr)
	fmt.Println(BubbleSort(arr))
	fmt.Println(SelectSort(arr))
	fmt.Println(InsertSort(arr))
	fmt.Println(ShellSort(arr))
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// BubbleSort 两个数比较大小，较大的数下沉，较小的数冒起来。
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// SelectSort 在长度为N的无序数组中，第一次遍历n-1个数，找到最小的数值与第一个元素交换，
// 第二次遍历n-2个数，找到最小的数值与第二个元素交换。。。第n-1次遍历，找到最小的数值与第n-1
// 个元素交换，排序完成。
func SelectSort(arr []int) []int {
	var min int
	for i := 0; i < len(arr)-1; i++ {
		min = 1
		for j := 1 + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != 1 {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

// InsertSort 在要排序的一组数中，假定前n-1个数已经排好序，现在将第n个数插到前面的有序数列中，
// 使得这n个数也是排好顺序的。如此反复循环，直到全部排号顺序。
func InsertSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}
	return arr
}

// QuickSort 选择一个基准数，通过一趟排序将要排序的数据分割成独立的两部分；其中一部分的所有数据
// 都比另外一部分的所有数据都要小。然后，再按此方法对这两部分数据分别进行快速排序，整个排序过程
// 可以递归进行，以此达到整个数据变成有序序列。
func QuickSort(arr []int, left, right int) {
	// 只剩一个元素时就返回了
	if left >= right {
		return
	}
	// 标记最左侧元素作为参考
	tmp := arr[left]
	// 两个游标分别从两端相向移动，寻找合适的"支点"
	i, j := left, right
	for i < j {
		// 右边的游标向左移动，直到找到比参考的元素值小的
		for arr[j] >= tmp && i < j {
			j--
		}
		// 左侧游标向右移动，直到找到比参考元素值大的
		for arr[i] <= tmp && i < j {
			i++
		}

		// 如果找到的两个游标位置不统一，就游标位置元素的值，并继续下一轮寻找
		// 此时交换的左右位置的值，右侧一定不大于左侧。可能相等但也会交换位置，所以才叫不稳定的排序算法
		arr[i], arr[j] = arr[j], arr[i]
	}

	// 这时的left位置已经是我们要找的支点了，交换位置
	arr[left], arr[i] = arr[i], tmp

	// 按支点位置吧原数列分成两段，再各自逐步缩小范围排序
	QuickSort(arr, left, i-1)
	QuickSort(arr, i+1, right)
}

// ShellSort 在要排序的一组数中，根据某一增量分为若干子序列，并对子序列分别进行插入排序。
// 然后逐渐将增量减小,并重复上述过程。直至增量为1,此时数据序列基本有序,最后进行插入排序。
func ShellSort(arr []int) []int {
	increment := len(arr)
	for {
		increment = increment / 2
		for i := 0; i < increment; i++ {
			for j := i + increment; j < len(arr); j = j + increment {
				for k := j; k > i; k = k - increment {
					if arr[k] < arr[k-increment] {
						arr[k], arr[k-increment] = arr[k-increment], arr[k]
					} else {
						break
					}
				}
			}
		}
		if increment == 1 {
			break
		}
	}
	return arr
}
