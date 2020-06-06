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

type Decrypter struct {
	l *log.Logger
	u *utils.Utils
}

type DecrypterInput struct {
	Data string `json:"data"`
	EncPass string	`json:"encpass"`
}

type DecrypterOutput struct {
	Data string `json:"data"`
	Error string	`json:"error"`
}

func NewDecrypter(l *log.Logger) *Decrypter{
	return &Decrypter{
		l: l,
		u: utils.New(),
	}
}

func (d *Decrypter) Decrypt(toDecrypt, password string) (string,error) {
	return d.u.Decrypt(toDecrypt, password)
}

func (d *Decrypter) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Language", "en")

	output := &DecrypterOutput{
		Data:  "",
		Error: "",
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, bodyBytesErr := ioutil.ReadAll(r.Body)

	if bodyBytesErr != nil {
		d.l.Errorf("Unable to read request Body. %s", bodyBytesErr.Error())
		output.Error = bodyBytesErr.Error()
		d.u.SendResponseJSON(output, w)
		return
	}

	input := &EncrypterInput{
		Data:    "",
		EncPass: "",
	}

	inputErr := json.Unmarshal(bodyBytes, input)
	if inputErr != nil {
		d.l.Errorf("Unable to parse request data. %s", inputErr.Error())
		output.Error = inputErr.Error()
		d.u.SendResponseJSON(output, w)
	}

	d.l.Infof("Unrmarshalled Request: %#v", input)

	outputData, outputErr := d.Decrypt(input.Data, input.EncPass)
	output.Data = outputData
	if outputErr != nil {
		output.Error = fmt.Sprintf("%s",outputErr.Error())
		d.l.Warnf("Unable to Encrypt. %s", outputErr.Error())
	} else {
		output.Error = "null"
	}

	d.u.SendResponseJSON(output, w)
}
