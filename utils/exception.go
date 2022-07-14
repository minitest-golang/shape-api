package utils

import (
	"runtime/debug"
)

type Block struct {
	Try     func()
	Catch   func(string)
	Finally func()
}

func (block Block) Do() {
	if block.Finally != nil {
		defer func() {
			recover()
			block.Finally()
		}()
	}
	if block.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				ex := UNKNOWN_ERROR
				ErrorLog("Error: %v\n", r)
				if AppMode == "debug" {
					debug.PrintStack()
				}
				if _ex, ok := r.(string); ok {
					ex = _ex
				}
				block.Catch(ex)
			}
		}()
	}
	block.Try()
}

func ThrowException(e string) {
	panic(e)
}
