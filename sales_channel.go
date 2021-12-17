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

// SalesChannelsReturn is to decode the json return
type SalesChannelsReturn struct {
	Total int `json:"total"`
	Data  []struct {
		TypeId                          string      `json:"typeId"`
		LanguageId                      string      `json:"languageId"`
		CurrencyId                      string      `json:"currencyId"`
		PaymentMethodId                 string      `json:"paymentMethodId"`
		ShippingMethodId                string      `json:"shippingMethodId"`
		CountryId                       string      `json:"countryId"`
		NavigationCategoryId            string      `json:"navigationCategoryId"`
		NavigationCategoryDepth         int         `json:"navigationCategoryDepth"`
		HomeSlotConfig                  interface{} `json:"homeSlotConfig"`
		HomeCmsPageId                   interface{} `json:"homeCmsPageId"`
		HomeCmsPage                     interface{} `json:"homeCmsPage"`
		HomeEnabled                     bool        `json:"homeEnabled"`
		HomeName                        interface{} `json:"homeName"`
		HomeMetaTitle                   interface{} `json:"homeMetaTitle"`
		HomeMetaDescription             interface{} `json:"homeMetaDescription"`
		HomeKeywords                    interface{} `json:"homeKeywords"`
		FooterCategoryId                interface{} `json:"footerCategoryId"`
		ServiceCategoryId               interface{} `json:"serviceCategoryId"`
		Name                            string      `json:"name"`
		ShortName                       interface{} `json:"shortName"`
		AccessKey                       string      `json:"accessKey"`
		Currencies                      interface{} `json:"currencies"`
		Languages                       interface{} `json:"languages"`
		Configuration                   interface{} `json:"configuration"`
		Active                          bool        `json:"active"`
		Maintenance                     bool        `json:"maintenance"`
		MaintenanceIpWhitelist          interface{} `json:"maintenanceIpWhitelist"`
		TaxCalculationType              string      `json:"taxCalculationType"`
		Type                            interface{} `json:"type"`
		Currency                        interface{} `json:"currency"`
		Language                        interface{} `json:"language"`
		PaymentMethod                   interface{} `json:"paymentMethod"`
		ShippingMethod                  interface{} `json:"shippingMethod"`
		Country                         interface{} `json:"country"`
		Orders                          interface{} `json:"orders"`
		Customers                       interface{} `json:"customers"`
		Countries                       interface{} `json:"countries"`
		PaymentMethods                  interface{} `json:"paymentMethods"`
		ShippingMethods                 interface{} `json:"shippingMethods"`
		Translations                    interface{} `json:"translations"`
		Domains                         interface{} `json:"domains"`
		SystemConfigs                   interface{} `json:"systemConfigs"`
		NavigationCategory              interface{} `json:"navigationCategory"`
		FooterCategory                  interface{} `json:"footerCategory"`
		ServiceCategory                 interface{} `json:"serviceCategory"`
		ProductVisibilities             interface{} `json:"productVisibilities"`
		MailHeaderFooterId              interface{} `json:"mailHeaderFooterId"`
		NumberRangeSalesChannels        interface{} `json:"numberRangeSalesChannels"`
		MailHeaderFooter                interface{} `json:"mailHeaderFooter"`
		CustomerGroupId                 string      `json:"customerGroupId"`
		CustomerGroup                   interface{} `json:"customerGroup"`
		NewsletterRecipients            interface{} `json:"newsletterRecipients"`
		PromotionSalesChannels          interface{} `json:"promotionSalesChannels"`
		DocumentBaseConfigSalesChannels interface{} `json:"documentBaseConfigSalesChannels"`
		ProductReviews                  interface{} `json:"productReviews"`
		SeoUrls                         interface{} `json:"seoUrls"`
		SeoUrlTemplates                 interface{} `json:"seoUrlTemplates"`
		MainCategories                  interface{} `json:"mainCategories"`
		PaymentMethodIds                []string    `json:"paymentMethodIds"`
		ProductExports                  interface{} `json:"productExports"`
		HreflangActive                  bool        `json:"hreflangActive"`
		HreflangDefaultDomainId         interface{} `json:"hreflangDefaultDomainId"`
		HreflangDefaultDomain           interface{} `json:"hreflangDefaultDomain"`
		AnalyticsId                     interface{} `json:"analyticsId"`
		Analytics                       interface{} `json:"analytics"`
		CustomerGroupsRegistrations     interface{} `json:"customerGroupsRegistrations"`
		EventActions                    interface{} `json:"eventActions"`
		BoundCustomers                  interface{} `json:"boundCustomers"`
		Wishlists                       interface{} `json:"wishlists"`
		LandingPages                    interface{} `json:"landingPages"`
		UniqueIdentifier                string      `json:"_uniqueIdentifier"`
		VersionId                       interface{} `json:"versionId"`
		Translated                      struct {
			Name         string `json:"name"`
			CustomFields struct {
			} `json:"customFields"`
			HomeSlotConfig      interface{} `json:"homeSlotConfig"`
			HomeEnabled         bool        `json:"homeEnabled"`
			HomeName            interface{} `json:"homeName"`
			HomeMetaTitle       interface{} `json:"homeMetaTitle"`
			HomeMetaDescription interface{} `json:"homeMetaDescription"`
			HomeKeywords        interface{} `json:"homeKeywords"`
		} `json:"translated"`
		CreatedAt  time.Time   `json:"createdAt"`
		UpdatedAt  interface{} `json:"updatedAt"`
		Extensions struct {
			ForeignKeys struct {
				ApiAlias   interface{}   `json:"apiAlias"`
				Extensions []interface{} `json:"extensions"`
			} `json:"foreignKeys"`
		} `json:"extensions"`
		Id                          string      `json:"id"`
		CustomFields                interface{} `json:"customFields"`
		NavigationCategoryVersionId string      `json:"navigationCategoryVersionId"`
		FooterCategoryVersionId     *string     `json:"footerCategoryVersionId"`
		ServiceCategoryVersionId    *string     `json:"serviceCategoryVersionId"`
		HomeCmsPageVersionId        *string     `json:"homeCmsPageVersionId"`
		ApiAlias                    string      `json:"apiAlias"`
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

// SalesChannels is to get a list of all sales channels
func SalesChannels(parameter map[string]string, r Request) (SalesChannelsReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/sales-channel",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return SalesChannelsReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return SalesChannelsReturn{}, err
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
		return SalesChannelsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode SalesChannelsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return SalesChannelsReturn{}, err
	}

	// Return data
	return decode, err

}
