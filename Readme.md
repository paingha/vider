<h1 align="center"> Vider </h1> <br>

Vider is a HTTP package with Javascript Fetch like API for making HTTP requests and handing responses. This project aims to help making HTTP requests in Go easier and fun. All responses are JSON encoded automatically.

## Table of Contents

* [Installation](#installation)
* [Importing & Usage](#usage)
* [Similarities](#similarities)
* [Author](#author)
  
## Installation

```
go get github.com/paingha/vider
```

## Importing & Usage

Import the package and setup the request. URL (string), Client (http.Client{} from the net/http package), Params (vider.Params).

### Vider Get Request Example
A Get Request using Vider

```golang
import(
    "github.com/paingha/vider"
    "net/http"
    "log"
)

resp, err := vider.Get(&vider.Request{
		URL:    "https://jsonplaceholder.typicode.com/todos/1",
		Client: &http.Client{},
		Params: vider.Params{
			"headers": {
				"Content-Type": "application/json",
			},
		},
	})
if err != nil {
	log.Error(err)
}
log.Println(resp)
```

### Vider Post Request Example
A Post Request using Vider is as simple and similar as the fetch API:

```golang
import(
    "github.com/paingha/vider"
    "net/http"
    "log"
)

resp, err := vider.Post(&vider.Request{
		URL:    "https://reqres.in/api/users",
		Client: &http.Client{},
		Params: vider.Params{
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
	log.Error(err)
}
log.Println(resp)
```

### Vider Put Request Example
A Put Request using Vider sends a PATCH HTTP Request:

```golang
import(
    "github.com/paingha/vider"
    "net/http"
    "log"
)

resp, err := vider.Put(&vider.Request{
		URL:    "https://reqres.in/api/users",
		Client: &http.Client{},
		Params: vider.Params{
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
	log.Error(err)
}
log.Println(resp)
```

### Vider Delete Request Example
A Delete Request using Vider sends a DELETE HTTP Request:

```golang
import(
    "github.com/paingha/vider"
    "net/http"
    "log"
)

resp, err := vider.Delete(&vider.Request{
		URL:    "https://jsonplaceholder.typicode.com/todos/1",
		Client: &http.Client{},
		Params: vider.Params{
			"headers": {
				"Content-Type": "application/json",
			},
		},
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

```golang
import(
    "github.com/paingha/vider"
    "net/http"
    "log"
)

resp, err := vider.Get(&vider.Request{
		URL:    "https://jsonplaceholder.typicode.com/todos/1",
		Client: &http.Client{},
		Params: vider.Params{
			"headers": {
				"Content-Type": "application/json",
			},
		},
	})
if err != nil {
	log.Error(err)
}
log.Println(resp)
```



## Author

Hi! I am Paingha Joe Alagoa Jnr. I am a recent college graduate with 4 years experience using Golang and Javascript to build scalable and user friendly applications.
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
