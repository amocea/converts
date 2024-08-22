package converts

import (
	"fmt"
	"reflect"
	"sync"
)

/*
 * @Author: Amour
 * @Email: chenyuhan@mobvista.com
 * @Date: 2024/8/22 14:40:30
 * @LastEditTime: 2024/8/22 14:40:30
 * @FilePath: /converts//resolve.go
 */

func init() {
	_ = RegisterKeyResolver("", &stringResolver{})
	_ = RegisterKeyResolver(0, &integerResolver{})
}

var (
	resolversCache []*resolverWrapper
	rwMutex        sync.RWMutex
)

type KeyResolver interface {
	SetKey(key reflect.Type)
	Support(v reflect.Type) bool
	Resolve(v []any) (*RConvert, error)
}

type resolverWrapper struct {
	kr KeyResolver
}

func RegisterKeyResolver(key any, kr KeyResolver) error {
	rwMutex.Lock()
	defer rwMutex.Unlock()

	keyVal := reflect.ValueOf(key)
	if IsPointerToPointer(keyVal) {
		return ErrUnsupportedType
	}

	kr.SetKey(keyVal.Type())
	target := resolverWrapper{kr}

	resolversCache = append(resolversCache, &target)

	return nil
}

func getKeyResolver(typ reflect.Type) KeyResolver {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	var res KeyResolver
	for _, r := range resolversCache {
		if r.kr.Support(typ) {
			res = r.kr
			break
		}
	}

	return res
}

// stringResolver: resolve the string type key
type stringResolver struct {
	key reflect.Type
}

func (sr *stringResolver) SetKey(key reflect.Type) {
	sr.key = key
}

func (sr *stringResolver) Support(v reflect.Type) bool {
	return sr.key == v
}

func (sr *stringResolver) Resolve(v []any) (*RConvert, error) {
	strArr := make([]string, 0, len(v))
	for _, tmp := range v {
		strArr = append(strArr, tmp.(string))
	}

	return &RConvert{strArr}, nil
}

// integerResolver: resolve the integer type key
type integerResolver struct {
}

func (ir *integerResolver) SetKey(key reflect.Type) {

}

// Support 支持所有整形值（包括有无整数的）
func (ir *integerResolver) Support(v reflect.Type) bool {
	var support bool

	NewCondChecker(nil, v).
		AllInt(func() {
			support = true
		})

	return support
}

// Resolve 按照不同的类型进行转换
func (ir *integerResolver) Resolve(v []any) (*RConvert, error) {
	fmt.Println(1)

	return &RConvert{nil}, nil
}
