package infracore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringTo(t *testing.T) {
	// 测试字符串转换
	var strVal string
	ConvertStringTo("test", &strVal)
	assert.Equal(t, "test", strVal)

	// 测试整数转换
	var intVal int
	ConvertStringTo("123", &intVal)
	assert.Equal(t, 123, intVal)

	// 测试浮点数转换
	var floatVal float64
	ConvertStringTo("1.23", &floatVal)
	assert.Equal(t, 1.23, floatVal)

	// 测试布尔值转换
	var boolVal bool
	ConvertStringTo("true", &boolVal)
	assert.Equal(t, true, boolVal)
}
