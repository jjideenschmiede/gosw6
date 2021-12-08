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

// CreatePropertyGroupBody is to structure the body data
type CreatePropertyGroupBody struct {
	Name                       string `json:"name"`
	DisplayType                string `json:"displayType"`
	SortingType                string `json:"sortingType"`
	Description                string `json:"description"`
	Position                   int    `json:"position"`
	Filterable                 bool   `json:"filterable"`
	VisibleOnProductDetailPage bool   `json:"visibleOnProductDetailPage"`
}

// CreatePropertyGroupBodyReturn is to decode the json data
type CreatePropertyGroupBodyReturn struct {
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

// UpdatePropertyGroupBody is to structure the body data
type UpdatePropertyGroupBody struct {
	Name                       string `json:"name"`
	DisplayType                string `json:"displayType"`
	SortingType                string `json:"sortingType"`
	Description                string `json:"description"`
	Position                   int    `json:"position"`
	Filterable                 bool   `json:"filterable"`
	VisibleOnProductDetailPage bool   `json:"visibleOnProductDetailPage"`
}

// UpdatePropertyGroupBodyReturn is to decode the json data
type UpdatePropertyGroupBodyReturn struct {
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

// DeletePropertyGroupBodyReturn is to decode the json data
type DeletePropertyGroupBodyReturn struct {
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
func PropertyGroups(parameter map[string]string, r Request) (PropertyGroupsReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/property-group",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return PropertyGroupsReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return PropertyGroupsReturn{}, err
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

// CreatePropertyGroup is to create a product property group
func CreatePropertyGroup(body CreatePropertyGroupBody, r Request) (CreatePropertyGroupBodyReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return CreatePropertyGroupBodyReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/property-group",
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CreatePropertyGroupBodyReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CreatePropertyGroupBodyReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return CreatePropertyGroupBodyReturn{}, err
		}
	}

	// Get location in header & set to return struct
	decode.Location = response.Header.Get("location")

	// Return data
	return decode, err

}

// UpdatePropertyGroup is to update a product property group
func UpdatePropertyGroup(id string, body UpdatePropertyGroupBody, r Request) (UpdatePropertyGroupBodyReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return UpdatePropertyGroupBodyReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/property-group/" + id,
		Method: "PATCH",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return UpdatePropertyGroupBodyReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode UpdatePropertyGroupBodyReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return UpdatePropertyGroupBodyReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// DeletePropertyGroup is to delete a product property group
func DeletePropertyGroup(id string, r Request) (DeletePropertyGroupBodyReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/property-group/" + id,
		Method: "DELETE",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return DeletePropertyGroupBodyReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode DeletePropertyGroupBodyReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return DeletePropertyGroupBodyReturn{}, err
		}
	}

	// Return data
	return decode, err

}
