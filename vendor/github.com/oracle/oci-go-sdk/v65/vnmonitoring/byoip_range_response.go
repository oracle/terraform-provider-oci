// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ByoipRangeResponse A ByoipRange, is an IP address prefix that the user owns and wishes to import into OCI.
type ByoipRangeResponse struct {

	// The address range the user is on-boarding.
	CidrBlock *string `mandatory:"true" json:"cidrBlock"`

	// The OCID of the compartment containing the Byoip Range.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Oracle ID (OCID) of the Byoip Range.
	Id *string `mandatory:"true" json:"id"`

	// The Byoip Range's current state.
	LifecycleState ByoipRangeResponseLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// This is an internally generated ASCII string that the user will then use as part of the validation process. Specifically, they will need to add the token string generated by the service to their Internet Registry record.
	ValidationToken *string `mandatory:"true" json:"validationToken"`

	// The date and time the public IP pool was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The Byoip Range's current substate.
	LifecycleDetails ByoipRangeResponseLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// This field indicates the Byoip Range is locked for further modification or deletion
	Locked *bool `mandatory:"false" json:"locked"`
}

func (m ByoipRangeResponse) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ByoipRangeResponse) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingByoipRangeResponseLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetByoipRangeResponseLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingByoipRangeResponseLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetByoipRangeResponseLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ByoipRangeResponseLifecycleDetailsEnum Enum with underlying type: string
type ByoipRangeResponseLifecycleDetailsEnum string

// Set of constants representing the allowable values for ByoipRangeResponseLifecycleDetailsEnum
const (
	ByoipRangeResponseLifecycleDetailsCreating    ByoipRangeResponseLifecycleDetailsEnum = "CREATING"
	ByoipRangeResponseLifecycleDetailsValidating  ByoipRangeResponseLifecycleDetailsEnum = "VALIDATING"
	ByoipRangeResponseLifecycleDetailsProvisioned ByoipRangeResponseLifecycleDetailsEnum = "PROVISIONED"
	ByoipRangeResponseLifecycleDetailsActive      ByoipRangeResponseLifecycleDetailsEnum = "ACTIVE"
	ByoipRangeResponseLifecycleDetailsFailed      ByoipRangeResponseLifecycleDetailsEnum = "FAILED"
	ByoipRangeResponseLifecycleDetailsDeleting    ByoipRangeResponseLifecycleDetailsEnum = "DELETING"
	ByoipRangeResponseLifecycleDetailsDeleted     ByoipRangeResponseLifecycleDetailsEnum = "DELETED"
)

var mappingByoipRangeResponseLifecycleDetailsEnum = map[string]ByoipRangeResponseLifecycleDetailsEnum{
	"CREATING":    ByoipRangeResponseLifecycleDetailsCreating,
	"VALIDATING":  ByoipRangeResponseLifecycleDetailsValidating,
	"PROVISIONED": ByoipRangeResponseLifecycleDetailsProvisioned,
	"ACTIVE":      ByoipRangeResponseLifecycleDetailsActive,
	"FAILED":      ByoipRangeResponseLifecycleDetailsFailed,
	"DELETING":    ByoipRangeResponseLifecycleDetailsDeleting,
	"DELETED":     ByoipRangeResponseLifecycleDetailsDeleted,
}

var mappingByoipRangeResponseLifecycleDetailsEnumLowerCase = map[string]ByoipRangeResponseLifecycleDetailsEnum{
	"creating":    ByoipRangeResponseLifecycleDetailsCreating,
	"validating":  ByoipRangeResponseLifecycleDetailsValidating,
	"provisioned": ByoipRangeResponseLifecycleDetailsProvisioned,
	"active":      ByoipRangeResponseLifecycleDetailsActive,
	"failed":      ByoipRangeResponseLifecycleDetailsFailed,
	"deleting":    ByoipRangeResponseLifecycleDetailsDeleting,
	"deleted":     ByoipRangeResponseLifecycleDetailsDeleted,
}

// GetByoipRangeResponseLifecycleDetailsEnumValues Enumerates the set of values for ByoipRangeResponseLifecycleDetailsEnum
func GetByoipRangeResponseLifecycleDetailsEnumValues() []ByoipRangeResponseLifecycleDetailsEnum {
	values := make([]ByoipRangeResponseLifecycleDetailsEnum, 0)
	for _, v := range mappingByoipRangeResponseLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetByoipRangeResponseLifecycleDetailsEnumStringValues Enumerates the set of values in String for ByoipRangeResponseLifecycleDetailsEnum
func GetByoipRangeResponseLifecycleDetailsEnumStringValues() []string {
	return []string{
		"CREATING",
		"VALIDATING",
		"PROVISIONED",
		"ACTIVE",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingByoipRangeResponseLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingByoipRangeResponseLifecycleDetailsEnum(val string) (ByoipRangeResponseLifecycleDetailsEnum, bool) {
	enum, ok := mappingByoipRangeResponseLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ByoipRangeResponseLifecycleStateEnum Enum with underlying type: string
type ByoipRangeResponseLifecycleStateEnum string

// Set of constants representing the allowable values for ByoipRangeResponseLifecycleStateEnum
const (
	ByoipRangeResponseLifecycleStateInactive ByoipRangeResponseLifecycleStateEnum = "INACTIVE"
	ByoipRangeResponseLifecycleStateUpdating ByoipRangeResponseLifecycleStateEnum = "UPDATING"
	ByoipRangeResponseLifecycleStateActive   ByoipRangeResponseLifecycleStateEnum = "ACTIVE"
	ByoipRangeResponseLifecycleStateDeleting ByoipRangeResponseLifecycleStateEnum = "DELETING"
	ByoipRangeResponseLifecycleStateDeleted  ByoipRangeResponseLifecycleStateEnum = "DELETED"
)

var mappingByoipRangeResponseLifecycleStateEnum = map[string]ByoipRangeResponseLifecycleStateEnum{
	"INACTIVE": ByoipRangeResponseLifecycleStateInactive,
	"UPDATING": ByoipRangeResponseLifecycleStateUpdating,
	"ACTIVE":   ByoipRangeResponseLifecycleStateActive,
	"DELETING": ByoipRangeResponseLifecycleStateDeleting,
	"DELETED":  ByoipRangeResponseLifecycleStateDeleted,
}

var mappingByoipRangeResponseLifecycleStateEnumLowerCase = map[string]ByoipRangeResponseLifecycleStateEnum{
	"inactive": ByoipRangeResponseLifecycleStateInactive,
	"updating": ByoipRangeResponseLifecycleStateUpdating,
	"active":   ByoipRangeResponseLifecycleStateActive,
	"deleting": ByoipRangeResponseLifecycleStateDeleting,
	"deleted":  ByoipRangeResponseLifecycleStateDeleted,
}

// GetByoipRangeResponseLifecycleStateEnumValues Enumerates the set of values for ByoipRangeResponseLifecycleStateEnum
func GetByoipRangeResponseLifecycleStateEnumValues() []ByoipRangeResponseLifecycleStateEnum {
	values := make([]ByoipRangeResponseLifecycleStateEnum, 0)
	for _, v := range mappingByoipRangeResponseLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetByoipRangeResponseLifecycleStateEnumStringValues Enumerates the set of values in String for ByoipRangeResponseLifecycleStateEnum
func GetByoipRangeResponseLifecycleStateEnumStringValues() []string {
	return []string{
		"INACTIVE",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingByoipRangeResponseLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingByoipRangeResponseLifecycleStateEnum(val string) (ByoipRangeResponseLifecycleStateEnum, bool) {
	enum, ok := mappingByoipRangeResponseLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
