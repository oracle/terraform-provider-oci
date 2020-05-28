// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ListingSummary The model for a summary of an Oracle Cloud Infrastructure Marketplace listing.
type ListingSummary struct {

	// The unique identifier for the listing in Marketplace.
	Id *string `mandatory:"false" json:"id"`

	// The name of the listing.
	Name *string `mandatory:"false" json:"name"`

	// A short description of the listing.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	// The tagline of the listing.
	Tagline *string `mandatory:"false" json:"tagline"`

	Icon *UploadData `mandatory:"false" json:"icon"`

	// The listing's package type.
	PackageType PackageTypeEnumEnum `mandatory:"false" json:"packageType,omitempty"`

	// Summary of the pricing types available across all packages in the listing.
	PricingTypes []ListingSummaryPricingTypesEnum `mandatory:"false" json:"pricingTypes,omitempty"`

	// The regions where the listing is eligible to be deployed.
	Regions []Region `mandatory:"false" json:"regions"`

	// Indicates whether the listing is featured.
	IsFeatured *bool `mandatory:"false" json:"isFeatured"`

	// Product categories that the listing belongs to.
	Categories []string `mandatory:"false" json:"categories"`

	Publisher *PublisherSummary `mandatory:"false" json:"publisher"`
}

func (m ListingSummary) String() string {
	return common.PointerString(m)
}

// ListingSummaryPricingTypesEnum Enum with underlying type: string
type ListingSummaryPricingTypesEnum string

// Set of constants representing the allowable values for ListingSummaryPricingTypesEnum
const (
	ListingSummaryPricingTypesFree  ListingSummaryPricingTypesEnum = "FREE"
	ListingSummaryPricingTypesByol  ListingSummaryPricingTypesEnum = "BYOL"
	ListingSummaryPricingTypesPaygo ListingSummaryPricingTypesEnum = "PAYGO"
)

var mappingListingSummaryPricingTypes = map[string]ListingSummaryPricingTypesEnum{
	"FREE":  ListingSummaryPricingTypesFree,
	"BYOL":  ListingSummaryPricingTypesByol,
	"PAYGO": ListingSummaryPricingTypesPaygo,
}

// GetListingSummaryPricingTypesEnumValues Enumerates the set of values for ListingSummaryPricingTypesEnum
func GetListingSummaryPricingTypesEnumValues() []ListingSummaryPricingTypesEnum {
	values := make([]ListingSummaryPricingTypesEnum, 0)
	for _, v := range mappingListingSummaryPricingTypes {
		values = append(values, v)
	}
	return values
}
