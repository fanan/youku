package youku

import (
	"io/ioutil"
	"testing"
)

func TestYoukuParse(t *testing.T) {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		t.Error(err)
	}
	r := new(Youku)
	err = r.Parse(&data)
	if err != nil {
		t.Error(err)
	}
	//logger.Printf("%v", *r)
}

func TestNewYouku (t *testing.T) {
    id := "XMjI2ODA4MTQw"
    y, err := NewYouku(id)
    if err != nil {
        t.Error(err)
    }
    if len(y.Videos) != 1 {
        t.Error("y.Videos length should be 1")
    }
    v := y.Videos[0]
    logger.Println(v.Title, v.Seed, v.Videoid)
    if v.VidEncoded != id {
        t.Error("ids shoule be same")
    }
    lists := v.GetDownloadList()
    logger.Println(lists, v.Previous, v.Next, v.Lists)
}

func TestChain (t *testing.T) {
    id := "XMjI2ODA4MTQw"
    y, err := NewYouku(id)
    if err != nil {
        t.Error(err)
    }
    first, err := y.GetFirst()
    if err != nil {
        t.Error(err)
    }
    current := first
    for  {
        logger.Println(current.Videos[0].Title)
        current, err = current.GetNext()
        if err != nil {
            t.Error(err)
            break
        }
        if current == nil {
            break
        }
    }
}
