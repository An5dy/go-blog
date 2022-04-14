package helpers

import (
	mathrand "math/rand"
	"reflect"
	"time"
)

// Empty 判断值是否为空，类似 PHP empty() 函数
func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	reflectVal := reflect.ValueOf(val)
	switch reflectVal.Kind() {
	case reflect.String, reflect.Array:
		return reflectVal.Len() == 0
	case reflect.Map, reflect.Slice:
		return reflectVal.Len() == 0 || reflectVal.IsNil()
	case reflect.Bool:
		return !reflectVal.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflectVal.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return reflectVal.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return reflectVal.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return reflectVal.IsNil()
	}

	return reflect.DeepEqual(val, reflect.Zero(reflectVal.Type()).Interface())
}

// RandomString 生成长度为 length 的随机字符串
func RandomString(length int) string {
	mathrand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[mathrand.Intn(len(letters))]
	}
	return string(b)
}
