package api

import (
	"bytes"
	"io"
	"net/http"
)

func ReadFileFromRequest(req *http.Request, formKey string) ([]byte, string, error) {
	file, header, err := req.FormFile(formKey)

	if err != nil {
		return nil, "", err
	}

	defer file.Close()

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)

	if err != nil {
		return nil, "", err
	}

	bytes := buf.Bytes()
	return bytes, header.Filename, nil
}
