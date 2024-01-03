// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Listing The model for an Oracle Cloud Infrastructure Marketplace listing.
type Listing struct {

	// The unique identifier for the listing in Marketplace.
	Id *string `mandatory:"false" json:"id"`

	// The name of the listing.
	Name *string `mandatory:"false" json:"name"`

	// The version of the listing.
	Version *string `mandatory:"false" json:"version"`

	// The tagline of the listing.
	Tagline *string `mandatory:"false" json:"tagline"`

	// Keywords associated with the listing.
	Keywords *string `mandatory:"false" json:"keywords"`

	// A short description of the listing.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	// Usage information for the listing.
	UsageInformation *string `mandatory:"false" json:"usageInformation"`

	// A long description of the listing.
	LongDescription *string `mandatory:"false" json:"longDescription"`

	// A description of the publisher's licensing model for the listing.
	LicenseModelDescription *string `mandatory:"false" json:"licenseModelDescription"`

	// System requirements for the listing.
	SystemRequirements *string `mandatory:"false" json:"systemRequirements"`

	// The release date of the listing.
	TimeReleased *common.SDKTime `mandatory:"false" json:"timeReleased"`

	// Release notes for the listing.
	ReleaseNotes *string `mandatory:"false" json:"releaseNotes"`

	// Categories that the listing belongs to.
	Categories []string `mandatory:"false" json:"categories"`

	Publisher *Publisher `mandatory:"false" json:"publisher"`

	// Languages supported by the listing.
	Languages []Item `mandatory:"false" json:"languages"`

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

	Icon *UploadData `mandatory:"false" json:"icon"`

	Banner *UploadData `mandatory:"false" json:"banner"`

	// The list of compatible architectures supported by the listing
	CompatibleArchitectures []ListingCompatibleArchitecturesEnum `mandatory:"false" json:"compatibleArchitectures,omitempty"`

	// The regions where you can deploy the listing. (Some listings have restrictions that limit their deployment to United States regions only.)
	Regions []Region `mandatory:"false" json:"regions"`

	// The listing's package type.
	PackageType PackageTypeEnumEnum `mandatory:"false" json:"packageType,omitempty"`

	// The default package version.
	DefaultPackageVersion *string `mandatory:"false" json:"defaultPackageVersion"`

	// Links to reference material.
	Links []Link `mandatory:"false" json:"links"`

	// Indicates whether the listing is included in Featured Listings.
	IsFeatured *bool `mandatory:"false" json:"isFeatured"`

	// The publisher category to which the listing belongs. The publisher category informs where the listing appears for use.
	ListingType ListingTypeEnum `mandatory:"false" json:"listingType,omitempty"`

	// List of operating systems supported by the listing.
	SupportedOperatingSystems []OperatingSystem `mandatory:"false" json:"supportedOperatingSystems"`
}

func (m Listing) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Listing) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.CompatibleArchitectures {
		if _, ok := GetMappingListingCompatibleArchitecturesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CompatibleArchitectures: %s. Supported values are: %s.", val, strings.Join(GetListingCompatibleArchitecturesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingPackageTypeEnumEnum(string(m.PackageType)); !ok && m.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", m.PackageType, strings.Join(GetPackageTypeEnumEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListingTypeEnum(string(m.ListingType)); !ok && m.ListingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListingType: %s. Supported values are: %s.", m.ListingType, strings.Join(GetListingTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListingCompatibleArchitecturesEnum Enum with underlying type: string
type ListingCompatibleArchitecturesEnum string

// Set of constants representing the allowable values for ListingCompatibleArchitecturesEnum
const (
	ListingCompatibleArchitecturesX86 ListingCompatibleArchitecturesEnum = "X86"
	ListingCompatibleArchitecturesArm ListingCompatibleArchitecturesEnum = "ARM"
)

var mappingListingCompatibleArchitecturesEnum = map[string]ListingCompatibleArchitecturesEnum{
	"X86": ListingCompatibleArchitecturesX86,
	"ARM": ListingCompatibleArchitecturesArm,
}

var mappingListingCompatibleArchitecturesEnumLowerCase = map[string]ListingCompatibleArchitecturesEnum{
	"x86": ListingCompatibleArchitecturesX86,
	"arm": ListingCompatibleArchitecturesArm,
}

// GetListingCompatibleArchitecturesEnumValues Enumerates the set of values for ListingCompatibleArchitecturesEnum
func GetListingCompatibleArchitecturesEnumValues() []ListingCompatibleArchitecturesEnum {
	values := make([]ListingCompatibleArchitecturesEnum, 0)
	for _, v := range mappingListingCompatibleArchitecturesEnum {
		values = append(values, v)
	}
	return values
}

// GetListingCompatibleArchitecturesEnumStringValues Enumerates the set of values in String for ListingCompatibleArchitecturesEnum
func GetListingCompatibleArchitecturesEnumStringValues() []string {
	return []string{
		"X86",
		"ARM",
	}
}

// GetMappingListingCompatibleArchitecturesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingCompatibleArchitecturesEnum(val string) (ListingCompatibleArchitecturesEnum, bool) {
	enum, ok := mappingListingCompatibleArchitecturesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListingPackageTypeEnum is an alias to type: PackageTypeEnumEnum
// Consider using PackageTypeEnumEnum instead
// Deprecated
type ListingPackageTypeEnum = PackageTypeEnumEnum

// Set of constants representing the allowable values for PackageTypeEnumEnum
// Deprecated
const (
	ListingPackageTypeOrchestration PackageTypeEnumEnum = "ORCHESTRATION"
	ListingPackageTypeImage         PackageTypeEnumEnum = "IMAGE"
	ListingPackageTypeContainer     PackageTypeEnumEnum = "CONTAINER"
	ListingPackageTypeKubernetes    PackageTypeEnumEnum = "KUBERNETES"
)

// GetListingPackageTypeEnumValues Enumerates the set of values for PackageTypeEnumEnum
// Consider using GetPackageTypeEnumEnumValue
// Deprecated
var GetListingPackageTypeEnumValues = GetPackageTypeEnumEnumValues
