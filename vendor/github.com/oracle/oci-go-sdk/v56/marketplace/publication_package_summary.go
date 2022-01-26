// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// PublicationPackageSummary The model for a summary of a publication package.
type PublicationPackageSummary struct {

	// The ID of the listing that the specified package belongs to.
	ListingId *string `mandatory:"true" json:"listingId"`

	// The version of the specified package.
	PackageVersion *string `mandatory:"true" json:"packageVersion"`

	// The specified package's type.
	PackageType PackageTypeEnumEnum `mandatory:"true" json:"packageType"`

	// The unique identifier for the package resource.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The date and time the publication package was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m PublicationPackageSummary) String() string {
	return common.PointerString(m)
}
