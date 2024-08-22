package converts

/*
 * @Author: Amour
 * @Email: chenyuhan@mobvista.com
 * @Date: 2024/8/22 14:09:46
 * @LastEditTime: 2024/8/22 14:09:46
 * @FilePath: /converts//map.go
 */

// MapKey2Array
// @Design: convert map's key(any type) to array
func MapKey2Array(v any) (res *RConvert, e error) {
	if v == nil {
		return
	}

	// step=1; 获取底层实际对象
	obj := getPointerTo(v)

	// step=2; 判断结构类型是否为 map 集合
	if !mustBeMap(obj) {
		e = ErrUnsupportedType
	} else {
		// step=3; 获取 map 集合 key 的类型
		keyTyp := getMapAttrType(obj, key)

		// step=4; 根据 keyType 类型选择解析器
		resolver := getKeyResolver(keyTyp)

		if resolver == nil {
			e = ErrUnsupportedType
		} else {
			var keys []any

			// step=5; 将 map 提取出来转化为一个统一数组
			for _, k := range obj.MapKeys() {
				keys = append(keys, k.Interface())
			}
			res, e = resolver.Resolve(keys)
		}
	}

	return
}
