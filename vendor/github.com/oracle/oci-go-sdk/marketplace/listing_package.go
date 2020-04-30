// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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
}

type listingpackage struct {
	JsonData    []byte
	ListingId   *string         `mandatory:"true" json:"listingId"`
	Version     *string         `mandatory:"true" json:"version"`
	Description *string         `mandatory:"false" json:"description"`
	Pricing     *PricingModel   `mandatory:"false" json:"pricing"`
	ResourceId  *string         `mandatory:"false" json:"resourceId"`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
	PackageType string          `json:"packageType"`
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
	case "ORCHESTRATION":
		mm := OrchestrationListingPackage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "IMAGE":
		mm := ImageListingPackage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetListingId returns ListingId
func (m listingpackage) GetListingId() *string {
	return m.ListingId
}

//GetVersion returns Version
func (m listingpackage) GetVersion() *string {
	return m.Version
}

//GetDescription returns Description
func (m listingpackage) GetDescription() *string {
	return m.Description
}

//GetPricing returns Pricing
func (m listingpackage) GetPricing() *PricingModel {
	return m.Pricing
}

//GetResourceId returns ResourceId
func (m listingpackage) GetResourceId() *string {
	return m.ResourceId
}

//GetTimeCreated returns TimeCreated
func (m listingpackage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

func (m listingpackage) String() string {
	return common.PointerString(m)
}
