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
	"time"
)

// MediaReturn is to decode json data
type MediaReturn struct {
	Total int `json:"total"`
	Data  []struct {
		UserId        string      `json:"userId"`
		MimeType      string      `json:"mimeType"`
		FileExtension string      `json:"fileExtension"`
		FileSize      int         `json:"fileSize"`
		Title         interface{} `json:"title"`
		MetaData      struct {
			Width  int    `json:"width"`
			Height int    `json:"height"`
			Type   int    `json:"type"`
			Hash   string `json:"hash"`
		} `json:"metaData"`
		MediaType struct {
			Name       string        `json:"name"`
			Flags      []string      `json:"flags"`
			Extensions []interface{} `json:"extensions"`
			ApiAlias   string        `json:"apiAlias"`
		} `json:"mediaType"`
		UploadedAt                  time.Time     `json:"uploadedAt"`
		Alt                         interface{}   `json:"alt"`
		Url                         string        `json:"url"`
		FileName                    string        `json:"fileName"`
		User                        interface{}   `json:"user"`
		Translations                interface{}   `json:"translations"`
		Categories                  interface{}   `json:"categories"`
		ProductManufacturers        interface{}   `json:"productManufacturers"`
		ProductMedia                interface{}   `json:"productMedia"`
		AvatarUser                  interface{}   `json:"avatarUser"`
		Thumbnails                  []interface{} `json:"thumbnails"`
		MediaFolderId               string        `json:"mediaFolderId"`
		MediaFolder                 interface{}   `json:"mediaFolder"`
		HasFile                     bool          `json:"hasFile"`
		Private                     bool          `json:"private"`
		PropertyGroupOptions        interface{}   `json:"propertyGroupOptions"`
		MailTemplateMedia           interface{}   `json:"mailTemplateMedia"`
		Tags                        interface{}   `json:"tags"`
		DocumentBaseConfigs         interface{}   `json:"documentBaseConfigs"`
		ShippingMethods             interface{}   `json:"shippingMethods"`
		PaymentMethods              interface{}   `json:"paymentMethods"`
		ProductConfiguratorSettings interface{}   `json:"productConfiguratorSettings"`
		OrderLineItems              interface{}   `json:"orderLineItems"`
		CmsBlocks                   interface{}   `json:"cmsBlocks"`
		CmsSections                 interface{}   `json:"cmsSections"`
		CmsPages                    interface{}   `json:"cmsPages"`
		Documents                   interface{}   `json:"documents"`
		AppPaymentMethods           interface{}   `json:"appPaymentMethods"`
		UniqueIdentifier            string        `json:"_uniqueIdentifier"`
		VersionId                   interface{}   `json:"versionId"`
		Translated                  struct {
			Alt          interface{} `json:"alt"`
			Title        interface{} `json:"title"`
			CustomFields struct {
			} `json:"customFields"`
		} `json:"translated"`
		CreatedAt  time.Time `json:"createdAt"`
		UpdatedAt  time.Time `json:"updatedAt"`
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

// CreateMediaBody is to structure the body data
type CreateMediaBody struct {
	MediaFolderId string `json:"mediaFolderId"`
}

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

// DeleteMediaReturn is to decode the json data
type DeleteMediaReturn struct {
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

// Media are to get a list of all media
func Media(parameter map[string]string, r Request) (MediaReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/media",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return MediaReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return MediaReturn{}, err
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
		return MediaReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode MediaReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return MediaReturn{}, err
	}

	// Return data
	return decode, err

}

// CreateMedia is to create a new media
func CreateMedia(body CreateMediaBody, r Request) (CreateMediaReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return CreateMediaReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/media",
		Method: "POST",
		Body:   convert,
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

// DeleteMedia is to delete a media
func DeleteMedia(id string, r Request) (DeleteMediaReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/media/" + id,
		Method: "DELETE",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return DeleteMediaReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode DeleteMediaReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return DeleteMediaReturn{}, err
		}
	}

	// Return data
	return decode, err

}
