package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.org/eventmodeling/ecommerce/pkg/support/commons"
)

/*
Struct responsible for configure the settings for the request.

SuccessCodes: is a list of HTTP Status Codes that are considered successful. If the response status code is not in this list, a new attempt will be made.

AttemptRetries: is the number of times the request will be retried if it fails.

TimeoutSeconds: is the number of seconds the request will wait for a response before timing out.

WaitingTime: is the number of seconds the request will wait between each attempt.

Debug: show request information more detailed.

Uuid: is the unique identifier of the request.
*/
type RequestSettings struct {
	SuccessCodes   []int
	AttemptRetries int
	TimeoutSeconds int
	WaitingTime    int
	Debug          bool
	Uuid           string
}

type Request struct {
	Method  string
	Url     string
	Body    string
	Headers map[string][]string
}

type Response struct {
	StatusCode int
	Body       string
	Headers    map[string][]string
	Request    Request
}

/*
This function will return a parsed Entity from a GET request.

url: is the url of the resource.

responseObject: is the object that will be parsed from the response body.

headers: is a map of headers to be sent with the request.

requestSettings: is the configuration of the request.
*/
func GetEntity(url string, responseObject interface{}, headers map[string]string, requestSettings RequestSettings) (response Response, err error) {
	response, err = createRequisition(http.MethodGet, url, nil, headers, requestSettings)
	if err != nil {
		return
	}

	return parseJsonResponseObject(responseObject, response)
}

/*
This function is responsible for return a parsed Entity from a POST request.

url: is the url of the resource.

body: is the object that will be sended in the request.

responseObject: is the object that will be parsed from the response body.

headers: is a map of headers to be sent with the request.

requestSettings: is the configuration of the request.
*/
func PostEntity(url string, body interface{}, responseObject interface{}, headers map[string]string, requestSettings RequestSettings) (response Response, err error) {
	response, err = createRequisition(http.MethodPost, url, body, headers, requestSettings)
	if err != nil {
		return
	}

	return parseJsonResponseObject(responseObject, response)
}

/*
This function is responsible for return a parsed Entity from a PUT request.

url: is the url of the resource.

body: is the object that will be sended in the request.

responseObject: is the object that will be parsed from the response body.

headers: is a map of headers to be sent with the request.

requestSettings: is the configuration of the request.
*/
func PutEntity(url string, body interface{}, responseObject interface{}, headers map[string]string, requestSettings RequestSettings) (response Response, err error) {
	response, err = createRequisition(http.MethodPut, url, body, headers, requestSettings)
	if err != nil {
		return
	}

	return parseJsonResponseObject(responseObject, response)
}

/*
This function is responsible for return a parsed Entity from a PATCH request.

url: is the url of the resource.

body: is the object that will be sended in the request.

responseObject: is the object that will be parsed from the response body.

headers: is a map of headers to be sent with the request.

requestSettings: is the configuration of the request.
*/
func PatchEntity(url string, body interface{}, responseObject interface{}, headers map[string]string, requestSettings RequestSettings) (response Response, err error) {
	response, err = createRequisition(http.MethodPatch, url, body, headers, requestSettings)
	if err != nil {
		return
	}

	return parseJsonResponseObject(responseObject, response)
}

/*
This function is responsible for return a parsed Entity from a DELETE request.

url: is the url of the resource.

body: is the object that will be sended in the request.

responseObject: is the object that will be parsed from the response body.

headers: is a map of headers to be sent with the request.

requestSettings: is the configuration of the request.
*/
func DeleteEntity(url string, body interface{}, responseObject interface{}, headers map[string]string, requestSettings RequestSettings) (response Response, err error) {
	response, err = createRequisition(http.MethodDelete, url, body, headers, requestSettings)
	if err != nil {
		return
	}

	return parseJsonResponseObject(responseObject, response)
}

func createRequisition(method string, url string, bodyObject interface{}, headers map[string]string, requestSettings RequestSettings) (response Response, err error) {
	if requestSettings.Uuid != "" {
		requestSettings.Uuid = commons.ConcatenateStrings(requestSettings.Uuid, " ")
	}

	if requestSettings.AttemptRetries == 0 {
		requestSettings.AttemptRetries = 1
	}

	client := &http.Client{Timeout: time.Duration(requestSettings.TimeoutSeconds) * time.Second}

	var req *http.Request

	var bodyRequest []byte
	var body io.Reader

	for i := 0; i < requestSettings.AttemptRetries; i++ {
		doLog(requestSettings.Debug, requestSettings.Uuid, "Attempt ", i+1, " URL: ", url)
		if bodyObject != nil {
			if _, ok := interface{}(bodyObject).(io.Reader); ok {
				body = bodyObject.(io.Reader)
			} else {
				bodyRequest, err = json.Marshal(bodyObject)
				if err != nil {
					return response, err
				}

				body = bytes.NewBuffer(bodyRequest)
			}

			req, err = http.NewRequest(method, url, body)

			if err != nil {
				return response, err
			}
		} else {
			req, err = http.NewRequest(method, url, nil)

			if err != nil {
				return response, err
			}
		}

		for key, value := range headers {
			req.Header.Set(key, value)
		}

		resp, err := client.Do(req)

		var bodyResponseString string
		if resp != nil {
			bodyBytes, _ := io.ReadAll(resp.Body)
			bodyResponseString = string(bodyBytes)
		}

		bodyRequestString := commons.IfThenElse(commons.StringIsNotEmpty(string(bodyRequest)), string(bodyRequest), "is empty or is a io.Reader").(string)
		doLog(requestSettings.Debug, requestSettings.Uuid, "Request: ", req.Method, " ", req.URL.String(), " Headers: ", req.Header, " Body: ", bodyRequestString)
		if resp != nil {
			doLog(requestSettings.Debug, requestSettings.Uuid, "Response: StatusCode: ", resp.StatusCode, " Headers: ", resp.Header, " Body: ", bodyResponseString)
		}

		if err != nil {
			return response, err
		}

		response = createResponse(resp, bodyResponseString, req, bodyRequest)

		defer resp.Body.Close()

		if commons.ContainsInt(requestSettings.SuccessCodes, resp.StatusCode) {
			return response, nil
		}

		doLog(requestSettings.Debug, requestSettings.Uuid, "Unexpected Status Code: ", resp.StatusCode)
		time.Sleep(time.Duration(requestSettings.WaitingTime) * time.Second)
	}

	doLog(requestSettings.Debug, requestSettings.Uuid, "Request failed after ", requestSettings.AttemptRetries, " attempts.")
	return response, errors.New(commons.ConcatenateStrings("Unexpected Status Code in request: ", url, " - Status Code: ", strconv.Itoa(response.StatusCode)))
}

func createResponse(resp *http.Response, bodyResponseString string, req *http.Request, jsonBody []byte) (response Response) {
	response.StatusCode = resp.StatusCode
	response.Body = bodyResponseString
	response.Headers = resp.Header
	response.Request = Request{
		Method:  req.Method,
		Url:     req.URL.String(),
		Body:    string(jsonBody),
		Headers: req.Header,
	}

	return response
}

func parseJsonResponseObject(target interface{}, resp Response) (Response, error) {
	if target == nil {
		return resp, nil
	}

	json.Unmarshal([]byte(resp.Body), &target)

	return resp, nil
}

func doLog(debug bool, m ...any) {
	if !debug {
		return
	}

	log.Print(m...)
}
