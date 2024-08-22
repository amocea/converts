package converts

import (
	"reflect"
)

/*
 * @Author: Amour
 * @Email: chenyuhan@mobvista.com
 * @Date: 2024/8/22 16:26:30
 * @LastEditTime: 2024/8/22 16:26:30
 * @FilePath: /converts//conds.go
 */

type CondChecker struct {
	flag bool
	v    any
	typ  reflect.Type
}

func NewCondChecker(v any, t reflect.Type) *CondChecker {
	var typ reflect.Type
	if t != nil {
		typ = t
	} else {
		typ = reflect.TypeOf(v)
	}

	return &CondChecker{
		flag: false,
		v:    v,
		typ:  typ,
	}
}

func (cc *CondChecker) AllInt(fn func()) *CondChecker {
	if !cc.flag {
		switch cc.typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fn()
			cc.flag = true
		default:
			cc.flag = false
		}

	}
	return cc
}

func (cc *CondChecker) Bool(fn func()) *CondChecker {
	if !cc.flag {
		switch cc.typ.Kind() {
		case reflect.Bool:
			fn()
			cc.flag = true
		default:
			cc.flag = false
		}

	}
	return cc
}

func (cc *CondChecker) Slice(fn func()) *CondChecker {
	if !cc.flag {
		switch cc.typ.Kind() {
		case reflect.Slice:
			fn()
			cc.flag = true
		default:
			cc.flag = false
		}

	}
	return cc
}

func (cc *CondChecker) Array(fn func()) *CondChecker {
	if !cc.flag {
		switch cc.typ.Kind() {
		case reflect.Array:
			fn()
			cc.flag = true
		default:
			cc.flag = false
		}

	}
	return cc
}

func (cc *CondChecker) Map(fn func()) *CondChecker {
	if !cc.flag {
		switch cc.typ.Kind() {
		case reflect.Map:
			fn()
			cc.flag = true
		default:
			cc.flag = false
		}

	}
	return cc
}

func (cc *CondChecker) Ptr(fn func()) *CondChecker {
	if !cc.flag {
		switch cc.typ.Kind() {
		case reflect.Ptr:
			fn()
			cc.flag = true
		default:
			cc.flag = false
		}

	}
	return cc
}

func (cc *CondChecker) Number(fn func()) *CondChecker {
	if !cc.flag {
		switch cc.typ.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64,
			reflect.Complex128, reflect.Complex64:
			fn()
			cc.flag = true
		default:
			cc.flag = false
		}

	}
	return cc
}

func (cc *CondChecker) Float(fn func()) *CondChecker {
	if !cc.flag {
		switch cc.typ.Kind() {
		case reflect.Float32, reflect.Float64:
			fn()
			cc.flag = true
		default:
			cc.flag = false
		}

	}
	return cc
}

func (cc *CondChecker) Complex(fn func()) *CondChecker {
	if !cc.flag {
		switch cc.typ.Kind() {
		case reflect.Complex64, reflect.Complex128:
			fn()
			cc.flag = true
		default:
			cc.flag = false
		}

	}
	return cc
}

func (cc *CondChecker) String(fn func()) *CondChecker {
	if !cc.flag {
		switch cc.typ.Kind() {
		case reflect.String:
			fn()
			cc.flag = true
		default:
			cc.flag = false
		}

	}
	return cc
}

func (cc *CondChecker) AllBuiltinBasicType(fn func()) *CondChecker {
	if !cc.flag {
		switch cc.typ.Kind() {
		case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64,
			reflect.Bool,
			reflect.Complex128, reflect.Complex64,
			reflect.Interface:
			fn()
			cc.flag = true
		default:

			// 判断一些 rune 类型
			_, cc.flag = cc.v.(rune)
		}

	}
	return cc
}

func (cc *CondChecker) Default(fn func()) *CondChecker {
	if !cc.flag {
		fn()
		cc.flag = true
	}
	return cc
}
