package internal

import (
	"github.com/guilhermealegre/go-clean-arch-core-lib/errors"
	"github.com/guilhermealegre/go-clean-arch-infrastructure-lib/errors/config"
)

// Generic error codes
var (
	ErrorGeneric                           = config.GetError("1", "%s", errors.Error)
	ErrorInvalidInputFields                = config.GetError("2", "The field: %s as invalid value: %v1", errors.Info)
	ErrorStateMachineCouldNotConvertObject = config.GetError("3", "State Machine: Could not convert object", errors.Info)
	ErrorNoTransaction                     = config.GetError("4", "No transaction has been initialized previously", errors.Error)
)
