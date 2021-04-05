// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vider

import (
	"net/http"
	"testing"
)

type TodoResponse struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
}

type TodoRequest struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
}

//TestGet - tests the get method
func TestGet(t *testing.T) {
	resp, err := Get(&Request{
		URL:    "https://jsonplaceholder.typicode.com/todos/1",
		Client: &http.Client{},
		Params: Params{
			"headers": {
				"Content-Type": "application/json",
			},
		},
		Data: TodoResponse{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	//You will need to type cast from interface to the needed data datatype
	t.Log(resp.Body.(TodoResponse).Title)
}

func TestPost(t *testing.T) {
	resp, err := Post(&Request{
		URL:    "https://jsonplaceholder.typicode.com/posts",
		Client: &http.Client{},
		Params: Params{
			"headers": {
				"Content-Type": "application/json",
			},
		},
		Body: &TodoRequest{
			Title:  "Test Vider",
			Body:   "Testing Vider",
			UserID: 1,
		},
		Data: TodoResponse{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	//You will need to type cast from interface to the needed data datatype
	t.Log(resp.Body.(TodoResponse).Title)
}

func TestUpdate(t *testing.T) {
	resp, err := Put(&Request{
		URL:    "https://jsonplaceholder.typicode.com/posts/1",
		Client: &http.Client{},
		Params: Params{
			"headers": {
				"Content-Type": "application/json",
			},
		},
		Body: &TodoRequest{
			ID:     1,
			Title:  "Test Vider",
			Body:   "Testing Vider",
			UserID: 1,
		},
		Data: TodoResponse{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	//You will need to type cast from interface to the needed data datatype
	t.Log(resp.Body.(TodoResponse).Title)
}

func TestDelete(t *testing.T) {
	resp, err := Get(&Request{
		URL:    "https://jsonplaceholder.typicode.com/todos/1",
		Client: &http.Client{},
		Params: Params{
			"headers": {
				"Content-Type": "application/json",
			},
		},
		Data: TodoResponse{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
