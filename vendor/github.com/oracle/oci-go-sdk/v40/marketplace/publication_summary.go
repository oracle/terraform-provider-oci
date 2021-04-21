// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/v40/common"
)

// PublicationSummary The model for a summary of an Oracle Cloud Infrastructure publication
type PublicationSummary struct {

	// the lifecycleState of the listing
	LifecycleState PublicationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The Compartment id where the listings exists
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique identifier for the listing in Marketplace.
	Id *string `mandatory:"true" json:"id"`

	// The name of the listing.
	Name *string `mandatory:"true" json:"name"`

	// In which catalog the listing should exist.
	ListingType ListingTypeEnum `mandatory:"true" json:"listingType"`

	// A short description of the listing.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	Icon *UploadData `mandatory:"false" json:"icon"`

	// The listing's package type.
	PackageType PackageTypeEnumEnum `mandatory:"false" json:"packageType,omitempty"`

	// List of operating systems supprted.
	SupportedOperatingSystems []OperatingSystem `mandatory:"false" json:"supportedOperatingSystems"`

	// The date and time this publication was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m PublicationSummary) String() string {
	return common.PointerString(m)
}
