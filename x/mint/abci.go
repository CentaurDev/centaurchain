package mint

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/CentaurDev/centaurchain/x/mint/keeper"
	"github.com/CentaurDev/centaurchain/x/mint/types"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// fetch stored minter & params
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)

	// recalculate inflation rate
	// totalStakingSupply := k.StakingTokenSupply(ctx)
	bondedRatio := k.BondedRatio(ctx)
	minter.Inflation = minter.NextInflationRate(params, bondedRatio)
	
	// minter.AnnualProvisions = minter.NextAnnualProvisions(params, totalStakingSupply, ctx)
	// if sdk.NewInt(int64(ctx.BlockHeight())).GT(sdk.NewInt(int64(10))) {
	// 	minter.AnnualProvisions = sdk.NewDec(int64(0));
	// } else {
	// 	minter.AnnualProvisions = sdk.NewDec(int64(163392857142857));
	// }

	if sdk.NewInt(int64(ctx.BlockHeight())).GT(sdk.NewInt(int64(params.BlocksPerYear)).Mul(sdk.NewInt(int64(3)))) {
		minter.AnnualProvisions = sdk.NewDec((sdk.NewInt(int64(params.BlocksPerYear)).Mul(sdk.NewInt(int64(5)))).Int64());
	} else if sdk.NewInt(int64(ctx.BlockHeight())).GT(sdk.NewInt(int64(params.BlocksPerYear)).Mul(sdk.NewInt(int64(2)))) {
		minter.AnnualProvisions = sdk.NewDec((sdk.NewInt(int64(params.BlocksPerYear)).Mul(sdk.NewInt(int64(10)))).Int64());
	} else if sdk.NewInt(int64(ctx.BlockHeight())).GT(sdk.NewInt(int64(params.BlocksPerYear)).Mul(sdk.NewInt(int64(1)))) {
		minter.AnnualProvisions = sdk.NewDec((sdk.NewInt(int64(params.BlocksPerYear)).Mul(sdk.NewInt(int64(20)))).Int64());
	} else {
		minter.AnnualProvisions = sdk.NewDec((sdk.NewInt(int64(params.BlocksPerYear)).Mul(sdk.NewInt(int64(40)))).Int64());
	}

	k.SetMinter(ctx, minter)

	// mint coins, update supply
	mintedCoin := minter.BlockProvision(params)
	mintedCoins := sdk.NewCoins(mintedCoin)

	err := k.MintCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	// send the minted coins to the fee collector account
	err = k.AddCollectedFees(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	if mintedCoin.Amount.IsInt64() {
		defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyBondedRatio, bondedRatio.String()),
			sdk.NewAttribute(types.AttributeKeyInflation, minter.Inflation.String()),
			sdk.NewAttribute(types.AttributeKeyAnnualProvisions, minter.AnnualProvisions.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
		),
	)
}
