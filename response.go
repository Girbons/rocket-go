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
	Content     string
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
		Content:     "",
	}

	return response, nil
}

// handleResponse build the response and return the parsed one.
func handleResponse(response *http.Response, err error) (*Response, error) {
	evaluatedResponse, resErr := buildResponse(response, err)

	if evaluatedResponse.Ok != true {
		return evaluatedResponse, err
	}
	parseContent(evaluatedResponse, resErr)

	return evaluatedResponse, resErr
}

// parseResponse check if the response is ok and return the body parsed.
func parseContent(response *Response, err error) (*Response, error) {
	defer response.RawResponse.Body.Close()

	responseBody, parsErr := ioutil.ReadAll(response.RawResponse.Body)
	if parsErr != nil {
		response.Content = "cannot read response body"
		return response, parsErr
	}
	response.Content = string(responseBody)

	return response, nil
}
