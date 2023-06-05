package consts

import "time"

const (
	ServerShutdownWaitTime = time.Second * 1 // 5s
	DebugMode              = "debug"
	TestMode               = "test"
	ReleaseMode            = "release"
)
