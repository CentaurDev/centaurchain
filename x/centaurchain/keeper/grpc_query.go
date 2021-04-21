package keeper

import (
	"github.com/CentaurDev/centaurchain/x/centaurchain/types"
)

var _ types.QueryServer = Keeper{}
