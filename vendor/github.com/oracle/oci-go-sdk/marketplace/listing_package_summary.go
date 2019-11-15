// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ListingPackageSummary The model for a summary of a package.
type ListingPackageSummary struct {

	// The id of the listing the specified package belongs to.
	ListingId *string `mandatory:"false" json:"listingId"`

	// The version of the specified package.
	PackageVersion *string `mandatory:"false" json:"packageVersion"`

	// The unique identifier for the package resource.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The date and time this listing package was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m ListingPackageSummary) String() string {
	return common.PointerString(m)
}
