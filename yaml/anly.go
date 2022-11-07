package yaml

/**
    @date: 2022/11/7
**/

import (
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
)

// Unmarshal 解析yaml
func Unmarshal(data []byte, res interface{}) error {
	var (
		conf interface{}
	)
	yaml.Unmarshal(data, &conf)
	return mapstructure.Decode(conf, res)
}
