// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net/http"
	"testing"
)

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
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestPost(t *testing.T) {
	resp, err := Post(&Request{
		URL:    "https://reqres.in/api/users",
		Client: &http.Client{},
		Params: Params{
			"headers": {
				"Content-Type": "application/json",
			},
			"body": {
				"name": "hello world",
				"job":  "Software Dev",
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestDelete(t *testing.T) {
	resp, err := Delete(&Request{
		URL:    "https://reqres.in/api/users/2",
		Client: &http.Client{},
		Params: Params{
			"headers": {
				"Content-Type": "application/json",
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
