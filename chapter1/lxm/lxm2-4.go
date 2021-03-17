package main

func main() {
	nums := []int{2, 3, 8, 6, 1}
	ret := findInversionPairCnt(nums)
	print(ret)
}

/**
设A[1..n]是一个包含N个非负整数的数组。
如果在i<j的情况下，有A[i]>A[j]，则(i,j)就称为A中的一个逆序对(inversion)。
[2,3,8,6,1]的所有逆序对为<2,1>,<3,1>,<8,6>,<8,1>,<6,1>
求n个元素的任何排列中逆序对数量
通过归并排序实现
*/
func findInversionPairCnt(nums []int) int {
	return process(0, len(nums)-1, nums)
}
func process(l int, r int, nums []int) int {
	if l == r {
		return 0
	}
	mid := (r + l) / 2
	cntL := process(l, mid, nums)
	cntR := process(mid+1, r, nums)
	cntN := merge(l, mid, r, nums)
	return cntL + cntR + cntN
}
func merge(l int, mid int, r int, nums []int) int {
	helpArr := make([]int, r-l+1)
	index := 0
	p := l
	q := mid + 1
	ret := 0
	for p <= mid && q <= r {
		if nums[p] < nums[q] {
			helpArr[index] = nums[p]
			p += 1
		} else {
			//A[i]>A[j] i<j 满足条件的所有i的范围为[p,mid]
			if nums[p] > nums[q] {
				ret += mid - p + 1
			}
			helpArr[index] = nums[q]
			q += 1
		}
		index += 1
	}
	for p <= mid {
		helpArr[index] = nums[p]
		p += 1
		index += 1
	}
	for q <= r {
		helpArr[index] = nums[q]
		q += 1
		index += 1
	}
	for i := l; i <= r; i++ {
		nums[i] = helpArr[i-l]
	}
	return ret
}
