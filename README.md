# gosw6

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jjideenschmiede/gosw6.svg)](https://golang.org/) [![Go](https://github.com/jjideenschmiede/gosw6/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/gosw6/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/jjideenschmiede/gosw6)](https://goreportcard.com/report/github.com/jjideenschmiede/gosw6) [![Go Doc](https://godoc.org/github.com/jjideenschmiede/gosw6?status.svg)](https://pkg.go.dev/github.com/jjideenschmiede/gosw6) ![Lines of code](https://img.shields.io/tokei/lines/github/jjideenschmiede/gosw6) [![Developed with <3](https://img.shields.io/badge/Developed%20with-%3C3-19ABFF)](https://jj-dev.de/)

Here you can find our library for shopware 6. We develop the API endpoints according to our demand and need. You are welcome to help us to further develop this library.
## Install

```console
go get github.com/jjideenschmiede/gosw6
```

## How to use?

Currently we have the following functions covered:

- [Access token](https://github.com/jjideenschmiede/gosw6#access-token)
- [Product](https://github.com/jjideenschmiede/gosw6#product)
- [Category](https://github.com/jjideenschmiede/gosw6#category)
- [Product manufacturer](https://github.com/jjideenschmiede/gosw6#product-manufacturer)
- [Property groups](https://github.com/jjideenschmiede/gosw6#property-groups)

## Access token

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

## Product

### Get all products

If you want to read all products, you can do it with the following function.

```go
// Define the request
r := gosw6.Request{
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
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get a product by id
product, err := gosw6.Product("a27be66bc743476089a0672290eed674", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(product)
}
```

## Category

### Get all categories

If you want to read all categories, you can do it with the following function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get all categories
categories, err := gosw6.Categories(r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(categories)
}
```

### Get a category

If you want to read a specific category, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get a category by id
category, err := gosw6.Category("a27be66bc743476089a0672290eed674", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(category)
}
```

## Product manufacturer

### Get all product manufacturers

If you want to read all product manufacturers, you can do it with the following function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get all product manufacturers
productManufacturers, err := gosw6.ProductManufacturers(r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(productManufacturers)
}
```

### Get a product manufacturer

If you want to read a specific product manufacturer, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get a product manufacturer by id
productManufacturer, err := gosw6.ProductManufacturer("a27be66bc743476089a0672290eed674", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(productManufacturer)
}
```

### Create a product manufacturer

If you want to create a product manufacturer, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := CreateProductManufacturerBody{
    MediaId:     "7ae758e4c3cb457e8b023b2c859616f7",
    Name:        "Afterbuy",
    Link:        "https://www.afterbuy.de",
    Description: "Nutzen Sie die führende All-in-One-Lösung für den E-Commerce",
}

// Create a product manufacturer
createProductManufacturer, err := CreateProductManufacturer(body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(createProductManufacturer)
}
```

### Update a product manufacturer

If you want to update a product manufacturer, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := UpdateProductManufacturerBody{
    MediaId:     "7ae758e4c3cb457e8b023b2c859616f7",
    Name:        "Afterbuy2",
    Link:        "https://www.afterbuy.de",
    Description: "Nutzen Sie die führende All-in-One-Lösung für den E-Commerce",
}

// Update a product manufacturer by id
updateProductManufacturer, err := UpdateProductManufacturer("cd38ccc066ff4400a80373ba86058df3", body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(updateProductManufacturer)
}
```

### Delete a product manufacturer

If you want to delete a product manufacturer, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Delete a product manufacturer
deleteProductManufacturer, err := DeleteProductManufacturer("cd38ccc066ff4400a80373ba86058df3", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(deleteProductManufacturer)
}
```

## Property groups

### Get all property groups

If you want to read all property groups, you can do it with the following function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get all property groups
propertyGroups, err := gosw6.PropertyGroups(r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(propertyGroups)
}
```

### Get a property group

If you want to read a specific property group, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get a property group by id
propertyGroup, err := gosw6.PropertyGroup("7c1ea10d2c3844f1ba2ab88fbcda1df2", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(propertyGroup)
}
```