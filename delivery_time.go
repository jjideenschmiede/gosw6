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

// DeliveryTimesReturn is to decode the json data
type DeliveryTimesReturn struct {
	Total int `json:"total"`
	Data  []struct {
		Name             string      `json:"name"`
		Min              int         `json:"min"`
		Max              int         `json:"max"`
		Unit             string      `json:"unit"`
		ShippingMethods  interface{} `json:"shippingMethods"`
		Translations     interface{} `json:"translations"`
		Products         interface{} `json:"products"`
		UniqueIdentifier string      `json:"_uniqueIdentifier"`
		VersionId        interface{} `json:"versionId"`
		Translated       struct {
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

// CreateDeliveryTimesBody is to structure the body data
type CreateDeliveryTimesBody struct {
	Name string `json:"name"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
	Unit string `json:"unit"`
}

// CreateDeliveryTimesReturn is to decode the json data
type CreateDeliveryTimesReturn struct {
	Errors []struct {
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

// DeliveryTimes is to get a list of all delivery times
func DeliveryTimes(parameter map[string]string, r Request) (DeliveryTimesReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/delivery-time",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return DeliveryTimesReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return DeliveryTimesReturn{}, err
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
		return DeliveryTimesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode DeliveryTimesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return DeliveryTimesReturn{}, err
	}

	// Return data
	return decode, err

}

// CreateDeliveryTimes is to create a delivery time
func CreateDeliveryTimes(body CreateDeliveryTimesBody, r Request) (CreateDeliveryTimesReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return CreateDeliveryTimesReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/delivery-time",
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CreateDeliveryTimesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CreateDeliveryTimesReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return CreateDeliveryTimesReturn{}, err
		}
	}

	// Return data
	return decode, err

}
