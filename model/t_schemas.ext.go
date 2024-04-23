package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type CheckService interface {
	check(client TClientInfo) bool
}

type GeoResponse struct {
	Organization    string  `json:"organization"`
	Longitude       float64 `json:"longitude"`
	Timezone        string  `json:"timezone"`
	Isp             string  `json:"isp"`
	Offset          int     `json:"offset"`
	Asn             int     `json:"asn"`
	AsnOrganization string  `json:"asn_organization"`
	Country         string  `json:"country"`
	Ip              string  `json:"ip"`
	Latitude        float64 `json:"latitude"`
	ContinentCode   string  `json:"continent_code"`
	CountryCode     string  `json:"country_code"`
}

func (s TSchema) checkGeo(c *TClientInfo) {
	// 调用接口获取ip地址
	url := "https://api.ip.sb/geoip/" + c.IP
	client := &http.Client{}
	fmt.Println(url)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla")
	response, err := client.Do(request)
	if err != nil {
		c.Status = 1
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)
	// 解析json
	res, err := io.ReadAll(response.Body)
	if err != nil {
		c.Status = 1
		return
	}
	fmt.Print(string(res))
	var geo GeoResponse
	err = json.Unmarshal(res, &geo)
	fmt.Printf("geo: %v\n", geo)

	// 判断是否允许访问
	c.CountryCode = geo.CountryCode
	c.ContinentCode = geo.ContinentCode
	c.Country = geo.Country
	c.Latitude = geo.Latitude
	c.Isp = geo.Isp
	c.Organization = geo.Organization
	c.Longitude = geo.Longitude
	c.Latitude = geo.Latitude
	status := strings.Contains(s.AllowRegion, geo.CountryCode)
	if status {
		c.Status = 1
		return
	}
	c.Status = 2
	return
}

func (s TSchema) Check(c *TClientInfo) bool {
	s.checkGeo(c)
	c.Insert()
	return c.Status == 1
}
