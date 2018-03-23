// +build !go1.9

package errors

import "runtime"

func Info(err error, cmd ...interface{}) *ErrorInfo {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		pc = 0
	}
	return &ErrorInfo{cmd: cmd, err: Err(err), pc: pc}
}

func InfoEx(calldepth int, err error, cmd ...interface{}) *ErrorInfo {
	pc, _, _, ok := runtime.Caller(calldepth)
	if !ok {
		pc = 0
	}
	return &ErrorInfo{cmd: cmd, err: Err(err), pc: pc}
}
