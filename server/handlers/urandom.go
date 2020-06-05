package handlers

import (
	"encoding/base64"
	"generator/utils"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"time"
)

// Original Concept by: https://github.com/backwardspy/dev-urandom-as-a-service/blob/master/rando.go

type Urandom struct {
	l *log.Logger
	u *utils.Utils
}

func NewUrandom(l *log.Logger) *Urandom{
	return &Urandom{
		l: l,
		u: utils.New(),
	}
}

func (u *Urandom) RandBytes(totalBytes int64) []byte {
	data := make([]byte, totalBytes)
	rand.Read(data)
	b64 := base64.StdEncoding.EncodeToString(data)
	return []byte(b64)
}

func (u *Urandom) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	maxBytes := u.u.ParseInt(r.URL.Query().Get("length"), 1024)

	w.Header().Add("Content-Language", "en")

	b64 := u.RandBytes(maxBytes)
	bytesWritten, err := w.Write(b64)
	if err != nil {
		u.l.Warn(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		u.l.Infof("Wrote %d bytes", bytesWritten)
	}
	time.Sleep(100 * time.Millisecond)

}
