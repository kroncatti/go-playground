package utils

import (
	"fmt"
	"runtime"
)

func TraceFunction() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fmt.Println(frame.Function)
}
