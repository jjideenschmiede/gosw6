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

// CmsPagesReturn is to decode the json data
type CmsPagesReturn struct {
	Total int `json:"total"`
	Data  []struct {
		Name              string      `json:"name"`
		Type              string      `json:"type"`
		Entity            interface{} `json:"entity"`
		Sections          interface{} `json:"sections"`
		Translations      interface{} `json:"translations"`
		Categories        interface{} `json:"categories"`
		Products          interface{} `json:"products"`
		Config            interface{} `json:"config"`
		PreviewMediaId    interface{} `json:"previewMediaId"`
		PreviewMedia      interface{} `json:"previewMedia"`
		Locked            bool        `json:"locked"`
		LandingPages      interface{} `json:"landingPages"`
		HomeSalesChannels interface{} `json:"homeSalesChannels"`
		UniqueIdentifier  string      `json:"_uniqueIdentifier"`
		VersionId         string      `json:"versionId"`
		Translated        struct {
			Name         string `json:"name"`
			CustomFields struct {
			} `json:"customFields"`
		} `json:"translated"`
		CreatedAt  time.Time  `json:"createdAt"`
		UpdatedAt  *time.Time `json:"updatedAt"`
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

// CmsPages are to get a list of all cms pages
func CmsPages(parameter map[string]string, r Request) (CmsPagesReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/cms-page",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return CmsPagesReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return CmsPagesReturn{}, err
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
		return CmsPagesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CmsPagesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return CmsPagesReturn{}, err
	}

	// Return data
	return decode, err

}
