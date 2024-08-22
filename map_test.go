package converts

import (
	"fmt"
	"testing"
)

/*
 * @Author: Amour
 * @Email: chenyuhan@mobvista.com
 * @Date: 2024/8/22 14:10:56
 * @LastEditTime: 2024/8/22 14:10:56
 * @FilePath: /converts//map_test.go
 */
type User struct {
	Username string
	Password string
}

func TestMapKey2Array(t *testing.T) {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	m[3] = 3

	res, err := MapKey2Array(m)
	if err != nil {
		t.Fatal(err)
	}

	for _, tmp := range res.StringSlice() {
		fmt.Print(tmp, ", ")
	}
}
