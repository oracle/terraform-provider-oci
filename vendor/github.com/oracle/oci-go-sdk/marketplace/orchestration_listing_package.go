// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// OrchestrationListingPackage A listing package for orchestration.
type OrchestrationListingPackage struct {

	// The id of the listing this package belongs to.
	ListingId *string `mandatory:"true" json:"listingId"`

	// The version of this package.
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

	// Link to the orchestration resource.
	ResourceLink *string `mandatory:"false" json:"resourceLink"`

	// List of variables for the orchestration resource.
	Variables []OrchestrationVariable `mandatory:"false" json:"variables"`
}

//GetDescription returns Description
func (m OrchestrationListingPackage) GetDescription() *string {
	return m.Description
}

//GetListingId returns ListingId
func (m OrchestrationListingPackage) GetListingId() *string {
	return m.ListingId
}

//GetVersion returns Version
func (m OrchestrationListingPackage) GetVersion() *string {
	return m.Version
}

//GetPricing returns Pricing
func (m OrchestrationListingPackage) GetPricing() *PricingModel {
	return m.Pricing
}

//GetResourceId returns ResourceId
func (m OrchestrationListingPackage) GetResourceId() *string {
	return m.ResourceId
}

//GetTimeCreated returns TimeCreated
func (m OrchestrationListingPackage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

func (m OrchestrationListingPackage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m OrchestrationListingPackage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOrchestrationListingPackage OrchestrationListingPackage
	s := struct {
		DiscriminatorParam string `json:"packageType"`
		MarshalTypeOrchestrationListingPackage
	}{
		"ORCHESTRATION",
		(MarshalTypeOrchestrationListingPackage)(m),
	}

	return json.Marshal(&s)
}
