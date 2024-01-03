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

// Address Address details model.
type Address struct {

	// Address identifier.
	AddressKey *string `mandatory:"false" json:"addressKey"`

	// Address line 1.
	Line1 *string `mandatory:"false" json:"line1"`

	// Address line 2.
	Line2 *string `mandatory:"false" json:"line2"`

	// Address line 3.
	Line3 *string `mandatory:"false" json:"line3"`

	// Address line 4.
	Line4 *string `mandatory:"false" json:"line4"`

	// Street name of the address.
	StreetName *string `mandatory:"false" json:"streetName"`

	// Street number of the address.
	StreetNumber *string `mandatory:"false" json:"streetNumber"`

	// Name of the city.
	City *string `mandatory:"false" json:"city"`

	// County of the address.
	County *string `mandatory:"false" json:"county"`

	// Country of the address.
	Country *string `mandatory:"false" json:"country"`

	// Province of the address.
	Province *string `mandatory:"false" json:"province"`

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

	// Middle name of the contact person.
	MiddleName *string `mandatory:"false" json:"middleName"`

	// Last name of the contact person.
	LastName *string `mandatory:"false" json:"lastName"`

	// Phone country code of the contact person.
	PhoneCountryCode *string `mandatory:"false" json:"phoneCountryCode"`

	// Phone number of the contact person.
	PhoneNumber *string `mandatory:"false" json:"phoneNumber"`

	// Job title of the contact person.
	JobTitle *string `mandatory:"false" json:"jobTitle"`

	// Department name of the customer company.
	DepartmentName *string `mandatory:"false" json:"departmentName"`

	// Internal number of the customer company.
	InternalNumber *string `mandatory:"false" json:"internalNumber"`

	// Contributor class of the customer company.
	ContributorClass *string `mandatory:"false" json:"contributorClass"`

	// State Inscription.
	StateInscription *string `mandatory:"false" json:"stateInscription"`

	// Municipal Inscription.
	MunicipalInscription *string `mandatory:"false" json:"municipalInscription"`
}

func (m Address) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Address) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
