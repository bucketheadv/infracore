package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartition(t *testing.T) {
	// 测试空数组
	arr1 := []int{}
	expected1 := [][]int{}
	result1 := Partition(arr1, 2)
	assert.Equal(t, expected1, result1)

	// 测试数组长度小于分割大小
	arr2 := []int{1}
	expected2 := [][]int{{1}}
	result2 := Partition(arr2, 2)
	assert.Equal(t, expected2, result2)

	// 测试正常分割
	arr3 := []int{1, 2, 3, 4, 5}
	expected3 := [][]int{{1, 2}, {3, 4}, {5}}
	result3 := Partition(arr3, 2)
	assert.Equal(t, expected3, result3)

	// 测试分割大小为 1
	arr4 := []int{1, 2, 3, 4, 5}
	expected4 := [][]int{{1}, {2}, {3}, {4}, {5}}
	result4 := Partition(arr4, 1)
	assert.Equal(t, expected4, result4)
}

func TestGroupBy(t *testing.T) {
	// 测试用例 1
	arr1 := []int{1, 2, 3, 4, 5}
	expected1 := map[int][]int{
		1: {1},
		2: {2},
		3: {3},
		4: {4},
		5: {5},
	}
	result1 := GroupBy(arr1, func(t int) int { return t })
	assert.Equal(t, expected1, result1)

	// 测试用例 2
	arr2 := []string{"apple", "banana", "cherry", "apple", "banana"}
	expected2 := map[string][]string{
		"apple":  {"apple", "apple"},
		"banana": {"banana", "banana"},
		"cherry": {"cherry"},
	}
	result2 := GroupBy(arr2, func(t string) string { return t })
	assert.Equal(t, expected2, result2)

	// 测试用例 3
	arr3 := []struct {
		id   int
		name string
	}{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
		{1, "David"},
	}
	expected3 := map[int][]struct {
		id   int
		name string
	}{
		1: {
			{1, "Alice"},
			{1, "David"},
		},
		2: {
			{2, "Bob"},
		},
		3: {
			{3, "Charlie"},
		},
	}
	result3 := GroupBy(arr3, func(t struct {
		id   int
		name string
	}) int {
		return t.id
	})
	assert.Equal(t, expected3, result3)
}

func TestArrayToMap(t *testing.T) {
	// 测试用例 1：正常情况
	arr1 := []int{1, 2, 3, 4, 5}
	keyFunc1 := func(t int) int { return t }
	expectedMap1 := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
	}
	resultMap1 := ArrayToMap(arr1, false, keyFunc1)
	assert.Equal(t, expectedMap1, resultMap1)

	// 测试用例 2：有冲突的情况
	arr2 := []int{1, 2, 2, 3, 3, 3}
	keyFunc2 := func(t int) int { return t }
	expectedMap2 := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	resultMap2 := ArrayToMap(arr2, true, keyFunc2)
	assert.Equal(t, expectedMap2, resultMap2)

	// 测试用例 3：空数组
	arr3 := []int{}
	keyFunc3 := func(t int) int { return t }
	expectedMap3 := map[int]int{}
	resultMap3 := ArrayToMap(arr3, false, keyFunc3)
	assert.Equal(t, expectedMap3, resultMap3)
}
