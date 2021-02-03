// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Request - struct containing request data
type Request struct {
	URL    string
	Client *http.Client
	Params map[string]map[string]interface{}
}

// Response - struct containing response data
type Response struct {
	StatusCode    int
	StatusMessage string
	Body          interface{}
}

// Client - http.Client
var Client http.Client

// Params - maps for allowed http parameters
type Params map[string]map[string]interface{}

// Get - Makes a HTTP Get Request to provided URL
func Get(request *Request) (*Response, error) {
	var response interface{}
	req, err := http.NewRequest("GET", request.URL, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range request.Params["headers"] {
		req.Header.Set(key, fmt.Sprintf("%v", value))
	}
	resp, err := request.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&response)
	currentResp := &Response{
		StatusCode:    resp.StatusCode,
		StatusMessage: http.StatusText(resp.StatusCode),
		Body:          response,
	}
	return currentResp, nil
}

// Post - Makes a HTTP Post Request to provided URL
func Post(request *Request) (*Response, error) {
	var response interface{}
	body, err := json.Marshal(request.Params["body"])
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", request.URL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for key, value := range request.Params["headers"] {
		req.Header.Set(key, fmt.Sprintf("%v", value))
	}
	resp, err := request.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&response)
	currentResp := &Response{
		StatusCode:    resp.StatusCode,
		StatusMessage: http.StatusText(resp.StatusCode),
		Body:          response,
	}
	return currentResp, nil
}

// Put - Makes a HTTP Patch Request to provided URL
func Put(request *Request) (*Response, error) {
	var response interface{}
	body, err := json.Marshal(request.Params["body"])
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", request.URL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for key, value := range request.Params["headers"] {
		req.Header.Set(key, fmt.Sprintf("%v", value))
	}
	resp, err := request.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&response)
	currentResp := &Response{
		StatusCode:    resp.StatusCode,
		StatusMessage: http.StatusText(resp.StatusCode),
		Body:          response,
	}
	return currentResp, nil
}

// Delete - Makes a HTTP Delete Request to provided URL
func Delete(request *Request) (*Response, error) {
	var response interface{}
	req, err := http.NewRequest("DELETE", request.URL, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range request.Params["headers"] {
		req.Header.Set(key, fmt.Sprintf("%v", value))
	}
	resp, err := request.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&response)
	currentResp := &Response{
		StatusCode:    resp.StatusCode,
		StatusMessage: http.StatusText(resp.StatusCode),
		Body:          response,
	}
	return currentResp, nil
}

// Encode - URL encodes provided string
func Encode(text string) string {
	return url.QueryEscape(text)
}
