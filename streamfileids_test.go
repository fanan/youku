package youku

import (
	"testing"
)

var data = []byte(`
    {
        "hd2": "33*17*33*33*33*43*33*8*33*33*45*43*37*55*55*30*8*36*47*43*27*45*33*27*43*36*36*36*13*13*50*13*55*33*50*13*17*17*9*13*45*27*50*9*22*13*13*8*9*27*47*47*8*9*43*13*55*43*36*53*27*55*13*30*43*37*",
        "mp4": "33*17*33*33*33*22*33*37*33*33*45*43*37*55*55*43*27*47*47*43*27*45*33*27*43*36*36*36*13*13*50*13*55*33*50*13*17*17*9*13*45*27*50*9*22*13*13*8*9*27*47*47*8*9*43*13*55*43*36*53*27*55*13*30*43*37*",
        "flv": "33*17*33*33*33*27*33*37*33*33*45*43*37*55*45*43*33*22*47*43*27*45*33*27*43*36*36*36*13*13*50*13*55*33*50*13*17*17*9*13*45*27*50*9*22*13*13*8*9*27*47*47*8*9*43*13*55*43*36*53*27*55*13*30*43*37*"
    }
    `)

func TestStreamFildIdsParse(t *testing.T) {
	sf := new(StreamFileIds)
	err := sf.Parse(&data)
	if err != nil {
		t.Error(err)
	}
	vt := sf.ChooseType()
	if vt != HD2 {
		t.Error("video type should be hd2")
	}
	mix := sf.GetFileIdMixedString(236)
	if mix != `oGLBIZ\vE-tqa9xkO3KJb/8nYwz2V_4rp0geD7uhHQd1y5PCMXAfUF:6NsmlWiRScTj.` {
		t.Error("GetFileId error")
	}
	logger.Println(sf.GetFileId(236))
}
