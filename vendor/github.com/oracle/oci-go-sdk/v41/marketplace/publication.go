// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// Publication The model for an Oracle Cloud Infrastructure Marketplace Publication
type Publication struct {

	// The Compartment id where the listings exists
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique identifier for the listing in Marketplace.
	Id *string `mandatory:"true" json:"id"`

	// The name of the listing.
	Name *string `mandatory:"true" json:"name"`

	// In which catalog the listing should exist.
	ListingType ListingTypeEnum `mandatory:"true" json:"listingType"`

	// The state of the listing in its lifecycle
	LifecycleState PublicationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A short description of the listing.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	// A long description of the listing.
	LongDescription *string `mandatory:"false" json:"longDescription"`

	// Contact information to use to get support from the publisher for the listing.
	SupportContacts []SupportContact `mandatory:"false" json:"supportContacts"`

	Icon *UploadData `mandatory:"false" json:"icon"`

	// The listing's package type.
	PackageType PackageTypeEnumEnum `mandatory:"false" json:"packageType,omitempty"`

	// List of operating systems supprted.
	SupportedOperatingSystems []OperatingSystem `mandatory:"false" json:"supportedOperatingSystems"`

	// The date and time this publication was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m Publication) String() string {
	return common.PointerString(m)
}

// PublicationLifecycleStateEnum Enum with underlying type: string
type PublicationLifecycleStateEnum string

// Set of constants representing the allowable values for PublicationLifecycleStateEnum
const (
	PublicationLifecycleStateCreating PublicationLifecycleStateEnum = "CREATING"
	PublicationLifecycleStateActive   PublicationLifecycleStateEnum = "ACTIVE"
	PublicationLifecycleStateDeleting PublicationLifecycleStateEnum = "DELETING"
	PublicationLifecycleStateDeleted  PublicationLifecycleStateEnum = "DELETED"
	PublicationLifecycleStateFailed   PublicationLifecycleStateEnum = "FAILED"
)

var mappingPublicationLifecycleState = map[string]PublicationLifecycleStateEnum{
	"CREATING": PublicationLifecycleStateCreating,
	"ACTIVE":   PublicationLifecycleStateActive,
	"DELETING": PublicationLifecycleStateDeleting,
	"DELETED":  PublicationLifecycleStateDeleted,
	"FAILED":   PublicationLifecycleStateFailed,
}

// GetPublicationLifecycleStateEnumValues Enumerates the set of values for PublicationLifecycleStateEnum
func GetPublicationLifecycleStateEnumValues() []PublicationLifecycleStateEnum {
	values := make([]PublicationLifecycleStateEnum, 0)
	for _, v := range mappingPublicationLifecycleState {
		values = append(values, v)
	}
	return values
}
