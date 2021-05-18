package lock

import (
	"fmt"
	"time"
)

var set = make(map[int]bool, 0)

func printOnce(num int) {
	if _, exist := set[num]; !exist {
		fmt.Println()
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go printOnce(100)
	}
	time.Sleep(time.Second)
}
//
//func NewGroup(s string, i int, getterFunc interface{}) {
//}
//
//func GetterFunc(f func(key string) ([]byte, error)) interface{} {
//
//}
