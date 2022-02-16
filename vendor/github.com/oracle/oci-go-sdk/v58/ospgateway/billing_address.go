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

// BillingAddress Billing address details model.
type BillingAddress struct {

	// Address identifier.
	AddressKey *string `mandatory:"false" json:"addressKey"`

	// Address line 1.
	Line1 *string `mandatory:"false" json:"line1"`

	// Address line 2.
	Line2 *string `mandatory:"false" json:"line2"`

	// Name of the city.
	City *string `mandatory:"false" json:"city"`

	// Country of the address.
	Country *string `mandatory:"false" json:"country"`

	// Post code of the address.
	PostalCode *string `mandatory:"false" json:"postalCode"`

	// State of the address.
	State *string `mandatory:"false" json:"state"`

	// Contact person email address.
	EmailAddress *string `mandatory:"false" json:"emailAddress"`

	// Name of the customer company.
	CompanyName *string `mandatory:"false" json:"companyName"`

	// First name of the contact person.
	FirstName *string `mandatory:"false" json:"firstName"`

	// Last name of the contact person.
	LastName *string `mandatory:"false" json:"lastName"`
}

func (m BillingAddress) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BillingAddress) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
