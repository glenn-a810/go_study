package initpkg

import "fmt"

var (
	a = c + b // a값은 c와 b가 초기화 된 다음 초기화
	b = f()   // b값은 4
	c = f()   // c값은 5
	d = 3     // d값은 초기화가 끝난 뒤 6
)

func init() {
	d++
	fmt.Println("init function", d)
}

func f() int {
	d++
	fmt.Println("f() d:", d)
	return d
}

func PrintD() {
	fmt.Println("d:", d)
}
