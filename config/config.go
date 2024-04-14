package config

import (
	"flag"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	TransactionLogger TransactionLoggerConfig `toml:"TransactionLogger"`
	Network           NetworkConfig           `toml:"Network"`
	Encryption        EncryptionConfig        `toml:"Encryption"`
	WebApp            WebAppConfig            `toml:"WebApp"`
	Balancer          BalancerConfig          `toml:"Balancer"`
	Security          SecurityConfig          `toml:"Security"`
	Logging           LoggingConfig           `toml:"Logging"`
}

type TransactionLoggerConfig struct {
	On              bool          `toml:"On"`
	BackupDirectory string        `toml:"BackupDirectory"`
	SyncBufferTime  time.Duration `toml:"SyncBufferTime"`
}

type NetworkConfig struct {
	GRPC string `toml:"GRPC"`
	REST string `toml:"FastHTTP"`
}

type EncryptionConfig struct {
	On  bool   `toml:"On"`
	Key string `toml:"Key"`
}

type WebAppConfig struct {
	On   bool   `toml:"On"`
	Host string `toml:"Host"`
	Logs bool   `toml:"Logs"`
}

type BalancerConfig struct {
	On           bool     `toml:"On"`
	BalancerOnly bool     `toml:"BalancerOnly"`
	Servers      []string `toml:"Servers"`
}

type SecurityConfig struct {
	MandatoryAuthorization bool `toml:"MandatoryAuthorization"`
	MandatoryEncryption    bool `toml:"MandatoryEncryption"`
}

type LoggingConfig struct {
	Level string `toml:"Level"`
}

var _noSecurity = SecurityConfig{
	MandatoryAuthorization: false,
	MandatoryEncryption:    false,
}

func New() (*Config, error) {
	flag.Parse()

	cfg := &Config{}
	_, err := toml.DecodeFile("config/config.toml", cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}

	return cfg, nil
}
