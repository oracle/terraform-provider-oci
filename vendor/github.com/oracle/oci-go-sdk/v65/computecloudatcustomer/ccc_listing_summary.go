// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CccListingSummary The model for a summary of an Oracle Cloud Infrastructure marketplace listing.
type CccListingSummary struct {

	// The unique identifier for the listing in Marketplace.
	Id *string `mandatory:"true" json:"id"`

	// Compute Cloud@Customer marketplace listing display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A description of the listing.
	Description *string `mandatory:"true" json:"description"`

	// The pricing type for the listing and the versions of the package within the listing.
	PricingType CccListingSummaryPricingTypeEnum `mandatory:"true" json:"pricingType"`

	// The listing's package type.
	PackageType PackageTypeEnum `mandatory:"true" json:"packageType"`

	// The publisher category to which the listing belongs. The publisher category informs where the listing appears for use.
	ListingType CccListingSummaryListingTypeEnum `mandatory:"true" json:"listingType"`

	// The create time of the listing, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The tagline of the listing.
	Tagline *string `mandatory:"false" json:"tagline"`

	// Product categories that the listing belongs to.
	Categories []string `mandatory:"false" json:"categories"`

	Publisher *Publisher `mandatory:"false" json:"publisher"`

	Icon *UploadData `mandatory:"false" json:"icon"`

	Logo *UploadData `mandatory:"false" json:"logo"`

	// The list of compatible architectures supported by the listing
	CompatibleArchitectures []CccListingSummaryCompatibleArchitecturesEnum `mandatory:"false" json:"compatibleArchitectures,omitempty"`

	// The regions where you can deploy the listing. (Some listings have restrictions that limit their deployment to United States regions only.)
	Regions []Region `mandatory:"false" json:"regions"`

	// Indicates whether the listing is included in Featured Listings.
	IsFeatured *bool `mandatory:"false" json:"isFeatured"`

	PrivateOffer *PrivateOffer `mandatory:"false" json:"privateOffer"`

	// The list of operating systems supported by the listing.
	SupportedOperatingSystems []OperatingSystem `mandatory:"false" json:"supportedOperatingSystems"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CccListingSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CccListingSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCccListingSummaryPricingTypeEnum(string(m.PricingType)); !ok && m.PricingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PricingType: %s. Supported values are: %s.", m.PricingType, strings.Join(GetCccListingSummaryPricingTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPackageTypeEnum(string(m.PackageType)); !ok && m.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", m.PackageType, strings.Join(GetPackageTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCccListingSummaryListingTypeEnum(string(m.ListingType)); !ok && m.ListingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListingType: %s. Supported values are: %s.", m.ListingType, strings.Join(GetCccListingSummaryListingTypeEnumStringValues(), ",")))
	}

	for _, val := range m.CompatibleArchitectures {
		if _, ok := GetMappingCccListingSummaryCompatibleArchitecturesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CompatibleArchitectures: %s. Supported values are: %s.", val, strings.Join(GetCccListingSummaryCompatibleArchitecturesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CccListingSummaryCompatibleArchitecturesEnum Enum with underlying type: string
type CccListingSummaryCompatibleArchitecturesEnum string

// Set of constants representing the allowable values for CccListingSummaryCompatibleArchitecturesEnum
const (
	CccListingSummaryCompatibleArchitecturesX86 CccListingSummaryCompatibleArchitecturesEnum = "X86"
	CccListingSummaryCompatibleArchitecturesArm CccListingSummaryCompatibleArchitecturesEnum = "ARM"
)

var mappingCccListingSummaryCompatibleArchitecturesEnum = map[string]CccListingSummaryCompatibleArchitecturesEnum{
	"X86": CccListingSummaryCompatibleArchitecturesX86,
	"ARM": CccListingSummaryCompatibleArchitecturesArm,
}

var mappingCccListingSummaryCompatibleArchitecturesEnumLowerCase = map[string]CccListingSummaryCompatibleArchitecturesEnum{
	"x86": CccListingSummaryCompatibleArchitecturesX86,
	"arm": CccListingSummaryCompatibleArchitecturesArm,
}

// GetCccListingSummaryCompatibleArchitecturesEnumValues Enumerates the set of values for CccListingSummaryCompatibleArchitecturesEnum
func GetCccListingSummaryCompatibleArchitecturesEnumValues() []CccListingSummaryCompatibleArchitecturesEnum {
	values := make([]CccListingSummaryCompatibleArchitecturesEnum, 0)
	for _, v := range mappingCccListingSummaryCompatibleArchitecturesEnum {
		values = append(values, v)
	}
	return values
}

// GetCccListingSummaryCompatibleArchitecturesEnumStringValues Enumerates the set of values in String for CccListingSummaryCompatibleArchitecturesEnum
func GetCccListingSummaryCompatibleArchitecturesEnumStringValues() []string {
	return []string{
		"X86",
		"ARM",
	}
}

// GetMappingCccListingSummaryCompatibleArchitecturesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCccListingSummaryCompatibleArchitecturesEnum(val string) (CccListingSummaryCompatibleArchitecturesEnum, bool) {
	enum, ok := mappingCccListingSummaryCompatibleArchitecturesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CccListingSummaryPricingTypeEnum Enum with underlying type: string
type CccListingSummaryPricingTypeEnum string

// Set of constants representing the allowable values for CccListingSummaryPricingTypeEnum
const (
	CccListingSummaryPricingTypeFree  CccListingSummaryPricingTypeEnum = "FREE"
	CccListingSummaryPricingTypeByol  CccListingSummaryPricingTypeEnum = "BYOL"
	CccListingSummaryPricingTypePaygo CccListingSummaryPricingTypeEnum = "PAYGO"
)

var mappingCccListingSummaryPricingTypeEnum = map[string]CccListingSummaryPricingTypeEnum{
	"FREE":  CccListingSummaryPricingTypeFree,
	"BYOL":  CccListingSummaryPricingTypeByol,
	"PAYGO": CccListingSummaryPricingTypePaygo,
}

var mappingCccListingSummaryPricingTypeEnumLowerCase = map[string]CccListingSummaryPricingTypeEnum{
	"free":  CccListingSummaryPricingTypeFree,
	"byol":  CccListingSummaryPricingTypeByol,
	"paygo": CccListingSummaryPricingTypePaygo,
}

// GetCccListingSummaryPricingTypeEnumValues Enumerates the set of values for CccListingSummaryPricingTypeEnum
func GetCccListingSummaryPricingTypeEnumValues() []CccListingSummaryPricingTypeEnum {
	values := make([]CccListingSummaryPricingTypeEnum, 0)
	for _, v := range mappingCccListingSummaryPricingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCccListingSummaryPricingTypeEnumStringValues Enumerates the set of values in String for CccListingSummaryPricingTypeEnum
func GetCccListingSummaryPricingTypeEnumStringValues() []string {
	return []string{
		"FREE",
		"BYOL",
		"PAYGO",
	}
}

// GetMappingCccListingSummaryPricingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCccListingSummaryPricingTypeEnum(val string) (CccListingSummaryPricingTypeEnum, bool) {
	enum, ok := mappingCccListingSummaryPricingTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CccListingSummaryListingTypeEnum Enum with underlying type: string
type CccListingSummaryListingTypeEnum string

// Set of constants representing the allowable values for CccListingSummaryListingTypeEnum
const (
	CccListingSummaryListingTypeCommunity CccListingSummaryListingTypeEnum = "COMMUNITY"
	CccListingSummaryListingTypePartner   CccListingSummaryListingTypeEnum = "PARTNER"
	CccListingSummaryListingTypePrivate   CccListingSummaryListingTypeEnum = "PRIVATE"
)

var mappingCccListingSummaryListingTypeEnum = map[string]CccListingSummaryListingTypeEnum{
	"COMMUNITY": CccListingSummaryListingTypeCommunity,
	"PARTNER":   CccListingSummaryListingTypePartner,
	"PRIVATE":   CccListingSummaryListingTypePrivate,
}

var mappingCccListingSummaryListingTypeEnumLowerCase = map[string]CccListingSummaryListingTypeEnum{
	"community": CccListingSummaryListingTypeCommunity,
	"partner":   CccListingSummaryListingTypePartner,
	"private":   CccListingSummaryListingTypePrivate,
}

// GetCccListingSummaryListingTypeEnumValues Enumerates the set of values for CccListingSummaryListingTypeEnum
func GetCccListingSummaryListingTypeEnumValues() []CccListingSummaryListingTypeEnum {
	values := make([]CccListingSummaryListingTypeEnum, 0)
	for _, v := range mappingCccListingSummaryListingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCccListingSummaryListingTypeEnumStringValues Enumerates the set of values in String for CccListingSummaryListingTypeEnum
func GetCccListingSummaryListingTypeEnumStringValues() []string {
	return []string{
		"COMMUNITY",
		"PARTNER",
		"PRIVATE",
	}
}

// GetMappingCccListingSummaryListingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCccListingSummaryListingTypeEnum(val string) (CccListingSummaryListingTypeEnum, bool) {
	enum, ok := mappingCccListingSummaryListingTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
