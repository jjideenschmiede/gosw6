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

// MediaFoldersReturn is to decode the json data
type MediaFoldersReturn struct {
	Total int `json:"total"`
	Data  []struct {
		Name                   string        `json:"name"`
		ParentId               string        `json:"parentId"`
		Parent                 interface{}   `json:"parent"`
		ChildCount             int           `json:"childCount"`
		Media                  interface{}   `json:"media"`
		ConfigurationId        string        `json:"configurationId"`
		Configuration          interface{}   `json:"configuration"`
		UseParentConfiguration bool          `json:"useParentConfiguration"`
		Children               interface{}   `json:"children"`
		DefaultFolder          interface{}   `json:"defaultFolder"`
		DefaultFolderId        string        `json:"defaultFolderId"`
		UniqueIdentifier       string        `json:"_uniqueIdentifier"`
		VersionId              interface{}   `json:"versionId"`
		Translated             []interface{} `json:"translated"`
		CreatedAt              time.Time     `json:"createdAt"`
		UpdatedAt              interface{}   `json:"updatedAt"`
		Extensions             struct {
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

// MediaFolderReturn is to decode the json data
type MediaFolderReturn struct {
	Data struct {
		Name                   string        `json:"name"`
		ParentId               interface{}   `json:"parentId"`
		Parent                 interface{}   `json:"parent"`
		ChildCount             int           `json:"childCount"`
		Media                  interface{}   `json:"media"`
		ConfigurationId        string        `json:"configurationId"`
		Configuration          interface{}   `json:"configuration"`
		UseParentConfiguration bool          `json:"useParentConfiguration"`
		Children               interface{}   `json:"children"`
		DefaultFolder          interface{}   `json:"defaultFolder"`
		DefaultFolderId        string        `json:"defaultFolderId"`
		UniqueIdentifier       string        `json:"_uniqueIdentifier"`
		VersionId              interface{}   `json:"versionId"`
		Translated             []interface{} `json:"translated"`
		CreatedAt              time.Time     `json:"createdAt"`
		UpdatedAt              interface{}   `json:"updatedAt"`
		Extensions             struct {
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

// MediaFolderBody is to structure the body data
type MediaFolderBody struct {
	ConfigurationId        string `json:"configurationId"`
	Name                   string `json:"name"`
	ParentId               string `json:"parentId"`
	UseParentConfiguration bool   `json:"useParentConfiguration"`
}

// CreateMediaFolderReturn is to decode the json data
type CreateMediaFolderReturn struct {
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

// UpdateMediaFolderReturn is to decode the json data
type UpdateMediaFolderReturn struct {
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

// DeleteMediaFolderReturn is to decode the json data
type DeleteMediaFolderReturn struct {
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

// MediaFolders are to get a list of all media folders
func MediaFolders(parameter map[string]string, r Request) (MediaFoldersReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/media-folder",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return MediaFoldersReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return MediaFoldersReturn{}, err
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
		return MediaFoldersReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode MediaFoldersReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return MediaFoldersReturn{}, err
	}

	// Return data
	return decode, err

}

// MediaFolder are to get a specific product by id
func MediaFolder(id string, r Request) (MediaFolderReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/media-folder/" + id,
		Method: "GET",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return MediaFolderReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode MediaFolderReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return MediaFolderReturn{}, err
	}

	// Return data
	return decode, err

}

// CreateMediaFolder is to create a media folder
func CreateMediaFolder(body MediaFolderBody, r Request) (CreateMediaFolderReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return CreateMediaFolderReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/media-folder",
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CreateMediaFolderReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CreateMediaFolderReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return CreateMediaFolderReturn{}, err
		}
	}

	// Get location in header & set to return struct
	decode.Location = response.Header.Get("location")

	// Return data
	return decode, err

}

// UpdateMediaFolder is to update a media folder
func UpdateMediaFolder(id string, body MediaFolderBody, r Request) (UpdateMediaFolderReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return UpdateMediaFolderReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/media-folder/" + id,
		Method: "PATCH",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return UpdateMediaFolderReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode UpdateMediaFolderReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return UpdateMediaFolderReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// DeleteMediaFolder is to delete a media folder
func DeleteMediaFolder(id string, r Request) (DeleteMediaFolderReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/media-folder/" + id,
		Method: "DELETE",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return DeleteMediaFolderReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode DeleteMediaFolderReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return DeleteMediaFolderReturn{}, err
		}
	}

	// Return data
	return decode, err

}
