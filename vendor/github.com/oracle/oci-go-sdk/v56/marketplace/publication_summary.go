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

// PublicationSummary The model for a summary of an Oracle Cloud Infrastructure publication.
type PublicationSummary struct {

	// The lifecycle state of the publication.
	LifecycleState PublicationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment where the publication exists.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique identifier for the publication in Marketplace.
	Id *string `mandatory:"true" json:"id"`

	// The name of the publication, which is also used in the listing.
	Name *string `mandatory:"true" json:"name"`

	// The publisher category to which the publication belongs. The publisher category informs where the listing appears for use.
	ListingType ListingTypeEnum `mandatory:"true" json:"listingType"`

	// A short description of the publication to use in the listing.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	Icon *UploadData `mandatory:"false" json:"icon"`

	// The listing's package type.
	PackageType PackageTypeEnumEnum `mandatory:"false" json:"packageType,omitempty"`

	// The list of operating systems supported by the listing.
	SupportedOperatingSystems []OperatingSystem `mandatory:"false" json:"supportedOperatingSystems"`

	// The date and time the publication was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m PublicationSummary) String() string {
	return common.PointerString(m)
}
