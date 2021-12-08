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

// PropertyGroupOptionsReturn is to decode the json return
type PropertyGroupOptionsReturn struct {
	Total int `json:"total"`
	Data  []struct {
		GroupId                     string      `json:"groupId"`
		Name                        string      `json:"name"`
		Position                    int         `json:"position"`
		ColorHexCode                interface{} `json:"colorHexCode"`
		MediaId                     interface{} `json:"mediaId"`
		Group                       interface{} `json:"group"`
		Translations                interface{} `json:"translations"`
		ProductConfiguratorSettings interface{} `json:"productConfiguratorSettings"`
		ProductProperties           interface{} `json:"productProperties"`
		ProductOptions              interface{} `json:"productOptions"`
		Media                       interface{} `json:"media"`
		UniqueIdentifier            string      `json:"_uniqueIdentifier"`
		VersionId                   interface{} `json:"versionId"`
		Translated                  struct {
			Name         string `json:"name"`
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

// PropertyGroupOptionReturn is to decode json data
type PropertyGroupOptionReturn struct {
	Data struct {
		GroupId                     string      `json:"groupId"`
		Name                        string      `json:"name"`
		Position                    int         `json:"position"`
		ColorHexCode                interface{} `json:"colorHexCode"`
		MediaId                     interface{} `json:"mediaId"`
		Group                       interface{} `json:"group"`
		Translations                interface{} `json:"translations"`
		ProductConfiguratorSettings interface{} `json:"productConfiguratorSettings"`
		ProductProperties           interface{} `json:"productProperties"`
		ProductOptions              interface{} `json:"productOptions"`
		Media                       interface{} `json:"media"`
		UniqueIdentifier            string      `json:"_uniqueIdentifier"`
		VersionId                   interface{} `json:"versionId"`
		Translated                  struct {
			Name         string `json:"name"`
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

// CreatePropertyGroupOptionBody is to structure the body data
type CreatePropertyGroupOptionBody struct {
	Name         string      `json:"name"`
	Position     int         `json:"position"`
	ColorHexCode interface{} `json:"colorHexCode"`
	MediaId      interface{} `json:"mediaId"`
}

// CreatePropertyGroupOptionReturn is to decode the json data
type CreatePropertyGroupOptionReturn struct {
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

// UpdatePropertyGroupOptionBody is to structure the body data
type UpdatePropertyGroupOptionBody struct {
	Name         string      `json:"name"`
	Position     int         `json:"position"`
	ColorHexCode interface{} `json:"colorHexCode"`
	MediaId      interface{} `json:"mediaId"`
}

// UpdatePropertyGroupOptionReturn is to decode the json data
type UpdatePropertyGroupOptionReturn struct {
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

// PropertyGroupOptions are to get a list of all property group options
func PropertyGroupOptions(id string, parameter map[string]string, r Request) (PropertyGroupOptionsReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/property-group/" + id + "/options",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return PropertyGroupOptionsReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return PropertyGroupOptionsReturn{}, err
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
		return PropertyGroupOptionsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode PropertyGroupOptionsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return PropertyGroupOptionsReturn{}, err
	}

	// Return data
	return decode, err

}

// PropertyGroupOption is to get a property group option by id
func PropertyGroupOption(id string, r Request) (PropertyGroupOptionReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/property-group-option/" + id,
		Method: "GET",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return PropertyGroupOptionReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode PropertyGroupOptionReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return PropertyGroupOptionReturn{}, err
	}

	// Return data
	return decode, err

}

// CreatePropertyGroupOption is to create a property group option
func CreatePropertyGroupOption(id string, body CreatePropertyGroupOptionBody, r Request) (CreatePropertyGroupOptionReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return CreatePropertyGroupOptionReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/property-group/" + id + "/options",
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CreatePropertyGroupOptionReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CreatePropertyGroupOptionReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return CreatePropertyGroupOptionReturn{}, err
		}
	}

	// Get location in header & set to return struct
	decode.Location = response.Header.Get("location")

	// Return data
	return decode, err

}

// UpdatePropertyGroupOption is to update a property group option
func UpdatePropertyGroupOption(id string, body UpdatePropertyGroupOptionBody, r Request) (UpdatePropertyGroupOptionReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return UpdatePropertyGroupOptionReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/property-group-option/" + id,
		Method: "PATCH",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return UpdatePropertyGroupOptionReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode UpdatePropertyGroupOptionReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return UpdatePropertyGroupOptionReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// DeletePropertyGroupOption is to delete a property group option
func DeletePropertyGroupOption(id string, r Request) (UpdatePropertyGroupOptionReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/property-group-option/" + id,
		Method: "DELETE",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return UpdatePropertyGroupOptionReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode UpdatePropertyGroupOptionReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return UpdatePropertyGroupOptionReturn{}, err
		}
	}

	// Return data
	return decode, err

}
