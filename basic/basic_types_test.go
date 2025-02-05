package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringTo(t *testing.T) {
	// 测试字符串转换
	var strVal string
	_ = ConvertStringTo("test", &strVal)
	assert.Equal(t, "test", strVal)

	// 测试整数转换
	var intVal int
	_ = ConvertStringTo("123", &intVal)
	assert.Equal(t, 123, intVal)

	var int32Val int32
	_ = ConvertStringTo("123", &int32Val)
	assert.Equal(t, int32(123), int32Val)

	var int64Val int64
	_ = ConvertStringTo("123", &int64Val)
	assert.Equal(t, int64(123), int64Val)

	// 测试浮点数转换
	var floatVal32 float32
	_ = ConvertStringTo("1.23", &floatVal32)
	assert.Equal(t, float32(1.23), floatVal32)

	// 测试浮点数转换
	var floatVal float64
	_ = ConvertStringTo("1.23", &floatVal)
	assert.Equal(t, 1.23, floatVal)

	// 测试布尔值转换
	var boolVal bool
	_ = ConvertStringTo("true", &boolVal)
	assert.Equal(t, true, boolVal)
}

func TestConvertArrayTo(t *testing.T) {
	// 测试数组转换
	arr, _ := ConvertArrayTo[int]([]string{"1", "2", "3"})
	assert.Equal(t, []int{1, 2, 3}, arr)

	arr2, _ := ConvertArrayTo[int8]([]string{"1", "2", "3"})
	assert.Equal(t, []int8{1, 2, 3}, arr2)

	arr3, _ := ConvertArrayTo[int16]([]string{"1", "2", "3"})
	assert.Equal(t, []int16{1, 2, 3}, arr3)

	arr4, _ := ConvertArrayTo[int32]([]string{"1", "2", "3"})
	assert.Equal(t, []int32{1, 2, 3}, arr4)

	arr5, _ := ConvertArrayTo[int64]([]string{"1", "2", "3"})
	assert.Equal(t, []int64{1, 2, 3}, arr5)

	arr6, _ := ConvertArrayTo[string]([]string{"1", "2", "3"})
	assert.Equal(t, []string{"1", "2", "3"}, arr6)

	arr7, _ := ConvertArrayTo[float32]([]string{"1", "2", "3"})
	assert.Equal(t, []float32{1, 2, 3}, arr7)

	arr8, _ := ConvertArrayTo[float64]([]string{"1", "2", "3"})
	assert.Equal(t, []float64{1, 2, 3}, arr8)

	arr9, _ := ConvertArrayTo[bool]([]string{"true", "false"})
	assert.Equal(t, []bool{true, false}, arr9)
}
