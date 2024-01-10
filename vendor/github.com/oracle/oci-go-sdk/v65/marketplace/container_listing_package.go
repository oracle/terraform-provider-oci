// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ContainerListingPackage A listing package for container.
type ContainerListingPackage struct {

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

	// The regions where you can deploy the listing package. (Some packages have restrictions that limit their deployment to United States regions only.)
	Regions []Region `mandatory:"false" json:"regions"`
}

// GetDescription returns Description
func (m ContainerListingPackage) GetDescription() *string {
	return m.Description
}

// GetListingId returns ListingId
func (m ContainerListingPackage) GetListingId() *string {
	return m.ListingId
}

// GetVersion returns Version
func (m ContainerListingPackage) GetVersion() *string {
	return m.Version
}

// GetPricing returns Pricing
func (m ContainerListingPackage) GetPricing() *PricingModel {
	return m.Pricing
}

// GetResourceId returns ResourceId
func (m ContainerListingPackage) GetResourceId() *string {
	return m.ResourceId
}

// GetTimeCreated returns TimeCreated
func (m ContainerListingPackage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetOperatingSystem returns OperatingSystem
func (m ContainerListingPackage) GetOperatingSystem() *OperatingSystem {
	return m.OperatingSystem
}

// GetRegions returns Regions
func (m ContainerListingPackage) GetRegions() []Region {
	return m.Regions
}

func (m ContainerListingPackage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerListingPackage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ContainerListingPackage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeContainerListingPackage ContainerListingPackage
	s := struct {
		DiscriminatorParam string `json:"packageType"`
		MarshalTypeContainerListingPackage
	}{
		"CONTAINER",
		(MarshalTypeContainerListingPackage)(m),
	}

	return json.Marshal(&s)
}
