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

// UpdatePublicationDetails The model for the parameters needed to update a publication.
type UpdatePublicationDetails struct {

	// The name of the publication, which is also used in the listing.
	Name *string `mandatory:"false" json:"name"`

	// A short description of the publication to use in the listing.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	// A long description of the publication to use in the listing.
	LongDescription *string `mandatory:"false" json:"longDescription"`

	// Contact information for getting support from the publisher for the listing.
	SupportContacts []SupportContact `mandatory:"false" json:"supportContacts"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m UpdatePublicationDetails) String() string {
	return common.PointerString(m)
}
