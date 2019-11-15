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

// AcceptedAgreementSummary The model for a summary of an accepted agreement.
type AcceptedAgreementSummary struct {

	// The unique identifier for the acceptance of the agreement within a specific compartment.
	Id *string `mandatory:"false" json:"id"`

	// A display name for the accepted agreement.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The unique identifier for the compartment where the agreement was accepted.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The unique identifier for the listing associated with the agreement.
	ListingId *string `mandatory:"false" json:"listingId"`

	// The package version associated with the agreement.
	PackageVersion *string `mandatory:"false" json:"packageVersion"`

	// The unique identifier for the terms of use agreement itself.
	AgreementId *string `mandatory:"false" json:"agreementId"`

	// The time the agreement was accepted.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`
}

func (m AcceptedAgreementSummary) String() string {
	return common.PointerString(m)
}
