package handlers

import (
	"github.com/z9905080/freebsdrandom"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type Freebsdrandom struct {
	l *log.Logger
}

func NewFreebsdrandom(l *log.Logger) *Freebsdrandom{
	return &Freebsdrandom{
		l: l,
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

	maxbytesStr := r.Header.Get("Max-Bytes")
	maxBytes, maxBytesErr := strconv.Atoi(maxbytesStr)
	if maxBytesErr != nil {
		u.l.Warnf("Error converting \"maxbytes\" + to int64. %s", maxBytesErr.Error())
		u.l.Warnf("Setting maxBytes = 1024")
		maxBytes = 1024
	}

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
