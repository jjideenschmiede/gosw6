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
	Aggregations []interface{} `json:"aggregations"`
}

// Categories is to get a list of all categories
func Categories(r Request) (ProductsReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/category",
		Method: "GET",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductsReturn{}, err
	}

	// Return data
	return decode, err

}
