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

// CreateAcceptedAgreementDetails The model for the parameters needed to accept a terms of use agreement.
type CreateAcceptedAgreementDetails struct {

	// The unique identifier for the compartment where the agreement will be accepted.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique identifier for the listing associated with the agreement.
	ListingId *string `mandatory:"true" json:"listingId"`

	// The package version associated with the agreement.
	PackageVersion *string `mandatory:"true" json:"packageVersion"`

	// The agreement to accept.
	AgreementId *string `mandatory:"true" json:"agreementId"`

	// A signature generated for the listing package agreements that you can retrieve
	// with GetAgreement (https://docs.cloud.oracle.com/api/#/en/marketplace/20181001/Agreement/GetAgreement).
	Signature *string `mandatory:"true" json:"signature"`

	// A display name for the accepted agreement.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreateAcceptedAgreementDetails) String() string {
	return common.PointerString(m)
}
