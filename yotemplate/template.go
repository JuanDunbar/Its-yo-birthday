package yotemplate

import (
	"bytes"
	"github.com/juandunbar/yobirthday/yodata"
	"text/template"
)

const defaultFilePath = "yotemplate/cannedresponses.gotext"

type CannedResponse struct {
	Template *template.Template
}

// Create a new CannedResponse object with our parsed template file
func NewResponse(filePath string) (*CannedResponse, error) {
	// Allow different file paths for unit testing
	var respFilePath string
	if len(filePath) > 0 {
		respFilePath = filePath
	} else {
		respFilePath = defaultFilePath
	}
	t, err := template.ParseFiles(respFilePath)
	if err != nil {
		return nil, err
	}
	return &CannedResponse{Template:t}, nil
}

// Render the specific template defined by our Email.Type
func (r *CannedResponse) Render(data *yodata.Email) (string, error) {
	buf := new(bytes.Buffer)
	err := r.Template.ExecuteTemplate(buf, data.Type, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}