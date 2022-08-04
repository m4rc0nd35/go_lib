package main

/*
#cgo CFLAGS: -I./src
#cgo LDFLAGS: -L./lib -lhello
#include "hello.h"
*/
import "C"

func main() {
	// #cgo CFLAGS: -DPNG_DEBUG=1
	// #cgo amd64 386 CFLAGS: -DX86=1
	// #cgo LDFLAGS: -lpng
	// #include <png.h>
	// cs := C.CString("Hello from stdio")
	C.hello()
	// C.free(unsafe.Pointer(cs))
}
