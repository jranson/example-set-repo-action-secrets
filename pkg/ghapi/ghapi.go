package ghapi

import (
	"bytes"
	"io"
	"net/http"
	"strings"
)

func Get(token, path string) (int, []byte, error) {
	return do(token, http.MethodGet, path, nil)
}

func Put(token, path string, body []byte) (int, []byte, error) {
	return do(token, http.MethodPut, path, body)
}

func do(token, method, path string, body []byte) (int, []byte, error) {
	path = strings.TrimPrefix(path, "/")
	var req *http.Request
	var err error
	if body != nil {
		req, err = http.NewRequest(method,
			"https://api.github.com/"+path, bytes.NewReader(body))
	} else {
		req, err = http.NewRequest(method,
			"https://api.github.com/"+path, nil)
	}
	if err != nil {
		return 500, nil, err
	}
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	code := 500
	var b []byte
	var err2 error
	if resp != nil {
		if resp.StatusCode > 0 {
			code = resp.StatusCode
		}
		if resp.Body != nil {
			b, err2 = io.ReadAll(resp.Body)
		}
	}
	if err != nil {
		return code, b, err
	}
	return code, b, err2
}
