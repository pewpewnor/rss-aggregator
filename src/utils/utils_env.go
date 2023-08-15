package utils

import "os"

func EnvIsProduction() bool {
	if os.Getenv("PRODUCTION") == "true" {
		return true
	}
	return false
}
