package datagen

import (
	"encoding/hex"
	"strconv"
	"time"

	"github.com/feiyuw/xxtea"
	json "github.com/json-iterator/go"
)

var (
	_Key = []byte("3ef71a5250fc40af2211b434d961ff5209")
)

type webRequest struct {
	Ua     string `json:"ua"`
	Cd     string `json:"cd"`
	Pr     string `json:"pr"`
	Res    string `json:"Res"`
	Ar     string `json:"ar"`
	To     string `json:"to"`
	Ss     string `json:"ss"`
	Ls     string `json:"ls"`
	Ind    string `json:"ind"`
	Od     string `json:"od"`
	Cc     string `json:"cc"`
	Dnt    string `json:"dnt"`
	Can    string `json:"can"`
	Web    string `json:"web"`
	Adb    string `json:"adb"`
	Hlb    string `json:"hlb"`
	Hll    string `json:"hll"`
	Hlr    string `json:"hlr"`
	Jf     string `json:"jf"`
	Rp     string `json:"rp"`
	Ts     string `json:"ts"`
	AppKey string `json:"appKey"`
	UserID string `json:"userId"`
	Lid    string `json:"lid"`
}

// RandomWebRequest generate a random constid request of web client
func RandomWebRequest(appKey string) ([]byte, error) {
	request := webRequest{
		Ua:     "96c87b12fc7f3abfcca429277d273bba",
		Cd:     "24",
		Pr:     "1.25",
		Res:    "1536;864",
		Ar:     "1536;824",
		To:     "-480",
		Ss:     "1",
		Ls:     "1",
		Ind:    "1",
		Od:     "1",
		Cc:     "win32",
		Dnt:    "unknown",
		Can:    "d8b25151d20c53301d153e04e5f4b1f4",
		Web:    randomUUID(),
		Adb:    "false",
		Hlb:    "false",
		Hll:    "false",
		Hlr:    "false",
		Jf:     randomUUID(),
		Rp:     randomUUID(),
		Ts:     "0;false;false",
		AppKey: appKey,
		UserID: "a12345",
		Lid:    randomServerLid(),
	}

	return json.Marshal(request)
}

func randomServerLid() string {
	src := make([]byte, 31)

	// timestamp 8 byte
	copy(src[:8], []byte(strconv.FormatInt(time.Now().Unix(), 16)))
	// random 23 byte
	copy(src[8:], randomBytes(23))

	dst, err := xxtea.Encrypt(src, _Key)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(dst)
}
