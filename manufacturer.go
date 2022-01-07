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

// ManufacturersReturn is to decode the json data
type ManufacturersReturn struct {
	Total int `json:"total"`
	Data  []struct {
		MediaId          interface{} `json:"mediaId"`
		Name             string      `json:"name"`
		Link             interface{} `json:"link"`
		Description      interface{} `json:"description"`
		Media            interface{} `json:"media"`
		Translations     interface{} `json:"translations"`
		Products         interface{} `json:"products"`
		UniqueIdentifier string      `json:"_uniqueIdentifier"`
		VersionId        string      `json:"versionId"`
		Translated       struct {
			Name         string      `json:"name"`
			Description  interface{} `json:"description"`
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

// ManufacturerReturn is to decode the json return
type ManufacturerReturn struct {
	Data struct {
		MediaId          interface{} `json:"mediaId"`
		Name             string      `json:"name"`
		Link             interface{} `json:"link"`
		Description      interface{} `json:"description"`
		Media            interface{} `json:"media"`
		Translations     interface{} `json:"translations"`
		Products         interface{} `json:"products"`
		UniqueIdentifier string      `json:"_uniqueIdentifier"`
		VersionId        string      `json:"versionId"`
		Translated       struct {
			Name         string      `json:"name"`
			Description  interface{} `json:"description"`
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

// ManufacturerBody is to structure the body data
type ManufacturerBody struct {
	MediaId     string `json:"mediaId,omitempty"`
	Name        string `json:"name"`
	Link        string `json:"link,omitempty"`
	Description string `json:"description,omitempty"`
}

// CreateManufacturerReturn is to decode the json data
type CreateManufacturerReturn struct {
	Location string `json:"location"`
	Errors   []struct {
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

// UpdateManufacturerReturn is to decode the json data
type UpdateManufacturerReturn struct {
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

// DeleteManufacturerReturn is to decode the json data
type DeleteManufacturerReturn struct {
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

// Manufacturers are to get a list of all product manufacturers
func Manufacturers(parameter map[string]string, r Request) (ManufacturersReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product-manufacturer",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return ManufacturersReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return ManufacturersReturn{}, err
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
		return ManufacturersReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ManufacturersReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ManufacturersReturn{}, err
	}

	// Return data
	return decode, err

}

// Manufacturer is to get a specific product manufacturer by id
func Manufacturer(id string, r Request) (ManufacturerReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product-manufacturer/" + id,
		Method: "GET",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ManufacturerReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ManufacturerReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ManufacturerReturn{}, err
	}

	// Return data
	return decode, err

}

// CreateManufacturer is to create a new product manufacturer
func CreateManufacturer(body ManufacturerBody, r Request) (CreateManufacturerReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return CreateManufacturerReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/product-manufacturer",
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CreateManufacturerReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CreateManufacturerReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return CreateManufacturerReturn{}, err
		}
	}

	// Get location in header & set to return struct
	decode.Location = response.Header.Get("location")

	// Return data
	return decode, err

}

// UpdateManufacturer is to update a new product manufacturer
func UpdateManufacturer(id string, body ManufacturerBody, r Request) (UpdateManufacturerReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return UpdateManufacturerReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/product-manufacturer/" + id,
		Method: "PATCH",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return UpdateManufacturerReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode UpdateManufacturerReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return UpdateManufacturerReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// DeleteManufacturer is to delete a new product manufacturer
func DeleteManufacturer(id string, r Request) (DeleteManufacturerReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product-manufacturer/" + id,
		Method: "DELETE",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return DeleteManufacturerReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode DeleteManufacturerReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return DeleteManufacturerReturn{}, err
		}
	}

	// Return data
	return decode, err

}
