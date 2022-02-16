// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ImageListingPackage A package for image listings.
type ImageListingPackage struct {

	// The ID of the listing this package belongs to.
	ListingId *string `mandatory:"true" json:"listingId"`

	// The package version.
	Version *string `mandatory:"true" json:"version"`

	// Description of this package.
	Description *string `mandatory:"false" json:"description"`

	Pricing *PricingModel `mandatory:"false" json:"pricing"`

	// The unique identifier for the package resource.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The date and time this listing package was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	OperatingSystem *OperatingSystem `mandatory:"false" json:"operatingSystem"`

	// The ID of the listing resource associated with this listing package. For more information, see AppCatalogListing (https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/latest/AppCatalogListing/) in the Core Services API.
	AppCatalogListingId *string `mandatory:"false" json:"appCatalogListingId"`

	// The resource version of the listing resource associated with this listing package.
	AppCatalogListingResourceVersion *string `mandatory:"false" json:"appCatalogListingResourceVersion"`

	// The ID of the image corresponding to the package.
	ImageId *string `mandatory:"false" json:"imageId"`

	// The regions where you can deploy the listing package. (Some packages have restrictions that limit their deployment to United States regions only.)
	Regions []Region `mandatory:"false" json:"regions"`
}

//GetDescription returns Description
func (m ImageListingPackage) GetDescription() *string {
	return m.Description
}

//GetListingId returns ListingId
func (m ImageListingPackage) GetListingId() *string {
	return m.ListingId
}

//GetVersion returns Version
func (m ImageListingPackage) GetVersion() *string {
	return m.Version
}

//GetPricing returns Pricing
func (m ImageListingPackage) GetPricing() *PricingModel {
	return m.Pricing
}

//GetResourceId returns ResourceId
func (m ImageListingPackage) GetResourceId() *string {
	return m.ResourceId
}

//GetTimeCreated returns TimeCreated
func (m ImageListingPackage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetOperatingSystem returns OperatingSystem
func (m ImageListingPackage) GetOperatingSystem() *OperatingSystem {
	return m.OperatingSystem
}

func (m ImageListingPackage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ImageListingPackage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ImageListingPackage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeImageListingPackage ImageListingPackage
	s := struct {
		DiscriminatorParam string `json:"packageType"`
		MarshalTypeImageListingPackage
	}{
		"IMAGE",
		(MarshalTypeImageListingPackage)(m),
	}

	return json.Marshal(&s)
}
