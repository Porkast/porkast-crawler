package worker

import (
	"testing"

	"github.com/gogf/gf/v2/encoding/gjson"
)

func Test_getXimalayaAlbumUrlList(t *testing.T) {
	var (
		sampleJson   *gjson.Json
		albumUrlList []string
		err          error
	)

	sampleJson, err = gjson.Load("./testdata/ximalaya_resp.json")
	if err != nil {
		t.Fatal("Parse ximalaya category response as json failed : ", err)
	}

	albumUrlList = getXimalayaAlbumUrlList(*sampleJson)
	if len(albumUrlList) == 0 {
		t.Fatal("Parse ximalaya album list is empty")
	}

}
