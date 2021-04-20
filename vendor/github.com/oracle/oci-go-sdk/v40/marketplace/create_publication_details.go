// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v40/common"
)

// CreatePublicationDetails Publication Creation Details
type CreatePublicationDetails struct {

	// In which catalog the listing should exist.
	ListingType ListingTypeEnum `mandatory:"true" json:"listingType"`

	// The name of the listing.
	Name *string `mandatory:"true" json:"name"`

	// short description of the catalog listing
	ShortDescription *string `mandatory:"true" json:"shortDescription"`

	// Contact information to use to get support from the publisher for the listing.
	SupportContacts []SupportContact `mandatory:"true" json:"supportContacts"`

	// The OCID of the compartment to create the resource within.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	PackageDetails CreatePublicationPackage `mandatory:"true" json:"packageDetails"`

	// Acknowledgement that invoker has the right and authority to share this Community Image in accordance with their agreement with Oracle applicable to the Services and the related Service Specifications
	IsAgreementAcknowledged *bool `mandatory:"true" json:"isAgreementAcknowledged"`

	// short description of the catalog listing
	LongDescription *string `mandatory:"false" json:"longDescription"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m CreatePublicationDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreatePublicationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		LongDescription         *string                           `json:"longDescription"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		ListingType             ListingTypeEnum                   `json:"listingType"`
		Name                    *string                           `json:"name"`
		ShortDescription        *string                           `json:"shortDescription"`
		SupportContacts         []SupportContact                  `json:"supportContacts"`
		CompartmentId           *string                           `json:"compartmentId"`
		PackageDetails          createpublicationpackage          `json:"packageDetails"`
		IsAgreementAcknowledged *bool                             `json:"isAgreementAcknowledged"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.LongDescription = model.LongDescription

	m.DefinedTags = model.DefinedTags

	m.FreeformTags = model.FreeformTags

	m.ListingType = model.ListingType

	m.Name = model.Name

	m.ShortDescription = model.ShortDescription

	m.SupportContacts = make([]SupportContact, len(model.SupportContacts))
	for i, n := range model.SupportContacts {
		m.SupportContacts[i] = n
	}

	m.CompartmentId = model.CompartmentId

	nn, e = model.PackageDetails.UnmarshalPolymorphicJSON(model.PackageDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.PackageDetails = nn.(CreatePublicationPackage)
	} else {
		m.PackageDetails = nil
	}

	m.IsAgreementAcknowledged = model.IsAgreementAcknowledged

	return
}
