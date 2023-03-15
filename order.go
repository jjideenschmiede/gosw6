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

var (
	orderTransactionState = map[int]string{
		0:  "authorize",
		1:  "paid",
		2:  "remind",
		3:  "refund",
		4:  "fail",
		5:  "do_pay",
		6:  "reopen",
		7:  "chargeback",
		8:  "paid_partially",
		9:  "refund_partially",
		10: "process_unconfirmed",
		11: "cancel",
	}
	orderDelieryState = map[int]string{
		0: "reopen",
		1: "retour",
		2: "retour_partially",
		3: "ship",
		4: "ship_partially",
		5: "cancel",
	}
	orderState = map[int]string{
		0: "reopen",
		1: "process",
		2: "complete",
		3: "cancel",
	}
)

// OrdersReturn is to decode the json data
type OrdersReturn struct {
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
				TaxRate    int           `json:"taxRate"`
				Price      float64       `json:"price"`
				Extensions []interface{} `json:"extensions"`
				ApiAlias   string        `json:"apiAlias"`
			} `json:"calculatedTaxes"`
			TaxRules []struct {
				TaxRate    int           `json:"taxRate"`
				Percentage int           `json:"percentage"`
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
			UnitPrice       int `json:"unitPrice"`
			Quantity        int `json:"quantity"`
			TotalPrice      int `json:"totalPrice"`
			CalculatedTaxes []struct {
				Tax        int           `json:"tax"`
				TaxRate    int           `json:"taxRate"`
				Price      int           `json:"price"`
				Extensions []interface{} `json:"extensions"`
				ApiAlias   string        `json:"apiAlias"`
			} `json:"calculatedTaxes"`
			TaxRules []struct {
				TaxRate    int           `json:"taxRate"`
				Percentage int           `json:"percentage"`
				Extensions []interface{} `json:"extensions"`
				ApiAlias   string        `json:"apiAlias"`
			} `json:"taxRules"`
			ReferencePrice interface{}   `json:"referencePrice"`
			ListPrice      interface{}   `json:"listPrice"`
			Extensions     []interface{} `json:"extensions"`
			ApiAlias       string        `json:"apiAlias"`
		} `json:"shippingCosts"`
		ShippingTotal int `json:"shippingTotal"`
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
		Currency          interface{} `json:"currency"`
		LanguageId        string      `json:"languageId"`
		Language          interface{} `json:"language"`
		SalesChannel      interface{} `json:"salesChannel"`
		Addresses         interface{} `json:"addresses"`
		BillingAddress    interface{} `json:"billingAddress"`
		Deliveries        interface{} `json:"deliveries"`
		LineItems         interface{} `json:"lineItems"`
		Transactions      interface{} `json:"transactions"`
		DeepLinkCode      string      `json:"deepLinkCode"`
		AutoIncrement     int         `json:"autoIncrement"`
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
				Name         string        `json:"name"`
				CustomFields []interface{} `json:"customFields"`
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
		CustomerComment interface{} `json:"customerComment"`
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

// OrderReturn is to decode the json data
type OrderReturn struct {
	Data struct {
		OrderNumber      string    `json:"orderNumber"`
		CurrencyId       string    `json:"currencyId"`
		CurrencyFactor   float64   `json:"currencyFactor"`
		SalesChannelId   string    `json:"salesChannelId"`
		BillingAddressId string    `json:"billingAddressId"`
		OrderDateTime    time.Time `json:"orderDateTime"`
		OrderDate        time.Time `json:"orderDate"`
		Price            struct {
			NetPrice        float64 `json:"netPrice"`
			TotalPrice      int     `json:"totalPrice"`
			CalculatedTaxes []struct {
				Tax        float64       `json:"tax"`
				TaxRate    int           `json:"taxRate"`
				Price      int           `json:"price"`
				Extensions []interface{} `json:"extensions"`
				ApiAlias   string        `json:"apiAlias"`
			} `json:"calculatedTaxes"`
			TaxRules []struct {
				TaxRate    int           `json:"taxRate"`
				Percentage float64       `json:"percentage"`
				Extensions []interface{} `json:"extensions"`
				ApiAlias   string        `json:"apiAlias"`
			} `json:"taxRules"`
			PositionPrice int           `json:"positionPrice"`
			TaxStatus     string        `json:"taxStatus"`
			RawTotal      int           `json:"rawTotal"`
			Extensions    []interface{} `json:"extensions"`
			ApiAlias      string        `json:"apiAlias"`
		} `json:"price"`
		AmountTotal   int     `json:"amountTotal"`
		AmountNet     float64 `json:"amountNet"`
		PositionPrice int     `json:"positionPrice"`
		TaxStatus     string  `json:"taxStatus"`
		ShippingCosts struct {
			UnitPrice       int `json:"unitPrice"`
			Quantity        int `json:"quantity"`
			TotalPrice      int `json:"totalPrice"`
			CalculatedTaxes []struct {
				Tax        int           `json:"tax"`
				TaxRate    int           `json:"taxRate"`
				Price      int           `json:"price"`
				Extensions []interface{} `json:"extensions"`
				ApiAlias   string        `json:"apiAlias"`
			} `json:"calculatedTaxes"`
			TaxRules []struct {
				TaxRate    int           `json:"taxRate"`
				Percentage float64       `json:"percentage"`
				Extensions []interface{} `json:"extensions"`
				ApiAlias   string        `json:"apiAlias"`
			} `json:"taxRules"`
			ReferencePrice interface{}   `json:"referencePrice"`
			ListPrice      interface{}   `json:"listPrice"`
			Extensions     []interface{} `json:"extensions"`
			ApiAlias       string        `json:"apiAlias"`
		} `json:"shippingCosts"`
		ShippingTotal int `json:"shippingTotal"`
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
			RemoteAddress    interface{}   `json:"remoteAddress"`
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
		Currency          interface{} `json:"currency"`
		LanguageId        string      `json:"languageId"`
		Language          interface{} `json:"language"`
		SalesChannel      interface{} `json:"salesChannel"`
		Addresses         interface{} `json:"addresses"`
		BillingAddress    interface{} `json:"billingAddress"`
		Deliveries        interface{} `json:"deliveries"`
		LineItems         interface{} `json:"lineItems"`
		Transactions      interface{} `json:"transactions"`
		DeepLinkCode      string      `json:"deepLinkCode"`
		AutoIncrement     int         `json:"autoIncrement"`
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
				Name         string        `json:"name"`
				CustomFields []interface{} `json:"customFields"`
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
		CustomerComment interface{} `json:"customerComment"`
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

// OrderTransactionStateBody is to structure the body data
type OrderTransactionStateBody struct {
	DocumentIds []interface{} `json:"documentIds"`
	SendMail    bool          `json:"sendMail"`
}

// OrderTransactionStateReturn is to decode the json data
type OrderTransactionStateReturn struct {
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
		Name         string        `json:"name"`
		CustomFields []interface{} `json:"customFields"`
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

// OrderDeliveryStateBody is to structure the body data
type OrderDeliveryStateBody struct {
	DocumentIds []interface{} `json:"documentIds"`
	SendMail    bool          `json:"sendMail"`
}

// OrderDeliveryStateReturn is to decode the json data
type OrderDeliveryStateReturn struct {
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
		Name         string        `json:"name"`
		CustomFields []interface{} `json:"customFields"`
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

// OrderStateBody is to structure the body data
type OrderStateBody struct {
	DocumentIds []interface{} `json:"documentIds"`
	SendMail    bool          `json:"sendMail"`
}

// OrderStateReturn is to decode the json data
type OrderStateReturn struct {
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
		Name         string        `json:"name"`
		CustomFields []interface{} `json:"customFields"`
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

// UpdateOrderTrackingNumberBody is to structure the body data
type UpdateOrderTrackingNumberBody struct {
	Deliveries []UpdateOrderTrackingNumberBodyDeliveries `json:"deliveries"`
}

type UpdateOrderTrackingNumberBodyDeliveries struct {
	Id            string   `json:"id"`
	TrackingCodes []string `json:"trackingCodes"`
}

// UpdateOrderTrackingNumberReturn is to decode the json data
type UpdateOrderTrackingNumberReturn struct {
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

// Orders are to get a list of all orders
func Orders(parameter map[string]string, r Request) (OrdersReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/order",
		Method: "GET",
		Body:   nil,
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return OrdersReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return OrdersReturn{}, err
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
		return OrdersReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode OrdersReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return OrdersReturn{}, err
	}

	// Return data
	return decode, err

}

// Order are to get a order by id
func Order(id string, r Request) (OrderReturn, error) {

	// Set config for request
	c := Config{
		Path:   "/api/order/" + id,
		Method: "GET",
		Body:   nil,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return OrderReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode OrderReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return OrderReturn{}, err
	}

	// Return data
	return decode, err

}

// OrderTransactionState are to change a transaction state of an order
func OrderTransactionState(id string, state int, body OrderTransactionStateBody, r Request) (OrderTransactionStateReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return OrderTransactionStateReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/_action/order_transaction/" + id + "/state/" + orderTransactionState[state],
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return OrderTransactionStateReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode OrderTransactionStateReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return OrderTransactionStateReturn{}, err
	}

	// Return data
	return decode, err

}

// OrderDeliveryState are to change a delivery state of an order
func OrderDeliveryState(id string, state int, body OrderDeliveryStateBody, r Request) (OrderDeliveryStateReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return OrderDeliveryStateReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/_action/order_delivery/" + id + "/state/" + orderDelieryState[state],
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return OrderDeliveryStateReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode OrderDeliveryStateReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return OrderDeliveryStateReturn{}, err
	}

	// Return data
	return decode, err

}

// OrderState are to change a state of an order
func OrderState(id string, state int, body OrderStateBody, r Request) (OrderStateReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return OrderStateReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/_action/order/" + id + "/state/" + orderState[state],
		Method: "POST",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return OrderStateReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode OrderStateReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return OrderStateReturn{}, err
	}

	// Return data
	return decode, err

}

// UpdateOrderTrackingNumber is to update the tracking number of an order
func UpdateOrderTrackingNumber(id string, body UpdateOrderTrackingNumberBody, r Request) (UpdateOrderTrackingNumberReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return UpdateOrderTrackingNumberReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:   "/api/order/" + id,
		Method: "PATCH",
		Body:   convert,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return UpdateOrderTrackingNumberReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode UpdateOrderTrackingNumberReturn

	// Check response header
	if response.Status != "204 No Content" {
		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			return UpdateOrderTrackingNumberReturn{}, err
		}
	}

	// Return data
	return decode, err

}
