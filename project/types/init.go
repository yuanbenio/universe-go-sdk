package types

import (
	kcfg "p/config"

	tlog "github.com/tendermint/tmlibs/log"

	"github.com/json-iterator/go"
)

var (
	json   = jsoniter.ConfigCompatibleWithStandardLibrary
	cfg    = kcfg.GetCfg()
	logger tlog.Logger
)