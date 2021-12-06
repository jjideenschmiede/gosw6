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

import "encoding/json"

// AccessTokenBody is to structure the data
type AccessTokenBody struct {
	ClientId  string `json:"client_id"`
	GrantType string `json:"grant_type"`
	Scopes    string `json:"scopes"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

// AccessTokenReturn is to decode the json return
type AccessTokenReturn struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// AccessToken is to get a bearer token from the system
func AccessToken(body AccessTokenBody, r Request) (AccessTokenReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return AccessTokenReturn{}, err
	}

	// Set config for request
	c := Config{
		Path:        "/api/oauth/token",
		Method:      "POST",
		Body:        convert,
		AccessToken: false,
	}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return AccessTokenReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode AccessTokenReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return AccessTokenReturn{}, err
	}

	// Return data
	return decode, err

}
