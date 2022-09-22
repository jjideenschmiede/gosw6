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

// TagsReturn is to decode the json data
type TagsReturn struct {
	Total int `json:"total"`
	Data  []struct {
		Name                 string        `json:"name"`
		Products             interface{}   `json:"products"`
		Media                interface{}   `json:"media"`
		Categories           interface{}   `json:"categories"`
		Customers            interface{}   `json:"customers"`
		Orders               interface{}   `json:"orders"`
		ShippingMethods      interface{}   `json:"shippingMethods"`
		NewsletterRecipients interface{}   `json:"newsletterRecipients"`
		LandingPages         interface{}   `json:"landingPages"`
		UniqueIdentifier     string        `json:"_uniqueIdentifier"`
		VersionId            interface{}   `json:"versionId"`
		Translated           []interface{} `json:"translated"`
		CreatedAt            time.Time     `json:"createdAt"`
		UpdatedAt            interface{}   `json:"updatedAt"`
		Extensions           struct {
			ForeignKeys struct {
				ApiAlias   interface{}   `json:"apiAlias"`
				Extensions []interface{} `json:"extensions"`
			} `json:"foreignKeys"`
		} `json:"extensions"`
		Id       string `json:"id"`
		ApiAlias string `json:"apiAlias"`
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

// Tags are to get a list of all tags
func Tags(parameter map[string]string, r Request) (TagsReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/tag",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return TagsReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return TagsReturn{}, err
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
		return TagsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode TagsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return TagsReturn{}, err
	}

	// Return data
	return decode, err

}

// DeleteTag is to delete an tag
func DeleteTag(id string, r Request) (TagsReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/tag/" + id,
		Method: "DELETE",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return TagsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode TagsReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return TagsReturn{}, err
		}
	}

	// Return data
	return decode, err

}
