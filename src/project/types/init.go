package types

import (
	tlog "github.com/tendermint/tmlibs/log"

	"github.com/json-iterator/go"
)

var (
	json   = jsoniter.ConfigCompatibleWithStandardLibrary
	logger tlog.Logger
)
