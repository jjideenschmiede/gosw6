# gosw6

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jjideenschmiede/gosw6.svg)](https://golang.org/) [![Go](https://github.com/jjideenschmiede/gosw6/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/gosw6/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/jjideenschmiede/gosw6)](https://goreportcard.com/report/github.com/jjideenschmiede/gosw6) [![Go Doc](https://godoc.org/github.com/jjideenschmiede/gosw6?status.svg)](https://pkg.go.dev/github.com/jjideenschmiede/gosw6) ![Lines of code](https://img.shields.io/tokei/lines/github/jjideenschmiede/gosw6) [![Developed with <3](https://img.shields.io/badge/Developed%20with-%3C3-19ABFF)](https://jj-dev.de/)

Here you can find our library for shopware 6. We develop the API endpoints according to our demand and need. You are welcome to help us to further develop this library.
## Install

```console
go get github.com/jjideenschmiede/gosw6
```

## How to use?

### Obtain an access token

If you want to create a new Bearer Token, you can do this with the following function. You get the Bearer Token and a Refresh Token back. With this Refresh Token, you can always generate a new Bearer Token before expiration. [Here](https://shopware.stoplight.io/docs/admin-api/ZG9jOjEwODA3NjQx-authentication-and-authorisation) you can find the documentation.

```go
// Define the request
r := gosw6.Request{
    BaseUrl: "https://shopware-demo.test.de",
}

// Define the body
body := gosw6.AccessTokenBody{
    ClientId:  "administration",
    GrantType: "password",
    Scopes:    "write",
    Username:  "gowizzward",
    Password:  "nicePassword!?2021",
}

// Get the access token
accessToken, err := gosw6.AccessToken(body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(accessToken)
}
```


### Obtain a refresh token

If you want to create a new Bearer Token with the Refresh Token, then you can do this with the following function. Please remember that you will also get a new Refresh Token back from Shopware. [Here](https://shopware.stoplight.io/docs/admin-api/ZG9jOjEwODA3NjQx-authentication-and-authorisation) you can find the documentation.

```go
// Define the request
r := gosw6.Request{
    BaseUrl: "https://shopware-demo.test.de",
}

// Define the body
body := gosw6.RefreshTokenBody{
    GrantType:    "refresh_token",
    ClientId:     "administration",
    RefreshToken: "def50200fa2...",
}

// Get the access token
accessToken, err := gosw6.RefreshToken(body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(accessToken)
}
```

### Get all products

If you want to read all products, you can do it with the following function.

```go
// Define the request
r := Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get all products
products, err := gosw6.Products(r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(products)
}
```

### Get a product

If you want to read a specific product, you can do it with this function.

```go
// Define the request
r := Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get all products
product, err := Product("a27be66bc743476089a0672290eed674", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(product)
}
```