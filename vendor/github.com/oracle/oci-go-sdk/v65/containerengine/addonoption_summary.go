// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddonoptionSummary The properties that define addon summary.
type AddonoptionSummary struct {

	// Name of the addon and it would be unique.
	Name *string `mandatory:"true" json:"name"`

	// The life cycle state of the addon.
	LifecycleState AddonoptionSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Is it an essential addon for cluster operation or not.
	IsEssential *bool `mandatory:"true" json:"isEssential"`

	// The resources this work request affects.
	Versions []AddonVersions `mandatory:"true" json:"versions"`

	// Addon definition schema version to validate addon.
	AddonSchemaVersion *string `mandatory:"false" json:"addonSchemaVersion"`

	// Addon group info, a namespace concept that groups addons with similar functionalities.
	AddonGroup *string `mandatory:"false" json:"addonGroup"`

	// Description on the addon.
	Description *string `mandatory:"false" json:"description"`

	// Defined tags of the addon.
	DefinedTags *string `mandatory:"false" json:"definedTags"`

	// FreeFrom tags of the addon.
	FreeformTags *string `mandatory:"false" json:"freeformTags"`

	// The time the work request was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m AddonoptionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddonoptionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddonoptionSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAddonoptionSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AddonoptionSummaryLifecycleStateEnum Enum with underlying type: string
type AddonoptionSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AddonoptionSummaryLifecycleStateEnum
const (
	AddonoptionSummaryLifecycleStateActive   AddonoptionSummaryLifecycleStateEnum = "ACTIVE"
	AddonoptionSummaryLifecycleStateInactive AddonoptionSummaryLifecycleStateEnum = "INACTIVE"
)

var mappingAddonoptionSummaryLifecycleStateEnum = map[string]AddonoptionSummaryLifecycleStateEnum{
	"ACTIVE":   AddonoptionSummaryLifecycleStateActive,
	"INACTIVE": AddonoptionSummaryLifecycleStateInactive,
}

var mappingAddonoptionSummaryLifecycleStateEnumLowerCase = map[string]AddonoptionSummaryLifecycleStateEnum{
	"active":   AddonoptionSummaryLifecycleStateActive,
	"inactive": AddonoptionSummaryLifecycleStateInactive,
}

// GetAddonoptionSummaryLifecycleStateEnumValues Enumerates the set of values for AddonoptionSummaryLifecycleStateEnum
func GetAddonoptionSummaryLifecycleStateEnumValues() []AddonoptionSummaryLifecycleStateEnum {
	values := make([]AddonoptionSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAddonoptionSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAddonoptionSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for AddonoptionSummaryLifecycleStateEnum
func GetAddonoptionSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingAddonoptionSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddonoptionSummaryLifecycleStateEnum(val string) (AddonoptionSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingAddonoptionSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
