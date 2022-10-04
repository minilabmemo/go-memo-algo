package _01_two_sum

import (
	"fmt"
	"reflect"
)

//依序尋找對應要的目標值放入map key中，並把自己的索引放在值， 下一次找時，檢查自身值是不是在map 中， 就可以拿到這次找的索引跟之前紀錄的索引了。
//時間複雜度 空間複雜度Time: O(n) Space: O(n)

func twoSum(nums []int, target int) []int {
	res := []int{}

	m := make(map[int]int, 0)
	for i, v := range nums {
		if mi, ok := m[v]; ok {
			return []int{mi, i}
		} else {
			m[target-v] = i //需額外儲存n個數的對應map ，空間複雜度為O(n)
		}
	}
	return res

}

func StartTwoSum() {

	var target int
	var input, output, res []int

	input, target, output = testcast1()
	res = twoSum(input, target)
	fmt.Printf("res:%v \n", res)
	fmt.Printf("result:%v \n", reflect.DeepEqual(res, output))

}

func testcast1() (input []int, target int, output []int) {
	input = []int{2, 7, 11, 15}
	target = 9
	output = []int{0, 1}
	return input, target, output
}
