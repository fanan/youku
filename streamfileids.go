package youku

import (
    "encoding/json"
    "math"
    "strconv"
    "strings"
)

type StreamFileIds struct {
    Hd2 string `json:"hd2"`
    Mp4 string `json:"mp4"`
    Flv string `json:"flv"`
}

func (s *StreamFileIds) Parse(data *[]byte) (err error) {
    return json.Unmarshal(*data, s)
}

func (s *StreamFileIds) ChooseType() int {
    if s.Hd2 != "" {
        return HD2
    }
    if s.Mp4 != "" {
        return MP4
    }
    return FLV
}

func (s *StreamFileIds) GetFileIdMixedString(seed int) (fileid string) {
    source := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ/\\:._-1234567890"
    n := len(source)
    for i:=0;i<n;i++ {
        seed = (seed * 211 + 30031) % 65536
        //index := int(float64(seed) * float64(65536) / float64(len(source)))
        index := int(math.Floor((float64(seed) / float64(65536) * float64(len(source)))))
        c := source[index]
        //logger.Println("seed:", seed, "index:", index, "c:", string(c))
        fileid += string(c)
        source = strings.Replace(source, string(c), "", -1)
        //logger.Println(source)
    }
    return fileid
}

func (s *StreamFileIds) GetFileId(seed int) (fileid string) {
    mixed := s.GetFileIdMixedString(seed)
    typeid := s.ChooseType()
    var indecies = ""
    switch typeid {
    case HD2:
        indecies = s.Hd2
    case MP4:
        indecies = s.Mp4
    default:
        indecies = s.Flv
    }
    items := make([]byte, 0)
    for _, _index := range strings.Split(indecies, "*") {
        index, _ := strconv.Atoi(_index)
        items = append(items, mixed[index])
    }
    return string(items)
}
