package xml

/**
    @date: 2022/11/7
**/

import (
	"encoding/xml"
)

// Unmarshal 解析xml
func Unmarshal(data []byte, res interface{}) error {
	return xml.Unmarshal(data, res)
}
