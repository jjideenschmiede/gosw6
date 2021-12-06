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
    RefreshToken: "def50200fa20fc9acf83f312a18f95760b25f17844afb3abe9423e0c17ad2d950695b0b051a77c1d942b69a5886ed323c4600d5ddbcdaa851a682a9e5c80f9457da1c3fd477aacf56d7932c355e96944750d0baec8d8edc73222c62adbca65f768c7a5e0947517918cd410863e1c24fd10f5282fd0d605b6b3322ec3c8abc56f4df3178641cadffa313e35eb7d1e67efb428207477fb48ebe8b61e20f21f61791150245f6d0d32ba3a241297f3bc6cd9bffe2fe7e317f4c6ee11f2abfcb851602561b6ef72b0290a000293358965156800e85188edaf94b866adf6b0ebc1cc92cf1590cdef1b272a99055ffc79c81799c805718fc5c14ddc93dfd14217724e6b947aca6a578deb6a0854abbd346432157595ef11656b2a54212f0c0becbd133ed78017a9d5a0f3928dbb4039d98c98ba2b97a4f55d7f1e0dbac3f55ee99f1f1fee93f61fd690437ed5f776d7c44a805275966a605221f5890dfee337bb13d478c940ec806e4c820f87c9b2a46d4969036c705ebc23aa23a9e6944b805d9b3deb1ae2e8b15f537ffb3717e2f11701665bf1fdde17866726ecabd5bc45fe82",
}

// Get the access token
accessToken, err := gosw6.RefreshToken(body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(accessToken)
}
```