package butil

import "strconv"

//
// 类型互转
//

// StringToUint : Parse string s to uint
func StringToUint(s string) (uint, error) {
	u64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	rst := uint(u64)
	return rst, nil
}
