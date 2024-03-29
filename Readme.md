<h1 align="center"> Vider </h1> <br>

Vider is a HTTP package with Javascript Fetch like API for making HTTP requests and handing responses. This project aims to help making HTTP requests in Go easier and fun. All responses are JSON encoded automatically.

Vider v1.0.1 has been released 🎉 🎉 🎉.


## Table of Contents

* [Installation](#installation)
* [Importing & Usage](#usage)
* [Similarities](#similarities)
* [Changelog](#chanelog)
* [Author](#author)
  
## Installation

```
go get github.com/paingha/vider@v1.0.1
```

## Importing & Usage

Import the package and setup the request. URL (string), Client (http.Client{} from the net/http package), Params (vider.Params).

### Vider Get Request Example
A Get Request using Vider

vider.Get - returns a response (which has a StatusCode and a StatusMessage) and an error

```golang
import(
    "github.com/paingha/vider"
    "net/http"
    "log"
)
type TodoResponse struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
}
resp, err := vider.Get(&vider.Request{
		URL:    "https://jsonplaceholder.typicode.com/todos/1",
		Client: &http.Client{},
		Params: vider.Params{
			"headers": {
				"Content-Type": "application/json",
			},
		},
		Data: TodoResponse{},
	})
if err != nil {
	log.Error(err)
}
log.Println(resp)
//You will need to type cast from interface to the needed data datatype
log.Println(resp.Body.(TodoResponse).Title)
```

### Vider Post Request Example
A Post Request using Vider is as simple and similar as the fetch API:

vider.Post - returns a response (which has a StatusCode and a StatusMessage) and an error

```golang
import(
    "github.com/paingha/vider"
    "net/http"
    "log"
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
resp, err := vider.Post(&vider.Request{
		URL:    "https://jsonplaceholder.typicode.com/posts",
		Client: &http.Client{},
		Params: vider.Params{
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
	log.Error(err)
}
log.Println(resp)
//You will need to type cast from interface to the needed data datatype
log.Println(resp.Body.(TodoResponse).Title)
```

### Vider Put Request Example
A Put Request using Vider sends a PATCH HTTP Request:

vider.Put - returns a response (which has a StatusCode and a StatusMessage) and an error

```golang
import(
    "github.com/paingha/vider"
    "net/http"
    "log"
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
resp, err := vider.Put(&vider.Request{
		URL:    "https://jsonplaceholder.typicode.com/posts/1",
		Client: &http.Client{},
		Params: vider.Params{
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
	log.Error(err)
}
log.Println(resp)
//You will need to type cast from interface to the needed data datatype
log.Println(resp.Body.(TodoResponse).Title)
```

### Vider Delete Request Example
A Delete Request using Vider sends a DELETE HTTP Request:

vider.Delete - returns a response (which has a StatusCode and a StatusMessage) and an error

```golang
import(
    "github.com/paingha/vider"
    "net/http"
    "log"
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
resp, err := vider.Delete(&vider.Request{
		URL:    "https://jsonplaceholder.typicode.com/todos/1",
		Client: &http.Client{},
		Params: vider.Params{
			"headers": {
				"Content-Type": "application/json",
			},
		},
		Data: TodoResponse{},
	})
if err != nil {
	log.Error(err)
}
log.Println(resp)
```


## Similarities

Javascript Fetch Request

```javascript
fetch('https://jsonplaceholder.typicode.com/todos/1', {
	method: 'GET',
	headers: {
    'Content-Type': 'application/json'
  	},
})
```

Vider Get Request


vider.Get - returns a response (which has a StatusCode and a StatusMessage) and an error

```golang
import(
    "github.com/paingha/vider"
    "net/http"
    "log"
)
type TodoResponse struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
}
resp, err := vider.Get(&vider.Request{
		URL:    "https://jsonplaceholder.typicode.com/todos/1",
		Client: &http.Client{},
		Params: vider.Params{
			"headers": {
				"Content-Type": "application/json",
			},
		},
		Data: TodoResponse{},
	})
if err != nil {
	log.Error(err)
}
log.Println(resp)
//You will need to type cast from interface to the needed data datatype
log.Println(resp.Body.(TodoResponse).Title)
```

## Changelog

### v1.0 - Official production release
<ul>
<li>Added JSON to Struct serialization for responses</li>
<li>Added Tests and Github workflow to run them</li>
</ul>

## Author

Hi! I am Paingha Joe Alagoa Jnr. I am a recent college graduate with 5 years experience using Golang and Javascript to build scalable and user friendly applications.
<br/>
<br />
<strong>Currently Open to New Full-time Software Developer positions in the United States (both remote and in person roles) for Golang and Javascript. </strong>
<br />
<br />
Contact me by Twitter or Linked In :)
<br/>
[Linked In](https://www.linkedin.com/in/paingha-alagoa-joe/)
<br />
[Twitter](https://twitter.com/painghajnr)
