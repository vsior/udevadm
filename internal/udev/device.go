package udev

import "encoding/json"

type Device map[string]string

func (d Device) MarshalJSON() ([]byte, error) {
	type tmp Device
	return json.Marshal(tmp(d))
}

func (d Device) String() string {
	buff, _ := json.MarshalIndent(d, "", "  ")
	return string(buff)
}

//type devices []device
