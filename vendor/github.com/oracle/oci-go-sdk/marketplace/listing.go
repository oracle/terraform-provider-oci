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

	// The regions where the listing is eligible to be deployed.
	Regions []Region `mandatory:"false" json:"regions"`

	// The listing's package type.
	PackageType PackageTypeEnumEnum `mandatory:"false" json:"packageType,omitempty"`

	// The default package version.
	DefaultPackageVersion *string `mandatory:"false" json:"defaultPackageVersion"`

	// Links to reference material.
	Links []Link `mandatory:"false" json:"links"`

	// Indicates whether the listing is included in Featured Listings.
	IsFeatured *bool `mandatory:"false" json:"isFeatured"`
}

func (m Listing) String() string {
	return common.PointerString(m)
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
)

// GetListingPackageTypeEnumValues Enumerates the set of values for PackageTypeEnumEnum
// Consider using GetPackageTypeEnumEnumValue
// Deprecated
var GetListingPackageTypeEnumValues = GetPackageTypeEnumEnumValues
