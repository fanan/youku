package youku

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type List struct {
	Seq        string `json:"seq"`
	Vid        string `json:"vid"`
	VidEncoded string `json:"vidEncoded"`
	Title      string `json:"title"`
	Vv         string `json:"vv"`
}

type Video struct {
	Seed        int           `json:"seed"`
	Lists       []List        `json:"list"`
	Videoid     string        `json:"videoid"`
	VidEncoded  string        `json:"vidEncoded"`
	Title       string        `json:"title"`
	Key1        string        `json:"key1"`
	Key2        string        `json:"key2"`
	Seconds     string        `json:"seconds"`
	StreamFiles StreamFileIds `json:"streamfileids"`
	Segs        Segs          `json:"segs"`
	Previous    List          `json:"list_pre"`
	Next        List          `json:"list_next"`
}

type Youku struct {
	Videos []Video `json:"data"`
	Id     string
}

func (r *Youku) Parse(data *[]byte) error {
	return json.Unmarshal(*data, r)
}

func NewYouku(id string) (y *Youku, err error) {
	y = new(Youku)
	y.Id = id
	url := fmt.Sprintf(FLV_DATA_SKELETON, id)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return y, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return y, err
	}
	err = y.Parse(&data)
	if err != nil {
		return y, err
	}
	return y, nil
}

func (v *Video) GetDownloadList() []string {
	downloadList := make([]string, 0)
	videoFormat := v.StreamFiles.ChooseType()
	format := ""
	var segs *[]Seg
	switch videoFormat {
	case MP4:
		segs = &v.Segs.Mp4
		format = "mp4"
	case FLV:
		segs = &v.Segs.Flv
		format = "flv"
	case HD2:
		segs = &v.Segs.Hd2
		format = "flv"
	}
	fileId := v.StreamFiles.GetFileId(v.Seed)
	for _, seg := range *segs {
		no := seg.No
        _no, _ := strconv.Atoi(no)
        no = strconv.FormatInt(int64(_no), 16)
        no = strings.ToUpper(no)
		if len(no) == 1 {
            no = "0" + no
        }
        url := fmt.Sprintf(FLV_URL_SKELETON, format, fileId[:8], no, fileId[10:], seg.K)
		downloadList = append(downloadList, url)
	}
	return downloadList
}

func (y *Youku) GetPrevious() (prev *Youku, err error) {
    if len(y.Videos) != 1{
        return nil, errors.New("youku data error!")
    }
    v := y.Videos[0]
    if v.Previous.VidEncoded != "" {
        prev, err = NewYouku(v.Previous.VidEncoded)
        if err != nil {
            return nil, err
        }
        return prev, nil
    }
    return nil, nil
}

func (y *Youku) GetFirst() (first *Youku, err error) {
    current := y
    for  {
        first, err = current.GetPrevious()
        if err != nil {
            return current, err
        }
        if first != nil {
            current = first
        } else {
            return current, nil
        }
    }
}


func (y *Youku) GetNext() (next *Youku, err error) {
    if len(y.Videos) != 1{
        return nil, errors.New("youku data error!")
    }
    v := y.Videos[0]
    if v.Next.VidEncoded != "" {
        next, err = NewYouku(v.Next.VidEncoded)
        if err != nil {
            return nil, err
        }
        return next, nil
    }
    return nil, nil
}
