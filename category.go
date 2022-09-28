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

// CategoriesReturn is to decode the json data
type CategoriesReturn struct {
	Total int `json:"total"`
	Data  []struct {
		ParentId                *string     `json:"parentId"`
		AutoIncrement           int         `json:"autoIncrement"`
		MediaId                 interface{} `json:"mediaId"`
		Name                    string      `json:"name"`
		Breadcrumb              []string    `json:"breadcrumb"`
		Path                    *string     `json:"path"`
		Level                   int         `json:"level"`
		Active                  bool        `json:"active"`
		ChildCount              int         `json:"childCount"`
		VisibleChildCount       int         `json:"visibleChildCount"`
		DisplayNestedProducts   bool        `json:"displayNestedProducts"`
		Parent                  interface{} `json:"parent"`
		Children                interface{} `json:"children"`
		Translations            interface{} `json:"translations"`
		Media                   interface{} `json:"media"`
		Products                interface{} `json:"products"`
		NestedProducts          interface{} `json:"nestedProducts"`
		AfterCategoryId         interface{} `json:"afterCategoryId"`
		Tags                    interface{} `json:"tags"`
		CmsPageId               string      `json:"cmsPageId"`
		CmsPage                 interface{} `json:"cmsPage"`
		ProductStreamId         interface{} `json:"productStreamId"`
		ProductStream           interface{} `json:"productStream"`
		SlotConfig              interface{} `json:"slotConfig"`
		NavigationSalesChannels interface{} `json:"navigationSalesChannels"`
		FooterSalesChannels     interface{} `json:"footerSalesChannels"`
		ServiceSalesChannels    interface{} `json:"serviceSalesChannels"`
		LinkType                interface{} `json:"linkType"`
		LinkNewTab              interface{} `json:"linkNewTab"`
		InternalLink            interface{} `json:"internalLink"`
		ExternalLink            interface{} `json:"externalLink"`
		Visible                 bool        `json:"visible"`
		Type                    string      `json:"type"`
		ProductAssignmentType   string      `json:"productAssignmentType"`
		Description             interface{} `json:"description"`
		MetaTitle               *string     `json:"metaTitle"`
		MetaDescription         *string     `json:"metaDescription"`
		Keywords                *string     `json:"keywords"`
		MainCategories          interface{} `json:"mainCategories"`
		SeoUrls                 interface{} `json:"seoUrls"`
		UniqueIdentifier        string      `json:"_uniqueIdentifier"`
		VersionId               string      `json:"versionId"`
		Translated              struct {
			Breadcrumb   []string `json:"breadcrumb"`
			Name         string   `json:"name"`
			CustomFields struct {
			} `json:"customFields"`
			SlotConfig      interface{} `json:"slotConfig"`
			LinkType        interface{} `json:"linkType"`
			InternalLink    interface{} `json:"internalLink"`
			ExternalLink    interface{} `json:"externalLink"`
			LinkNewTab      interface{} `json:"linkNewTab"`
			Description     interface{} `json:"description"`
			MetaTitle       *string     `json:"metaTitle"`
			MetaDescription *string     `json:"metaDescription"`
			Keywords        *string     `json:"keywords"`
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

// CategoryReturn is to decode the json return
type CategoryReturn struct {
	Data struct {
		ParentId                interface{} `json:"parentId"`
		AutoIncrement           int         `json:"autoIncrement"`
		MediaId                 interface{} `json:"mediaId"`
		Name                    string      `json:"name"`
		Breadcrumb              []string    `json:"breadcrumb"`
		Path                    interface{} `json:"path"`
		Level                   int         `json:"level"`
		Active                  bool        `json:"active"`
		ChildCount              int         `json:"childCount"`
		VisibleChildCount       int         `json:"visibleChildCount"`
		DisplayNestedProducts   bool        `json:"displayNestedProducts"`
		Parent                  interface{} `json:"parent"`
		Children                interface{} `json:"children"`
		Translations            interface{} `json:"translations"`
		Media                   interface{} `json:"media"`
		Products                interface{} `json:"products"`
		NestedProducts          interface{} `json:"nestedProducts"`
		AfterCategoryId         interface{} `json:"afterCategoryId"`
		Tags                    interface{} `json:"tags"`
		CmsPageId               string      `json:"cmsPageId"`
		CmsPage                 interface{} `json:"cmsPage"`
		ProductStreamId         interface{} `json:"productStreamId"`
		ProductStream           interface{} `json:"productStream"`
		SlotConfig              interface{} `json:"slotConfig"`
		NavigationSalesChannels interface{} `json:"navigationSalesChannels"`
		FooterSalesChannels     interface{} `json:"footerSalesChannels"`
		ServiceSalesChannels    interface{} `json:"serviceSalesChannels"`
		LinkType                interface{} `json:"linkType"`
		LinkNewTab              interface{} `json:"linkNewTab"`
		InternalLink            interface{} `json:"internalLink"`
		ExternalLink            interface{} `json:"externalLink"`
		Visible                 bool        `json:"visible"`
		Type                    string      `json:"type"`
		ProductAssignmentType   string      `json:"productAssignmentType"`
		Description             interface{} `json:"description"`
		MetaTitle               interface{} `json:"metaTitle"`
		MetaDescription         interface{} `json:"metaDescription"`
		Keywords                interface{} `json:"keywords"`
		MainCategories          interface{} `json:"mainCategories"`
		SeoUrls                 interface{} `json:"seoUrls"`
		UniqueIdentifier        string      `json:"_uniqueIdentifier"`
		VersionId               string      `json:"versionId"`
		Translated              struct {
			Breadcrumb   []string `json:"breadcrumb"`
			Name         string   `json:"name"`
			CustomFields struct {
			} `json:"customFields"`
			SlotConfig      interface{} `json:"slotConfig"`
			LinkType        interface{} `json:"linkType"`
			InternalLink    interface{} `json:"internalLink"`
			ExternalLink    interface{} `json:"externalLink"`
			LinkNewTab      interface{} `json:"linkNewTab"`
			Description     interface{} `json:"description"`
			MetaTitle       interface{} `json:"metaTitle"`
			MetaDescription interface{} `json:"metaDescription"`
			Keywords        interface{} `json:"keywords"`
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

// CategoryBody is to structure the body data
type CategoryBody struct {
	ParentId              string      `json:"parentId,omitempty"`
	MediaId               interface{} `json:"mediaId,omitempty"`
	Name                  string      `json:"name"`
	Active                bool        `json:"active"`
	VisibleChildCount     int         `json:"visibleChildCount"`
	DisplayNestedProducts bool        `json:"displayNestedProducts"`
	CmsPageId             string      `json:"cmsPageId,omitempty"`
	LinkType              interface{} `json:"linkType,omitempty"`
	LinkNewTab            interface{} `json:"linkNewTab,omitempty"`
	InternalLink          interface{} `json:"internalLink,omitempty"`
	ExternalLink          interface{} `json:"externalLink,omitempty"`
	Visible               bool        `json:"visible"`
	Type                  string      `json:"type"`
	ProductAssignmentType string      `json:"productAssignmentType"`
	Description           string      `json:"description,omitempty"`
	MetaTitle             string      `json:"metaTitle,omitempty"`
	MetaDescription       string      `json:"metaDescription,omitempty"`
	Keywords              string      `json:"keywords,omitempty"`
}

// CreateCategoryReturn is to decode the json data
type CreateCategoryReturn struct {
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

// UpdateCategoryReturn is to decode the json data
type UpdateCategoryReturn struct {
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

// Categories are to get a list of all categories
func Categories(parameter map[string]string, r Request) (CategoriesReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/category",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return CategoriesReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return CategoriesReturn{}, err
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
		return CategoriesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CategoriesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return CategoriesReturn{}, err
	}

	// Return data
	return decode, err

}

// Category is to get a specific category by id
func Category(id string, r Request) (CategoryReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/category/" + id,
		Method: "GET",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CategoryReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CategoryReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return CategoryReturn{}, err
	}

	// Return data
	return decode, err

}

// CreateCategory is to create a new category
func CreateCategory(body CategoryBody, r Request) (CreateCategoryReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return CreateCategoryReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/category",
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CreateCategoryReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CreateCategoryReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return CreateCategoryReturn{}, err
		}
	}

	// Get location in header & set to return struct
	decode.Location = response.Header.Get("location")

	// Return data
	return decode, err

}

// UpdateCategory is to update an update category
func UpdateCategory(id string, body CategoryBody, r Request) (UpdateCategoryReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return UpdateCategoryReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/category/" + id,
		Method: "PATCH",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return UpdateCategoryReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode UpdateCategoryReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return UpdateCategoryReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// DeleteCategory is to delete an update category
func DeleteCategory(id string, r Request) (UpdateCategoryReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/category/" + id,
		Method: "DELETE",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return UpdateCategoryReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode UpdateCategoryReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return UpdateCategoryReturn{}, err
		}
	}

	// Return data
	return decode, err

}
