package init

import (
	"encoding/xml"
	. "pms/models/base"
)

type InitGroups struct {
	XMLName xml.Name `xml:"Groups"`
	Groups  []Group  `xml:"group"`
}

type InitUsers struct {
	XMLName xml.Name `xml:"Users"`
	Users   []User   `xml:"user"`
}
type InitCountries struct {
	XMLName   xml.Name  `xml:"Countries"`
	Countries []Country `xml:"country"`
}

type InitProvince struct {
	ID   uint   `xml:"ID,attr"`
	Name string `xml:"ProvinceName,attr"`
	PID  uint   `xml:"PID,attr"`
}
type InitProvinces struct {
	XMLName   xml.Name       `xml:"Provinces"`
	Provinces []InitProvince `xml:"Province"`
}
type InitCity struct {
	ID   uint   `xml:"ID,attr"`
	Name string `xml:"CityName,attr"`
	PID  uint   `xml:"PID,attr"`
}
type InitCities struct {
	XMLName xml.Name   `xml:"Cities"`
	Cities  []InitCity `xml:"City"`
}

type InitDistrict struct {
	ID   uint   `xml:"ID,attr"`
	Name string `xml:"DistrictName,attr"`
	PID  uint   `xml:"CID,attr"`
}
type InitDistricts struct {
	XMLName   xml.Name       `xml:"Districts"`
	Districts []InitDistrict `xml:"District"`
}
