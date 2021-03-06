package config

import (
	"fmt"
	"io/ioutil"

	"github.com/pelletier/go-toml"

	"github.com/rs/zerolog/log"
)

var (
	DefaultConfigPath = "./config.toml"
)

// Config defines all necessary configuration parameters.
type Config struct {
	RPC    *RPCConfig    `toml:"rpc"`
	GRPC   *GRPCConfig   `toml:"grpc"`
	Custom *CustomConfig `toml:"custom"`
}

// RPCConfig contains the configuration of the RPC endpoint.
type RPCConfig struct {
	Address string `toml:"address"`
}

// GRPCConfig contains the configuration of the gRPC endpoint.
type GRPCConfig struct {
	Address string `toml:"address"`
}
type CustomConfig struct {
	Mnemonics     []string `toml:"mnemonics"`
	GasLimit      int64    `toml:"gas_limit"`
	FeeDenom      string   `toml:"fee_denom"`
	FeeAmount     int64    `toml:"fee_amount"`
	Memo          string   `toml:"memo"`
	ValidatorAddr string   `toml:"validatoraddr"`
}

// SetupConfig takes the path to a configuration file and returns the properly parsed configuration.
func Read(configPath string) (*Config, error) {
	if configPath == "" {
		return nil, fmt.Errorf("empty configuration path")
	}

	log.Debug().Msg("reading config file")

	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %s", err)
	}

	return ParseString(configData)
}

// ParseString attempts to read and parse  config from the given string bytes.
// An error reading or parsing the config results in a panic.
func ParseString(configData []byte) (*Config, error) {
	var cfg Config

	log.Debug().Msg("parsing config data")

	err := toml.Unmarshal(configData, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config: %s", err)
	}

	return &cfg, nil
}
