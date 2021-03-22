// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vider

import (
	"net/http"
	"testing"
)

type TodoResponse struct {
	Title  string
	ID     int
	UserID int
}

type TodoRequest struct {
	Title  string
	ID     int
	UserID int
}

//TestGet - tests the get method
func TestGet(t *testing.T) {
	mine := &TodoResponse{}
	_, err := Get(&Request{
		URL:    "https://jsonplaceholder.typicode.com/todos/1",
		Client: &http.Client{},
		Params: Params{
			"headers": {
				"Content-Type": "application/json",
			},
		},
		Data: mine,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(mine)
}

func TestPost(t *testing.T) {
	newTodoRequest := &TodoRequest{}
	newTodoResponse := &TodoResponse{}
	resp, err := Post(&Request{
		URL:    "https://reqres.in/api/users",
		Client: &http.Client{},
		Params: Params{
			"headers": {
				"Content-Type": "application/json",
			},
		},
		Body: newTodoRequest,
		Data: newTodoResponse,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestUpdate(t *testing.T) {
	dataResponse := &TodoResponse{}
	dataRequest := &TodoRequest{}
	resp, err := Put(&Request{
		URL:    "https://reqres.in/api/users",
		Client: &http.Client{},
		Params: Params{
			"headers": {
				"Content-Type": "application/json",
			},
		},
		Body: dataRequest,
		Data: dataResponse,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	t.Log(dataResponse)
}

func TestDelete(t *testing.T) {
	dataResponse := &TodoResponse{}
	resp, err := Get(&Request{
		URL:    "https://jsonplaceholder.typicode.com/todos/1",
		Client: &http.Client{},
		Params: Params{
			"headers": {
				"Content-Type": "application/json",
			},
		},
		Data: dataResponse,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
