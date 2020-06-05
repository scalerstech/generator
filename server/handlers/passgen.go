package handlers

import (
	"encoding/json"
	"fmt"
	"generator/utils"
	"github.com/sethvargo/go-password/password"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Original Concept by: https://github.com/backwardspy/dev-urandom-as-a-service/blob/master/rando.go

type Passgen struct {
	l *log.Logger
	u *utils.Utils
}

type PassgenOutput struct {
	Password string `json:"password"`
	Error string	`json:"error"`
}

func NewPassgen(l *log.Logger) *Passgen{
	return &Passgen{
		l: l,
		u: utils.New(),
	}
}

func (p *Passgen) Generate(length, numDigits, numSymbols int, noUpper, allowRepeat bool) (string,error) {
	return password.Generate(length, numDigits, numSymbols, noUpper, allowRepeat)
}

func (p *Passgen) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Language", "en")

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	outFmt := r.URL.Query().Get("format")
	p.l.Infof("Output format provided: %s", outFmt)
	switch outFmt {
	case "":
		outFmt = "json"
	case "json":
		outFmt = "json"
	case "raw":
		outFmt = "raw"
	default:
		outFmt = "json"
	}

	p.l.Infof("Will use output format: %s", outFmt)

	passLen := p.u.Atoi(r.URL.Query().Get("length"), 24)
	digits := p.u.Atoi(r.URL.Query().Get("digits"), 9)
	symbols := p.u.Atoi(r.URL.Query().Get("symbol"), 9)
	lowercase := p.u.ParseBool(r.URL.Query().Get("lowercase"), false)
	repeat := p.u.ParseBool(r.URL.Query().Get("repeated"), false)

	output := &PassgenOutput{
		Password: "",
		Error:    "",
	}
	var passwdErr error
	output.Password, passwdErr = p.Generate(passLen, digits, symbols, lowercase, repeat)
	if passwdErr != nil {
		p.l.Warnf("Error Generating password. %s", passwdErr.Error())
		output.Error = fmt.Sprintf("%s", passwdErr.Error())
	}

	out, outErr := json.Marshal(output)
	if outErr != nil {
		p.l.Errorf("Error converting output data to JSON. %s", outErr.Error())
		http.Error(w, outErr.Error(), http.StatusInternalServerError)
		return
	} else {
		var bytesWritten int
		var err error
		if outFmt == "json" {
			w.Header().Add("Content-Type", "application/json")
			bytesWritten, err = w.Write(out)
		} else {
			w.Header().Add("Content-Type", "text/plain")
			bytesWritten, err = w.Write([]byte(output.Password))
		}

		if err != nil {
			p.l.Warn(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			p.l.Infof("Wrote %d bytes", bytesWritten)
		}
	}
}
