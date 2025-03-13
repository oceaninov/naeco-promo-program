package gvars

import (
	logger "github.com/go-kit/kit/log"
	"sync"
)

const HashKeyMap = `hash_km`

var Log logger.Logger
var SyncMapHashStorage sync.Map
