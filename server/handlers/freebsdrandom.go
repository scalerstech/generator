package handlers

import (
	"generator/utils"
	log "github.com/sirupsen/logrus"
	"github.com/z9905080/freebsdrandom"
	"net/http"
	"time"
)

type Freebsdrandom struct {
	l *log.Logger
	u *utils.Utils
}

func NewFreebsdrandom(l *log.Logger) *Freebsdrandom{
	return &Freebsdrandom{
		l: l,
		u: utils.New(),
	}
}

func (u *Freebsdrandom) RandBytes(totalBytes int) []byte {

	return freebsdrandom.Bytes(totalBytes)
}

func (u *Freebsdrandom) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	maxBytes := u.u.Atoi(r.URL.Query().Get("length"), 1024)

	w.Header().Add("Content-Language", "en")
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Content-Disposition", "inline; filename=\"freebsdrandom.log\"")

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
