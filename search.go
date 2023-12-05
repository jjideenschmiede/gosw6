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
	"time"
)

// CustomOrderSearchBody is to decode the json data
type CustomOrderSearchBody struct {
	Page         int                           `json:"page"`
	Limit        int                           `json:"limit"`
	Filter       []CustomOrderSearchBodyFilter `json:"filter"`
	Sort         []CustomOrderSearchBodySort   `json:"sort"`
	Associations struct {
		Currency struct {
		} `json:"currency"`
		LineItems struct {
			Associations struct {
				Product struct {
				} `json:"product"`
			} `json:"associations"`
		} `json:"lineItems"`
		Addresses struct {
			Associations struct {
				Country struct {
				} `json:"country"`
				CountryState struct {
				} `json:"countryState"`
				Salutation struct {
				} `json:"salutation"`
			} `json:"associations"`
		} `json:"addresses"`
		Transactions struct {
			Associations struct {
				PaymentMethod struct {
				} `json:"paymentMethod"`
			} `json:"associations"`
		} `json:"transactions"`
		Deliveries struct {
			Associations struct {
				ShippingMethod struct {
				} `json:"shippingMethod"`
			} `json:"associations"`
		} `json:"deliveries"`
	} `json:"associations"`
	Includes struct {
		Product []string `json:"product"`
	} `json:"includes"`
}

type CustomOrderSearchBodyFilter struct {
	Type       string                          `json:"type"`
	Field      string                          `json:"field"`
	Parameters CustomOrderSearchBodyParameters `json:"parameters"`
}

type CustomOrderSearchBodyParameters struct {
	Gte string `json:"gte"`
}

type CustomOrderSearchBodySort struct {
	Field string `json:"field"`
	Order string `json:"order"`
}

// CustomOrderSearchReturn is to decode the json data
type CustomOrderSearchReturn struct {
	Total int `json:"total"`
	Data  []struct {
		OrderNumber      string    `json:"orderNumber"`
		CurrencyId       string    `json:"currencyId"`
		CurrencyFactor   float64   `json:"currencyFactor"`
		SalesChannelId   string    `json:"salesChannelId"`
		BillingAddressId string    `json:"billingAddressId"`
		OrderDateTime    time.Time `json:"orderDateTime"`
		OrderDate        time.Time `json:"orderDate"`
		Price            struct {
			NetPrice        float64 `json:"netPrice"`
			TotalPrice      float64 `json:"totalPrice"`
			CalculatedTaxes []struct {
				Tax        float64       `json:"tax"`
				TaxRate    float64       `json:"taxRate"`
				Price      float64       `json:"price"`
				Extensions []interface{} `json:"extensions"`
				ApiAlias   string        `json:"apiAlias"`
			} `json:"calculatedTaxes"`
			TaxRules []struct {
				TaxRate    float64       `json:"taxRate"`
				Percentage float64       `json:"percentage"`
				Extensions []interface{} `json:"extensions"`
				ApiAlias   string        `json:"apiAlias"`
			} `json:"taxRules"`
			PositionPrice float64       `json:"positionPrice"`
			TaxStatus     string        `json:"taxStatus"`
			RawTotal      float64       `json:"rawTotal"`
			Extensions    []interface{} `json:"extensions"`
			ApiAlias      string        `json:"apiAlias"`
		} `json:"price"`
		AmountTotal   float64 `json:"amountTotal"`
		AmountNet     float64 `json:"amountNet"`
		PositionPrice float64 `json:"positionPrice"`
		TaxStatus     string  `json:"taxStatus"`
		ShippingCosts struct {
			UnitPrice       float64 `json:"unitPrice"`
			Quantity        int     `json:"quantity"`
			TotalPrice      float64 `json:"totalPrice"`
			CalculatedTaxes []struct {
				Tax        float64       `json:"tax"`
				TaxRate    float64       `json:"taxRate"`
				Price      float64       `json:"price"`
				Extensions []interface{} `json:"extensions"`
				ApiAlias   string        `json:"apiAlias"`
			} `json:"calculatedTaxes"`
			TaxRules []struct {
				TaxRate    float64       `json:"taxRate"`
				Percentage float64       `json:"percentage"`
				Extensions []interface{} `json:"extensions"`
				ApiAlias   string        `json:"apiAlias"`
			} `json:"taxRules"`
			ReferencePrice interface{}   `json:"referencePrice"`
			ListPrice      interface{}   `json:"listPrice"`
			Extensions     []interface{} `json:"extensions"`
			ApiAlias       string        `json:"apiAlias"`
		} `json:"shippingCosts"`
		ShippingTotal float64 `json:"shippingTotal"`
		OrderCustomer struct {
			Email            string        `json:"email"`
			OrderId          string        `json:"orderId"`
			SalutationId     string        `json:"salutationId"`
			FirstName        string        `json:"firstName"`
			LastName         string        `json:"lastName"`
			Title            interface{}   `json:"title"`
			VatIds           []interface{} `json:"vatIds"`
			Company          string        `json:"company"`
			CustomerNumber   string        `json:"customerNumber"`
			CustomerId       string        `json:"customerId"`
			Customer         interface{}   `json:"customer"`
			Salutation       interface{}   `json:"salutation"`
			Order            interface{}   `json:"order"`
			RemoteAddress    string        `json:"remoteAddress"`
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
				InternalMappingStorage struct {
					ApiAlias   interface{}   `json:"apiAlias"`
					Extensions []interface{} `json:"extensions"`
				} `json:"internal_mapping_storage"`
			} `json:"extensions"`
			Id             string      `json:"id"`
			CustomFields   interface{} `json:"customFields"`
			OrderVersionId string      `json:"orderVersionId"`
			ApiAlias       string      `json:"apiAlias"`
		} `json:"orderCustomer"`
		Currency struct {
			IsoCode                        string      `json:"isoCode"`
			Factor                         float64     `json:"factor"`
			Symbol                         string      `json:"symbol"`
			ShortName                      string      `json:"shortName"`
			Name                           string      `json:"name"`
			Position                       int         `json:"position"`
			Translations                   interface{} `json:"translations"`
			Orders                         interface{} `json:"orders"`
			SalesChannels                  interface{} `json:"salesChannels"`
			SalesChannelDefaultAssignments interface{} `json:"salesChannelDefaultAssignments"`
			SalesChannelDomains            interface{} `json:"salesChannelDomains"`
			ShippingMethodPrices           interface{} `json:"shippingMethodPrices"`
			PromotionDiscountPrices        interface{} `json:"promotionDiscountPrices"`
			IsSystemDefault                bool        `json:"isSystemDefault"`
			ProductExports                 interface{} `json:"productExports"`
			CountryRoundings               interface{} `json:"countryRoundings"`
			ItemRounding                   struct {
				Decimals    int           `json:"decimals"`
				Interval    float64       `json:"interval"`
				RoundForNet bool          `json:"roundForNet"`
				Extensions  []interface{} `json:"extensions"`
				ApiAlias    string        `json:"apiAlias"`
			} `json:"itemRounding"`
			TotalRounding struct {
				Decimals    int           `json:"decimals"`
				Interval    float64       `json:"interval"`
				RoundForNet bool          `json:"roundForNet"`
				Extensions  []interface{} `json:"extensions"`
				ApiAlias    string        `json:"apiAlias"`
			} `json:"totalRounding"`
			TaxFreeFrom      int         `json:"taxFreeFrom"`
			UniqueIdentifier string      `json:"_uniqueIdentifier"`
			VersionId        interface{} `json:"versionId"`
			Translated       struct {
				ShortName    string      `json:"shortName"`
				Name         string      `json:"name"`
				CustomFields interface{} `json:"customFields"`
			} `json:"translated"`
			CreatedAt  time.Time   `json:"createdAt"`
			UpdatedAt  interface{} `json:"updatedAt"`
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
		} `json:"currency"`
		LanguageId   string      `json:"languageId"`
		Language     interface{} `json:"language"`
		SalesChannel interface{} `json:"salesChannel"`
		Addresses    []struct {
			CountryId              string      `json:"countryId"`
			CountryStateId         string      `json:"countryStateId"`
			SalutationId           string      `json:"salutationId"`
			FirstName              string      `json:"firstName"`
			LastName               string      `json:"lastName"`
			Street                 string      `json:"street"`
			Zipcode                string      `json:"zipcode"`
			City                   string      `json:"city"`
			Company                string      `json:"company"`
			Department             string      `json:"department"`
			Title                  interface{} `json:"title"`
			VatId                  interface{} `json:"vatId"`
			PhoneNumber            string      `json:"phoneNumber"`
			AdditionalAddressLine1 interface{} `json:"additionalAddressLine1"`
			AdditionalAddressLine2 interface{} `json:"additionalAddressLine2"`
			Country                struct {
				Name                       string `json:"name"`
				Iso                        string `json:"iso"`
				Position                   int    `json:"position"`
				TaxFree                    bool   `json:"taxFree"`
				Active                     bool   `json:"active"`
				ShippingAvailable          bool   `json:"shippingAvailable"`
				Iso3                       string `json:"iso3"`
				DisplayStateInRegistration bool   `json:"displayStateInRegistration"`
				ForceStateInRegistration   bool   `json:"forceStateInRegistration"`
				CompanyTaxFree             bool   `json:"companyTaxFree"`
				CheckVatIdPattern          bool   `json:"checkVatIdPattern"`
				VatIdPattern               string `json:"vatIdPattern"`
				VatIdRequired              bool   `json:"vatIdRequired"`
				CustomerTax                struct {
					Enabled    bool          `json:"enabled"`
					CurrencyId string        `json:"currencyId"`
					Amount     int           `json:"amount"`
					Extensions []interface{} `json:"extensions"`
					ApiAlias   string        `json:"apiAlias"`
				} `json:"customerTax"`
				CompanyTax struct {
					Enabled    bool          `json:"enabled"`
					CurrencyId string        `json:"currencyId"`
					Amount     int           `json:"amount"`
					Extensions []interface{} `json:"extensions"`
					ApiAlias   string        `json:"apiAlias"`
				} `json:"companyTax"`
				States                         interface{} `json:"states"`
				Translations                   interface{} `json:"translations"`
				OrderAddresses                 interface{} `json:"orderAddresses"`
				CustomerAddresses              interface{} `json:"customerAddresses"`
				SalesChannelDefaultAssignments interface{} `json:"salesChannelDefaultAssignments"`
				SalesChannels                  interface{} `json:"salesChannels"`
				TaxRules                       interface{} `json:"taxRules"`
				CurrencyCountryRoundings       interface{} `json:"currencyCountryRoundings"`
				UniqueIdentifier               string      `json:"_uniqueIdentifier"`
				VersionId                      interface{} `json:"versionId"`
				Translated                     struct {
					Name         string      `json:"name"`
					CustomFields interface{} `json:"customFields"`
				} `json:"translated"`
				CreatedAt  time.Time   `json:"createdAt"`
				UpdatedAt  interface{} `json:"updatedAt"`
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
			} `json:"country"`
			CountryState struct {
				CountryId         string      `json:"countryId"`
				ShortCode         string      `json:"shortCode"`
				Name              string      `json:"name"`
				Position          int         `json:"position"`
				Active            bool        `json:"active"`
				Country           interface{} `json:"country"`
				Translations      interface{} `json:"translations"`
				CustomerAddresses interface{} `json:"customerAddresses"`
				OrderAddresses    interface{} `json:"orderAddresses"`
				UniqueIdentifier  string      `json:"_uniqueIdentifier"`
				VersionId         interface{} `json:"versionId"`
				Translated        struct {
					Name         string      `json:"name"`
					CustomFields interface{} `json:"customFields"`
				} `json:"translated"`
				CreatedAt  time.Time   `json:"createdAt"`
				UpdatedAt  interface{} `json:"updatedAt"`
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
			} `json:"countryState"`
			Order      interface{} `json:"order"`
			Salutation struct {
				SalutationKey        string      `json:"salutationKey"`
				DisplayName          string      `json:"displayName"`
				LetterName           string      `json:"letterName"`
				Translations         interface{} `json:"translations"`
				Customers            interface{} `json:"customers"`
				CustomerAddresses    interface{} `json:"customerAddresses"`
				OrderCustomers       interface{} `json:"orderCustomers"`
				OrderAddresses       interface{} `json:"orderAddresses"`
				NewsletterRecipients interface{} `json:"newsletterRecipients"`
				UniqueIdentifier     string      `json:"_uniqueIdentifier"`
				VersionId            interface{} `json:"versionId"`
				Translated           struct {
					DisplayName  string      `json:"displayName"`
					LetterName   string      `json:"letterName"`
					CustomFields interface{} `json:"customFields"`
				} `json:"translated"`
				CreatedAt  time.Time   `json:"createdAt"`
				UpdatedAt  interface{} `json:"updatedAt"`
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
			} `json:"salutation"`
			OrderDeliveries  interface{}   `json:"orderDeliveries"`
			OrderId          string        `json:"orderId"`
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
			Id             string      `json:"id"`
			CustomFields   interface{} `json:"customFields"`
			OrderVersionId string      `json:"orderVersionId"`
			ApiAlias       string      `json:"apiAlias"`
		} `json:"addresses"`
		BillingAddress interface{} `json:"billingAddress"`
		Deliveries     []struct {
			OrderId                string    `json:"orderId"`
			ShippingOrderAddressId string    `json:"shippingOrderAddressId"`
			ShippingMethodId       string    `json:"shippingMethodId"`
			TrackingCodes          []string  `json:"trackingCodes"`
			ShippingDateEarliest   time.Time `json:"shippingDateEarliest"`
			ShippingDateLatest     time.Time `json:"shippingDateLatest"`
			ShippingCosts          struct {
				UnitPrice       float64 `json:"unitPrice"`
				Quantity        int     `json:"quantity"`
				TotalPrice      float64 `json:"totalPrice"`
				CalculatedTaxes []struct {
					Tax        float64       `json:"tax"`
					TaxRate    float64       `json:"taxRate"`
					Price      float64       `json:"price"`
					Extensions []interface{} `json:"extensions"`
					ApiAlias   string        `json:"apiAlias"`
				} `json:"calculatedTaxes"`
				TaxRules []struct {
					TaxRate    float64       `json:"taxRate"`
					Percentage float64       `json:"percentage"`
					Extensions []interface{} `json:"extensions"`
					ApiAlias   string        `json:"apiAlias"`
				} `json:"taxRules"`
				ReferencePrice interface{}   `json:"referencePrice"`
				ListPrice      interface{}   `json:"listPrice"`
				Extensions     []interface{} `json:"extensions"`
				ApiAlias       string        `json:"apiAlias"`
			} `json:"shippingCosts"`
			ShippingOrderAddress struct {
				CountryId              string      `json:"countryId"`
				CountryStateId         string      `json:"countryStateId"`
				SalutationId           string      `json:"salutationId"`
				FirstName              string      `json:"firstName"`
				LastName               string      `json:"lastName"`
				Street                 string      `json:"street"`
				Zipcode                string      `json:"zipcode"`
				City                   string      `json:"city"`
				Company                interface{} `json:"company"`
				Department             interface{} `json:"department"`
				Title                  interface{} `json:"title"`
				VatId                  interface{} `json:"vatId"`
				PhoneNumber            interface{} `json:"phoneNumber"`
				AdditionalAddressLine1 interface{} `json:"additionalAddressLine1"`
				AdditionalAddressLine2 interface{} `json:"additionalAddressLine2"`
				Country                interface{} `json:"country"`
				CountryState           interface{} `json:"countryState"`
				Order                  interface{} `json:"order"`
				Salutation             struct {
					SalutationKey        string      `json:"salutationKey"`
					DisplayName          string      `json:"displayName"`
					LetterName           string      `json:"letterName"`
					Translations         interface{} `json:"translations"`
					Customers            interface{} `json:"customers"`
					CustomerAddresses    interface{} `json:"customerAddresses"`
					OrderCustomers       interface{} `json:"orderCustomers"`
					OrderAddresses       interface{} `json:"orderAddresses"`
					NewsletterRecipients interface{} `json:"newsletterRecipients"`
					UniqueIdentifier     string      `json:"_uniqueIdentifier"`
					VersionId            interface{} `json:"versionId"`
					Translated           struct {
						DisplayName  string      `json:"displayName"`
						LetterName   string      `json:"letterName"`
						CustomFields interface{} `json:"customFields"`
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
				} `json:"salutation"`
				OrderDeliveries  interface{}   `json:"orderDeliveries"`
				OrderId          string        `json:"orderId"`
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
					InternalMappingStorage struct {
						ApiAlias   interface{}   `json:"apiAlias"`
						Extensions []interface{} `json:"extensions"`
					} `json:"internal_mapping_storage"`
				} `json:"extensions"`
				Id             string      `json:"id"`
				CustomFields   interface{} `json:"customFields"`
				OrderVersionId string      `json:"orderVersionId"`
				ApiAlias       string      `json:"apiAlias"`
			} `json:"shippingOrderAddress"`
			StateId           string `json:"stateId"`
			StateMachineState struct {
				Name                           string      `json:"name"`
				TechnicalName                  string      `json:"technicalName"`
				StateMachineId                 string      `json:"stateMachineId"`
				StateMachine                   interface{} `json:"stateMachine"`
				FromStateMachineTransitions    interface{} `json:"fromStateMachineTransitions"`
				ToStateMachineTransitions      interface{} `json:"toStateMachineTransitions"`
				Translations                   interface{} `json:"translations"`
				Orders                         interface{} `json:"orders"`
				OrderTransactions              interface{} `json:"orderTransactions"`
				OrderDeliveries                interface{} `json:"orderDeliveries"`
				FromStateMachineHistoryEntries interface{} `json:"fromStateMachineHistoryEntries"`
				ToStateMachineHistoryEntries   interface{} `json:"toStateMachineHistoryEntries"`
				UniqueIdentifier               string      `json:"_uniqueIdentifier"`
				VersionId                      interface{} `json:"versionId"`
				Translated                     struct {
					Name         string      `json:"name"`
					CustomFields interface{} `json:"customFields"`
				} `json:"translated"`
				CreatedAt  time.Time   `json:"createdAt"`
				UpdatedAt  interface{} `json:"updatedAt"`
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
			} `json:"stateMachineState"`
			ShippingMethod struct {
				Name           string      `json:"name"`
				Active         bool        `json:"active"`
				Description    interface{} `json:"description"`
				TrackingUrl    interface{} `json:"trackingUrl"`
				DeliveryTimeId string      `json:"deliveryTimeId"`
				DeliveryTime   struct {
					Name             string      `json:"name"`
					Min              int         `json:"min"`
					Max              int         `json:"max"`
					Unit             string      `json:"unit"`
					ShippingMethods  interface{} `json:"shippingMethods"`
					Translations     interface{} `json:"translations"`
					Products         interface{} `json:"products"`
					UniqueIdentifier string      `json:"_uniqueIdentifier"`
					VersionId        interface{} `json:"versionId"`
					Translated       struct {
						Name         string      `json:"name"`
						CustomFields interface{} `json:"customFields"`
					} `json:"translated"`
					CreatedAt  time.Time   `json:"createdAt"`
					UpdatedAt  interface{} `json:"updatedAt"`
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
				} `json:"deliveryTime"`
				Translations                   interface{}   `json:"translations"`
				OrderDeliveries                interface{}   `json:"orderDeliveries"`
				SalesChannelDefaultAssignments interface{}   `json:"salesChannelDefaultAssignments"`
				SalesChannels                  interface{}   `json:"salesChannels"`
				AvailabilityRule               interface{}   `json:"availabilityRule"`
				AvailabilityRuleId             string        `json:"availabilityRuleId"`
				Prices                         []interface{} `json:"prices"`
				MediaId                        interface{}   `json:"mediaId"`
				TaxId                          interface{}   `json:"taxId"`
				Media                          interface{}   `json:"media"`
				Tags                           interface{}   `json:"tags"`
				TaxType                        string        `json:"taxType"`
				Tax                            interface{}   `json:"tax"`
				UniqueIdentifier               string        `json:"_uniqueIdentifier"`
				VersionId                      interface{}   `json:"versionId"`
				Translated                     struct {
					Name         string      `json:"name"`
					CustomFields interface{} `json:"customFields"`
					Description  interface{} `json:"description"`
					TrackingUrl  interface{} `json:"trackingUrl"`
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
			} `json:"shippingMethod"`
			Order            interface{}   `json:"order"`
			Positions        interface{}   `json:"positions"`
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
			Id                            string      `json:"id"`
			CustomFields                  interface{} `json:"customFields"`
			OrderVersionId                string      `json:"orderVersionId"`
			ShippingOrderAddressVersionId string      `json:"shippingOrderAddressVersionId"`
			ApiAlias                      string      `json:"apiAlias"`
		} `json:"deliveries"`
		LineItems []struct {
			OrderId      string  `json:"orderId"`
			Identifier   string  `json:"identifier"`
			ReferencedId string  `json:"referencedId"`
			ProductId    string  `json:"productId"`
			Quantity     int     `json:"quantity"`
			UnitPrice    float64 `json:"unitPrice"`
			TotalPrice   float64 `json:"totalPrice"`
			Label        string  `json:"label"`
			Description  string  `json:"description"`
			Good         bool    `json:"good"`
			Removable    bool    `json:"removable"`
			CoverId      string  `json:"coverId"`
			Stackable    bool    `json:"stackable"`
			Position     int     `json:"position"`
			Price        struct {
				UnitPrice       float64 `json:"unitPrice"`
				Quantity        int     `json:"quantity"`
				TotalPrice      float64 `json:"totalPrice"`
				CalculatedTaxes []struct {
					Tax        float64       `json:"tax"`
					TaxRate    float64       `json:"taxRate"`
					Price      float64       `json:"price"`
					Extensions []interface{} `json:"extensions"`
					ApiAlias   string        `json:"apiAlias"`
				} `json:"calculatedTaxes"`
				TaxRules []struct {
					TaxRate    float64       `json:"taxRate"`
					Percentage float64       `json:"percentage"`
					Extensions []interface{} `json:"extensions"`
					ApiAlias   string        `json:"apiAlias"`
				} `json:"taxRules"`
				ReferencePrice interface{}   `json:"referencePrice"`
				ListPrice      interface{}   `json:"listPrice"`
				Extensions     []interface{} `json:"extensions"`
				ApiAlias       string        `json:"apiAlias"`
			} `json:"price"`
			PriceDefinition struct {
				Price    float64 `json:"price"`
				TaxRules []struct {
					TaxRate    float64       `json:"taxRate"`
					Percentage float64       `json:"percentage"`
					Extensions []interface{} `json:"extensions"`
					ApiAlias   string        `json:"apiAlias"`
				} `json:"taxRules"`
				Quantity                 int           `json:"quantity"`
				IsCalculated             bool          `json:"isCalculated"`
				ReferencePriceDefinition interface{}   `json:"referencePriceDefinition"`
				ListPrice                interface{}   `json:"listPrice"`
				Extensions               []interface{} `json:"extensions"`
				Type                     string        `json:"type"`
				ApiAlias                 string        `json:"apiAlias"`
			} `json:"priceDefinition"`
			Payload                json.RawMessage `json:"payload"`
			ParentId               interface{}     `json:"parentId"`
			Parent                 interface{}     `json:"parent"`
			Type                   string          `json:"type"`
			Order                  interface{}     `json:"order"`
			OrderDeliveryPositions interface{}     `json:"orderDeliveryPositions"`
			Cover                  interface{}     `json:"cover"`
			Children               interface{}     `json:"children"`
			Product                struct {
				ProductNumber string `json:"productNumber"`
				ApiAlias      string `json:"apiAlias"`
			} `json:"product"`
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
			Id               string                 `json:"id"`
			CustomFields     map[string]interface{} `json:"customFields"`
			OrderVersionId   string                 `json:"orderVersionId"`
			ProductVersionId string                 `json:"productVersionId"`
			ParentVersionId  string                 `json:"parentVersionId"`
			ApiAlias         string                 `json:"apiAlias"`
		} `json:"lineItems"`
		Transactions []struct {
			OrderId         string `json:"orderId"`
			PaymentMethodId string `json:"paymentMethodId"`
			Amount          struct {
				UnitPrice       float64 `json:"unitPrice"`
				Quantity        int     `json:"quantity"`
				TotalPrice      float64 `json:"totalPrice"`
				CalculatedTaxes []struct {
					Tax        float64       `json:"tax"`
					TaxRate    float64       `json:"taxRate"`
					Price      float64       `json:"price"`
					Extensions []interface{} `json:"extensions"`
					ApiAlias   string        `json:"apiAlias"`
				} `json:"calculatedTaxes"`
				TaxRules []struct {
					TaxRate    float64       `json:"taxRate"`
					Percentage float64       `json:"percentage"`
					Extensions []interface{} `json:"extensions"`
					ApiAlias   string        `json:"apiAlias"`
				} `json:"taxRules"`
				ReferencePrice interface{}   `json:"referencePrice"`
				ListPrice      interface{}   `json:"listPrice"`
				Extensions     []interface{} `json:"extensions"`
				ApiAlias       string        `json:"apiAlias"`
			} `json:"amount"`
			PaymentMethod struct {
				PluginId                       interface{} `json:"pluginId"`
				Name                           string      `json:"name"`
				DistinguishableName            string      `json:"distinguishableName"`
				Description                    string      `json:"description"`
				Position                       int         `json:"position"`
				Active                         bool        `json:"active"`
				AfterOrderEnabled              bool        `json:"afterOrderEnabled"`
				Plugin                         interface{} `json:"plugin"`
				Translations                   interface{} `json:"translations"`
				OrderTransactions              interface{} `json:"orderTransactions"`
				Customers                      interface{} `json:"customers"`
				SalesChannelDefaultAssignments interface{} `json:"salesChannelDefaultAssignments"`
				SalesChannels                  interface{} `json:"salesChannels"`
				AvailabilityRule               interface{} `json:"availabilityRule"`
				AvailabilityRuleId             interface{} `json:"availabilityRuleId"`
				MediaId                        interface{} `json:"mediaId"`
				Media                          interface{} `json:"media"`
				FormattedHandlerIdentifier     string      `json:"formattedHandlerIdentifier"`
				ShortName                      string      `json:"shortName"`
				AppPaymentMethod               interface{} `json:"appPaymentMethod"`
				UniqueIdentifier               string      `json:"_uniqueIdentifier"`
				VersionId                      interface{} `json:"versionId"`
				Translated                     struct {
					Name                string      `json:"name"`
					DistinguishableName string      `json:"distinguishableName"`
					Description         string      `json:"description"`
					CustomFields        interface{} `json:"customFields"`
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
			} `json:"paymentMethod"`
			Order             interface{} `json:"order"`
			StateMachineState struct {
				Name                           string      `json:"name"`
				TechnicalName                  string      `json:"technicalName"`
				StateMachineId                 string      `json:"stateMachineId"`
				StateMachine                   interface{} `json:"stateMachine"`
				FromStateMachineTransitions    interface{} `json:"fromStateMachineTransitions"`
				ToStateMachineTransitions      interface{} `json:"toStateMachineTransitions"`
				Translations                   interface{} `json:"translations"`
				Orders                         interface{} `json:"orders"`
				OrderTransactions              interface{} `json:"orderTransactions"`
				OrderDeliveries                interface{} `json:"orderDeliveries"`
				FromStateMachineHistoryEntries interface{} `json:"fromStateMachineHistoryEntries"`
				ToStateMachineHistoryEntries   interface{} `json:"toStateMachineHistoryEntries"`
				UniqueIdentifier               string      `json:"_uniqueIdentifier"`
				VersionId                      interface{} `json:"versionId"`
				Translated                     struct {
					Name         string      `json:"name"`
					CustomFields interface{} `json:"customFields"`
				} `json:"translated"`
				CreatedAt  time.Time   `json:"createdAt"`
				UpdatedAt  interface{} `json:"updatedAt"`
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
			} `json:"stateMachineState"`
			StateId          string        `json:"stateId"`
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
			Id             string                 `json:"id"`
			CustomFields   map[string]interface{} `json:"customFields"`
			OrderVersionId string                 `json:"orderVersionId"`
			ApiAlias       string                 `json:"apiAlias"`
		} `json:"transactions"`
		DeepLinkCode      string `json:"deepLinkCode"`
		AutoIncrement     int    `json:"autoIncrement"`
		StateMachineState struct {
			Name                           string      `json:"name"`
			TechnicalName                  string      `json:"technicalName"`
			StateMachineId                 string      `json:"stateMachineId"`
			StateMachine                   interface{} `json:"stateMachine"`
			FromStateMachineTransitions    interface{} `json:"fromStateMachineTransitions"`
			ToStateMachineTransitions      interface{} `json:"toStateMachineTransitions"`
			Translations                   interface{} `json:"translations"`
			Orders                         interface{} `json:"orders"`
			OrderTransactions              interface{} `json:"orderTransactions"`
			OrderDeliveries                interface{} `json:"orderDeliveries"`
			FromStateMachineHistoryEntries interface{} `json:"fromStateMachineHistoryEntries"`
			ToStateMachineHistoryEntries   interface{} `json:"toStateMachineHistoryEntries"`
			UniqueIdentifier               string      `json:"_uniqueIdentifier"`
			VersionId                      interface{} `json:"versionId"`
			Translated                     struct {
				Name         string      `json:"name"`
				CustomFields interface{} `json:"customFields"`
			} `json:"translated"`
			CreatedAt  time.Time   `json:"createdAt"`
			UpdatedAt  interface{} `json:"updatedAt"`
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
		} `json:"stateMachineState"`
		StateId         string      `json:"stateId"`
		Documents       interface{} `json:"documents"`
		Tags            interface{} `json:"tags"`
		AffiliateCode   interface{} `json:"affiliateCode"`
		CampaignCode    interface{} `json:"campaignCode"`
		CustomerComment string      `json:"customerComment"`
		RuleIds         []string    `json:"ruleIds"`
		CreatedById     string      `json:"createdById"`
		CreatedBy       interface{} `json:"createdBy"`
		UpdatedById     string      `json:"updatedById"`
		UpdatedBy       interface{} `json:"updatedBy"`
		ItemRounding    struct {
			Decimals    int           `json:"decimals"`
			Interval    float64       `json:"interval"`
			RoundForNet bool          `json:"roundForNet"`
			Extensions  []interface{} `json:"extensions"`
			ApiAlias    string        `json:"apiAlias"`
		} `json:"itemRounding"`
		TotalRounding struct {
			Decimals    int           `json:"decimals"`
			Interval    float64       `json:"interval"`
			RoundForNet bool          `json:"roundForNet"`
			Extensions  []interface{} `json:"extensions"`
			ApiAlias    string        `json:"apiAlias"`
		} `json:"totalRounding"`
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
		Id                      string      `json:"id"`
		CustomFields            interface{} `json:"customFields"`
		BillingAddressVersionId string      `json:"billingAddressVersionId"`
		ApiAlias                string      `json:"apiAlias"`
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

// CustomOrderTransactionsSearchBody is to define the filter for the request
type CustomOrderTransactionsSearchBody struct {
	Ids          []string `json:"ids"`
	Associations struct {
		Transactions struct {
			Associations struct {
				PaymentMethod struct {
				} `json:"paymentMethod"`
			} `json:"associations"`
		} `json:"transactions"`
	} `json:"associations"`
	Includes CustomOrderTransactionsSearchBodyIncludes `json:"includes"`
}

type CustomOrderTransactionsSearchBodyIncludes struct {
	Order             []string `json:"order"`
	OrderTransaction  []string `json:"order_transaction"`
	StateMachineState []string `json:"state_machine_state"`
	PaymentMethod     []string `json:"payment_method"`
}

// CustomOrderTransactionsSearchReturn is to decode the json data
type CustomOrderTransactionsSearchReturn struct {
	Total int `json:"total"`
	Data  []struct {
		AmountTotal  float64 `json:"amountTotal"`
		Transactions []struct {
			PaymentMethod struct {
				UpdatedAt time.Time `json:"updatedAt"`
				Name      string    `json:"name"`
				ApiAlias  string    `json:"apiAlias"`
			} `json:"paymentMethod"`
			StateMachineState struct {
				TechnicalName string `json:"technicalName"`
				ApiAlias      string `json:"apiAlias"`
			} `json:"stateMachineState"`
			CustomFields map[string]interface{} `json:"customFields"`
			ApiAlias     string                 `json:"apiAlias"`
		} `json:"transactions"`
		CustomFields map[string]interface{} `json:"customFields"`
		ApiAlias     string                 `json:"apiAlias"`
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

// CustomOrderSearch is to search orders from date x with currency, lineItems, addresses & transactions
func CustomOrderSearch(page, limit int, date string, r Request) (CustomOrderSearchReturn, error) {

	// Define body
	body := CustomOrderSearchBody{
		Page:   page,
		Limit:  limit,
		Filter: []CustomOrderSearchBodyFilter{},
		Sort:   []CustomOrderSearchBodySort{},
	}

	body.Filter = append(body.Filter, CustomOrderSearchBodyFilter{
		Type:  "range",
		Field: "orderDate",
		Parameters: CustomOrderSearchBodyParameters{
			Gte: date + "T00:00:00.000+00:00",
		},
	})

	body.Sort = append(body.Sort, CustomOrderSearchBodySort{
		Field: "orderNumber",
		Order: "ASC",
	})

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return CustomOrderSearchReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/search/order",
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CustomOrderSearchReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CustomOrderSearchReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return CustomOrderSearchReturn{}, err
		}
	}

	// Return data
	return decode, err

}

// CustomOrderTransactionsSearch is to search an order for transactions and payment method
func CustomOrderTransactionsSearch(id string, r Request) (CustomOrderTransactionsSearchReturn, error) {

	// Define body
	body := CustomOrderTransactionsSearchBody{
		Ids: []string{id},
		Includes: CustomOrderTransactionsSearchBodyIncludes{
			Order: []string{
				"amountTotal",
				"transactions",
				"customFields",
			},
			OrderTransaction: []string{
				"paymentMethod",
				"stateMachineState",
				"customFields",
			},
			StateMachineState: []string{
				"technicalName",
			},
			PaymentMethod: []string{
				"name",
				"updatedAt",
			},
		},
	}

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return CustomOrderTransactionsSearchReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/search/order",
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return CustomOrderTransactionsSearchReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode CustomOrderTransactionsSearchReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return CustomOrderTransactionsSearchReturn{}, err
		}
	}

	// Return data
	return decode, err

}
