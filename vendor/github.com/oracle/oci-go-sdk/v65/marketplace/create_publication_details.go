// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreatePublicationDetails The model for the parameters needed to create a publication.
type CreatePublicationDetails struct {

	// The publisher category to which the publication belongs. The publisher category informs where the listing appears for use.
	ListingType ListingTypeEnum `mandatory:"true" json:"listingType"`

	// The name of the publication, which is also used in the listing.
	Name *string `mandatory:"true" json:"name"`

	// A short description of the publication to use in the listing.
	ShortDescription *string `mandatory:"true" json:"shortDescription"`

	// Contact information for getting support from the publisher for the listing.
	SupportContacts []SupportContact `mandatory:"true" json:"supportContacts"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the publication.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	PackageDetails CreatePublicationPackage `mandatory:"true" json:"packageDetails"`

	// Whether the publisher acknowledged that they have the right and authority to share the contents of the publication and that they accepted the Oracle terms of use agreements required to create a publication.
	IsAgreementAcknowledged *bool `mandatory:"true" json:"isAgreementAcknowledged"`

	// A long description of the publication to use in the listing.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePublicationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListingTypeEnum(string(m.ListingType)); !ok && m.ListingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListingType: %s. Supported values are: %s.", m.ListingType, strings.Join(GetListingTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
	copy(m.SupportContacts, model.SupportContacts)
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
