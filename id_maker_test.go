package id_maker

import (
	"fmt"
	"testing"
)

func TestGenID(t *testing.T) {
	m := make(map[int64]int32, 100)
	flake := NewSnowFlake(1024416752)
	for i := 0; i < 100; i++ {
		//time.Sleep(1 * time.Second)
		id := flake.GenSnowID()
		if _, ok := m[id]; ok {
			fmt.Println("重复了id：", id)
		} else {
			fmt.Println("初次id：", id)
		}
		m[id] = 1
	}
	fmt.Println("--------------------------")
	fmt.Println(len(m))
}
