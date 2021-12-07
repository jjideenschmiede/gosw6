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
	"time"
)

// ProductManufacturersReturn is to decode the json data
type ProductManufacturersReturn struct {
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

// ProductManufacturerReturn is to decode the json return
type ProductManufacturerReturn struct {
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

// ProductManufacturers are to get a list of all product manufacturers
func ProductManufacturers(r Request) (ProductManufacturersReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product-manufacturer",
		Method: "GET",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductManufacturersReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductManufacturersReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductManufacturersReturn{}, err
	}

	// Return data
	return decode, err

}

// ProductManufacturer is to get a specific product manufacturer by id
func ProductManufacturer(id string, r Request) (ProductManufacturerReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product-manufacturer/" + id,
		Method: "GET",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductManufacturerReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductManufacturerReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductManufacturerReturn{}, err
	}

	// Return data
	return decode, err

}
