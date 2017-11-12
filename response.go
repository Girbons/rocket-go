package rocket

import (
	"io/ioutil"
	"net/http"
)

// Response struct.
type Response struct {
	Ok          bool
	Error       error
	Message     string
	StatusCode  int
	RawResponse *http.Response
}

// buildResponse.
func buildResponse(resp *http.Response, err error) (*Response, error) {
	if err != nil {
		return &Response{Error: err}, err
	}

	response := &Response{
		Ok:          resp.StatusCode >= 200 && resp.StatusCode < 300,
		Error:       nil,
		Message:     resp.Status,
		StatusCode:  resp.StatusCode,
		RawResponse: resp,
	}

	return response, nil
}

// handleResponse build the response and return the parsed one.
func handleResponse(response *http.Response, err error) (*Response, error) {
	evaluatedResponse, resErr := buildResponse(response, err)

	if evaluatedResponse.Ok != true {
		return evaluatedResponse, err
	}

	return evaluatedResponse, resErr
}

// Content return parsed body.
func (response *Response) Content() string {
	defer response.RawResponse.Body.Close()

	responseBody, parsErr := ioutil.ReadAll(response.RawResponse.Body)
	if parsErr != nil {
		return "Cannot read body"
	}

	return string(responseBody)
}
