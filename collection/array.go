package collection

import "cmp"

func Partition[T any](arr []T, size int) [][]T {
	result := make([][]T, 0)
	length := len(arr)
	if length <= 0 {
		return result
	}
	outSize := length / size
	for i := 0; i <= outSize; i++ {
		innerArr := make([]T, 0)
		for j := i * size; j < min((i+1)*size, length); j++ {
			innerArr = append(innerArr, arr[j])
		}
		if len(innerArr) > 0 {
			result = append(result, innerArr)
		}
	}
	return result
}

func GroupBy[T any, R cmp.Ordered](arr []T, function func(T) R) map[R][]T {
	result := map[R][]T{}
	for _, ele := range arr {
		k := function(ele)
		result[k] = append(result[k], ele)
	}
	return result
}

// ArrayToMap 将数组转换为 Map
// arr 数组
// keyFunc 取数组中哪个字段作为 map 的 key
// coverExists 如果有相同的 key 元素, 是否用后面的覆盖前面的
func ArrayToMap[T any, R cmp.Ordered](arr []T, coverExists bool, keyFunc func(T) R) map[R]T {
	result := map[R]T{}
	for _, ele := range arr {
		k := keyFunc(ele)
		_, ok := result[k]
		if !ok || coverExists {
			result[k] = ele
		}
	}
	return result
}
