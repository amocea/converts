package converts

import "reflect"

/*
 * @Author: Amour
 * @Email: chenyuhan@mobvista.com
 * @Date: 2024/8/22 14:20:35
 * @LastEditTime: 2024/8/22 14:20:35
 * @FilePath: /converts//internal.go
 */
const (
	key = iota
	value
)

func getPointerTo(v any) reflect.Value {
	target := reflect.ValueOf(v)
	for target.Kind() == reflect.Ptr {
		target = target.Elem()
	}

	return target
}

func mustBeMap(val reflect.Value) bool {
	return val.Kind() == reflect.Map
}

func getMapAttrType(val reflect.Value, t int) reflect.Type {
	typ := val.Type()

	var res reflect.Type
	switch t {
	case key:
		res = typ.Key()
	case value:
		res = typ.Elem()
	}

	return res
}

func IsPointerToPointer(v reflect.Value) bool {
	i := 1
	for v.Kind() == reflect.Ptr {
		i++
		v = v.Elem()
	}

	return i > 1
}

func mustBeSlice(v reflect.Value) bool {
	return v.Kind() == reflect.Slice
}

func GetArrEleType(v any) (reflect.Type, error) {
	val := reflect.ValueOf(v)
	if !mustBeSlice(val) {
		return nil, ErrUnsupportedType
	}

	return val.Type().Elem(), nil
}
