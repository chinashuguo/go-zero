package rescue

import "github.com/chinashuguo/go-zero/core/logx"

func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		logx.ErrorStack(p)
	}
}
