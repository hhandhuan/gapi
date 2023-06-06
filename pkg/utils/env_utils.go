package utils

import (
	"gapi/pkg/conf"
	"gapi/pkg/consts"
)

func IsDebugMode() bool {
	return conf.GetConfig().System.Env == consts.DebugMode
}

func IsTestMode() bool {
	return conf.GetConfig().System.Env == consts.TestMode
}

func IsReleaseMode() bool {
	return conf.GetConfig().System.Env == consts.ReleaseMode
}
