package basic

import (
	"cmp"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func StringTo[T cmp.Ordered | bool](v string) (T, error) {
	var t T
	typ := reflect.TypeOf(&t)
	val := reflect.ValueOf(&t)
	eleType := typ.Elem().Kind()
	switch eleType {
	case reflect.String:
		val.Elem().SetString(v)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return t, err
		}
		val.Elem().SetInt(i)
	case reflect.Float32:
		f, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return t, err
		}
		val.Elem().SetFloat(f)
	case reflect.Float64:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return t, err
		}
		val.Elem().SetFloat(f)
	case reflect.Bool:
		b, err := strconv.ParseBool(v)
		if err != nil {
			return t, err
		}
		val.Elem().SetBool(b)
	default:
		return t, errors.New(fmt.Sprintf("不支持的数据类型: %s\n", eleType))
	}
	return t, nil
}

func ArrayElemTo[T cmp.Ordered | bool](v []string) ([]T, error) {
	result := make([]T, 0)
	for _, v := range v {
		tmp, err := StringTo[T](v)
		if err != nil {
			return result, err
		}
		result = append(result, tmp)
	}
	return result, nil
}
