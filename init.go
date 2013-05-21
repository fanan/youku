package youku

import (
	"io"
	"log"
	"os"
)

const (
	FLV_URL_SKELETON = `http://f.youku.com/player/getFlvPath/sid/00_00/st/%s/fileid/%s%s%s?K=%s`
    FLV_DATA_SKELETON = `http://v.youku.com/player/getPlayList/VideoIDS/%s`
)

type LogLevel int

const (
	Trace = iota
	Debug
	Info
	Warn
	Critical
)

//video quality
const (
	FLV = iota
	MP4
	HD2
)

var (
	logger = log.New(os.Stdout, "[Youku logger] ", log.LstdFlags|log.Lshortfile)
)

func initLogger(wr io.Writer) {
	logger = log.New(wr, "[Youku logger]", log.LstdFlags)
}
