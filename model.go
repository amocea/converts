package converts

import (
	"fmt"
	"reflect"
)

/*
 * @Author: Amour
 * @Email: chenyuhan@mobvista.com
 * @Date: 2024/8/22 14:14:48
 * @LastEditTime: 2024/8/22 14:14:48
 * @FilePath: /converts//model.go
 */

type RConvert struct {
	v any
}

// Interface 直接返回原值
func (r *RConvert) Interface() interface{} {
	return r.v
}

// InterfaceSlice 返回 []interface{}
func (r *RConvert) InterfaceSlice() []interface{} {
	var res []interface{}

	if r.v != nil {
		val := reflect.ValueOf(r.v)

		fn := func() {
			res = append(res, val.Interface())
		}

		arrFn := func() {
			res = make([]interface{}, val.Len())
			for i := 0; i < val.Len(); i++ {
				res[i] = val.Index(i).Interface()
			}
		}

		NewCondChecker(r.v, nil).
			AllBuiltinBasicType(fn).
			Ptr(fn).
			Slice(arrFn).
			Array(arrFn)
	}

	return res
}

// StringSlice 返回 []string
func (r *RConvert) StringSlice() []string {
	var res []string

	arr := r.InterfaceSlice()

	res = make([]string, len(arr))

	for idx, v := range arr {
		tmp := RConvert{v}

		str, err := tmp.String()
		if err != nil {
			continue
		}

		res[idx] = str
	}

	return res
}

// 返回 string 类型
func (r *RConvert) String() (string, error) {
	var res string
	var err error
	target := r.v

	NewCondChecker(r.v, nil).
		String(func() {
			res = target.(string)
		}).
		AllInt(func() {
			res = fmt.Sprintf("%d", target)
		}).
		Float(func() {
			res = fmt.Sprintf("%f", target)
		}).
		Bool(func() {
			res = fmt.Sprintf("%t", target)
		}).
		Complex(func() {
			res = fmt.Sprintf("%c", target)
		}).
		Default(func() {
			typ := reflect.TypeOf(res)
			val := reflect.ValueOf(res)
			var isEnd bool
			for !isEnd {
				switch typ.Kind() {
				case reflect.Slice, reflect.Array, reflect.Map:
					err = ErrUnsupportedType
					isEnd = true
				case reflect.Ptr:
					val = getPointerTo(target)
					typ = val.Type()
				case reflect.Struct:
					// 调用 string 方法
					callResult := val.MethodByName("String").Call([]reflect.Value{})

					if len(callResult) > 0 {
						var ok bool
						res, ok = callResult[0].Interface().(string)
						if !ok {
							err = ErrUnsupportedType
						}
					}
					isEnd = true
				default:
					isEnd = true
					err = ErrUnsupportedType
				}
			}
		})

	return res, err
}
