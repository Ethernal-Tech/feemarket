// Package networksuite provides a base test suite for tests that need a local network instance
package networksuite

import (
	"math/rand"

	pruningtypes "cosmossdk.io/store/pruning/types"
	tmrand "github.com/cometbft/cometbft/libs/rand"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	"github.com/cosmos/gogoproto/proto"

	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	"cosmossdk.io/core/appconfig"
	"cosmossdk.io/depinject"
	"github.com/cosmos/cosmos-sdk/testutil/configurator"
	"github.com/skip-mev/chaintestutil/network"
	"github.com/skip-mev/chaintestutil/sample"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	feemarketmodule "github.com/skip-mev/feemarket/api/feemarket/feemarket/module/v1"
	"github.com/skip-mev/feemarket/tests/app"
	feemarkettypes "github.com/skip-mev/feemarket/x/feemarket/types"
)

var (
	chainID = "chain-" + tmrand.NewRand().Str(6)

	DefaultAppConstructor = func(val network.ValidatorI) servertypes.Application {
		return app.New(
			val.GetCtx().Logger,
			dbm.NewMemDB(),
			nil,
			true,
			simtestutil.EmptyAppOptions{},
			baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(val.GetAppConfig().Pruning)),
			baseapp.SetMinGasPrices(val.GetAppConfig().MinGasPrices),
			baseapp.SetChainID(chainID),
		)
	}
)

// NetworkTestSuite is a test suite for query tests that initializes a network instance.
type NetworkTestSuite struct {
	suite.Suite

	Network        *network.Network
	FeeMarketState feemarkettypes.GenesisState
}

// SetupSuite setups the local network with a genesis state.
func (nts *NetworkTestSuite) SetupSuite() {
	var (
		r = sample.Rand()
		// TODO: check this
		cfg = network.NewConfig(appConfig())
	)

	updateGenesisConfigState := func(moduleName string, moduleState proto.Message) {
		buf, err := cfg.Codec.MarshalJSON(moduleState)
		require.NoError(nts.T(), err)
		cfg.GenesisState[moduleName] = buf
	}

	// initialize fee market
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[feemarkettypes.ModuleName], &nts.FeeMarketState))
	nts.FeeMarketState = populateFeeMarket(r, nts.FeeMarketState)
	updateGenesisConfigState(feemarkettypes.ModuleName, &nts.FeeMarketState)

	nts.Network = network.New(nts.T(), cfg)
}

func populateFeeMarket(_ *rand.Rand, feeMarketState feemarkettypes.GenesisState) feemarkettypes.GenesisState {
	// TODO intercept and populate state randomly if desired
	return feeMarketState
}

func appConfig() depinject.Config {
	return configurator.NewAppConfig(
		configurator.AuthModule(),
		configurator.GenutilModule(),
		configurator.BankModule(),
		configurator.StakingModule(),
		configurator.MintModule(),
		configurator.DistributionModule(),
		configurator.GovModule(),
		configurator.ParamsModule(),
		configurator.SlashingModule(),
		configurator.FeegrantModule(),
		configurator.EvidenceModule(),
		configurator.AuthzModule(),
		configurator.GroupModule(),
		configurator.VestingModule(),
		configurator.NFTModule(),
		configurator.ConsensusModule(),
		configurator.TxModule(),
		FeemarketModule(),
	)
}

func FeemarketModule() configurator.ModuleOption {
	return func(config *configurator.Config) {
		config.ModuleConfigs[feemarkettypes.ModuleName] = &appv1alpha1.ModuleConfig{
			Name:   feemarkettypes.ModuleName,
			Config: appconfig.WrapAny(&feemarketmodule.Module{}),
		}
	}
}
