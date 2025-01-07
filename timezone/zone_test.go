package timezone

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTimeZone(t *testing.T) {
	// 准备测试数据
	timeZoneName := "Asia/Taipei"
	expectedLocation, _ := time.LoadLocation(timeZoneName)

	// 添加时区到映射
	timeZoneMap[timeZoneName] = expectedLocation

	// 执行测试函数
	actualLocation := GetTimeZone(timeZoneName)

	// 断言结果
	assert.Equal(t, expectedLocation, actualLocation)

	// 测试不存在的时区
	nonExistentTimeZoneName := "NonExistentTimeZone"
	actualLocation = GetTimeZone(nonExistentTimeZoneName)
	assert.Nil(t, actualLocation)
}
