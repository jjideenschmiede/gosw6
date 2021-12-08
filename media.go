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
	"io"
	"net/url"
	"os"
)

// CreateMediaReturn is to decode the json return
type CreateMediaReturn struct {
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

// UploadUrlMediaBody is to structure the body data
type UploadUrlMediaBody struct {
	Url string `json:"url"`
}

// UploadUrlMediaReturn is to decode json data
type UploadUrlMediaReturn struct {
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

// UploadLocalMediaReturn is to decode json data
type UploadLocalMediaReturn struct {
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

// CreateMedia is to create a new media
func CreateMedia(r Request) (CreateMediaReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/media",
		Method: "POST",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CreateMediaReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CreateMediaReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return CreateMediaReturn{}, err
		}
	}

	// Get location in header & set to return struct
	decode.Location = response.Header.Get("location")

	// Return data
	return decode, err

}

// UploadUrlMedia is to upload a file to media via url
func UploadUrlMedia(name, extension, id string, body UploadUrlMediaBody, r Request) (UploadUrlMediaReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return UploadUrlMediaReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/_action/media/" + id + "/upload",
		Method: "POST",
		Body:   convert,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return UploadUrlMediaReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return UploadUrlMediaReturn{}, err
	}

	newUrl.Add("extension", extension)
	newUrl.Add("fileName", name)

	// Set new url
	parse.RawQuery = newUrl.Encode()
	c.Path = fmt.Sprintf("%s", parse)

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return UploadUrlMediaReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode UploadUrlMediaReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return UploadUrlMediaReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// UploadLocalMedia is to upload a file to media via local
func UploadLocalMedia(file *os.File, name, extension, id string, r Request) (UploadLocalMediaReturn, error) {

	// Read all data from file
	read, err := io.ReadAll(file)
	if err != nil {
		return UploadLocalMediaReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:        "/api/_action/media/" + id + "/upload",
		Method:      "POST",
		ContentType: "image/" + extension,
		Body:        read,
		UploadMedia: true,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return UploadLocalMediaReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return UploadLocalMediaReturn{}, err
	}

	newUrl.Add("extension", extension)
	newUrl.Add("fileName", name)

	// Set new url
	parse.RawQuery = newUrl.Encode()
	c.Path = fmt.Sprintf("%s", parse)

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return UploadLocalMediaReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode UploadLocalMediaReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return UploadLocalMediaReturn{}, err
		}
	}

	// Return data
	return decode, err

}
