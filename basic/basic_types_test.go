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
