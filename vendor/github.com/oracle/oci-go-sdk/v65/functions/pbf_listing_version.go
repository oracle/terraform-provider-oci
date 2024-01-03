// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PbfListingVersion This represents a version of a PbfListing. Each new update from the publisher or the change in the image will
// result in the creation of new PbfListingVersion resource creation. This is a sub-resource of a PbfListing.
type PbfListingVersion struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the PbfListing this resource version belongs to.
	PbfListingId *string `mandatory:"true" json:"pbfListingId"`

	// Semantic version
	Name *string `mandatory:"true" json:"name"`

	Requirements *RequirementDetails `mandatory:"true" json:"requirements"`

	// Details changes are included in this version.
	ChangeSummary *string `mandatory:"true" json:"changeSummary"`

	// An array of Trigger. A list of triggers that may activate the PBF.
	Triggers []Trigger `mandatory:"true" json:"triggers"`

	// The time the PbfListingVersion was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The last time the PbfListingVersion was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the PBF resource.
	LifecycleState PbfListingVersionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Details about the required and optional Function configurations needed for proper performance of the PBF.
	Config []ConfigDetails `mandatory:"false" json:"config"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m PbfListingVersion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PbfListingVersion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPbfListingVersionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPbfListingVersionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PbfListingVersionLifecycleStateEnum Enum with underlying type: string
type PbfListingVersionLifecycleStateEnum string

// Set of constants representing the allowable values for PbfListingVersionLifecycleStateEnum
const (
	PbfListingVersionLifecycleStateActive   PbfListingVersionLifecycleStateEnum = "ACTIVE"
	PbfListingVersionLifecycleStateInactive PbfListingVersionLifecycleStateEnum = "INACTIVE"
	PbfListingVersionLifecycleStateDeleted  PbfListingVersionLifecycleStateEnum = "DELETED"
)

var mappingPbfListingVersionLifecycleStateEnum = map[string]PbfListingVersionLifecycleStateEnum{
	"ACTIVE":   PbfListingVersionLifecycleStateActive,
	"INACTIVE": PbfListingVersionLifecycleStateInactive,
	"DELETED":  PbfListingVersionLifecycleStateDeleted,
}

var mappingPbfListingVersionLifecycleStateEnumLowerCase = map[string]PbfListingVersionLifecycleStateEnum{
	"active":   PbfListingVersionLifecycleStateActive,
	"inactive": PbfListingVersionLifecycleStateInactive,
	"deleted":  PbfListingVersionLifecycleStateDeleted,
}

// GetPbfListingVersionLifecycleStateEnumValues Enumerates the set of values for PbfListingVersionLifecycleStateEnum
func GetPbfListingVersionLifecycleStateEnumValues() []PbfListingVersionLifecycleStateEnum {
	values := make([]PbfListingVersionLifecycleStateEnum, 0)
	for _, v := range mappingPbfListingVersionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPbfListingVersionLifecycleStateEnumStringValues Enumerates the set of values in String for PbfListingVersionLifecycleStateEnum
func GetPbfListingVersionLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}

// GetMappingPbfListingVersionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPbfListingVersionLifecycleStateEnum(val string) (PbfListingVersionLifecycleStateEnum, bool) {
	enum, ok := mappingPbfListingVersionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
