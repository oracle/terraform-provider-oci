// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ListingSummary The model for a summary of an Oracle Cloud Infrastructure Marketplace listing.
type ListingSummary struct {

	// The unique identifier for the listing in Marketplace.
	Id *string `mandatory:"false" json:"id"`

	// The name of the listing.
	Name *string `mandatory:"false" json:"name"`

	// A short description of the listing.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	// True if this application is Rover exportable
	IsRoverExportable *bool `mandatory:"false" json:"isRoverExportable"`

	// The tagline of the listing.
	Tagline *string `mandatory:"false" json:"tagline"`

	Icon *UploadData `mandatory:"false" json:"icon"`

	// The listing's package type.
	PackageType PackageTypeEnumEnum `mandatory:"false" json:"packageType,omitempty"`

	// Summary of the pricing types available across all packages in the listing.
	PricingTypes []ListingSummaryPricingTypesEnum `mandatory:"false" json:"pricingTypes,omitempty"`

	// The regions where you can deploy the listing. (Some listings have restrictions that limit their deployment to United States regions only.)
	Regions []Region `mandatory:"false" json:"regions"`

	// Indicates whether the listing is featured.
	IsFeatured *bool `mandatory:"false" json:"isFeatured"`

	// Product categories that the listing belongs to.
	Categories []string `mandatory:"false" json:"categories"`

	Publisher *PublisherSummary `mandatory:"false" json:"publisher"`

	// The list of operating systems supported by the listing.
	SupportedOperatingSystems []OperatingSystem `mandatory:"false" json:"supportedOperatingSystems"`

	// The publisher category to which the listing belongs. The publisher category informs where the listing appears for use.
	ListingType ListingTypeEnum `mandatory:"false" json:"listingType,omitempty"`
}

func (m ListingSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ListingSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPackageTypeEnumEnum(string(m.PackageType)); !ok && m.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", m.PackageType, strings.Join(GetPackageTypeEnumEnumStringValues(), ",")))
	}
	for _, val := range m.PricingTypes {
		if _, ok := GetMappingListingSummaryPricingTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PricingTypes: %s. Supported values are: %s.", val, strings.Join(GetListingSummaryPricingTypesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListingTypeEnum(string(m.ListingType)); !ok && m.ListingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListingType: %s. Supported values are: %s.", m.ListingType, strings.Join(GetListingTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListingSummaryPricingTypesEnum Enum with underlying type: string
type ListingSummaryPricingTypesEnum string

// Set of constants representing the allowable values for ListingSummaryPricingTypesEnum
const (
	ListingSummaryPricingTypesFree  ListingSummaryPricingTypesEnum = "FREE"
	ListingSummaryPricingTypesByol  ListingSummaryPricingTypesEnum = "BYOL"
	ListingSummaryPricingTypesPaygo ListingSummaryPricingTypesEnum = "PAYGO"
)

var mappingListingSummaryPricingTypesEnum = map[string]ListingSummaryPricingTypesEnum{
	"FREE":  ListingSummaryPricingTypesFree,
	"BYOL":  ListingSummaryPricingTypesByol,
	"PAYGO": ListingSummaryPricingTypesPaygo,
}

// GetListingSummaryPricingTypesEnumValues Enumerates the set of values for ListingSummaryPricingTypesEnum
func GetListingSummaryPricingTypesEnumValues() []ListingSummaryPricingTypesEnum {
	values := make([]ListingSummaryPricingTypesEnum, 0)
	for _, v := range mappingListingSummaryPricingTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetListingSummaryPricingTypesEnumStringValues Enumerates the set of values in String for ListingSummaryPricingTypesEnum
func GetListingSummaryPricingTypesEnumStringValues() []string {
	return []string{
		"FREE",
		"BYOL",
		"PAYGO",
	}
}

// GetMappingListingSummaryPricingTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingSummaryPricingTypesEnum(val string) (ListingSummaryPricingTypesEnum, bool) {
	mappingListingSummaryPricingTypesEnumIgnoreCase := make(map[string]ListingSummaryPricingTypesEnum)
	for k, v := range mappingListingSummaryPricingTypesEnum {
		mappingListingSummaryPricingTypesEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListingSummaryPricingTypesEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
