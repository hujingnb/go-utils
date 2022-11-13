// Package hcommon 工具方法
package hcommon

import (
	"github.com/hujingnb/go-utils/harray"
	"reflect"
)

// Copy 实现对任意类型进行深度复制
func Copy[T any](source T) T {
	structType := reflect.TypeOf(source)
	structValue := reflect.ValueOf(source)
	if structType == nil {
		return source
	}
	// 结果值
	resultValue := reflect.New(structType)
	changeValue := resultValue.Elem() // 层级与 structValue 保持一致
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
		structValue = structValue.Elem()
		// 若元素为指针, 对指针变量赋值
		changeValue.Set(reflect.New(structValue.Type()))
		changeValue = changeValue.Elem()
	}
	// 向结果中写入值
	switch structType.Kind() {
	case reflect.Struct: // 结构体
		for i := 0; i < structType.NumField(); i++ {
			tmpStructField := structType.Field(i)
			fieldName := tmpStructField.Name
			tmpStructFieldValue := structValue.FieldByName(fieldName)
			tmpRetField := changeValue.FieldByName(fieldName)
			// 空指针跳过
			if harray.InArray([]reflect.Kind{
				reflect.Ptr, reflect.Interface, reflect.Array, reflect.Slice, reflect.Map,
			}, tmpStructField.Type.Kind()) &&
				tmpStructFieldValue.IsNil() {
				continue
			}
			// 复制值
			tmpRet := Copy(tmpStructFieldValue.Interface())
			tmpRetField.Set(reflect.ValueOf(tmpRet))
		}
	case reflect.Interface:
		tmpRet := Copy(structValue.Interface())
		changeValue.Set(reflect.ValueOf(tmpRet))
	case reflect.Slice:
		changeValue.Set(reflect.MakeSlice(structType, structValue.Len(), structValue.Cap()))
		for i := 0; i < structValue.Len(); i++ {
			tmpRet := Copy(structValue.Index(i).Interface())
			changeValue.Index(i).Set(reflect.ValueOf(tmpRet))
		}
	case reflect.Map:
		changeValue.Set(reflect.MakeMap(structType))
		for _, key := range structValue.MapKeys() {
			tmpValue := Copy(structValue.MapIndex(key).Interface())
			tmpKey := Copy(key.Interface())
			changeValue.SetMapIndex(reflect.ValueOf(tmpKey), reflect.ValueOf(tmpValue))
		}
	case reflect.Chan:
		changeValue.Set(reflect.MakeChan(structType, structValue.Cap()))
	default:
		changeValue.Set(structValue)
	}
	return resultValue.Elem().Interface().(T)
}
