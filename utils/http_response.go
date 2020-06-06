package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (u *Utils) SendResponse(i interface{}, w http.ResponseWriter, mimeType string) {
	switch mimeType {
	case "application/json":
		w.Header().Add("Content-Type", mimeType)
		out, outErr := json.Marshal(i)
		if outErr != nil {
			u.l.Errorf("Error converting output data to JSON. %s", outErr.Error())
			http.Error(w, outErr.Error(), http.StatusInternalServerError)
		} else {
			var bytesWritten int
			var err error
			bytesWritten, err = w.Write(out)

			if err != nil {
				u.l.Warn(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			} else {
				u.l.Infof("Wrote %d bytes", bytesWritten)
			}
		}
	default:
		w.Header().Add("Content-Type", "application/octet-stream")
		var bytesWritten int
		var err error
		bytesWritten, err = w.Write([]byte(fmt.Sprintf("%#v", i))) // https://stackoverflow.com/a/56816239/6670698

		if err != nil {
			u.l.Warn(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			u.l.Infof("Wrote %d bytes", bytesWritten)
		}
	}
}

func (u *Utils) SendResponseJSON(i interface{}, w http.ResponseWriter) {
	u.SendResponse(i, w, "application/json")
}