package rocket

import (
	"fmt"
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
func handleResponse(response *http.Response, err error) (string, error) {
	evaluatedResponse, resErr := buildResponse(response, err)

	if evaluatedResponse.Ok != true {
		return fmt.Sprintf(`{"status_code": %d, "message": "%s"}`, evaluatedResponse.StatusCode, evaluatedResponse.Message), nil
	}

	return parseResponse(evaluatedResponse, resErr)
}

// parseResponse check if the response is ok and return the body parsed.
func parseResponse(response *Response, err error) (string, error) {
	defer response.RawResponse.Body.Close()

	responseBody, parsErr := ioutil.ReadAll(response.RawResponse.Body)
	if parsErr != nil {
		msg := fmt.Sprintf("cannot read response body")
		return msg, parsErr
	}

	return string(responseBody), nil
}
