// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AppCatalogListingResourceVersionAgreements Agreements for a listing resource version.
type AppCatalogListingResourceVersionAgreements struct {

	// The OCID of the listing associated with these agreements.
	ListingId *string `mandatory:"false" json:"listingId"`

	// Listing resource version associated with these agreements.
	ListingResourceVersion *string `mandatory:"false" json:"listingResourceVersion"`

	// Oracle TOU link
	OracleTermsOfUseLink *string `mandatory:"false" json:"oracleTermsOfUseLink"`

	// EULA link
	EulaLink *string `mandatory:"false" json:"eulaLink"`

	// Date and time the agreements were retrieved, in RFC3339 format.
	// Example: `2018-03-20T12:32:53.532Z`
	TimeRetrieved *common.SDKTime `mandatory:"false" json:"timeRetrieved"`

	// A generated signature for this agreement retrieval operation which should be used in the create subscription call.
	Signature *string `mandatory:"false" json:"signature"`
}

func (m AppCatalogListingResourceVersionAgreements) String() string {
	return common.PointerString(m)
}
