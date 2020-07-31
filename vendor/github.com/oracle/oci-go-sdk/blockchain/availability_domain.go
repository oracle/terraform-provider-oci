// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AvailabilityDomain Availability Domains
type AvailabilityDomain struct {

	// Availability Domain Identifiers
	Ads AvailabilityDomainAdsEnum `mandatory:"false" json:"ads,omitempty"`
}

func (m AvailabilityDomain) String() string {
	return common.PointerString(m)
}

// AvailabilityDomainAdsEnum Enum with underlying type: string
type AvailabilityDomainAdsEnum string

// Set of constants representing the allowable values for AvailabilityDomainAdsEnum
const (
	AvailabilityDomainAdsAd1 AvailabilityDomainAdsEnum = "AD1"
	AvailabilityDomainAdsAd2 AvailabilityDomainAdsEnum = "AD2"
	AvailabilityDomainAdsAd3 AvailabilityDomainAdsEnum = "AD3"
)

var mappingAvailabilityDomainAds = map[string]AvailabilityDomainAdsEnum{
	"AD1": AvailabilityDomainAdsAd1,
	"AD2": AvailabilityDomainAdsAd2,
	"AD3": AvailabilityDomainAdsAd3,
}

// GetAvailabilityDomainAdsEnumValues Enumerates the set of values for AvailabilityDomainAdsEnum
func GetAvailabilityDomainAdsEnumValues() []AvailabilityDomainAdsEnum {
	values := make([]AvailabilityDomainAdsEnum, 0)
	for _, v := range mappingAvailabilityDomainAds {
		values = append(values, v)
	}
	return values
}
