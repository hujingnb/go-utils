// Package hstruct 结构体操作
package hstruct

import (
	"errors"
	"reflect"
)

/*
ToMap 转为 map

参数如下:
 data: 结构体数据
 tag: 使用指定 tag 作为结果的 key
     若为空, 则使用字段名
*/
func ToMap(data interface{}, tag string) (map[string]interface{}, error) {
	structType, structValue, err := getReflect(data)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	// 遍历所有字段
	numField := structType.NumField()
	for i := 0; i < numField; i++ {
		field := structType.Field(i)
		fieldName := field.Name
		value := structValue.FieldByName(fieldName).Interface()
		key := fieldName
		// 指定 tag
		if tag != "" {
			tagValue := field.Tag.Get(tag)
			// 标签未定义, 跳过
			if tagValue == "" || tagValue == "-" {
				continue
			}
			key = tagValue
		}
		result[key] = value
	}
	return result, nil
}

// Name 获取结构体名称
func Name(data interface{}) (string, error) {
	structType, _, err := getReflect(data)
	if err != nil {
		return "", err
	}
	name := structType.Name()
	// 匿名结构体
	if name == "" {
		return "", errors.New("is anonymous struct")
	}
	return name, nil
}

// getReflect 获取结构体的反射类型
func getReflect(data interface{}) (reflect.Type, reflect.Value, error) {
	structType := reflect.TypeOf(data)
	structValue := reflect.ValueOf(data)
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
		structValue = structValue.Elem()
	}
	if structType.Kind() != reflect.Struct {
		return nil, reflect.Value{}, errors.New("is not struct")
	}
	return structType, structValue, nil
}
