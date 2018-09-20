package config

import (
	"os"
)

var IS_PRODUCTION = (os.Getenv("IS_PRODUCTION") == "true")
var CACHE_PASSWORD = os.Getenv("CACHE_PASSWORD")
