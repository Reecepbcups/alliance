package app

import (
	"github.com/cosmos/cosmos-sdk/std"

	appparams "github.com/terra-money/alliance/app/params"
)

var encodingConfig = makeEncodingConfig()

func GetEncodingConfig() appparams.EncodingConfig {
	return encodingConfig
}

// makeEncodingConfig creates an EncodingConfig.
func makeEncodingConfig() appparams.EncodingConfig {
	encodingConfig := appparams.MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
