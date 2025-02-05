package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringTo(t *testing.T) {
	// 测试字符串转换
	strVal, _ := StringTo[string]("test")
	assert.Equal(t, "test", strVal)

	// 测试整数转换
	intVal, _ := StringTo[int]("123")
	assert.Equal(t, 123, intVal)

	int32Val, _ := StringTo[int32]("123")
	assert.Equal(t, int32(123), int32Val)

	int64Val, _ := StringTo[int64]("123")
	assert.Equal(t, int64(123), int64Val)

	// 测试浮点数转换
	floatVal32, _ := StringTo[float32]("1.23")
	assert.Equal(t, float32(1.23), floatVal32)

	// 测试浮点数转换
	floatVal, _ := StringTo[float64]("1.23")
	assert.Equal(t, 1.23, floatVal)

	// 测试布尔值转换
	boolVal, _ := StringTo[bool]("true")
	assert.Equal(t, true, boolVal)
}

func TestConvertArrayTo(t *testing.T) {
	// 测试数组转换
	arr, _ := ArrayElemTo[int]([]string{"1", "2", "3"})
	assert.Equal(t, []int{1, 2, 3}, arr)

	arr2, _ := ArrayElemTo[int8]([]string{"1", "2", "3"})
	assert.Equal(t, []int8{1, 2, 3}, arr2)

	arr3, _ := ArrayElemTo[int16]([]string{"1", "2", "3"})
	assert.Equal(t, []int16{1, 2, 3}, arr3)

	arr4, _ := ArrayElemTo[int32]([]string{"1", "2", "3"})
	assert.Equal(t, []int32{1, 2, 3}, arr4)

	arr5, _ := ArrayElemTo[int64]([]string{"1", "2", "3"})
	assert.Equal(t, []int64{1, 2, 3}, arr5)

	arr6, _ := ArrayElemTo[string]([]string{"1", "2", "3"})
	assert.Equal(t, []string{"1", "2", "3"}, arr6)

	arr7, _ := ArrayElemTo[float32]([]string{"1", "2", "3"})
	assert.Equal(t, []float32{1, 2, 3}, arr7)

	arr8, _ := ArrayElemTo[float64]([]string{"1", "2", "3"})
	assert.Equal(t, []float64{1, 2, 3}, arr8)

	arr9, _ := ArrayElemTo[bool]([]string{"true", "false"})
	assert.Equal(t, []bool{true, false}, arr9)
}
