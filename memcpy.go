package main

//#include <string.h>
import (
	"C"
	"unsafe"
)

func memcpy(dest, src []byte) int {
	n := len(src)
	if len(dest) < len(src) {
		n = len(dest)
	}
	if n == 0 {
		return 0
	}
	C.memcpy(unsafe.Pointer(&dest[0]), unsafe.Pointer(&src[0]), C.size_t(n))
	return n
}
