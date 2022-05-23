//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2022 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of gosw6.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package gosw6

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// CurrenciesReturn is to decode the json return
type CurrenciesReturn struct {
	Total int `json:"total"`
	Data  []struct {
		IsoCode                        string      `json:"isoCode"`
		Factor                         float64     `json:"factor"`
		Symbol                         string      `json:"symbol"`
		ShortName                      string      `json:"shortName"`
		Name                           string      `json:"name"`
		Position                       int         `json:"position"`
		Translations                   interface{} `json:"translations"`
		Orders                         interface{} `json:"orders"`
		SalesChannels                  interface{} `json:"salesChannels"`
		SalesChannelDefaultAssignments interface{} `json:"salesChannelDefaultAssignments"`
		SalesChannelDomains            interface{} `json:"salesChannelDomains"`
		ShippingMethodPrices           interface{} `json:"shippingMethodPrices"`
		PromotionDiscountPrices        interface{} `json:"promotionDiscountPrices"`
		IsSystemDefault                bool        `json:"isSystemDefault"`
		ProductExports                 interface{} `json:"productExports"`
		CountryRoundings               interface{} `json:"countryRoundings"`
		ItemRounding                   struct {
			Decimals    int           `json:"decimals"`
			Interval    float64       `json:"interval"`
			RoundForNet bool          `json:"roundForNet"`
			Extensions  []interface{} `json:"extensions"`
			ApiAlias    string        `json:"apiAlias"`
		} `json:"itemRounding"`
		TotalRounding struct {
			Decimals    int           `json:"decimals"`
			Interval    float64       `json:"interval"`
			RoundForNet bool          `json:"roundForNet"`
			Extensions  []interface{} `json:"extensions"`
			ApiAlias    string        `json:"apiAlias"`
		} `json:"totalRounding"`
		TaxFreeFrom      int         `json:"taxFreeFrom"`
		UniqueIdentifier string      `json:"_uniqueIdentifier"`
		VersionId        interface{} `json:"versionId"`
		Translated       struct {
			ShortName    string `json:"shortName"`
			Name         string `json:"name"`
			CustomFields struct {
			} `json:"customFields"`
		} `json:"translated"`
		CreatedAt  time.Time   `json:"createdAt"`
		UpdatedAt  interface{} `json:"updatedAt"`
		Extensions struct {
			ForeignKeys struct {
				ApiAlias   interface{}   `json:"apiAlias"`
				Extensions []interface{} `json:"extensions"`
			} `json:"foreignKeys"`
		} `json:"extensions"`
		Id           string      `json:"id"`
		CustomFields interface{} `json:"customFields"`
		ApiAlias     string      `json:"apiAlias"`
	} `json:"data"`
	Aggregations []interface{} `json:"aggregations"`
	Errors       []struct {
		Code   string `json:"code"`
		Status string `json:"status"`
		Title  string `json:"title"`
		Detail string `json:"detail"`
		Meta   struct {
			Trace []struct {
				File     string `json:"file"`
				Line     int    `json:"line"`
				Function string `json:"function"`
				Class    string `json:"class"`
				Type     string `json:"type"`
			} `json:"trace"`
			File string `json:"file"`
			Line int    `json:"line"`
		} `json:"meta"`
	} `json:"errors"`
}

// Currencies are to get a list of all currencies
func Currencies(parameter map[string]string, r Request) (CurrenciesReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/currency",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return CurrenciesReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return CurrenciesReturn{}, err
	}

	for index, value := range parameter {
		newUrl.Add(index, value)
	}

	// Set new url
	parse.RawQuery = newUrl.Encode()
	c.Path = fmt.Sprintf("%s", parse)

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CurrenciesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CurrenciesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return CurrenciesReturn{}, err
	}

	// Return data
	return decode, err

}
