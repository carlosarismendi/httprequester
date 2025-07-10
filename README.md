# HTTP Requester

This package provides an HTTP wrapper on top of [net/http](https://pkg.go.dev/net/http) to make http request easier. It's intended purpose is testing, so ease of use and functionality are priorities over security or any other considerations.

**NOTE**: Websocket connections were not considered when developing this package, so may not work.

## Install
```bash
go get github.com/carlosarismendi/httprequester
```
## Usage

```Go
// Create a *Requester with a URL predefined. Following usages of
// this requester will call to the configured URL. It is possible
// to create a *Requester with multiple options predefined.
//    r := NewRequester(
//        URL("https://catfact.ninja"),
//        ContentType("application/json"),
//    )
// From now on, every use of this requester will have the same base
// URL and content type header.
// An alternative way of using a requester would be to create an empty
// *Requester and defining the whole path each time we use the requester.
//    r := NewRequester()
//    r.Send(Get("https://catfact.ninja"))
//    r.Send(Get("https://google.com"))
r := NewRequester(
URL("https://catfact.ninja"),
)

// Send a request with method GET to path /fact.
// resp is an *http.Response object from "net/http".
// body is a byte array containing the response body.
// err will only have an error if there was an error running the
// request(e.g. the url doesn't exist.). In case the url exists
// but the path doesn't, for example, the err output will be nil
// and the error will be found in body and resp (and it will be
// the error returned by the host that processed the request).
resp, body, err := r.Send(Get("/fact"))

type CatFact struct {
Fact   string `json:"fact"`
Length int    `json:"length"`
}
var cf CatFact
// If you pass a reference to an object as first parameter, the
// requester will try to unmarshal the response body into it by
// calling json.Unmarshal(...) from "encoding/json".
resp, body, err := r.Send(&cf, Get("/fact"))

// It is possible to pass multiple options in a single call.
resp, body, err := r.Send(&cf,
Get("/api"),
ContentType("application/json"),
)
```

## Options

| Option          | Description                                                                                                              | Example                                       |
| --------------- | ------------------------------------------------------------------------------------------------------------------------ | --------------------------------------------- |
| `AppendPath`    | Appends a path to the current path and url.                                                                              | `AppendPath("/api/user")`                     |
| `Authorization` | Sets the HTTP header Authorization.                                                                                      | `Authorization("Bearer <token>")`             |
| `RequestBody`   | Sets the request body. Currently only marshalls body into json <br> by using `json.Encoder(...)` from `"encoding/json"`. | `Body(<an struct type object>)`               |
| `ContentType`   | Sets the HTTP header Content-Type.                                                                                       | `ContentType("application/json")`             |
| `Delete`        | Sets the HTTP method DELETE.                                                                                             | `Delete("/api/user"))`                        |
| `Doer`          | Sets the HTTP client. Default is `http.DefaultClient` from `"net/http"`.                                                 | `Doer(http.DefaultClient)`                    |
| `Get`           | Sets the HTTP method GET.                                                                                                | `Get("/api/user"))`                           |
| `Header`        | Sets an HTTP header. Overwrites previous values for the given key.                                                       | `Header("Content-Type", "application/json"))` |
| `Method`        | Sets the HTTP method.                                                                                                    | `Get("POST", "/api/user"))`                   |
| `Post`          | Sets the HTTP method POST.                                                                                               | `Post("/api/user"))`                          |
| `Put`           | Sets the HTTP method PUT.                                                                                                | `Put("/api/user"))`                           |
| `Patch`         | Sets the HTTP method PATCH.                                                                                              | `Patch("/api/user"))`                         |
| `QueryParam`    | Appends a query parameter to the path.                                                                                   | `QueryParam("color", "green")`                |
| `QueryParams`   | Appends multiple query parameters to the path.                                                                           | `QueryParams(url.Values{})`                   |
| `URL`           | Sets the base URL.                                                                                                       | `URL("https://catfact.ninja")`                |

**NOTE** Options will be processed sequentially from left to right, so in case of conflicting options (e.g. a `Post`
after a `Get`, both path defined in both options will be considered but the used method will be POST).
