// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreateAcceptedAgreementDetails) String() string {
	return common.PointerString(m)
}
