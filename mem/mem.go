package main

/*
	#include <stdlib.h>
	#include <stdio.h>
	#include <string.h>
	struct Test {
		char *name;
		int age;
	};
 */
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	p := C.malloc(C.size_t(1))
	if p == nil {
		C.perror(C.CString("malloc failure"))
	}
	defer C.free(p)
	C.strcpy((*C.char)(p), (*C.char)(C.CString("123")))
	C.strcat((*C.char)(p), C.CString("233"))
	fmt.Printf(C.GoString((*C.char)(p)))

	ts := C.malloc(C.size_t(unsafe.Sizeof(C.struct_Test{})))
	if ts == nil {
		C.perror(C.CString("malloc struct"))
	}
	defer C.free(ts)
}