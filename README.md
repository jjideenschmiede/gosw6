# Library for Shopware 6

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jjideenschmiede/gosw6.svg)](https://golang.org/) [![Go](https://github.com/jjideenschmiede/gosw6/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/gosw6/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/jjideenschmiede/gosw6)](https://goreportcard.com/report/github.com/jjideenschmiede/gosw6) [![Go Doc](https://godoc.org/github.com/jjideenschmiede/gosw6?status.svg)](https://pkg.go.dev/github.com/jjideenschmiede/gosw6) ![Lines of code](https://img.shields.io/tokei/lines/github/jjideenschmiede/gosw6) [![Developed with <3](https://img.shields.io/badge/Developed%20with-%3C3-19ABFF)](https://jj-dev.de/)

Here you can find our library for shopware 6. We develop the API endpoints according to our demand and need. You are welcome to help us to further develop this library.
## Install

```console
go get github.com/jjideenschmiede/gosw6
```

## How to use?

Currently we have the following functions covered:

- [Access token](https://github.com/jjideenschmiede/gosw6#access-token)
- [Search](https://github.com/jjideenschmiede/gosw6#search)
- [Product](https://github.com/jjideenschmiede/gosw6#product)
- [Category](https://github.com/jjideenschmiede/gosw6#category)
- [Cms Page](https://github.com/jjideenschmiede/gosw6#cms-page)
- [Manufacturer](https://github.com/jjideenschmiede/gosw6#manufacturer)
- [Property groups](https://github.com/jjideenschmiede/gosw6#property-group)
- [Property groups option](https://github.com/jjideenschmiede/gosw6#property-group-option)
- [Tax](https://github.com/jjideenschmiede/gosw6#tax)
- [Tag](https://github.com/jjideenschmiede/gosw6#tag)
- [Unit](https://github.com/jjideenschmiede/gosw6#unit)
- [Sale channels](https://github.com/jjideenschmiede/gosw6#sale-channels)
- [Currencies](https://github.com/jjideenschmiede/gosw6#currencies)
- [Delivery time](https://github.com/jjideenschmiede/gosw6#delivery-time)
- [Media](https://github.com/jjideenschmiede/gosw6#media)
- [Media folder](https://github.com/jjideenschmiede/gosw6#media-folder)
- [Order](https://github.com/jjideenschmiede/gosw6#order)

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

## Search

## Custom order search 

With this function you get all information you need for an order. We use the `associations` for that. Here is the request body:

```json
{
    "page": 1,
    "limit": 200,
    "filter": [
        {
            "type": "range",
            "field": "orderDate",
            "parameters": {
                "gte": "2021-01-09T00:00:00.000+00:00"
            }
        }
    ],
    "sort": [
        {
            "field": "orderDate",
            "order": "ASC"
        }
    ],
    "associations": {
        "currency": {},
        "lineItems": {},
        "addresses": {
            "associations": {
                "country": {},
                "countryState": {},
                "salutation": {}
            }
        },
        "transactions": {
            "associations": {
                "paymentMethod": {}
            }
        }
    }
}
```

And here you can find an example for the function. You can define the page, limit & the date since which all orders should be recorded. 

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get all orders since date (with all information you need)
orders, err := CustomOrderSearch(1, 200, "2021-12-01", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(orders)
}
```

## Custom order transactions search

With this function you get all information you need for an order. We use the `associations` for that. Here is the request body:

```json
{
  "ids": [
    "0041285a9fd64bdaa5c22718f5a15bc0"
  ],
  "associations": {
    "transactions": {
      "associations": {
        "paymentMethod": {}
      }
    }
  },
  "includes": {
    "order": [
      "transactions",
      "customFields"
    ],
    "order_transaction": [
      "paymentMethod",
      "stateMachineState",
      "customFields"
    ],
    "state_machine_state": [
      "technicalName"
    ],
    "payment_method": [
      "name",
      "updatedAt"
    ]
  }
}
```

And here you can find an example for the function. You can define the page, limit & the date since which all orders should be recorded.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get an order by id (with all information you need) 
orders, err := CustomOrderTransactionsSearch("0041285a9fd64bdaa5c22718f5a15bc0", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(orders)
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

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all products
products, err := gosw6.Products(parameter, r)
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

### Create a product

If you want to create a product, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.ProductBody{
    ParentId:           nil,
    ManufacturerId:     "6b4df46cb3bc48c5a9b7cb1ecd896ae8",
    Active:             true,
    Price:              []*gosw6.ProductBodyPrice{},
    ManufacturerNumber: "JJAW2021",
    Ean:                "1234894781274382",
    ProductNumber:      "SW1202001",
    Stock:              1000,
    DeliveryTimeId:     "b40a246351694e4089d79c4f0a7dec68",
    RestockTime:        7,
    IsCloseout:         false,
    PurchaseSteps:      1,
    MaxPurchase:        100,
    MinPurchase:        1,
    ShippingFree:       true,
    PurchasePrices:     []*gosw6.ProductBodyPurchasePrices{},
    MarkAsTopseller:    true,
    Weight:             0,
    Width:              0,
    Height:             0,
    Length:             0,
    UnitId:             "367da3d1aef34cc3a442f68143fa8ede",
    PurchaseUnit:       187,
    ReferenceUnit:      1,
    ReleaseDate:        time.Now().Format("2006-01-02T15:04:05.000+00:00"),
    Name:               "J&J Afterware6",
    Tags:               []*gosw6.ProductBodyTags{},
    Description:        "",
    MetaDescription:    "",
    MetaTitle:          "",
	Keywords:           "",
    TaxId:              "",
    Properties:         []*gosw6.ProductBodyProperties{},
	Options:            []*gosw6.ProductBodyOptions{},
    Categories:         []*gosw6.ProductBodyCategories{},
    CustomSearchKeywords: []string{
        "New custom search keyword",
    },
}

// Add price
body.Price = append(body.Price, &gosw6.ProductBodyPrice{
    CurrencyId: "b7d2554b0ce847cd82f3ac9bd1c0dfca",
    Net:        16.806722689076,
    Gross:      20,
    Linked:     true,
    ListPrice: gosw6.ProductBodyListPrice{
        CurrencyId: "b7d2554b0ce847cd82f3ac9bd1c0dfca",
        Net:        9.2436974789916,
        Gross:      11,
        Linked:     false,
    },
	RegulationPrice: gosw6.ProductBodyRegulationPrice{
        CurrencyId: "b7d2554b0ce847cd82f3ac9bd1c0dfca",
        Net:        81,
        Gross:      100,
        Linked:     false,
    },
})

// Add purchase price
body.PurchasePrices = append(body.PurchasePrices, &gosw6.ProductBodyPurchasePrices{
    CurrencyId: "b7d2554b0ce847cd82f3ac9bd1c0dfca",
    Net:        11000,
    Gross:      13090,
    Linked:     true,
})

// Add tags
body.Tags = append(body.Tags, &gosw6.ProductBodyTags{
    Id: "013a333c30e84059ba1b0025e70edec6",
})

// Add category
body.Categories = append(body.Categories, &gosw6.ProductBodyCategories{
    Id: "d9cc3869c4b64b15ad19c50088e339be",
})

// Create a product
createProduct, err := gosw6.Product(body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(createProduct)
}
```

### Update a product

If you want to update a product, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.ProductBody{
    ParentId:           nil,
    ManufacturerId:     "6b4df46cb3bc48c5a9b7cb1ecd896ae8",
    Active:             true,
    Price:              []*gosw6.ProductBodyPrice{},
    ManufacturerNumber: "JJAW2021",
    Ean:                "1234894781274382",
    ProductNumber:      "SW1202001",
    Stock:              1000,
    DeliveryTimeId:     "b40a246351694e4089d79c4f0a7dec68",
    RestockTime:        7,
    IsCloseout:         false,
    PurchaseSteps:      1,
    MaxPurchase:        100,
    MinPurchase:        1,
    ShippingFree:       true,
    PurchasePrices:     []*gosw6.ProductBodyPurchasePrices{},
    MarkAsTopseller:    true,
    Weight:             nil,
    Width:              nil,
    Height:             nil,
    Length:             nil,
    UnitId:             "367da3d1aef34cc3a442f68143fa8ede",
    PurchaseUnit:       187,
    ReferenceUnit:      1,
    ReleaseDate:        time.Now().Format("2006-01-02T15:04:05.000+00:00"),
    Name:               "J&J Afterware6 v2",
    Tags:               []*gosw6.ProductBodyTags{},
    Description:        "",
    MetaDescription:    "",
    MetaTitle:          "",
    TaxId:              "",
    Properties:         []*gosw6.ProductBodyProperties{},
    Options:            []*gosw6.ProductBodyOptions{},
    Categories:         []*gosw6.ProductBodyCategories{},
    CustomSearchKeywords: []string{
        "New custom search keyword",
    },
}

// Add price
body.Price = append(body.Price, &gosw6.ProductBodyPrice{
    CurrencyId: "b7d2554b0ce847cd82f3ac9bd1c0dfca",
    Net:        16.806722689076,
    Gross:      20,
    Linked:     true,
    ListPrice: gosw6.ProductBodyListPrice{
        CurrencyId: "b7d2554b0ce847cd82f3ac9bd1c0dfca",
        Net:        9.2436974789916,
        Gross:      11,
        Linked:     false,
    },
    RegulationPrice: gosw6.ProductBodyRegulationPrice{
        CurrencyId: "b7d2554b0ce847cd82f3ac9bd1c0dfca",
        Net:        81,
        Gross:      100,
        Linked:     false,
    },
})

// Add purchase price
body.PurchasePrices = append(body.PurchasePrices, &gosw6.ProductBodyPurchasePrices{
    CurrencyId: "b7d2554b0ce847cd82f3ac9bd1c0dfca",
    Net:        11000,
    Gross:      13090,
    Linked:     true,
})

// Add tags
body.Tags = append(body.Tags, &gosw6.ProductBodyTags{
    Id: "013a333c30e84059ba1b0025e70edec6",
})

// Add category
body.Categories = append(body.Categories, &gosw6.ProductBodyCategories{
    Id: "d9cc3869c4b64b15ad19c50088e339be",
})

// Update a product
updateProduct, err := gosw6.UpdateProduct("07caca75115742a3bbc3692963c7e2d3", body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(updateProduct)
}
```

### Delete a product

If you want to delete a product, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Delete a product by id
deleteProduct, err := gosw6.DeleteProduct("529eadb5f9a34946ad94caa8441f267a", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(deleteProduct)
}
```

### Get all product configuration settings

If you want to read all product configuration settings, you can do it with the following function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all product configuration settings
productConfigurationSettings, err := gosw6.ProductConfiguratorSettings(parameter, "c448b82288914946b1ca4a93bcbf824f", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(productConfigurationSettings)
}
```

### Update product configuration settings

If you want to update product configuration settings, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
updateProductConfiguratorSettingBody := UpdateProductConfiguratorSettingBody{
    ConfiguratorSettings: []UpdateProductConfiguratorSettingBodyConfiguratorSettings{},
}

updateProductConfiguratorSettingBody.ConfiguratorSettings = append(updateProductConfiguratorSettingBody.ConfiguratorSettings, UpdateProductConfiguratorSettingBodyConfiguratorSettings{"83f39523e6af4e248b5fb0bbb7b8837b"})
updateProductConfiguratorSettingBody.ConfiguratorSettings = append(updateProductConfiguratorSettingBody.ConfiguratorSettings, UpdateProductConfiguratorSettingBodyConfiguratorSettings{"c41d245ae1c6416f94267dba8104437a"})
updateProductConfiguratorSettingBody.ConfiguratorSettings = append(updateProductConfiguratorSettingBody.ConfiguratorSettings, UpdateProductConfiguratorSettingBodyConfiguratorSettings{"9513190a41a44347b054c47eb62e247d"})

// Update product configuration settings
updateProductConfiguratorSettings, err := UpdateProductConfiguratorSetting("9553fd97f5af4d77a2e8befcb83ed697", updateProductConfiguratorSettingBody, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(updateProductConfiguratorSettings)
}
```

### Delete a product configuration setting

If you want to delete a product configuration setting, you can do it with the following function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all product configuration settings
productConfigurationSettings, err := gosw6.ProductConfiguratorSettings(parameter, "c448b82288914946b1ca4a93bcbf824f", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(productConfigurationSettings)
}
```

### Update product stock

If you want to update product stock, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.UpdateProductStockBody{
    Stock:       31,
    MaxPurchase: 31,
}

// Update product stock
updateProductStock, err := gosw6.UpdateProductStock("24c3b6ebf71746b08ec8a6a1283ff0cb", body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(updateProductStock)
}
```

### Get all product properties

If you want to get all product properties, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all product properties
productProperties, err := gosw6.ProductProperties(parameter, "c1ed726029374656bdf22f551dd7fbea", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(productProperties)
}
```

### Delete a product property

If you want to delete a product property, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Delete a product property
delteProductProperty, err := gosw6.DeleteProductProperty("c1ed726029374656bdf22f551dd7fbea", "177996f2c2674d49ac3c863d4f643c5f", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(delteProductProperty)
}
```

### Get all product categories

If you want to get all product categories, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all product categories
productCategories, err := gosw6.ProductCategories(parameter, "c1ed726029374656bdf22f551dd7fbea", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(productCategories)
}
```

### Delete a product category

If you want to delete a product categories, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Delete a product category
delteProductCategory, err := gosw6.DeleteProductCategory("c1ed726029374656bdf22f551dd7fbea", "12aee528abbb446cb46444a0c1fd3361", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(delteProductCategory)
}
```

### Get all product tags

If you want to get all product tags, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all product tags
productTags, err := gosw6.ProductTags(parameter, "c1ed726029374656bdf22f551dd7fbea", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(productTags)
}
```

### Delete a product tag

If you want to delete a product tag, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Delete a product tag
delteProductTags, err := gosw6.DeleteProductTag("c1ed726029374656bdf22f551dd7fbea", "08cf4f62c02a48858928db7a40377c18", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(delteProductTags)
}
```

### Get all product manufacturers

If you want to get all product manufacturers, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all product manufacturers
productManufacturer, err := gosw6.ProductManufacturers(parameter, "b398f53c50f243aaacc38a68659f71e1", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(productManufacturer)
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
delteProductManufacturer, err := gosw6.DeleteProductManufacturer("b398f53c50f243aaacc38a68659f71e1", "a1b5a538e2ff4ac78034741873cee064", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(delteProductManufacturer)
}
```

### Get all product media

If you want to get all product media, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all product media
productMedia, err := gosw6.ProductMedia(parameter, "529eadb5f9a34946ad94caa8441f267a", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(productMedia)
}
```

### Create product media

If you want to create a product media to an product, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
createProductMediaBody := ProductMediaBody{
    MediaId:  "9be5fdf920454ed78a629195c86936fe",
    Position: 1,
}

// Crate product media
createProductMedia, err := gosw6.CreateProductMedia("07caca75115742a3bbc3692963c7e2d3", body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(createProductMedia)
}
```

### Get all product covers

If you want to get all covers of a product, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get all product covers
productCover, err := ProductCover(map[string]string{}, "dddbf38937df464690f25b18a32cdeb0", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(productCover)
}
```

### Create product cover

If you want to create a cover to an product, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
createProductCoverBody := CreateProductCoverBody{
    ProductId: "dddbf38937df464690f25b18a32cdeb0",
    MediaId:   "c97dc00288574ed394c5f9ee6012cd9a",
	Position:  1,
}

// Create a product cover
createProductCover, err := gosw6.CreateProductCover("dddbf38937df464690f25b18a32cdeb0", createProductCoverBody, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(createProductCover)
}
```

### Get all product visibilities

If you want to get all visibilities of a product, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get all product visibilities
productVisibilities, err := gosw6.ProductVisibilities(map[string]string{}, "a88079a02bf146179d706c58d7294031", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(productVisibilities)
}
```

### Create product visibility

If you want to create a visibility to an product, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
createProductVisibilityBody := gosw6.ProductVisibilityBody{
    Visibility:     30,
    SalesChannelId: "1ac37e86b4f1408f8319c99cee52d772",
}

// Create product visibility
createProductVisibility, err := gosw6.CreateProductVisibility("a88079a02bf146179d706c58d7294031", createProductVisibilityBody, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(createProductVisibility)
}
```

### Get all product prices 

If you want to get all prices of a product, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get all product prices
productPrices, err := gosw6.ProductPrices(map[string]string{}, "a88079a02bf146179d706c58d7294031", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(productPrices)
}
```

### Create product price 

If you want to create a price to a product, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
createProductPriceBody := gosw6.ProductPriceBody{
    QuantityStart: 11,
    QuantityEnd:   20,
    RuleId:        "26a7fc74519d4622a22b8f4dbcb015a3",
    Price: []gosw6.ProductPriceBodyPrice{
        {
            CurrencyId: "b7d2554b0ce847cd82f3ac9bd1c0dfca",
            Net:        19.99,
            Gross:      12.99,
            Linked:     true,
        },
    },
}

// Create product price
createProductPrice, err := gosw6.CreateProductPrice("f58d1ed3c59642f08fc882897f3e9758", createProductPriceBody, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(createProductPrice)
}
```

### Delete a product price 

If you want to delete a product price, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Delete a product price 
delteProductPrice, err := gosw6.DeleteProductPrice("b398f53c50f243aaacc38a68659f71e1", "a1b5a538e2ff4ac78034741873cee064", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(delteProductPrice)
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

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all categories
categories, err := gosw6.Categories(parameter, r)
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

### Create a category

If you want to create a category, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.CategoryBody{
    ParentId:              "ec1dc66781f04ec3bd7e5df94cb74b5a",
    MediaId:               nil,
    Name:                  "New category",
    Active:                true,
    VisibleChildCount:     0,
    DisplayNestedProducts: true,
    CmsPageId:             "6de66fe08d4544e397ed99d7efd4c32d",
    LinkType:              nil,
    LinkNewTab:            nil,
    InternalLink:          nil,
    ExternalLink:          nil,
    Visible:               false,
    Type:                  "page",
    ProductAssignmentType: "product",
    Description:           "New Description!",
    MetaTitle:             "Nice meta title",
    MetaDescription:       "Nice meta description!",
    Keywords:              "category, nice",
}

// Create a category
createCategory, err := gosw6.CreateCategory(body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(createCategory)
}
```

### Update a category

If you want to update a category, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.CategoryBody{
    ParentId:              "ec1dc66781f04ec3bd7e5df94cb74b5a",
    MediaId:               nil,
    Name:                  "New category v2",
    Active:                true,
    VisibleChildCount:     0,
    DisplayNestedProducts: true,
    CmsPageId:             "6de66fe08d4544e397ed99d7efd4c32d",
    LinkType:              nil,
    LinkNewTab:            nil,
    InternalLink:          nil,
    ExternalLink:          nil,
    Visible:               false,
    Type:                  "page",
    ProductAssignmentType: "product",
    Description:           "New Description v2!",
    MetaTitle:             "Nice meta title",
    MetaDescription:       "Nice meta description!",
    Keywords:              "category, nice",
}

// Update a category by id
updateCategory, err := gosw6.UpdateCategory("5967746e6cbd444bb4d3d7f060c58aea", body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(updateCategory)
}
```

### Delete a category

If you want to delete a category, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Delete a category by id
deleteCategory, err := gosw6.DeleteCategory("5967746e6cbd444bb4d3d7f060c58aea", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(deleteCategory)
}
```

## Cms Page

### Get all cms pages

If you want to read all cms pages, you can do it with the following function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all cms pages
cmsPages, err := gosw6.CmsPages(parameter, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(cmsPages)
}
```

## Manufacturer

### Get all manufacturers

If you want to read all manufacturers, you can do it with the following function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all manufacturers
manufacturers, err := gosw6.Manufacturers(parameter, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(manufacturers)
}
```

### Get a manufacturer

If you want to read a specific manufacturer, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get a manufacturer by id
manufacturer, err := gosw6.Manufacturer("a27be66bc743476089a0672290eed674", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(manufacturer)
}
```

### Create a manufacturer

If you want to create a manufacturer, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.ManufacturerBody{
    MediaId:     "7ae758e4c3cb457e8b023b2c859616f7",
    Name:        "Afterbuy",
    Link:        "https://www.afterbuy.de",
    Description: "Nutzen Sie die führende All-in-One-Lösung für den E-Commerce",
}

// Create a manufacturer
createManufacturer, err := gosw6.CreateManufacturer(body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(createManufacturer)
}
```

### Update a manufacturer

If you want to update a manufacturer, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.ManufacturerBody{
    MediaId:     "7ae758e4c3cb457e8b023b2c859616f7",
    Name:        "Afterbuy2",
    Link:        "https://www.afterbuy.de",
    Description: "Nutzen Sie die führende All-in-One-Lösung für den E-Commerce",
}

// Update a product manufacturer by id
updateProductManufacturer, err := gosw6.UpdateProductManufacturer("cd38ccc066ff4400a80373ba86058df3", body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(updateProductManufacturer)
}
```

### Delete a manufacturer

If you want to delete a manufacturer, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Delete a manufacturer by id
deleteManufacturer, err := gosw6.DeleteManufacturer("cd38ccc066ff4400a80373ba86058df3", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(deleteManufacturer)
}
```

## Property group

### Get all property groups

If you want to read all property groups, you can do it with the following function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all property groups
propertyGroups, err := gosw6.PropertyGroups(parameter, r)
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

### Create a property group

If you want to create a property group, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.PropertyGroupBody{
    Name:                       "Size",
    DisplayType:                "text",
    SortingType:                "alphanumeric",
    Description:                "To define the shoe size.",
    Position:                   1,
    Filterable:                 true,
    VisibleOnProductDetailPage: true,
}

// Create a property group
createPropertyGroup, err := gosw6.CreatePropertyGroup(body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(createPropertyGroup)
}
```

### Update a property group

If you want to update a property group, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.PropertyGroupBody{
    Name:                       "Size",
    DisplayType:                "text",
    SortingType:                "alphanumeric",
    Description:                "To define the shoe size.",
    Position:                   1,
    Filterable:                 true,
    VisibleOnProductDetailPage: true,
}

// Update a property group by id
updatePropertyGroup, err := gosw6.UpdatePropertyGroup("272c3f25999e4f779c28db479ad0af5c", body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(updatePropertyGroup)
}
```

### Delete a property group

If you want to delete a property group, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Update a property group by id
deletePropertyGroup, err := gosw6.DeletePropertyGroup("272c3f25999e4f779c28db479ad0af5c", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(deletePropertyGroup)
}
```

## Property group option

### Get all property group options

If you want to read all property group options, you can do it with the following function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all property group options
propertyGroupOptions, err := gosw6.PropertyGroupOptions("7c1ea10d2c3844f1ba2ab88fbcda1df2", parameter, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(propertyGroupOptions)
}
```

### Get a property group option

If you want to read a specific property group option, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get all property group option by id
propertyGroupOption, err := gosw6.PropertyGroupOption("aee4401b771748319f2651c2e113f267", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(propertyGroupOption)
}
```

### Create a property group option

If you want to create a property group option, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.PropertyGroupOptionBody{
    Name:         "Test1",
    Position:     1,
    ColorHexCode: "",
    MediaId:      "",
}

// Create a property group option
createPropertyGroupOption, err := gosw6.CreatePropertyGroupOption("7c1ea10d2c3844f1ba2ab88fbcda1df2", body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(createPropertyGroupOption)
}
```

### Update a property group option

If you want to update a property group option, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.PropertyGroupOptionBody{
    Name:         "Test2",
    Position:     1,
    ColorHexCode: "",
    MediaId:      "",
    }

// Update a property group option
updatePropertyGroupOption, err := gosw6.UpdatePropertyGroupOption("830cbe62c1b943c5b5ffa4895b48efd6", body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(updatePropertyGroupOption)
}
```

### Delete a property group option

If you want to delete a property group option, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Delete a property group option 
deletePropertyGroupOption, err := gosw6.DeletePropertyGroupOption("830cbe62c1b943c5b5ffa4895b48efd6", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(deletePropertyGroupOption)
}
```

## Tax

### Get all taxes

If you want to read all taxes, you can do it with the following function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all taxes
taxes, err := gosw6.Taxes(parameter, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(taxes)
}
```

## Tag

### Get all tags

If you want to read all tags, you can do it with the following function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all tags
tags, err := gosw6.Tags(parameter, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(tags)
}
```

### Delete a tag

If you want to delete a tag, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Delete a tag 
deleteTag, err := gosw6.DeleteTag("33fa1a797a0b49a3b90f85160521cafc", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(deleteTag)
}
```

## Unit

### Get all units

If you want to read all units, you can do it with the following function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all units
units, err := gosw6.Units(parameter, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(units)
}
```

## Rule

### Get all rules

If you want to read all rules, you can do it with the following function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all rules 
rules, err := gosw6.Rules(parameter, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(rules)
}
```

## Sale channels

### Get all sale channels

If you want to read all sale channels, you can do it with the following function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all sale channels
saleChannels, err := gosw6.SalesChannels(parameter, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(saleChannels)
}
```

## Currencies

### Get all currencies

If you want to read all currencies, you can do it with the following function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all currencies
saleChannels, err := gosw6.Currencies(parameter, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(saleChannels)
}
```

## Delivery time

### Get all delivery times

If you want to read all delivery times, you can do it with the following function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all delivery times
deliveryTimes, err := gosw6.DeliveryTimes(parameter, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(deliveryTimes)
}
```

### Create a delivery time

If you want to create a delivery time, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.CreateDeliveryTimesBody{
    Name: "Nicht angegeben",
    Min:  1,
    Max:  7,
    Unit: "day",
}

// Create a delivery time
createDeliveryTime, err := gosw6.CreateDeliveryTimes(body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(createDeliveryTime)
}
```

## Media

### Get all media

If you want to get a list of all media, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all media 
media, err := gosw6.Media(parameter, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(media)
}
```

### Create a media

If you want to create a media, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.CreateMediaBody{
    MediaFolderId: "bce5a570d3004c9c86471c7acb661593",
}

// Create a media
createMedia, err := gosw6.CreateMedia(body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(createMedia)
}
```

### Upload a media from url

If you want to create a media from url, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.UploadUrlMediaBody{
    Url: "https://jj-ideenschmiede.de/img/jan-droste.de11537e.jpg",
}

// Upload a media via url
uploadMedia, err := gosw6.UploadUrlMedia("jan-droste", "jpg", "1e3c74f2e6f448eeb660251a27e41165", body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(uploadMedia)
}
```

### Upload a media from local

If you want to upload a media from local, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Open file
file, err := os.Open("/Users/jonaskwiedor/Downloads/weihnachts-postkarte.pdf")
if err != nil {
    log.Fatalln(err)
}

// Upload a media from local
uploadMedia, err := gosw6.UploadLocalMedia(file, "weihnachts-postkarte", "pdf", "e747cc2d5a684131b972469c1e2bc815", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(uploadMedia)
}
```

### Delete a media

If you want to delete a media, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Delete a media 
deleteMedia, err := gosw6.DeleteMedia("33fa1a797a0b49a3b90f85160521cafc", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(deleteMedia)
}
```

## Media folder

### Get all media folders

If you want to get a list of all media folders, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get all media folders
mediaFolders, err := gosw6.MediaFolders(parameter, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(mediaFolders)
}
```

### Get a media folder

If you want to read a specific media folder, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get a media folder by id
mediaFolder, err := gosw6.MediaFolder("bce5a570d3004c9c86471c7acb661593", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(mediaFolder)
}
```

### Create a media folder

If you want to create a media folder option, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.MediaFolderBody{
    ConfigurationId:        "a0ae2934fc094b3086b06196a1125320",
    Name:                   "test-product",
    ParentId:               "bce5a570d3004c9c86471c7acb661593",
    UseParentConfiguration: true,
}

// Create a media folder
createMediaFolder, err := gosw6.CreateMediaFolder(body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(createMediaFolder)
}
```

### Update a media folder

If you want to update a media folder, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.MediaFolderBody{
    ConfigurationId:        "a0ae2934fc094b3086b06196a1125320",
    Name:                   "test-product-v2",
    ParentId:               "bce5a570d3004c9c86471c7acb661593",
    UseParentConfiguration: true,
}

// Update a media folder
updateMediaFolder, err := gosw6.UpdateMediaFolder("035b6e7440d04fc4b0a982cbce4a27c7", body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(updateMediaFolder)
}
```

### Delete a media folder

If you want to delete a media folder, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Delete a media folder
deleteMediaFolder, err := gosw6.DeleteMediaFolder("035b6e7440d04fc4b0a982cbce4a27c7", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(deleteMediaFolder)
}
```

## Order

### Get all orders

If you want to get a list of all orders, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Use parameter
parameter := make(map[string]string)
parameter["limit"] = "5"
parameter["page"] = "1"

// Get a list of all orders
orders, err := gosw6.Orders(parameter, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(orders)
}
```

### Get a order

If you want to get an order by id, you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Get an order by id
order, err := gosw6.Order("82cb57d2e4f645f2b5c39c9913f9354f", r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(order)
}
```

### Change order transaction state

If you want to change the order transaction state by id, you can do it with this function.

The State is deposited here as follows. Please always specify the int.

| Integer |                Value                 |
|---------|:------------------------------------:|
| 0       |              authorize               |
| 1       |                 paid                 |
| 2       |                remind                |
| 3       |                refund                |
| 4       |                 fail                 |
| 5       |                do_pay                |
| 6       |                reopen                |
| 7       |              chargeback              |
| 8       |            paid_partially            |
| 9       |           refund_partially           |
| 10      |         process_unconfirmed          |
| 11      |                cancel                |

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := OrderTransactionStateBody{
    DocumentIds: nil,
    SendMail:    false,
}

// Change order transaction state
orderTransactionState, err := OrderTransactionState("10f36343f5f2440ca439569598b9a0ff", 11, body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(orderTransactionState)
}
```

### Change order delivery state

If you want to change the order delivery state by id, you can do it with this function.

The State is deposited here as follows. Please always specify the int.

| Integer |           Value            |
|---------|:--------------------------:|
| 0       |           reopen           |
| 1       |           retour           |
| 2       |      retour_partially      |
| 3       |            ship            |
| 4       |       ship_partially       |
| 5       |           cancel           |

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.OrderDeliveryStateBody{
    DocumentIds: nil,
    SendMail:    false,
}

// Change order delivery state
orderDeliveryState, err := gosw6.OrderDeliveryState("8d406c6191204dc69e14048409be74fc", 3, body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(orderDeliveryState)
}
```

### Change order state

If you want to change the order state by id, you can do it with this function.

The State is deposited here as follows. Please always specify the int.

| Integer |  Value   |
|---------|:--------:|
| 0       |  reopen  |
| 1       | process  |
| 2       | complete |
| 3       |  cancel  |

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := gosw6.OrderStateBody{
    DocumentIds: nil,
    SendMail:    false,
}

// Change order state
orderState, err := gosw6.OrderState("82cb57d2e4f645f2b5c39c9913f9354f", 1, body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(orderState)
}
```

### Update order tracking number

If you want to update the order tracking number you can do it with this function.

```go
// Define the request
r := gosw6.Request{
    BaseUrl:     "https://shopware-demo.test.de",
    BearerToken: "eyJ0eXAiOiJKV1QiLCJhbGciOiJ...",
}

// Define body
body := UpdateOrderTrackingNumberBody{}

updateOrderBody.Deliveries = append(body.Deliveries, UpdateOrderTrackingNumberBodyDeliveries{
    Id:            "407928d7e48047c5a97c31bc55ca0498",
    TrackingCodes: []string{"187DE"},
})

// Update order tracking number
updateOrder, err := UpdateOrderTrackingNumber("61e17e50646c43d5b0c77ab23bedb008", body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(updateOrder)
}
```