package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedMapTraversal(t *testing.T) {
	// 测试正常情况
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	var result []string
	SortedMapTraversal(m, false, func(k int, v string) {
		result = append(result, v)
	})
	assert.Equal(t, []string{"one", "two", "three"}, result)

	// 测试反转情况
	result = []string{}
	SortedMapTraversal(m, true, func(k int, v string) {
		result = append(result, v)
	})
	assert.Equal(t, []string{"three", "two", "one"}, result)

	// 测试空 map
	m = map[int]string{}
	result = []string{}
	SortedMapTraversal(m, false, func(k int, v string) {
		result = append(result, v)
	})
	assert.Equal(t, []string{}, result)
}
