package ipipdontent

import (
	tp "github.com/henrylee2cn/teleport"
	"github.com/ipipdotnet/ipdb-go"
)

/**
    @date: 2022/9/15
**/

func GetUserCityToIp(ip string) []string {
	var (
		str []string
	)
	db, err := ipdb.NewCity("./ipipfreedb/ipipfree.ipdb")
	if err != nil {
		tp.Errorf("err-> %v", err)
		return str
	}
	cityInfo, err := db.FindInfo(ip, "CN")
	if err != nil {
		tp.Errorf("FindInfo err-> %v, ip-> %s", err, ip)
		return str
	}
	str = append(str, cityInfo.CityName, cityInfo.RegionName, cityInfo.CountryName)
	tp.Infof("CountryName-> %v,RegionName-> %v,CityName-> %v", cityInfo.CountryName, cityInfo.RegionName, cityInfo.CityName)
	return str
}
