package youku

import (
	"encoding/json"
)

type Seg struct {
	No      string `json:"no"`
	Size    string `json:"size"`
	Seconds string `json:"seconds"`
	K       string `json:"k"`
	K2      string `json:"k2"`
}

type Segs struct {
	Hd2 []Seg `json:"hd2"`
	Mp4 []Seg `json:"mp4"`
	Flv []Seg `json:"flv"`
}

func (s *Segs) Parse(data *[]byte) (err error) {
	err = json.Unmarshal(*data, s)
	if err != nil {
		return err
	}
	return nil
}
