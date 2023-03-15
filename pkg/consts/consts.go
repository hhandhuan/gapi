package consts

import "time"

const (
	ServerShutdownWaitTime = time.Second * 5 // 5s
	DebugMode              = "debug"
	TestMode               = "test"
	ReleaseMode            = "release"
)
