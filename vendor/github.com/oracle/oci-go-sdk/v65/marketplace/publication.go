// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Publication The model for an Oracle Cloud Infrastructure Marketplace publication.
type Publication struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment where the publication exists.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique identifier for the publication in Marketplace.
	Id *string `mandatory:"true" json:"id"`

	// The name of the publication, which is also used in the listing.
	Name *string `mandatory:"true" json:"name"`

	// The publisher category to which the publication belongs. The publisher category informs where the listing appears for use.
	ListingType ListingTypeEnum `mandatory:"true" json:"listingType"`

	// The lifecycle state of the publication.
	LifecycleState PublicationLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A short description of the publication to use in the listing.
	ShortDescription *string `mandatory:"false" json:"shortDescription"`

	// A long description of the publication to use in the listing.
	LongDescription *string `mandatory:"false" json:"longDescription"`

	// Contact information for getting support from the publisher for the listing.
	SupportContacts []SupportContact `mandatory:"false" json:"supportContacts"`

	Icon *UploadData `mandatory:"false" json:"icon"`

	// The listing's package type.
	PackageType PackageTypeEnumEnum `mandatory:"false" json:"packageType,omitempty"`

	// The list of operating systems supprted by the listing.
	SupportedOperatingSystems []OperatingSystem `mandatory:"false" json:"supportedOperatingSystems"`

	// The date and time the publication was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
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

	// The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Publication) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Publication) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListingTypeEnum(string(m.ListingType)); !ok && m.ListingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListingType: %s. Supported values are: %s.", m.ListingType, strings.Join(GetListingTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPublicationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPublicationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPackageTypeEnumEnum(string(m.PackageType)); !ok && m.PackageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PackageType: %s. Supported values are: %s.", m.PackageType, strings.Join(GetPackageTypeEnumEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingPublicationLifecycleStateEnum = map[string]PublicationLifecycleStateEnum{
	"CREATING": PublicationLifecycleStateCreating,
	"ACTIVE":   PublicationLifecycleStateActive,
	"DELETING": PublicationLifecycleStateDeleting,
	"DELETED":  PublicationLifecycleStateDeleted,
	"FAILED":   PublicationLifecycleStateFailed,
}

var mappingPublicationLifecycleStateEnumLowerCase = map[string]PublicationLifecycleStateEnum{
	"creating": PublicationLifecycleStateCreating,
	"active":   PublicationLifecycleStateActive,
	"deleting": PublicationLifecycleStateDeleting,
	"deleted":  PublicationLifecycleStateDeleted,
	"failed":   PublicationLifecycleStateFailed,
}

// GetPublicationLifecycleStateEnumValues Enumerates the set of values for PublicationLifecycleStateEnum
func GetPublicationLifecycleStateEnumValues() []PublicationLifecycleStateEnum {
	values := make([]PublicationLifecycleStateEnum, 0)
	for _, v := range mappingPublicationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPublicationLifecycleStateEnumStringValues Enumerates the set of values in String for PublicationLifecycleStateEnum
func GetPublicationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingPublicationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPublicationLifecycleStateEnum(val string) (PublicationLifecycleStateEnum, bool) {
	enum, ok := mappingPublicationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
