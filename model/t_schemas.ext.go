package model

import (
	"encoding/json"
	"fmt"
	"hs_short_link/common"
	"io"
	"net/http"
	"strings"
	"time"
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

type CallCloakDTO struct {
	ClientIp       string `json:"clientIp"`
	UserAgent      string `json:"userAgent"`
	VisitUrl       string `json:"visitUrl"`
	ClientLanguage string `json:"clientLanguage"`
	Referer        string `json:"referer"`
	Timestamp      int64  `json:"timestamp"`
	Sign           string `json:"sign"`
}

const CloakUrl = "https://api-visitor.fangyu.io/check/1451/vA7WLaoIjFfDpBL6Y8"
const CloakKey = "BWdix4qbXz45RJ5rDme5a"

type CloakResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Status    bool        `json:"status"`
		Message   interface{} `json:"message"`
		Jump      string      `json:"jump"`
		ErrorCode string      `json:"errorCode"`
		Custom    struct {
			WebId int    `json:"webId"`
			Name  string `json:"name"`
		} `json:"custom"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (s TSchema) callCloak(c *TClientInfo) {
	log := common.Logger.WithField("调用cloak", "开始")
	callCloakDTO := CallCloakDTO{
		ClientIp:       c.IP,
		UserAgent:      c.Ua,
		VisitUrl:       "https://a2site.xyz/" + c.No,
		ClientLanguage: c.Lang,
		Referer:        c.Referer,
		Timestamp:      time.Now().Unix(),
		Sign:           CloakKey,
	}
	jsonStr, _ := json.Marshal(callCloakDTO)
	req, _ := http.NewRequest("POST", CloakUrl, strings.NewReader(string(jsonStr)))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	res, _ := io.ReadAll(resp.Body)
	c.CloakResponse = string(res)
	var cloakResponse CloakResponse
	err := json.Unmarshal(res, &cloakResponse)
	if err != nil {
		c.Status = 1
		return
	}
	log.Infof("cloak response: %v", cloakResponse)
	if cloakResponse.Data.Status {
		c.Status = 1
		return
	}
	c.Status = 2
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
	s.callCloak(c)
	c.Insert()
	return c.Status == 1
}
