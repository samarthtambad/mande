package keeper

import (
	"github.com/mande-labs/mande/x/mande/types"
)

var _ types.QueryServer = Keeper{}
