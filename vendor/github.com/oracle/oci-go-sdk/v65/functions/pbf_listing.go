// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// PbfListing PbfListing resources provide details about the available PBFs for consumption by the user.
// This resource contains details about PBF's functionality, policies required, configuration parameters expected
// etc.
type PbfListing struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// A brief descriptive name for the PBF listing. The PBF listing name must be unique, and not match and existing
	// PBF.
	Name *string `mandatory:"true" json:"name"`

	// A short overview of the PBF Listing: the purpose of the PBF and and associated information.
	Description *string `mandatory:"true" json:"description"`

	PublisherDetails *PublisherDetails `mandatory:"true" json:"publisherDetails"`

	// The time the PbfListing was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The last time the PbfListing was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the PBF resource.
	LifecycleState PbfListingLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// An array of Trigger. A list of triggers that may activate the PBF.
	Triggers []Trigger `mandatory:"false" json:"triggers"`

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

func (m PbfListing) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PbfListing) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPbfListingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPbfListingLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PbfListingLifecycleStateEnum Enum with underlying type: string
type PbfListingLifecycleStateEnum string

// Set of constants representing the allowable values for PbfListingLifecycleStateEnum
const (
	PbfListingLifecycleStateActive   PbfListingLifecycleStateEnum = "ACTIVE"
	PbfListingLifecycleStateInactive PbfListingLifecycleStateEnum = "INACTIVE"
	PbfListingLifecycleStateDeleted  PbfListingLifecycleStateEnum = "DELETED"
)

var mappingPbfListingLifecycleStateEnum = map[string]PbfListingLifecycleStateEnum{
	"ACTIVE":   PbfListingLifecycleStateActive,
	"INACTIVE": PbfListingLifecycleStateInactive,
	"DELETED":  PbfListingLifecycleStateDeleted,
}

var mappingPbfListingLifecycleStateEnumLowerCase = map[string]PbfListingLifecycleStateEnum{
	"active":   PbfListingLifecycleStateActive,
	"inactive": PbfListingLifecycleStateInactive,
	"deleted":  PbfListingLifecycleStateDeleted,
}

// GetPbfListingLifecycleStateEnumValues Enumerates the set of values for PbfListingLifecycleStateEnum
func GetPbfListingLifecycleStateEnumValues() []PbfListingLifecycleStateEnum {
	values := make([]PbfListingLifecycleStateEnum, 0)
	for _, v := range mappingPbfListingLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPbfListingLifecycleStateEnumStringValues Enumerates the set of values in String for PbfListingLifecycleStateEnum
func GetPbfListingLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
	}
}

// GetMappingPbfListingLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPbfListingLifecycleStateEnum(val string) (PbfListingLifecycleStateEnum, bool) {
	enum, ok := mappingPbfListingLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
