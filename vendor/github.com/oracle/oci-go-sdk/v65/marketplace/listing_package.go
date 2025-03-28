// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.oracle.com/iaas/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ListingPackage A base object for all types of listing packages.
type ListingPackage interface {

	// The ID of the listing this package belongs to.
	GetListingId() *string

	// The package version.
	GetVersion() *string

	// Description of this package.
	GetDescription() *string

	GetPricing() *PricingModel

	// The unique identifier for the package resource.
	GetResourceId() *string

	// The date and time this listing package was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

	GetOperatingSystem() *OperatingSystem

	// The regions where you can deploy the listing package. (Some packages have restrictions that limit their deployment to United States regions only.)
	GetRegions() []Region
}

type listingpackage struct {
	JsonData        []byte
	Description     *string          `mandatory:"false" json:"description"`
	Pricing         *PricingModel    `mandatory:"false" json:"pricing"`
	ResourceId      *string          `mandatory:"false" json:"resourceId"`
	TimeCreated     *common.SDKTime  `mandatory:"false" json:"timeCreated"`
	OperatingSystem *OperatingSystem `mandatory:"false" json:"operatingSystem"`
	Regions         []Region         `mandatory:"false" json:"regions"`
	ListingId       *string          `mandatory:"true" json:"listingId"`
	Version         *string          `mandatory:"true" json:"version"`
	PackageType     string           `json:"packageType"`
}

// UnmarshalJSON unmarshals json
func (m *listingpackage) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerlistingpackage listingpackage
	s := struct {
		Model Unmarshalerlistingpackage
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ListingId = s.Model.ListingId
	m.Version = s.Model.Version
	m.Description = s.Model.Description
	m.Pricing = s.Model.Pricing
	m.ResourceId = s.Model.ResourceId
	m.TimeCreated = s.Model.TimeCreated
	m.OperatingSystem = s.Model.OperatingSystem
	m.Regions = s.Model.Regions
	m.PackageType = s.Model.PackageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *listingpackage) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PackageType {
	case "CONTAINER":
		mm := ContainerListingPackage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORCHESTRATION":
		mm := OrchestrationListingPackage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IMAGE":
		mm := ImageListingPackage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KUBERNETES":
		mm := KubernetesListingPackage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ListingPackage: %s.", m.PackageType)
		return *m, nil
	}
}

// GetDescription returns Description
func (m listingpackage) GetDescription() *string {
	return m.Description
}

// GetPricing returns Pricing
func (m listingpackage) GetPricing() *PricingModel {
	return m.Pricing
}

// GetResourceId returns ResourceId
func (m listingpackage) GetResourceId() *string {
	return m.ResourceId
}

// GetTimeCreated returns TimeCreated
func (m listingpackage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetOperatingSystem returns OperatingSystem
func (m listingpackage) GetOperatingSystem() *OperatingSystem {
	return m.OperatingSystem
}

// GetRegions returns Regions
func (m listingpackage) GetRegions() []Region {
	return m.Regions
}

// GetListingId returns ListingId
func (m listingpackage) GetListingId() *string {
	return m.ListingId
}

// GetVersion returns Version
func (m listingpackage) GetVersion() *string {
	return m.Version
}

func (m listingpackage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m listingpackage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
