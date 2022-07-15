package sort

import "fmt"

func ShellSort(array []int) {
	if nil == array || len(array) == 0 {
		return
	}

	size := len(array)
	step := size
	for {
		if step == 1 {
			break
		}
		step = step / 2
	}

}

func shellSort(arr []int) []int {
	length := len(arr)
	gap := 1
	for gap < length/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
	return arr
}

func QuickSort(array []int, begin, end int) {
	if nil == array || len(array) == 0 {
		return
	}
	if begin >= end {
		return
	}
	//make a temporary choice 5, 2, 4, 6, 1, 3
	mid := begin + (end-begin)/2
	i := begin
	j := begin

	pivot := array[mid]
	array[mid], array[end] = array[end], array[mid]

	for j = begin; j < end; j++ {
		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i = i + 1
		}
	}
	array[i], array[end] = array[end], array[i]

	QuickSort(array, begin, i-1)
	QuickSort(array, i+1, end)

}

func SelectSort(array []int) {
	if nil == array || len(array) == 0 {
		return
	}
	length := len(array)
	for i := 0; i < length-1; i++ {
		min := i
		for k := i + 1; k < length; k++ {
			if array[k] < array[min] {
				min = k
			}
		}
		array[i], array[min] = array[min], array[i]
	}

}

//本人解答
func InsertSort(array []int) {
	if nil == array || len(array) == 0 {
		return
	}
	length := len(array)
	for i := 1; i < length; i++ {
		current := array[i]
		k := i - 1
		for ; k > 0; k-- {
			if array[k] > current {
				array[k+1] = array[k-1]
			}
		}
		array[k+1] = current
	}

}

//1,2,3,5,5,5,8,9)找到第一个>=5   2)最后一个<5  3)第一个 >5  4)最后一个<=5
func BinarySearch1(array []int) int {
	if nil == array || len(array) == 0 {
		return -1
	}
	l := -1
	n := len(array)
	for (l + 1) != n {
		mid := l + (n-l)/2
		if array[mid] < 5 {
			l = mid
		} else {
			n = mid
		}
	}
	return n
}

//1,2,3,5,5,5,8,9)找到第一个>=5   2)最后一个<5  3)第一个 >5  4)最后一个<=5
func BinarySearch2(array []int) int {
	if nil == array || len(array) == 0 {
		return -1
	}
	l := -1
	n := len(array)
	for (l + 1) != n {
		mid := l + (n-l)/2
		if array[mid] < 5 {
			l = mid
		} else {
			n = mid
		}
	}
	return l
}

//1,2,3,5,5,5,8,9)找到第一个>=5   2)最后一个<5  3)第一个 >5  4)最后一个<=5
func BinarySearch3(array []int) int {
	if nil == array || len(array) == 0 {
		return -1
	}
	l := -1
	n := len(array)
	for (l + 1) != n {
		mid := l + (n-l)/2
		if array[mid] <= 5 {
			l = mid
		} else {
			n = mid
		}
	}
	return n
}

//1,2,3,5,5,5,8,9)找到第一个>=5   2)最后一个<5  3)第一个 >5  4)最后一个<=5
func BinarySearch4(array []int) int {
	if nil == array || len(array) == 0 {
		return -1
	}
	l := -1
	n := len(array)
	for (l + 1) != n {
		mid := l + (n-l)/2
		if array[mid] <= 5 {
			l = mid
		} else {
			n = mid
		}
	}
	return l
}

//网上解答 5, 2, 4, 6, 1, 3
func InsertionSort2(arr []int) []int {
	for i := range arr {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex = preIndex - 1
		}
		arr[preIndex+1] = current
	}
	return arr
}

func Bubble_Sort(arr []int) {
	if nil == arr || len(arr) == 0 {
		return
	}
	len := len(arr)
	for i := 0; i < len; i++ {
		for k := 0; k < len-1-i; k++ {
			if arr[k] > arr[k+1] {
				arr[k], arr[k+1] = arr[k+1], arr[k]
			}
		}
	}
}
func TestSort() {
	array := []int{9, 4, 1}
	Bubble_Sort(array)
	fmt.Println(array)

	x := []int{5, 2, 4, 6, 1, 3}
	InsertionSort2(x)
	fmt.Println(x)

	x1 := []int{5, 2, 4, 6, 1, 3}
	InsertSort(x1)
	fmt.Println(x1)

	x2 := []int{1, 2, 3, 5, 5, 5, 8, 9}
	ind := BinarySearch1(x2)
	fmt.Println(ind)
	//最后一个<5
	ind = BinarySearch2(x2)
	fmt.Println(ind)
	//第一个 >5
	ind = BinarySearch3(x2)
	fmt.Println(ind)
	//第一个 >5
	ind = BinarySearch4(x2)
	fmt.Println(ind)
}
