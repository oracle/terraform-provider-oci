// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
// Use this pattern to construct endpoints for identity domains: `https://<domainURL>/admin/v1/`. See Finding an Identity Domain URL (https://docs.oracle.com/en-us/iaas/Content/Identity/api-getstarted/locate-identity-domain-url.htm) to locate the domain URL you need.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Addresses A physical mailing address for this User, as described in (address Element). Canonical Type Values of work, home, and other. The value attribute is a complex type with the following sub-attributes.
type Addresses struct {

	// A label indicating the attribute's function; e.g., 'work' or 'home'.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Type AddressesTypeEnum `mandatory:"true" json:"type"`

	// The full mailing address, formatted for display or use with a mailing label. This attribute MAY contain newlines.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Formatted *string `mandatory:"false" json:"formatted"`

	// The full street address component, which may include house number, street name, PO BOX, and multi-line extended street address information. This attribute MAY contain newlines.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	StreetAddress *string `mandatory:"false" json:"streetAddress"`

	// The city or locality component.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Locality *string `mandatory:"false" json:"locality"`

	// The state or region component.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Region *string `mandatory:"false" json:"region"`

	// The zipcode or postal code component.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	PostalCode *string `mandatory:"false" json:"postalCode"`

	// The country name component.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsCanonicalValueSourceFilter: attrName eq "countries" and attrValues.value eq "upper($(country))"
	//  - idcsCanonicalValueSourceResourceType: AllowedValue
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Country *string `mandatory:"false" json:"country"`

	// A Boolean value indicating the 'primary' or preferred attribute value for this attribute. The primary attribute value 'true' MUST appear no more than once.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	Primary *bool `mandatory:"false" json:"primary"`
}

func (m Addresses) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Addresses) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddressesTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAddressesTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AddressesTypeEnum Enum with underlying type: string
type AddressesTypeEnum string

// Set of constants representing the allowable values for AddressesTypeEnum
const (
	AddressesTypeWork  AddressesTypeEnum = "work"
	AddressesTypeHome  AddressesTypeEnum = "home"
	AddressesTypeOther AddressesTypeEnum = "other"
)

var mappingAddressesTypeEnum = map[string]AddressesTypeEnum{
	"work":  AddressesTypeWork,
	"home":  AddressesTypeHome,
	"other": AddressesTypeOther,
}

var mappingAddressesTypeEnumLowerCase = map[string]AddressesTypeEnum{
	"work":  AddressesTypeWork,
	"home":  AddressesTypeHome,
	"other": AddressesTypeOther,
}

// GetAddressesTypeEnumValues Enumerates the set of values for AddressesTypeEnum
func GetAddressesTypeEnumValues() []AddressesTypeEnum {
	values := make([]AddressesTypeEnum, 0)
	for _, v := range mappingAddressesTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAddressesTypeEnumStringValues Enumerates the set of values in String for AddressesTypeEnum
func GetAddressesTypeEnumStringValues() []string {
	return []string{
		"work",
		"home",
		"other",
	}
}

// GetMappingAddressesTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddressesTypeEnum(val string) (AddressesTypeEnum, bool) {
	enum, ok := mappingAddressesTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
