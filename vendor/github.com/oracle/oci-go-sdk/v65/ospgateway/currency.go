// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OSP Gateway API
//
// This site describes all the Rest endpoints of OSP Gateway.
//

package ospgateway

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Currency Currency details model
type Currency struct {

	// Currency code
	CurrencyCode *string `mandatory:"false" json:"currencyCode"`

	// Currency symbol
	CurrencySymbol *string `mandatory:"false" json:"currencySymbol"`

	// Name of the currency
	Name *string `mandatory:"false" json:"name"`

	// USD conversion rate of the currency
	UsdConversion *float32 `mandatory:"false" json:"usdConversion"`

	// Round decimal point
	RoundDecimalPoint *float32 `mandatory:"false" json:"roundDecimalPoint"`
}

func (m Currency) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Currency) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
