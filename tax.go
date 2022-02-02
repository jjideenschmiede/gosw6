//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
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

// TaxesReturn is to decode the json data
type TaxesReturn struct {
	Total int `json:"total"`
	Data  []struct {
		TaxRate          float64       `json:"taxRate"`
		Name             string        `json:"name"`
		Position         int           `json:"position"`
		Products         interface{}   `json:"products"`
		Rules            interface{}   `json:"rules"`
		ShippingMethods  interface{}   `json:"shippingMethods"`
		UniqueIdentifier string        `json:"_uniqueIdentifier"`
		VersionId        interface{}   `json:"versionId"`
		Translated       []interface{} `json:"translated"`
		CreatedAt        time.Time     `json:"createdAt"`
		UpdatedAt        interface{}   `json:"updatedAt"`
		Extensions       struct {
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

// Taxes are to get a list of all taxes
func Taxes(parameter map[string]string, r Request) (TaxesReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/tax",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return TaxesReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return TaxesReturn{}, err
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
		return TaxesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode TaxesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return TaxesReturn{}, err
	}

	// Return data
	return decode, err

}
