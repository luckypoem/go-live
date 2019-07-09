package conf

import "github.com/BurntSushi/toml"

type Config struct {
	MySQLDSN string `toml:"MYSQL_DSN"`
}

func LoadConfigFromFile(filepath string) (*Config, error) {
	var conf Config

	_, err := toml.DecodeFile(filepath, &conf)

	if err != nil {
		return nil, err
	}

	return &conf, nil
}
