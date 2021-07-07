package ai

// #cgo LDFLAGS: -L. -lailib
// #include "cpplib/ailib.h"
import "C"
import "unsafe"

type AI struct {
	ptr unsafe.Pointer
}

func NewAI(smartness int) AI {
	var cnChessAI AI
	cnChessAI.ptr = C.LIB_NewCNChessAI(C.int(smartness))
	return cnChessAI
}

func (ai AI) Plus(value int) int {
	return int(C.LIB_Plus(ai.ptr, C.int(value)))
}