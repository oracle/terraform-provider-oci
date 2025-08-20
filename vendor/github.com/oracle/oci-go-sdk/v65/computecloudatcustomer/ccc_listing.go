// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CccListing The model for an Oracle Compute Cloud@Customer marketplace listing.
type CccListing struct {

	// The unique identifier an Oracle Compute Cloud@Customer marketplace listing.
	Id *string `mandatory:"true" json:"id"`

	// Compute Cloud@Customer marketplace listing display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A description of the listing.
	Description *string `mandatory:"true" json:"description"`

	// The release date of the listing, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	Publisher *Publisher `mandatory:"true" json:"publisher"`

	// The pricing type for the listing and the versions of the package within the listing.
	PricingType CccListingPricingTypeEnum `mandatory:"true" json:"pricingType"`

	// The listing's package type.
	PackageType PackageTypeEnum `mandatory:"true" json:"packageType"`

	// The publisher category to which the listing belongs. The publisher category informs where the listing appears for use.
	ListingType CccListingListingTypeEnum `mandatory:"true" json:"listingType"`

	// The version of the listing, this is not the version of the package itself.
	Version *string `mandatory:"false" json:"version"`

	// The tagline of the listing.
	Tagline *string `mandatory:"false" json:"tagline"`

	// Keywords associated with the listing.
	Keywords *string `mandatory:"false" json:"keywords"`

	// Usage information for the listing.
	UsageInformation *string `mandatory:"false" json:"usageInformation"`

	// A long description of the listing.
	LongDescription *string `mandatory:"false" json:"longDescription"`

	// A description of the publisher's licensing model for the listing.
	LicenseModelDescription *string `mandatory:"false" json:"licenseModelDescription"`

	// System requirements for the listing.
	SystemRequirements *string `mandatory:"false" json:"systemRequirements"`

	// Release notes for the listing.
	ReleaseNotes *string `mandatory:"false" json:"releaseNotes"`

	// Categories that the listing belongs to.
	Categories []string `mandatory:"false" json:"categories"`

	// Languages supported by the listing.
	Languages []Language `mandatory:"false" json:"languages"`

	// Screenshots of the listing.
	Screenshots []Screenshot `mandatory:"false" json:"screenshots"`

	// Videos of the listing.
	Videos []NamedLink `mandatory:"false" json:"videos"`

	// Contact information to use to get support from the publisher for the listing.
	SupportContacts []SupportContact `mandatory:"false" json:"supportContacts"`

	// Links to support resources for the listing.
	SupportLinks []NamedLink `mandatory:"false" json:"supportLinks"`

	// Links to additional documentation provided by the publisher specifically for the listing.
	DocumentationLinks []DocumentationLink `mandatory:"false" json:"documentationLinks"`

	Disclaimer *Link `mandatory:"false" json:"disclaimer"`

	Icon *UploadData `mandatory:"false" json:"icon"`

	Logo *UploadData `mandatory:"false" json:"logo"`

	Banner *UploadData `mandatory:"false" json:"banner"`

	// The list of compatible architectures supported by the listing
	CompatibleArchitectures []CccListingCompatibleArchitecturesEnum `mandatory:"false" json:"compatibleArchitectures,omitempty"`

	// The regions where you can deploy the listing. (Some listings have restrictions that
	// limit their deployment to United States regions only.)
	Regions []Region `mandatory:"false" json:"regions"`

	// The default package id.
	DefaultPackageId *string `mandatory:"false" json:"defaultPackageId"`

	// Links to reference material.
	Links []Link `mandatory:"false" json:"links"`

	// Indicates whether the listing is included in Featured Listings.
	IsFeatured *bool `mandatory:"false" json:"isFeatured"`

	PrivateOffer *PrivateOffer `mandatory:"false" json:"privateOffer"`

	// List of operating systems supported by the listing.
	SupportedOperatingSystems []OperatingSystem `mandatory:"false" json:"supportedOperatingSystems"`
}

func (m CccListing) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CccListing) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCccListingPricingTypeEnum(string(m.PricingType)); !ok && m.PricingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PricingType: %s. Supported values are: %s.", m.PricingType, strings.Join(GetCccListingPricingTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPackageTypeEnum(string(m.PackageType)); !ok && m.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", m.PackageType, strings.Join(GetPackageTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCccListingListingTypeEnum(string(m.ListingType)); !ok && m.ListingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListingType: %s. Supported values are: %s.", m.ListingType, strings.Join(GetCccListingListingTypeEnumStringValues(), ",")))
	}

	for _, val := range m.CompatibleArchitectures {
		if _, ok := GetMappingCccListingCompatibleArchitecturesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CompatibleArchitectures: %s. Supported values are: %s.", val, strings.Join(GetCccListingCompatibleArchitecturesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CccListingCompatibleArchitecturesEnum Enum with underlying type: string
type CccListingCompatibleArchitecturesEnum string

// Set of constants representing the allowable values for CccListingCompatibleArchitecturesEnum
const (
	CccListingCompatibleArchitecturesX86 CccListingCompatibleArchitecturesEnum = "X86"
	CccListingCompatibleArchitecturesArm CccListingCompatibleArchitecturesEnum = "ARM"
)

var mappingCccListingCompatibleArchitecturesEnum = map[string]CccListingCompatibleArchitecturesEnum{
	"X86": CccListingCompatibleArchitecturesX86,
	"ARM": CccListingCompatibleArchitecturesArm,
}

var mappingCccListingCompatibleArchitecturesEnumLowerCase = map[string]CccListingCompatibleArchitecturesEnum{
	"x86": CccListingCompatibleArchitecturesX86,
	"arm": CccListingCompatibleArchitecturesArm,
}

// GetCccListingCompatibleArchitecturesEnumValues Enumerates the set of values for CccListingCompatibleArchitecturesEnum
func GetCccListingCompatibleArchitecturesEnumValues() []CccListingCompatibleArchitecturesEnum {
	values := make([]CccListingCompatibleArchitecturesEnum, 0)
	for _, v := range mappingCccListingCompatibleArchitecturesEnum {
		values = append(values, v)
	}
	return values
}

// GetCccListingCompatibleArchitecturesEnumStringValues Enumerates the set of values in String for CccListingCompatibleArchitecturesEnum
func GetCccListingCompatibleArchitecturesEnumStringValues() []string {
	return []string{
		"X86",
		"ARM",
	}
}

// GetMappingCccListingCompatibleArchitecturesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCccListingCompatibleArchitecturesEnum(val string) (CccListingCompatibleArchitecturesEnum, bool) {
	enum, ok := mappingCccListingCompatibleArchitecturesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CccListingPricingTypeEnum Enum with underlying type: string
type CccListingPricingTypeEnum string

// Set of constants representing the allowable values for CccListingPricingTypeEnum
const (
	CccListingPricingTypeFree  CccListingPricingTypeEnum = "FREE"
	CccListingPricingTypeByol  CccListingPricingTypeEnum = "BYOL"
	CccListingPricingTypePaygo CccListingPricingTypeEnum = "PAYGO"
)

var mappingCccListingPricingTypeEnum = map[string]CccListingPricingTypeEnum{
	"FREE":  CccListingPricingTypeFree,
	"BYOL":  CccListingPricingTypeByol,
	"PAYGO": CccListingPricingTypePaygo,
}

var mappingCccListingPricingTypeEnumLowerCase = map[string]CccListingPricingTypeEnum{
	"free":  CccListingPricingTypeFree,
	"byol":  CccListingPricingTypeByol,
	"paygo": CccListingPricingTypePaygo,
}

// GetCccListingPricingTypeEnumValues Enumerates the set of values for CccListingPricingTypeEnum
func GetCccListingPricingTypeEnumValues() []CccListingPricingTypeEnum {
	values := make([]CccListingPricingTypeEnum, 0)
	for _, v := range mappingCccListingPricingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCccListingPricingTypeEnumStringValues Enumerates the set of values in String for CccListingPricingTypeEnum
func GetCccListingPricingTypeEnumStringValues() []string {
	return []string{
		"FREE",
		"BYOL",
		"PAYGO",
	}
}

// GetMappingCccListingPricingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCccListingPricingTypeEnum(val string) (CccListingPricingTypeEnum, bool) {
	enum, ok := mappingCccListingPricingTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CccListingListingTypeEnum Enum with underlying type: string
type CccListingListingTypeEnum string

// Set of constants representing the allowable values for CccListingListingTypeEnum
const (
	CccListingListingTypeCommunity CccListingListingTypeEnum = "COMMUNITY"
	CccListingListingTypePartner   CccListingListingTypeEnum = "PARTNER"
	CccListingListingTypePrivate   CccListingListingTypeEnum = "PRIVATE"
)

var mappingCccListingListingTypeEnum = map[string]CccListingListingTypeEnum{
	"COMMUNITY": CccListingListingTypeCommunity,
	"PARTNER":   CccListingListingTypePartner,
	"PRIVATE":   CccListingListingTypePrivate,
}

var mappingCccListingListingTypeEnumLowerCase = map[string]CccListingListingTypeEnum{
	"community": CccListingListingTypeCommunity,
	"partner":   CccListingListingTypePartner,
	"private":   CccListingListingTypePrivate,
}

// GetCccListingListingTypeEnumValues Enumerates the set of values for CccListingListingTypeEnum
func GetCccListingListingTypeEnumValues() []CccListingListingTypeEnum {
	values := make([]CccListingListingTypeEnum, 0)
	for _, v := range mappingCccListingListingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCccListingListingTypeEnumStringValues Enumerates the set of values in String for CccListingListingTypeEnum
func GetCccListingListingTypeEnumStringValues() []string {
	return []string{
		"COMMUNITY",
		"PARTNER",
		"PRIVATE",
	}
}

// GetMappingCccListingListingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCccListingListingTypeEnum(val string) (CccListingListingTypeEnum, bool) {
	enum, ok := mappingCccListingListingTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
