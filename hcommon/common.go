// Package hcommon 工具方法
package hcommon

import (
	"context"
	"github.com/hujingnb/go-utils/harray"
	"reflect"
	"unsafe"
)

// GetContextKeys 获取context中的所有key
// 注意, 获取的是系统 valueCtx 类中的内容, 若是自定义结构体无法识别
func GetContextKeys(ctx context.Context) []interface{} {
	ctxType := reflect.TypeOf(ctx)
	ctxValue := reflect.ValueOf(ctx)
	if ctxType == nil {
		return nil
	}
	if ctxType.Kind() == reflect.Ptr {
		ctxType = ctxType.Elem()
		ctxValue = ctxValue.Elem()
	}
	// emptyCtx 的 kind 为 int
	if ctxType.Kind() != reflect.Struct {
		return nil
	}
	// 当前结构体中存在 KV
	keys := make([]interface{}, 0)
	// 从 value 结构体中取出 KV
	if ctxType.Name() == "valueCtx" {
		tmpValue := ctxValue.FieldByName("key")
		if tmpValue.Kind() != reflect.Invalid { // 确保 key 存在
			// 因为 key 为结构体私有, 直接使用其地址
			tmpValue = reflect.NewAt(tmpValue.Type(), unsafe.Pointer(tmpValue.UnsafeAddr())).Elem()
			keys = append(keys, tmpValue.Interface())
		}
	}
	// 处理父类
	parentValue := ctxValue.FieldByName("Context")
	if parentValue.Kind() != reflect.Invalid {
		if parentCtx, ok := parentValue.Interface().(context.Context); ok {
			keys = append(keys, GetContextKeys(parentCtx)...)
		}
	}
	// 结果元素去重
	result := make([]interface{}, 0, len(keys))
	tmpMap := make(map[interface{}]bool)
	for _, item := range keys {
		if _, ok := tmpMap[item]; !ok {
			tmpMap[item] = true
			result = append(result, item)
		}
	}
	return result
}

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

// ReadChannelList 从 channel 列表中读取一个值
// 此方法用于 channel 数量不确定的情况
// 会使用反射机制, 比 select 效率要慢一些
// 参数:
// 	- channels: channel 列表
// 	- isBlock: 是否阻塞
// 返回值:
// 	- data: 读取到的值
// 	- ok: 是否成功读到数据
func ReadChannelList[T any](channels []chan T, isBlock bool) (data T, ok bool) {
	// 创建recv case
	var cases []reflect.SelectCase
	for _, ch := range channels {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}
	//  添加 default, 防止阻塞
	if !isBlock {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectDefault,
			Chan: reflect.Value{},
			Send: reflect.Value{},
		})
	}
	// 随机选择一个case
	_, recv, recvOK := reflect.Select(cases)
	if recvOK {
		var success bool
		data, success = recv.Interface().(T)
		if !success {
			panic("reflect error")
		}
		ok = true
		return
	} else {
		ok = false
		return
	}
}
