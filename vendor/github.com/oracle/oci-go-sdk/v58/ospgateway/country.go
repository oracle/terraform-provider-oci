// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSP Gateway API
//
// This site describes all the Rest endpoints of OSP Gateway.
//

package ospgateway

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Country Country details model
type Country struct {

	// Indentifier of the country. This is a DB side unique id which was generated when the entity was created in the table
	CountryId *float32 `mandatory:"false" json:"countryId"`

	// Country code in ISO-3166-1 2-letter format
	CountryCode *string `mandatory:"false" json:"countryCode"`

	// Name of the country
	CountryName *string `mandatory:"false" json:"countryName"`

	// Language identifier
	LanguageId *float32 `mandatory:"false" json:"languageId"`

	// Country code in ISO-3166-1 3-letter format
	Ascii3CountryCode *string `mandatory:"false" json:"ascii3CountryCode"`
}

func (m Country) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Country) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
