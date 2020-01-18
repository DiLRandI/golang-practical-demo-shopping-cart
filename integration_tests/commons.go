package integration_tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
)

var c *http.Client = &http.Client{
	Timeout: 30 * time.Second,
}

func assertStringsMatch(name, actual, expected string) error {
	if actual != expected {
		return fmt.Errorf("Expecting %q to be %q, but actual is %q", name, expected, actual)
	}
	return nil
}

func assertIntsMatch(name string, actual, expected int) error {
	if actual != expected {
		return fmt.Errorf("Expecting %q to be %v, but actual is %v", name, expected, actual)
	}
	return nil
}

// helper function to execute PUT / GET / DELETE requests
func SendHttp(method string, endPoint string, server string, in interface{}, out interface{}, headers map[string]string) (int, error) {

	buf := &bytes.Buffer{}
	defer buf.Reset()

	if err := json.NewEncoder(buf).Encode(in); err != nil {
		return 0, errors.New(fmt.Sprintf("Error when encoding request at url %s using request %s: %s", endPoint, in, err.Error()))
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", server, endPoint), buf)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Error when creating http request: %s", err))
	}

	appendHeaders(req, headers)
	resp, err := c.Do(req)

	if err != nil {
		return 0, errors.New(fmt.Sprintf("Error when attempting to send request: %s", err))
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		out = nil
		return resp.StatusCode, nil
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusCreated {
		return resp.StatusCode, errors.New(fmt.Sprintf("Unexpected status code got %s", resp.Status))
	}

	// if out exists, parse the result into it
	if out != nil {
		if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
			return resp.StatusCode, errors.New(fmt.Sprintf("Error decoding response: %s", err))
		}
	}

	return resp.StatusCode, nil
}

func getServiceHeaders() map[string]string {
	headers := make(map[string]string)
	headers["x-device-type"] = "postman"
	headers["x-client-version"] = "5.16"
	headers["x-device-language"] = "en-GB"
	return headers
}

func appendHeaders(r *http.Request, h map[string]string) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")

	for k, v := range h {
		r.Header.Add(k, v)
	}
}

// help function to check test environment containers are listening on expected ports
func serverActive(serverPort string) bool {

	ln, err := net.Listen("tcp", fmt.Sprintf(serverPort))

	if err == nil {
		_ = ln.Close()
		return false
	}

	// we couldn't open it - it's active
	return true
}

// help function to check test environment containers are listening on expected ports
func portActive(port int) bool {

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err == nil {
		_ = ln.Close()
		return false
	}

	// we couldn't open it - it's active
	return true
}

func makeLogError(format string, args ...interface{}) error {
	err := fmt.Errorf(format, args...)
	return err
}
