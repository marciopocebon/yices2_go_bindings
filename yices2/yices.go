package yices2

// #cgo CFLAGS: -g -fPIC
// #cgo LDFLAGS:  -lyices -lgmp
// #include <yices.h>
import "C"

func Version() string {
	return "just little steps for now"
}
