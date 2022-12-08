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

// RulesReturn is to decode the json data
type RulesReturn struct {
	Total int `json:"total"`
	Data  []struct {
		Name        string  `json:"name"`
		Description *string `json:"description"`
		Priority    int     `json:"priority"`
		ModuleTypes *struct {
			Types []string `json:"types"`
		} `json:"moduleTypes"`
		ProductPrices                   interface{}   `json:"productPrices"`
		ShippingMethods                 interface{}   `json:"shippingMethods"`
		PaymentMethods                  interface{}   `json:"paymentMethods"`
		EventActions                    interface{}   `json:"eventActions"`
		Conditions                      interface{}   `json:"conditions"`
		Invalid                         bool          `json:"invalid"`
		ShippingMethodPrices            interface{}   `json:"shippingMethodPrices"`
		PromotionDiscounts              interface{}   `json:"promotionDiscounts"`
		PromotionSetGroups              interface{}   `json:"promotionSetGroups"`
		ShippingMethodPriceCalculations interface{}   `json:"shippingMethodPriceCalculations"`
		PersonaPromotions               interface{}   `json:"personaPromotions"`
		FlowSequences                   interface{}   `json:"flowSequences"`
		Tags                            interface{}   `json:"tags"`
		OrderPromotions                 interface{}   `json:"orderPromotions"`
		CartPromotions                  interface{}   `json:"cartPromotions"`
		UniqueIdentifier                string        `json:"_uniqueIdentifier"`
		VersionId                       interface{}   `json:"versionId"`
		Translated                      []interface{} `json:"translated"`
		CreatedAt                       time.Time     `json:"createdAt"`
		UpdatedAt                       interface{}   `json:"updatedAt"`
		Extensions                      struct {
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

// Rules are to get a list of all rules
func Rules(parameter map[string]string, r Request) (RulesReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/rule",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return RulesReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return RulesReturn{}, err
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
		return RulesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode RulesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return RulesReturn{}, err
	}

	// Return data
	return decode, err

}
