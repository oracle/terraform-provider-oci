// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Demand Signal API
//
// Use the OCI Control Center Demand Signal API to manage Demand Signals.
//

package demandsignal

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OccDemandSignal An OccDemandSignal is a forecast created for different Resource Types.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type OccDemandSignal struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OccDemandSignal.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OccDemandSignal data.
	OccDemandSignals []OccDemandSignalData `mandatory:"true" json:"occDemandSignals"`

	// Indicator of whether to share the data with Oracle.
	IsActive *bool `mandatory:"true" json:"isActive"`

	// The current state of the OccDemandSignal.
	LifecycleState OccDemandSignalLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the OccDemandSignal was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The date and time the OccDemandSignal was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the OccDemandSignal in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OccDemandSignal) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccDemandSignal) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOccDemandSignalLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOccDemandSignalLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OccDemandSignalLifecycleStateEnum Enum with underlying type: string
type OccDemandSignalLifecycleStateEnum string

// Set of constants representing the allowable values for OccDemandSignalLifecycleStateEnum
const (
	OccDemandSignalLifecycleStateCreating OccDemandSignalLifecycleStateEnum = "CREATING"
	OccDemandSignalLifecycleStateUpdating OccDemandSignalLifecycleStateEnum = "UPDATING"
	OccDemandSignalLifecycleStateActive   OccDemandSignalLifecycleStateEnum = "ACTIVE"
	OccDemandSignalLifecycleStateDeleting OccDemandSignalLifecycleStateEnum = "DELETING"
	OccDemandSignalLifecycleStateDeleted  OccDemandSignalLifecycleStateEnum = "DELETED"
	OccDemandSignalLifecycleStateFailed   OccDemandSignalLifecycleStateEnum = "FAILED"
)

var mappingOccDemandSignalLifecycleStateEnum = map[string]OccDemandSignalLifecycleStateEnum{
	"CREATING": OccDemandSignalLifecycleStateCreating,
	"UPDATING": OccDemandSignalLifecycleStateUpdating,
	"ACTIVE":   OccDemandSignalLifecycleStateActive,
	"DELETING": OccDemandSignalLifecycleStateDeleting,
	"DELETED":  OccDemandSignalLifecycleStateDeleted,
	"FAILED":   OccDemandSignalLifecycleStateFailed,
}

var mappingOccDemandSignalLifecycleStateEnumLowerCase = map[string]OccDemandSignalLifecycleStateEnum{
	"creating": OccDemandSignalLifecycleStateCreating,
	"updating": OccDemandSignalLifecycleStateUpdating,
	"active":   OccDemandSignalLifecycleStateActive,
	"deleting": OccDemandSignalLifecycleStateDeleting,
	"deleted":  OccDemandSignalLifecycleStateDeleted,
	"failed":   OccDemandSignalLifecycleStateFailed,
}

// GetOccDemandSignalLifecycleStateEnumValues Enumerates the set of values for OccDemandSignalLifecycleStateEnum
func GetOccDemandSignalLifecycleStateEnumValues() []OccDemandSignalLifecycleStateEnum {
	values := make([]OccDemandSignalLifecycleStateEnum, 0)
	for _, v := range mappingOccDemandSignalLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOccDemandSignalLifecycleStateEnumStringValues Enumerates the set of values in String for OccDemandSignalLifecycleStateEnum
func GetOccDemandSignalLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOccDemandSignalLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccDemandSignalLifecycleStateEnum(val string) (OccDemandSignalLifecycleStateEnum, bool) {
	enum, ok := mappingOccDemandSignalLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
