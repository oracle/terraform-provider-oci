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

// AppCatalogListing Listing details.
type AppCatalogListing struct {

	// Listing's contact URL.
	ContactUrl *string `mandatory:"false" json:"contactUrl"`

	// Description of the listing.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the listing.
	ListingId *string `mandatory:"false" json:"listingId"`

	// Name of the listing.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Date and time the listing was published, in RFC3339 format.
	// Example: `2018-03-20T12:32:53.532Z`
	TimePublished *common.SDKTime `mandatory:"false" json:"timePublished"`

	// Publisher's logo URL.
	PublisherLogoUrl *string `mandatory:"false" json:"publisherLogoUrl"`

	// Name of the publisher who published this listing.
	PublisherName *string `mandatory:"false" json:"publisherName"`

	// Summary of the listing.
	Summary *string `mandatory:"false" json:"summary"`
}

func (m AppCatalogListing) String() string {
	return common.PointerString(m)
}
