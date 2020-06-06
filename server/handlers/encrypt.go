package handlers

import (
	"encoding/json"
	"fmt"
	"generator/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// Original Concept by: https://github.com/backwardspy/dev-urandom-as-a-service/blob/master/rando.go

type Encrypter struct {
	l *log.Logger
	u *utils.Utils
}

type EncrypterInput struct {
	Data string `json:"data"`
	EncPass string	`json:"encpass"`
}

type EncrypterOutput struct {
	Data string `json:"data"`
	Error string	`json:"error"`
}

func NewEncrypter(l *log.Logger) *Encrypter{
	return &Encrypter{
		l: l,
		u: utils.New(),
	}
}

func (e *Encrypter) Encrypt(toEncrypt, password string) (string,error) {
	return e.u.EncDec.Encrypt(toEncrypt, password)
}

func (e *Encrypter) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Language", "en")

	output := &EncrypterOutput{
		Data:  "",
		Error: "",
	}

	if r.Method != http.MethodPost {
		e.l.Errorf("API accessed with method %s. Sending HTTP 405", r.Method)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, bodyBytesErr := ioutil.ReadAll(r.Body)

	if bodyBytesErr != nil {
		e.l.Errorf("Unable to read request Body. %s", bodyBytesErr.Error())
		output.Error = bodyBytesErr.Error()
		e.u.SendResponseJSON(output, w)
		return
	}

	input := &EncrypterInput{
		Data:    "",
		EncPass: "",
	}

	inputErr := json.Unmarshal(bodyBytes, input)
	if inputErr != nil {
		e.l.Errorf("Unable to parse request data. %s", inputErr.Error())
		output.Error = inputErr.Error()
		e.u.SendResponseJSON(output, w)
	}

	e.l.Infof("Unrmarshalled Request: %#v", input)

	outputData, outputErr := e.Encrypt(input.Data, input.EncPass)
	output.Data = outputData
	if outputErr != nil {
		output.Error = fmt.Sprintf("%s",outputErr.Error())
		e.l.Warnf("Unable to Encrypt. %s", outputErr.Error())
	} else {
		output.Error = "null"
	}

	e.u.SendResponseJSON(output, w)
}
