package keeper_test

import (
	"testing"

	testkeeper "github.com/mande-labs/mande/testutil/keeper"
	"github.com/mande-labs/mande/x/mande/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MandeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
