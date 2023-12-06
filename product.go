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
			CurrencyId string  `json:"currencyId"`
			Net        float64 `json:"net"`
			Gross      float64 `json:"gross"`
			Linked     bool    `json:"linked"`
			ListPrice  struct {
				CurrencyId      string        `json:"currencyId"`
				Net             float64       `json:"net"`
				Gross           float64       `json:"gross"`
				Linked          bool          `json:"linked"`
				ListPrice       interface{}   `json:"listPrice"`
				Percentage      interface{}   `json:"percentage"`
				RegulationPrice interface{}   `json:"regulationPrice"`
				Extensions      []interface{} `json:"extensions"`
				ApiAlias        string        `json:"apiAlias"`
			} `json:"listPrice"`
			Percentage struct {
				Net   float64 `json:"net"`
				Gross float64 `json:"gross"`
			} `json:"percentage"`
			RegulationPrice struct {
				CurrencyId      string        `json:"currencyId"`
				Net             float64       `json:"net"`
				Gross           float64       `json:"gross"`
				Linked          bool          `json:"linked"`
				ListPrice       interface{}   `json:"listPrice"`
				Percentage      interface{}   `json:"percentage"`
				RegulationPrice interface{}   `json:"regulationPrice"`
				Extensions      []interface{} `json:"extensions"`
				ApiAlias        string        `json:"apiAlias"`
			} `json:"regulationPrice"`
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
			Net        float64       `json:"net"`
			Gross      float64       `json:"gross"`
			Linked     bool          `json:"linked"`
			ListPrice  interface{}   `json:"listPrice"`
			Percentage interface{}   `json:"percentage"`
			Extensions []interface{} `json:"extensions"`
			ApiAlias   string        `json:"apiAlias"`
		} `json:"purchasePrices"`
		MarkAsTopseller         bool          `json:"markAsTopseller"`
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
				Net        float64       `json:"net"`
				Gross      float64       `json:"gross"`
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
			Value      interface{}   `json:"value"`
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
		Id           string         `json:"id"`
		CustomFields map[string]any `json:"customFields"`
		ApiAlias     string         `json:"apiAlias"`
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

// ProductBody is to structure the body data
type ProductBody struct {
	ParentId           interface{}                  `json:"parentId,omitempty"`
	ManufacturerId     string                       `json:"manufacturerId,omitempty"`
	Active             bool                         `json:"active"`
	Price              []*ProductBodyPrice          `json:"price"`
	ManufacturerNumber string                       `json:"manufacturerNumber,omitempty"`
	Ean                string                       `json:"ean,omitempty"`
	ProductNumber      string                       `json:"productNumber"`
	Stock              int                          `json:"stock"`
	DeliveryTimeId     string                       `json:"deliveryTimeId"`
	RestockTime        int                          `json:"restockTime,omitempty"`
	IsCloseout         bool                         `json:"isCloseout"`
	PurchaseSteps      int                          `json:"purchaseSteps,omitempty"`
	MaxPurchase        int                          `json:"maxPurchase"`
	MinPurchase        int                          `json:"minPurchase"`
	ShippingFree       bool                         `json:"shippingFree"`
	PurchasePrices     []*ProductBodyPurchasePrices `json:"purchasePrices"`
	MarkAsTopseller    bool                         `json:"markAsTopseller,omitempty"`
	Weight             float64                      `json:"weight"`
	Width              float64                      `json:"width"`
	Height             float64                      `json:"height"`
	Length             float64                      `json:"length"`
	UnitId             string                       `json:"unitId,omitempty"`
	PurchaseUnit       float64                      `json:"purchaseUnit"`
	ReferenceUnit      float64                      `json:"referenceUnit"`
	ReleaseDate        string                       `json:"releaseDate,omitempty"`
	Name               string                       `json:"name"`
	Tags               []*ProductBodyTags           `json:"tags"`
	Description        string                       `json:"description,omitempty"`
	MetaDescription    string                       `json:"metaDescription"`
	MetaTitle          string                       `json:"metaTitle"`
	Keywords           string                       `json:"keywords"`
	TaxId              string                       `json:"taxId,omitempty"`
	CmsPageId          string                       `json:"cmsPageId,omitempty"`
	Properties         []*ProductBodyProperties     `json:"properties"`
	Options            []*ProductBodyOptions        `json:"options"`
	Categories         []*ProductBodyCategories     `json:"categories"`
	CustomFields       *map[string]string           `json:"customFields,omitempty"`
}

type ProductBodyPrice struct {
	CurrencyId      string                     `json:"currencyId"`
	Net             float64                    `json:"net"`
	Gross           float64                    `json:"gross"`
	Linked          bool                       `json:"linked"`
	ListPrice       ProductBodyListPrice       `json:"listPrice"`
	RegulationPrice ProductBodyRegulationPrice `json:"regulationPrice"`
}

type ProductBodyListPrice struct {
	CurrencyId string  `json:"currencyId"`
	Net        float64 `json:"net"`
	Gross      float64 `json:"gross"`
	Linked     bool    `json:"linked"`
}

type ProductBodyRegulationPrice struct {
	CurrencyId string  `json:"currencyId"`
	Net        float64 `json:"net"`
	Gross      float64 `json:"gross"`
	Linked     bool    `json:"linked"`
}

type ProductBodyPurchasePrices struct {
	CurrencyId string  `json:"currencyId"`
	Net        float64 `json:"net"`
	Gross      float64 `json:"gross"`
	Linked     bool    `json:"linked"`
}

type ProductBodyTags struct {
	Id string `json:"id"`
}

type ProductBodyProperties struct {
	Id string `json:"id"`
}

type ProductBodyOptions struct {
	Id string `json:"id"`
}

type ProductBodyCategories struct {
	Id string `json:"id"`
}

// ProductReturn is to decode the json return
type ProductReturn struct {
	Data struct {
		ParentId       interface{} `json:"parentId"`
		ChildCount     int         `json:"childCount"`
		AutoIncrement  int         `json:"autoIncrement"`
		TaxId          string      `json:"taxId"`
		ManufacturerId interface{} `json:"manufacturerId"`
		UnitId         interface{} `json:"unitId"`
		Active         bool        `json:"active"`
		DisplayGroup   string      `json:"displayGroup"`
		Price          []struct {
			CurrencyId string  `json:"currencyId"`
			Net        float64 `json:"net"`
			Gross      float64 `json:"gross"`
			Linked     bool    `json:"linked"`
			ListPrice  struct {
				CurrencyId      string        `json:"currencyId"`
				Net             float64       `json:"net"`
				Gross           float64       `json:"gross"`
				Linked          bool          `json:"linked"`
				ListPrice       interface{}   `json:"listPrice"`
				Percentage      interface{}   `json:"percentage"`
				RegulationPrice interface{}   `json:"regulationPrice"`
				Extensions      []interface{} `json:"extensions"`
				ApiAlias        string        `json:"apiAlias"`
			} `json:"listPrice"`
			Percentage struct {
				Net   float64 `json:"net"`
				Gross float64 `json:"gross"`
			} `json:"percentage"`
			RegulationPrice struct {
				CurrencyId      string        `json:"currencyId"`
				Net             float64       `json:"net"`
				Gross           float64       `json:"gross"`
				Linked          bool          `json:"linked"`
				ListPrice       interface{}   `json:"listPrice"`
				Percentage      interface{}   `json:"percentage"`
				RegulationPrice interface{}   `json:"regulationPrice"`
				Extensions      []interface{} `json:"extensions"`
				ApiAlias        string        `json:"apiAlias"`
			} `json:"regulationPrice"`
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
			Net        float64       `json:"net"`
			Gross      float64       `json:"gross"`
			Linked     bool          `json:"linked"`
			ListPrice  interface{}   `json:"listPrice"`
			Percentage interface{}   `json:"percentage"`
			Extensions []interface{} `json:"extensions"`
			ApiAlias   string        `json:"apiAlias"`
		} `json:"purchasePrices"`
		MarkAsTopseller         bool          `json:"markAsTopseller"`
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
				} `json:"a27be66bc743476089a0672290eed674"`
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
		Id           string         `json:"id"`
		CustomFields map[string]any `json:"customFields"`
		ApiAlias     string         `json:"apiAlias"`
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

// CreateProductReturn is decode the json data
type CreateProductReturn struct {
	Location string `json:"location"`
	Errors   []struct {
		Code     string `json:"code"`
		Status   string `json:"status"`
		Detail   string `json:"detail"`
		Template string `json:"template"`
		Meta     struct {
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
		Source struct {
			Pointer string `json:"pointer"`
		} `json:"source"`
	} `json:"errors"`
}

// UpdateProductReturn is decode the json data
type UpdateProductReturn struct {
	Errors []struct {
		Code     string `json:"code"`
		Status   string `json:"status"`
		Detail   string `json:"detail"`
		Template string `json:"template"`
		Meta     struct {
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
		Source struct {
			Pointer string `json:"pointer"`
		} `json:"source"`
	} `json:"errors"`
}

// DeleteProductReturn is to decode the json data
type DeleteProductReturn struct {
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

// ProductConfiguratorSettingsReturn is to decode the json data
type ProductConfiguratorSettingsReturn struct {
	Total int `json:"total"`
	Data  []struct {
		ProductId        string        `json:"productId"`
		OptionId         string        `json:"optionId"`
		MediaId          interface{}   `json:"mediaId"`
		Position         int           `json:"position"`
		Price            interface{}   `json:"price"`
		Option           interface{}   `json:"option"`
		Media            interface{}   `json:"media"`
		Selected         bool          `json:"selected"`
		Product          interface{}   `json:"product"`
		UniqueIdentifier string        `json:"_uniqueIdentifier"`
		VersionId        string        `json:"versionId"`
		Translated       []interface{} `json:"translated"`
		CreatedAt        time.Time     `json:"createdAt"`
		UpdatedAt        interface{}   `json:"updatedAt"`
		Extensions       struct {
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

// UpdateProductConfiguratorSettingBody is to structure the body data
type UpdateProductConfiguratorSettingBody struct {
	ConfiguratorSettings []UpdateProductConfiguratorSettingBodyConfiguratorSettings `json:"configuratorSettings"`
}

type UpdateProductConfiguratorSettingBodyConfiguratorSettings struct {
	OptionId string `json:"optionId"`
}

// UpdateProductConfiguratorSettingReturn is to decode the json data
type UpdateProductConfiguratorSettingReturn struct {
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

// DeleteProductConfiguratorSettingReturn is to decode the json data
type DeleteProductConfiguratorSettingReturn struct {
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

// UpdateProductStockBody is to structure the body data
type UpdateProductStockBody struct {
	Stock       int `json:"stock"`
	MaxPurchase int `json:"maxPurchase"`
}

// UpdateProductStockReturn is to decode the json data
type UpdateProductStockReturn struct {
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

// UpdateProductPriceBody is to structure the body data
type UpdateProductPriceBody struct {
	Price          []*UpdateProductPriceBodyPrice          `json:"price"`
	PurchasePrices []*UpdateProductPriceBodyPurchasePrices `json:"purchasePrices"`
}

type UpdateProductPriceBodyPrice struct {
	CurrencyId      string                                `json:"currencyId"`
	Net             float64                               `json:"net"`
	Gross           float64                               `json:"gross"`
	Linked          bool                                  `json:"linked"`
	ListPrice       UpdateProductPriceBodyListPrice       `json:"listPrice"`
	RegulationPrice UpdateProductPriceBodyRegulationPrice `json:"regulationPrice"`
}

type UpdateProductPriceBodyListPrice struct {
	CurrencyId string  `json:"currencyId"`
	Net        float64 `json:"net"`
	Gross      float64 `json:"gross"`
	Linked     bool    `json:"linked"`
}

type UpdateProductPriceBodyRegulationPrice struct {
	CurrencyId string  `json:"currencyId"`
	Net        float64 `json:"net"`
	Gross      float64 `json:"gross"`
	Linked     bool    `json:"linked"`
}

type UpdateProductPriceBodyPurchasePrices struct {
	CurrencyId string  `json:"currencyId"`
	Net        float64 `json:"net"`
	Gross      float64 `json:"gross"`
	Linked     bool    `json:"linked"`
}

// UpdateProductPriceReturn is to decode the json data
type UpdateProductPriceReturn struct {
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

// ProductPropertiesReturn is to decode the json data
type ProductPropertiesReturn struct {
	Total int `json:"total"`
	Data  []struct {
		GroupId                     string      `json:"groupId"`
		Name                        string      `json:"name"`
		Position                    int         `json:"position"`
		ColorHexCode                interface{} `json:"colorHexCode"`
		MediaId                     interface{} `json:"mediaId"`
		Group                       interface{} `json:"group"`
		Translations                interface{} `json:"translations"`
		ProductConfiguratorSettings interface{} `json:"productConfiguratorSettings"`
		ProductProperties           interface{} `json:"productProperties"`
		ProductOptions              interface{} `json:"productOptions"`
		Media                       interface{} `json:"media"`
		UniqueIdentifier            string      `json:"_uniqueIdentifier"`
		VersionId                   interface{} `json:"versionId"`
		Translated                  struct {
			Name         string `json:"name"`
			Position     int    `json:"position"`
			CustomFields struct {
			} `json:"customFields"`
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

// DeleteProductPropertyReturn is to decode the json data
type DeleteProductPropertyReturn struct {
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

// ProductCategoriesReturn is to decode the json data
type ProductCategoriesReturn struct {
	Total int `json:"total"`
	Data  []struct {
		ParentId                string      `json:"parentId"`
		AutoIncrement           int         `json:"autoIncrement"`
		MediaId                 interface{} `json:"mediaId"`
		Name                    string      `json:"name"`
		Breadcrumb              []string    `json:"breadcrumb"`
		Path                    string      `json:"path"`
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

// DeleteProductCategoryReturn is to decode the json data
type DeleteProductCategoryReturn struct {
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

// ProductTagsReturn is to decode the json data
type ProductTagsReturn struct {
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

// DeleteProductTagReturn is to decode the json data
type DeleteProductTagReturn struct {
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

// ProductManufacturersReturn is to decode the json data
type ProductManufacturersReturn struct {
	Total int `json:"total"`
	Data  []struct {
		MediaId          interface{} `json:"mediaId"`
		Name             string      `json:"name"`
		Link             interface{} `json:"link"`
		Description      interface{} `json:"description"`
		Media            interface{} `json:"media"`
		Translations     interface{} `json:"translations"`
		Products         interface{} `json:"products"`
		UniqueIdentifier string      `json:"_uniqueIdentifier"`
		VersionId        string      `json:"versionId"`
		Translated       struct {
			Name         string      `json:"name"`
			Description  interface{} `json:"description"`
			CustomFields struct {
			} `json:"customFields"`
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

// DeleteProductManufacturerReturn is to decode the json data
type DeleteProductManufacturerReturn struct {
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

// ProductMediaReturn is to decode the json data
type ProductMediaReturn struct {
	Total int `json:"total"`
	Data  []struct {
		ProductId string `json:"productId"`
		MediaId   string `json:"mediaId"`
		Position  int    `json:"position"`
		Media     struct {
			UserId        string      `json:"userId"`
			MimeType      string      `json:"mimeType"`
			FileExtension string      `json:"fileExtension"`
			FileSize      int         `json:"fileSize"`
			Title         interface{} `json:"title"`
			MetaData      struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Type   int `json:"type"`
			} `json:"metaData"`
			MediaType struct {
				Name       string        `json:"name"`
				Flags      []string      `json:"flags"`
				Extensions []interface{} `json:"extensions"`
				ApiAlias   string        `json:"apiAlias"`
			} `json:"mediaType"`
			UploadedAt           time.Time   `json:"uploadedAt"`
			Alt                  interface{} `json:"alt"`
			Url                  string      `json:"url"`
			FileName             string      `json:"fileName"`
			User                 interface{} `json:"user"`
			Translations         interface{} `json:"translations"`
			Categories           interface{} `json:"categories"`
			ProductManufacturers interface{} `json:"productManufacturers"`
			ProductMedia         interface{} `json:"productMedia"`
			AvatarUser           interface{} `json:"avatarUser"`
			Thumbnails           []struct {
				Width            int           `json:"width"`
				Height           int           `json:"height"`
				Url              string        `json:"url"`
				MediaId          string        `json:"mediaId"`
				Media            interface{}   `json:"media"`
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
				} `json:"extensions"`
				Id           string      `json:"id"`
				CustomFields interface{} `json:"customFields"`
				ApiAlias     string      `json:"apiAlias"`
			} `json:"thumbnails"`
			MediaFolderId               string      `json:"mediaFolderId"`
			MediaFolder                 interface{} `json:"mediaFolder"`
			HasFile                     bool        `json:"hasFile"`
			Private                     bool        `json:"private"`
			PropertyGroupOptions        interface{} `json:"propertyGroupOptions"`
			MailTemplateMedia           interface{} `json:"mailTemplateMedia"`
			Tags                        interface{} `json:"tags"`
			DocumentBaseConfigs         interface{} `json:"documentBaseConfigs"`
			ShippingMethods             interface{} `json:"shippingMethods"`
			PaymentMethods              interface{} `json:"paymentMethods"`
			ProductConfiguratorSettings interface{} `json:"productConfiguratorSettings"`
			OrderLineItems              interface{} `json:"orderLineItems"`
			CmsBlocks                   interface{} `json:"cmsBlocks"`
			CmsSections                 interface{} `json:"cmsSections"`
			CmsPages                    interface{} `json:"cmsPages"`
			Documents                   interface{} `json:"documents"`
			AppPaymentMethods           interface{} `json:"appPaymentMethods"`
			UniqueIdentifier            string      `json:"_uniqueIdentifier"`
			VersionId                   interface{} `json:"versionId"`
			Translated                  struct {
				Alt          interface{}   `json:"alt"`
				Title        interface{}   `json:"title"`
				CustomFields []interface{} `json:"customFields"`
			} `json:"translated"`
			CreatedAt  time.Time `json:"createdAt"`
			UpdatedAt  time.Time `json:"updatedAt"`
			Extensions struct {
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
		} `json:"media"`
		Product          interface{}   `json:"product"`
		UniqueIdentifier string        `json:"_uniqueIdentifier"`
		VersionId        string        `json:"versionId"`
		Translated       []interface{} `json:"translated"`
		CreatedAt        time.Time     `json:"createdAt"`
		UpdatedAt        interface{}   `json:"updatedAt"`
		Extensions       struct {
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

// ProductMediaBody is to structure the body data
type ProductMediaBody struct {
	MediaId  string `json:"mediaId"`
	Position int    `json:"position"`
}

// CreateProductMediaReturn is to decode the json data
type CreateProductMediaReturn struct {
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

// ProductCoverReturn is to decode the json return
type ProductCoverReturn struct {
	Total int `json:"total"`
	Data  []struct {
		ProductId string `json:"productId"`
		MediaId   string `json:"mediaId"`
		Position  int    `json:"position"`
		Media     struct {
			UserId        string      `json:"userId"`
			MimeType      string      `json:"mimeType"`
			FileExtension string      `json:"fileExtension"`
			FileSize      int         `json:"fileSize"`
			Title         interface{} `json:"title"`
			MetaData      interface{} `json:"metaData"`
			MediaType     struct {
				Name       string        `json:"name"`
				Flags      []interface{} `json:"flags"`
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
				Alt          interface{}   `json:"alt"`
				Title        interface{}   `json:"title"`
				CustomFields []interface{} `json:"customFields"`
			} `json:"translated"`
			CreatedAt  time.Time `json:"createdAt"`
			UpdatedAt  time.Time `json:"updatedAt"`
			Extensions struct {
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
		} `json:"media"`
		Product          interface{}   `json:"product"`
		UniqueIdentifier string        `json:"_uniqueIdentifier"`
		VersionId        string        `json:"versionId"`
		Translated       []interface{} `json:"translated"`
		CreatedAt        time.Time     `json:"createdAt"`
		UpdatedAt        interface{}   `json:"updatedAt"`
		Extensions       struct {
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

// CreateProductCoverBody is to structure the body data
type CreateProductCoverBody struct {
	ProductId string `json:"productId"`
	MediaId   string `json:"mediaId"`
	Position  int    `json:"position"`
}

// CreateProductCoverReturn is to decode the json data
type CreateProductCoverReturn struct {
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

// ProductVisibilitiesReturn is to decode the json return
type ProductVisibilitiesReturn struct {
	Total int `json:"total"`
	Data  []struct {
		Visibility       int           `json:"visibility"`
		ProductId        string        `json:"productId"`
		SalesChannelId   string        `json:"salesChannelId"`
		Product          interface{}   `json:"product"`
		SalesChannel     interface{}   `json:"salesChannel"`
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

// ProductVisibilityBody is to structure the data
type ProductVisibilityBody struct {
	Visibility     int    `json:"visibility"`
	SalesChannelId string `json:"salesChannelId"`
}

// CreateProductVisibilityReturn is to decode the json return
type CreateProductVisibilityReturn struct {
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

// ProductPricesReturn is to decode the json return
type ProductPricesReturn struct {
	Total int `json:"total"`
	Data  []struct {
		ProductId     string      `json:"productId"`
		QuantityStart int         `json:"quantityStart"`
		QuantityEnd   int         `json:"quantityEnd"`
		Product       interface{} `json:"product"`
		Rule          interface{} `json:"rule"`
		RuleId        string      `json:"ruleId"`
		Price         []struct {
			CurrencyId string  `json:"currencyId"`
			Net        float64 `json:"net"`
			Gross      float64 `json:"gross"`
			Linked     bool    `json:"linked"`
			ListPrice  struct {
				CurrencyId      string        `json:"currencyId"`
				Net             float64       `json:"net"`
				Gross           float64       `json:"gross"`
				Linked          bool          `json:"linked"`
				ListPrice       interface{}   `json:"listPrice"`
				Percentage      interface{}   `json:"percentage"`
				RegulationPrice interface{}   `json:"regulationPrice"`
				Extensions      []interface{} `json:"extensions"`
				ApiAlias        string        `json:"apiAlias"`
			} `json:"listPrice"`
			Percentage struct {
				Net   float64 `json:"net"`
				Gross float64 `json:"gross"`
			} `json:"percentage"`
			RegulationPrice struct {
				CurrencyId      string        `json:"currencyId"`
				Net             float64       `json:"net"`
				Gross           float64       `json:"gross"`
				Linked          bool          `json:"linked"`
				ListPrice       interface{}   `json:"listPrice"`
				Percentage      interface{}   `json:"percentage"`
				RegulationPrice interface{}   `json:"regulationPrice"`
				Extensions      []interface{} `json:"extensions"`
				ApiAlias        string        `json:"apiAlias"`
			} `json:"regulationPrice"`
			Extensions []interface{} `json:"extensions"`
			ApiAlias   string        `json:"apiAlias"`
		} `json:"price"`
		UniqueIdentifier string        `json:"_uniqueIdentifier"`
		VersionId        string        `json:"versionId"`
		Translated       []interface{} `json:"translated"`
		CreatedAt        time.Time     `json:"createdAt"`
		UpdatedAt        time.Time     `json:"updatedAt"`
		Extensions       struct {
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

// ProductPriceBody is to structure the data
type ProductPriceBody struct {
	QuantityStart int                     `json:"quantityStart"`
	QuantityEnd   int                     `json:"quantityEnd,omitempty"`
	RuleId        string                  `json:"ruleId"`
	Price         []ProductPriceBodyPrice `json:"price"`
}

type ProductPriceBodyPrice struct {
	CurrencyId      string                          `json:"currencyId"`
	Net             float64                         `json:"net"`
	Gross           float64                         `json:"gross"`
	Linked          bool                            `json:"linked"`
	ListPrice       ProductPriceBodyListPrice       `json:"listPrice"`
	RegulationPrice ProductPriceBodyRegulationPrice `json:"regulationPrice"`
}

type ProductPriceBodyListPrice struct {
	CurrencyId string  `json:"currencyId"`
	Net        float64 `json:"net"`
	Gross      float64 `json:"gross"`
	Linked     bool    `json:"linked"`
}

type ProductPriceBodyRegulationPrice struct {
	CurrencyId string  `json:"currencyId"`
	Net        float64 `json:"net"`
	Gross      float64 `json:"gross"`
	Linked     bool    `json:"linked"`
}

// CreateProductPriceReturn is to decode the json return
type CreateProductPriceReturn struct {
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

// DeleteProductPriceReturn is to decode the json data
type DeleteProductPriceReturn struct {
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

// Products are to get a list of all products
func Products(parameter map[string]string, r Request) (ProductsReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return ProductsReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return ProductsReturn{}, err
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

// Product is to get a specific product by id
func Product(id string, r Request) (ProductReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id,
		Method: "GET",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductReturn{}, err
	}

	// Return data
	return decode, err

}

// CreateProduct is to create a product
func CreateProduct(body ProductBody, r Request) (CreateProductReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return CreateProductReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/product",
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CreateProductReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CreateProductReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return CreateProductReturn{}, err
		}
	}

	// Get location in header & set to return struct
	decode.Location = response.Header.Get("location")

	// Return data
	return decode, err

}

// UpdateProduct is to update a product
func UpdateProduct(id string, body ProductBody, r Request) (UpdateProductReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return UpdateProductReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id,
		Method: "PATCH",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return UpdateProductReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode UpdateProductReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return UpdateProductReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// DeleteProduct is to delete a specific product by id
func DeleteProduct(id string, r Request) (DeleteProductReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id,
		Method: "DELETE",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return DeleteProductReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode DeleteProductReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return DeleteProductReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// ProductConfiguratorSettings is to get all product configuration settings
func ProductConfiguratorSettings(parameter map[string]string, id string, r Request) (ProductConfiguratorSettingsReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id + "/configuratorSettings",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return ProductConfiguratorSettingsReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return ProductConfiguratorSettingsReturn{}, err
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
		return ProductConfiguratorSettingsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductConfiguratorSettingsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductConfiguratorSettingsReturn{}, err
	}

	// Return data
	return decode, err

}

// UpdateProductConfiguratorSetting is to update the configurator settings
func UpdateProductConfiguratorSetting(id string, body UpdateProductConfiguratorSettingBody, r Request) (UpdateProductConfiguratorSettingReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return UpdateProductConfiguratorSettingReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id,
		Method: "PATCH",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return UpdateProductConfiguratorSettingReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode UpdateProductConfiguratorSettingReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return UpdateProductConfiguratorSettingReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// DeleteProductConfiguratorSetting is to delete all product configuration setting
func DeleteProductConfiguratorSetting(productId, configurationSettingId string, r Request) (DeleteProductConfiguratorSettingReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product/" + productId + "/configuratorSettings/" + configurationSettingId,
		Method: "DELETE",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return DeleteProductConfiguratorSettingReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode DeleteProductConfiguratorSettingReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return DeleteProductConfiguratorSettingReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// UpdateProductStock is to update the stock of an product
func UpdateProductStock(id string, body UpdateProductStockBody, r Request) (UpdateProductStockReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return UpdateProductStockReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id,
		Method: "PATCH",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return UpdateProductStockReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode UpdateProductStockReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return UpdateProductStockReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// UpdateProductPrice is to update the price of a product
func UpdateProductPrice(id string, body UpdateProductPriceBody, r Request) (UpdateProductPriceReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return UpdateProductPriceReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id,
		Method: "PATCH",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return UpdateProductPriceReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode UpdateProductPriceReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return UpdateProductPriceReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// ProductProperties are to get a list of all product properties
func ProductProperties(parameter map[string]string, id string, r Request) (ProductPropertiesReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id + "/properties",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return ProductPropertiesReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return ProductPropertiesReturn{}, err
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
		return ProductPropertiesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductPropertiesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductPropertiesReturn{}, err
	}

	// Return data
	return decode, err

}

// DeleteProductProperty is to delete a product property
func DeleteProductProperty(productId, propertyId string, r Request) (DeleteProductPropertyReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product/" + productId + "/properties/" + propertyId,
		Method: "DELETE",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return DeleteProductPropertyReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode DeleteProductPropertyReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return DeleteProductPropertyReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// ProductCategories are to get a list of all product categories
func ProductCategories(parameter map[string]string, id string, r Request) (ProductCategoriesReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id + "/categories",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return ProductCategoriesReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return ProductCategoriesReturn{}, err
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
		return ProductCategoriesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductCategoriesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductCategoriesReturn{}, err
	}

	// Return data
	return decode, err

}

// DeleteProductCategory is to delete a product category
func DeleteProductCategory(productId, categoryId string, r Request) (DeleteProductCategoryReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product/" + productId + "/categories/" + categoryId,
		Method: "DELETE",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return DeleteProductCategoryReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode DeleteProductCategoryReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return DeleteProductCategoryReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// ProductTags are to get a list of all product tags
func ProductTags(parameter map[string]string, id string, r Request) (ProductTagsReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id + "/tags",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return ProductTagsReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return ProductTagsReturn{}, err
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
		return ProductTagsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductTagsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductTagsReturn{}, err
	}

	// Return data
	return decode, err

}

// DeleteProductTag is to delete a product tag
func DeleteProductTag(productId, tagId string, r Request) (DeleteProductTagReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product/" + productId + "/tags/" + tagId,
		Method: "DELETE",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return DeleteProductTagReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode DeleteProductTagReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return DeleteProductTagReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// ProductManufacturers are to get a list of all manufacturers
func ProductManufacturers(parameter map[string]string, id string, r Request) (ProductManufacturersReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id + "/manufacturer",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return ProductManufacturersReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return ProductManufacturersReturn{}, err
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
		return ProductManufacturersReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductManufacturersReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductManufacturersReturn{}, err
	}

	// Return data
	return decode, err

}

// DeleteProductManufacturer is to delete a product manufacturer
func DeleteProductManufacturer(productId, manufacturerId string, r Request) (DeleteProductManufacturerReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product/" + productId + "/manufacturer/" + manufacturerId,
		Method: "DELETE",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return DeleteProductManufacturerReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode DeleteProductManufacturerReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return DeleteProductManufacturerReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// ProductMedia are to get a list of all product media
func ProductMedia(parameter map[string]string, id string, r Request) (ProductMediaReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id + "/media",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return ProductMediaReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return ProductMediaReturn{}, err
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
		return ProductMediaReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductMediaReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductMediaReturn{}, err
	}

	// Return data
	return decode, err

}

// CreateProductMedia is to create a product media
func CreateProductMedia(id string, body ProductMediaBody, r Request) (CreateProductMediaReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return CreateProductMediaReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id + "/media",
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CreateProductMediaReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CreateProductMediaReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return CreateProductMediaReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// ProductCover are to get a list of all product covers
func ProductCover(parameter map[string]string, id string, r Request) (ProductCoverReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id + "/cover",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return ProductCoverReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return ProductCoverReturn{}, err
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
		return ProductCoverReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductCoverReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductCoverReturn{}, err
	}

	// Return data
	return decode, err

}

// CreateProductCover is to create a product cover
func CreateProductCover(id string, body CreateProductCoverBody, r Request) (CreateProductCoverReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return CreateProductCoverReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id + "/cover",
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CreateProductCoverReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CreateProductCoverReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return CreateProductCoverReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// ProductVisibilities are to get a list of all product visibilities
func ProductVisibilities(parameter map[string]string, id string, r Request) (ProductVisibilitiesReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id + "/visibilities",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return ProductVisibilitiesReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return ProductVisibilitiesReturn{}, err
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
		return ProductVisibilitiesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductVisibilitiesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductVisibilitiesReturn{}, err
	}

	// Return data
	return decode, err

}

// CreateProductVisibility is to create a product visibility
func CreateProductVisibility(id string, body ProductVisibilityBody, r Request) (CreateProductVisibilityReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return CreateProductVisibilityReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id + "/visibilities",
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CreateProductVisibilityReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CreateProductVisibilityReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return CreateProductVisibilityReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// ProductPrices are to get a list of all product prices
func ProductPrices(parameter map[string]string, id string, r Request) (ProductPricesReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id + "/prices",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return ProductPricesReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return ProductPricesReturn{}, err
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
		return ProductPricesReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductPricesReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductPricesReturn{}, err
	}

	// Return data
	return decode, err

}

// CreateProductPrice is to create a product price
func CreateProductPrice(id string, body ProductPriceBody, r Request) (CreateProductPriceReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return CreateProductPriceReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/product/" + id + "/prices",
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CreateProductPriceReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CreateProductPriceReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return CreateProductPriceReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// DeleteProductPrice is to delete a product price
func DeleteProductPrice(productId, priceId string, r Request) (DeleteProductPriceReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/product/" + productId + "/prices/" + priceId,
		Method: "DELETE",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return DeleteProductPriceReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode DeleteProductPriceReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return DeleteProductPriceReturn{}, err
		}
	}

	// Return data
	return decode, err

}
