// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secure Desktops API
//
// Create and manage cloud-hosted desktops which can be accessed from a web browser or installed client.
//

package desktops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DesktopPoolVolumeSummary Provides information about a volume within the desktop pool.
type DesktopPoolVolumeSummary struct {

	// The OCID of the desktop pool volume.
	Id *string `mandatory:"true" json:"id"`

	// The name of the desktop pool volume.
	Name *string `mandatory:"true" json:"name"`

	// The state of the desktop pool volume.
	LifecycleState DesktopPoolVolumeSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The owner of the desktop pool volume.
	UserName *string `mandatory:"true" json:"userName"`

	// The availability domain of the desktop pool.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the desktop pool to which this volume belongs.
	PoolId *string `mandatory:"false" json:"poolId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m DesktopPoolVolumeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DesktopPoolVolumeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDesktopPoolVolumeSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDesktopPoolVolumeSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DesktopPoolVolumeSummaryLifecycleStateEnum Enum with underlying type: string
type DesktopPoolVolumeSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for DesktopPoolVolumeSummaryLifecycleStateEnum
const (
	DesktopPoolVolumeSummaryLifecycleStateActive   DesktopPoolVolumeSummaryLifecycleStateEnum = "ACTIVE"
	DesktopPoolVolumeSummaryLifecycleStateInactive DesktopPoolVolumeSummaryLifecycleStateEnum = "INACTIVE"
)

var mappingDesktopPoolVolumeSummaryLifecycleStateEnum = map[string]DesktopPoolVolumeSummaryLifecycleStateEnum{
	"ACTIVE":   DesktopPoolVolumeSummaryLifecycleStateActive,
	"INACTIVE": DesktopPoolVolumeSummaryLifecycleStateInactive,
}

var mappingDesktopPoolVolumeSummaryLifecycleStateEnumLowerCase = map[string]DesktopPoolVolumeSummaryLifecycleStateEnum{
	"active":   DesktopPoolVolumeSummaryLifecycleStateActive,
	"inactive": DesktopPoolVolumeSummaryLifecycleStateInactive,
}

// GetDesktopPoolVolumeSummaryLifecycleStateEnumValues Enumerates the set of values for DesktopPoolVolumeSummaryLifecycleStateEnum
func GetDesktopPoolVolumeSummaryLifecycleStateEnumValues() []DesktopPoolVolumeSummaryLifecycleStateEnum {
	values := make([]DesktopPoolVolumeSummaryLifecycleStateEnum, 0)
	for _, v := range mappingDesktopPoolVolumeSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDesktopPoolVolumeSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for DesktopPoolVolumeSummaryLifecycleStateEnum
func GetDesktopPoolVolumeSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingDesktopPoolVolumeSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDesktopPoolVolumeSummaryLifecycleStateEnum(val string) (DesktopPoolVolumeSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingDesktopPoolVolumeSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
