package yotemplate

import (
	"bytes"
	"github.com/juandunbar/yobirthday/yodata"
	"text/template"
)

type CannedResponse struct {
	template *template.Template
}

// Create a new CannedResponse object with our parsed template file
func NewResponse() (*CannedResponse, error) {
	t, err := template.ParseFiles("yotemplate/cannedresponses.gotext")
	if err != nil {
		return nil, err
	}
	return &CannedResponse{template:t}, nil
}

// Render the specific template defined by our Email.Type
func (r *CannedResponse) Render(data *yodata.Email) (string, error) {
	buf := new(bytes.Buffer)
	err := r.template.ExecuteTemplate(buf, data.Type, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}