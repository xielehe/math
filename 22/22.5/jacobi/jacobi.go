package jacobi

import (
	"errors"
	"fmt"

	"../../../lib/mymath"
)

//Jacobi 输入两个整数，输出bool值,true=>1,false=>-1
func Jacobi(up int64, down uint64) (int8, error) {
	if up == 1 {
		return 1, nil
	}
	if mymath.Even(down) {
		return -1, errors.New("b不能是偶数")
	}
	if up < 0 && up != -1 {
		r1, err1 := Jacobi(-1, down)
		r2, err2 := Jacobi(-up, down)
		if err1 == nil && err2 == nil {
			return r1 * r2, nil
		}
		fmt.Println(err1, err2)
	}
	if up != 2 && mymath.Even(uint64(up)) {
		next := up / 2
		r1, err1 := Jacobi(2, down)
		r2, err2 := Jacobi(next, down)
		if err1 == nil && err2 == nil {
			return r1 * r2, nil
		}
		fmt.Println(err1, err2)
	}
	if up == -1 {
		if down%4 == 1 {
			return 1, nil
		}
		return -1, nil
	} else if up == 2 {
		if r := down % 8; r == 1 || r == 7 {
			return 1, nil
		}
		return -1, nil
	} else if up%4 == 1 || down%4 == 1 {
		return Jacobi(int64(down%uint64(up)), uint64(up))
	} else if up%4 == 3 && down%4 == 3 {
		r1, err1 := Jacobi(-1, uint64(up))
		r2, err2 := Jacobi(int64(down%uint64(up)), uint64(up))
		if err1 == nil && err2 == nil {
			return r1 * r2, nil
		}
		fmt.Println(err1, err2)
	}
	return -1, errors.New("somewhere wrong")
}
