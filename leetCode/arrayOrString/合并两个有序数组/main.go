package main

//输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//输出：[1,2,2,3,5,6]
//解释：需要合并 [1,2,3] 和 [2,5,6] 。
//合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。

//func merge(nums1 []int, m int, nums2 []int, n int) {
//	num := make([]int, 0)
//	num = append(num, nums1...)
//	num = append(num, nums2...)
//	sort.Ints(num)
//	fmt.Println(num)
//}
//
//func main() {
//	nums1 := []int{1, 2, 3, 0, 0, 0}
//	nums2 := []int{2, 5, 6}
//	m := 0
//	n := 0
//	n1 := []int{}
//	n2 := []int{}
//	for _, v := range nums1 {
//		if v != 0 {
//			m++
//			n1 = append(n1, v)
//		}
//	}
//	for _, v := range nums2 {
//		if v != 0 {
//			n++
//			n2 = append(n2, v)
//		}
//	}
//
//	merge(n1, m, n2, n)
//}

import (
	"fmt"
)

//func merge(nums1 []int, m int, nums2 []int, n int) {
//	copy(nums1[m:], nums2)
//	sort.Ints(nums1)
//}

// 逆向双指针
func merge(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}

func main() {
	//var str string
	//var str2 string
	//fmt.Scanln(&str)
	//fmt.Scanln(&str2)

	str := []int{1, 2, 3, 0, 0, 0}
	str2 := []int{2, 5, 6}
	var a int
	var b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	merge(str, a, str2, b)
}
