package core

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestModelParse(t *testing.T) {
	FfmLoadModel("../data/train.log.model")
}

func TestSizeof(t *testing.T) {
	a := "abc"
	b := len(a)
	c := unsafe.Sizeof(a)
	fmt.Println(int(c))
	fmt.Println(a, b, c)
}
