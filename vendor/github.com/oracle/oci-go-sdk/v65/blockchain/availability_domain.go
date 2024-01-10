// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AvailabilityDomain Availability Domains
type AvailabilityDomain struct {

	// Availability Domain Identifiers
	Ads AvailabilityDomainAdsEnum `mandatory:"false" json:"ads,omitempty"`
}

func (m AvailabilityDomain) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AvailabilityDomain) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAvailabilityDomainAdsEnum(string(m.Ads)); !ok && m.Ads != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Ads: %s. Supported values are: %s.", m.Ads, strings.Join(GetAvailabilityDomainAdsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AvailabilityDomainAdsEnum Enum with underlying type: string
type AvailabilityDomainAdsEnum string

// Set of constants representing the allowable values for AvailabilityDomainAdsEnum
const (
	AvailabilityDomainAdsAd1 AvailabilityDomainAdsEnum = "AD1"
	AvailabilityDomainAdsAd2 AvailabilityDomainAdsEnum = "AD2"
	AvailabilityDomainAdsAd3 AvailabilityDomainAdsEnum = "AD3"
)

var mappingAvailabilityDomainAdsEnum = map[string]AvailabilityDomainAdsEnum{
	"AD1": AvailabilityDomainAdsAd1,
	"AD2": AvailabilityDomainAdsAd2,
	"AD3": AvailabilityDomainAdsAd3,
}

var mappingAvailabilityDomainAdsEnumLowerCase = map[string]AvailabilityDomainAdsEnum{
	"ad1": AvailabilityDomainAdsAd1,
	"ad2": AvailabilityDomainAdsAd2,
	"ad3": AvailabilityDomainAdsAd3,
}

// GetAvailabilityDomainAdsEnumValues Enumerates the set of values for AvailabilityDomainAdsEnum
func GetAvailabilityDomainAdsEnumValues() []AvailabilityDomainAdsEnum {
	values := make([]AvailabilityDomainAdsEnum, 0)
	for _, v := range mappingAvailabilityDomainAdsEnum {
		values = append(values, v)
	}
	return values
}

// GetAvailabilityDomainAdsEnumStringValues Enumerates the set of values in String for AvailabilityDomainAdsEnum
func GetAvailabilityDomainAdsEnumStringValues() []string {
	return []string{
		"AD1",
		"AD2",
		"AD3",
	}
}

// GetMappingAvailabilityDomainAdsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAvailabilityDomainAdsEnum(val string) (AvailabilityDomainAdsEnum, bool) {
	enum, ok := mappingAvailabilityDomainAdsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
