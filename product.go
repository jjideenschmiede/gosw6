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

// ProductsReturn is to decode the json data
type ProductsReturn struct {
	Total int `json:"total"`
	Data  []struct {
		ParentId       interface{} `json:"parentId"`
		ChildCount     int         `json:"childCount"`
		AutoIncrement  int         `json:"autoIncrement"`
		TaxId          string      `json:"taxId"`
		ManufacturerId interface{} `json:"manufacturerId"`
		UnitId         interface{} `json:"unitId"`
		Active         bool        `json:"active"`
		DisplayGroup   string      `json:"displayGroup"`
		Price          []struct {
			CurrencyId string        `json:"currencyId"`
			Net        int           `json:"net"`
			Gross      int           `json:"gross"`
			Linked     bool          `json:"linked"`
			ListPrice  interface{}   `json:"listPrice"`
			Percentage interface{}   `json:"percentage"`
			Extensions []interface{} `json:"extensions"`
			ApiAlias   string        `json:"apiAlias"`
		} `json:"price"`
		ManufacturerNumber interface{} `json:"manufacturerNumber"`
		Ean                interface{} `json:"ean"`
		Sales              int         `json:"sales"`
		ProductNumber      string      `json:"productNumber"`
		Stock              int         `json:"stock"`
		AvailableStock     int         `json:"availableStock"`
		Available          bool        `json:"available"`
		DeliveryTimeId     interface{} `json:"deliveryTimeId"`
		DeliveryTime       interface{} `json:"deliveryTime"`
		RestockTime        interface{} `json:"restockTime"`
		IsCloseout         bool        `json:"isCloseout"`
		PurchaseSteps      int         `json:"purchaseSteps"`
		MaxPurchase        interface{} `json:"maxPurchase"`
		MinPurchase        int         `json:"minPurchase"`
		PurchaseUnit       interface{} `json:"purchaseUnit"`
		ReferenceUnit      interface{} `json:"referenceUnit"`
		ShippingFree       bool        `json:"shippingFree"`
		PurchasePrices     []struct {
			CurrencyId string        `json:"currencyId"`
			Net        int           `json:"net"`
			Gross      int           `json:"gross"`
			Linked     bool          `json:"linked"`
			ListPrice  interface{}   `json:"listPrice"`
			Percentage interface{}   `json:"percentage"`
			Extensions []interface{} `json:"extensions"`
			ApiAlias   string        `json:"apiAlias"`
		} `json:"purchasePrices"`
		MarkAsTopseller         *bool         `json:"markAsTopseller"`
		Weight                  interface{}   `json:"weight"`
		Width                   interface{}   `json:"width"`
		Height                  interface{}   `json:"height"`
		Length                  interface{}   `json:"length"`
		ReleaseDate             interface{}   `json:"releaseDate"`
		CategoryTree            interface{}   `json:"categoryTree"`
		StreamIds               interface{}   `json:"streamIds"`
		OptionIds               interface{}   `json:"optionIds"`
		PropertyIds             interface{}   `json:"propertyIds"`
		Name                    string        `json:"name"`
		Keywords                interface{}   `json:"keywords"`
		Description             string        `json:"description"`
		MetaDescription         interface{}   `json:"metaDescription"`
		MetaTitle               interface{}   `json:"metaTitle"`
		PackUnit                interface{}   `json:"packUnit"`
		PackUnitPlural          interface{}   `json:"packUnitPlural"`
		VariantRestrictions     interface{}   `json:"variantRestrictions"`
		ConfiguratorGroupConfig interface{}   `json:"configuratorGroupConfig"`
		MainVariantId           interface{}   `json:"mainVariantId"`
		Variation               []interface{} `json:"variation"`
		Tax                     struct {
			TaxRate          int           `json:"taxRate"`
			Name             string        `json:"name"`
			Position         int           `json:"position"`
			Products         interface{}   `json:"products"`
			Rules            interface{}   `json:"rules"`
			ShippingMethods  interface{}   `json:"shippingMethods"`
			UniqueIdentifier string        `json:"_uniqueIdentifier"`
			VersionId        interface{}   `json:"versionId"`
			Translated       []interface{} `json:"translated"`
			CreatedAt        time.Time     `json:"createdAt"`
			UpdatedAt        interface{}   `json:"updatedAt"`
			Extensions       struct {
				ForeignKeys struct {
					ApiAlias   interface{}   `json:"apiAlias"`
					Extensions []interface{} `json:"extensions"`
				} `json:"foreignKeys"`
				InternalMappingStorage struct {
					ApiAlias   interface{}   `json:"apiAlias"`
					Extensions []interface{} `json:"extensions"`
				} `json:"internal_mapping_storage"`
			} `json:"extensions"`
			Id           string      `json:"id"`
			CustomFields interface{} `json:"customFields"`
			ApiAlias     string      `json:"apiAlias"`
		} `json:"tax"`
		Manufacturer  interface{}   `json:"manufacturer"`
		Unit          interface{}   `json:"unit"`
		Prices        []interface{} `json:"prices"`
		CheapestPrice struct {
			HasRange  bool        `json:"hasRange"`
			VariantId string      `json:"variantId"`
			ParentId  string      `json:"parentId"`
			RuleId    interface{} `json:"ruleId"`
			Purchase  interface{} `json:"purchase"`
			Reference interface{} `json:"reference"`
			UnitId    interface{} `json:"unitId"`
			Price     []struct {
				CurrencyId string        `json:"currencyId"`
				Net        int           `json:"net"`
				Gross      int           `json:"gross"`
				Linked     bool          `json:"linked"`
				ListPrice  interface{}   `json:"listPrice"`
				Percentage interface{}   `json:"percentage"`
				Extensions []interface{} `json:"extensions"`
				ApiAlias   string        `json:"apiAlias"`
			} `json:"price"`
			Extensions []interface{} `json:"extensions"`
			ApiAlias   string        `json:"apiAlias"`
		} `json:"cheapestPrice"`
		Cover                         interface{} `json:"cover"`
		Parent                        interface{} `json:"parent"`
		Children                      interface{} `json:"children"`
		Media                         interface{} `json:"media"`
		CmsPageId                     interface{} `json:"cmsPageId"`
		CmsPage                       interface{} `json:"cmsPage"`
		SlotConfig                    interface{} `json:"slotConfig"`
		SearchKeywords                interface{} `json:"searchKeywords"`
		Translations                  interface{} `json:"translations"`
		Categories                    interface{} `json:"categories"`
		CustomFieldSets               interface{} `json:"customFieldSets"`
		Tags                          interface{} `json:"tags"`
		Properties                    interface{} `json:"properties"`
		Options                       interface{} `json:"options"`
		ConfiguratorSettings          interface{} `json:"configuratorSettings"`
		CategoriesRo                  interface{} `json:"categoriesRo"`
		CoverId                       interface{} `json:"coverId"`
		BlacklistIds                  interface{} `json:"blacklistIds"`
		WhitelistIds                  interface{} `json:"whitelistIds"`
		Visibilities                  interface{} `json:"visibilities"`
		TagIds                        interface{} `json:"tagIds"`
		CategoryIds                   interface{} `json:"categoryIds"`
		ProductReviews                interface{} `json:"productReviews"`
		RatingAverage                 interface{} `json:"ratingAverage"`
		MainCategories                interface{} `json:"mainCategories"`
		SeoUrls                       interface{} `json:"seoUrls"`
		OrderLineItems                interface{} `json:"orderLineItems"`
		CrossSellings                 interface{} `json:"crossSellings"`
		CrossSellingAssignedProducts  interface{} `json:"crossSellingAssignedProducts"`
		FeatureSetId                  interface{} `json:"featureSetId"`
		FeatureSet                    interface{} `json:"featureSet"`
		CustomFieldSetSelectionActive interface{} `json:"customFieldSetSelectionActive"`
		CustomSearchKeywords          interface{} `json:"customSearchKeywords"`
		Wishlists                     interface{} `json:"wishlists"`
		CanonicalProductId            interface{} `json:"canonicalProductId"`
		CanonicalProduct              interface{} `json:"canonicalProduct"`
		CheapestPriceContainer        struct {
			Value struct {
				D8865E0973F47F397Dacf5Ee8802C46 struct {
					Default struct {
						ParentId  string      `json:"parent_id"`
						VariantId string      `json:"variant_id"`
						RuleId    interface{} `json:"rule_id"`
						IsRanged  string      `json:"is_ranged"`
						Price     struct {
							Cb7D2554B0Ce847Cd82F3Ac9Bd1C0Dfca struct {
								CurrencyId string `json:"currencyId"`
								Gross      int    `json:"gross"`
								Net        int    `json:"net"`
								Linked     bool   `json:"linked"`
							} `json:"cb7d2554b0ce847cd82f3ac9bd1c0dfca"`
						} `json:"price"`
						MinPurchase   string      `json:"min_purchase"`
						UnitId        interface{} `json:"unit_id"`
						PurchaseUnit  interface{} `json:"purchase_unit"`
						ReferenceUnit interface{} `json:"reference_unit"`
						ChildCount    string      `json:"child_count"`
					} `json:"default"`
				} `json:"3d8865e0973f47f397dacf5ee8802c46,omitempty"`
				A27Be66Bc743476089A0672290Eed674 struct {
					Default struct {
						ParentId  string      `json:"parent_id"`
						VariantId string      `json:"variant_id"`
						RuleId    interface{} `json:"rule_id"`
						IsRanged  string      `json:"is_ranged"`
						Price     struct {
							Cb7D2554B0Ce847Cd82F3Ac9Bd1C0Dfca struct {
								CurrencyId string `json:"currencyId"`
								Gross      int    `json:"gross"`
								Net        int    `json:"net"`
								Linked     bool   `json:"linked"`
							} `json:"cb7d2554b0ce847cd82f3ac9bd1c0dfca"`
						} `json:"price"`
						MinPurchase   string      `json:"min_purchase"`
						UnitId        interface{} `json:"unit_id"`
						PurchaseUnit  interface{} `json:"purchase_unit"`
						ReferenceUnit interface{} `json:"reference_unit"`
						ChildCount    string      `json:"child_count"`
					} `json:"default"`
				} `json:"a27be66bc743476089a0672290eed674,omitempty"`
			} `json:"value"`
			Default    interface{}   `json:"default"`
			Extensions []interface{} `json:"extensions"`
			ApiAlias   string        `json:"apiAlias"`
		} `json:"cheapestPriceContainer"`
		Streams          interface{} `json:"streams"`
		UniqueIdentifier string      `json:"_uniqueIdentifier"`
		VersionId        string      `json:"versionId"`
		Translated       struct {
			MetaDescription interface{} `json:"metaDescription"`
			Name            string      `json:"name"`
			Keywords        interface{} `json:"keywords"`
			Description     string      `json:"description"`
			MetaTitle       interface{} `json:"metaTitle"`
			PackUnit        interface{} `json:"packUnit"`
			PackUnitPlural  interface{} `json:"packUnitPlural"`
			CustomFields    struct {
			} `json:"customFields"`
			SlotConfig           interface{} `json:"slotConfig"`
			CustomSearchKeywords interface{} `json:"customSearchKeywords"`
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
}

// Products is to get a list of all products
func Products(r Request) (ProductsReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product",
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
