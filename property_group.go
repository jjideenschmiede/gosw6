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

// PropertyGroupsReturn is to decode the json data
type PropertyGroupsReturn struct {
	Total int `json:"total"`
	Data  []struct {
		Name                       string      `json:"name"`
		DisplayType                string      `json:"displayType"`
		SortingType                string      `json:"sortingType"`
		Description                string      `json:"description"`
		Position                   int         `json:"position"`
		Filterable                 bool        `json:"filterable"`
		VisibleOnProductDetailPage bool        `json:"visibleOnProductDetailPage"`
		Options                    interface{} `json:"options"`
		Translations               interface{} `json:"translations"`
		UniqueIdentifier           string      `json:"_uniqueIdentifier"`
		VersionId                  interface{} `json:"versionId"`
		Translated                 struct {
			Name         string `json:"name"`
			Description  string `json:"description"`
			Position     int    `json:"position"`
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

// PropertyGroupReturn is to decode the json data
type PropertyGroupReturn struct {
	Data struct {
		Name                       string      `json:"name"`
		DisplayType                string      `json:"displayType"`
		SortingType                string      `json:"sortingType"`
		Description                string      `json:"description"`
		Position                   int         `json:"position"`
		Filterable                 bool        `json:"filterable"`
		VisibleOnProductDetailPage bool        `json:"visibleOnProductDetailPage"`
		Options                    interface{} `json:"options"`
		Translations               interface{} `json:"translations"`
		UniqueIdentifier           string      `json:"_uniqueIdentifier"`
		VersionId                  interface{} `json:"versionId"`
		Translated                 struct {
			Name         string `json:"name"`
			Description  string `json:"description"`
			Position     int    `json:"position"`
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

// PropertyGroups are to get a list of all property groups
func PropertyGroups(r Request) (PropertyGroupsReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/property-group",
		Method: "GET",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return PropertyGroupsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode PropertyGroupsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return PropertyGroupsReturn{}, err
	}

	// Return data
	return decode, err

}

// PropertyGroup is to get a specific product property group by id
func PropertyGroup(id string, r Request) (PropertyGroupReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/property-group/" + id,
		Method: "GET",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return PropertyGroupReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode PropertyGroupReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return PropertyGroupReturn{}, err
	}

	// Return data
	return decode, err

}
