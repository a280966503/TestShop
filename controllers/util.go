package controllers

import (
	"runtime"
)

func JsonResponse(flag bool,message string) map[string]interface{}  {
	resp := make(map[string]interface{})

	resp["flag"]=flag
	resp["message"]=message

	return resp
}

func IsLinux() bool  {
	switch runtime.GOOS {
	case "linux":
		return true
	default:
		return false
	}
}