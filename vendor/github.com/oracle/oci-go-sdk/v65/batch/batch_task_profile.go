// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BatchTaskProfile A batch task profile resource describes the minimum hardware requirements requested for a task.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type BatchTaskProfile struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the batch task profile.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The minimum required OCPUs.
	MinOcpus *int `mandatory:"true" json:"minOcpus"`

	// The minimum required memory.
	MinMemoryInGBs *int `mandatory:"true" json:"minMemoryInGBs"`

	// The current state of the batch task profile.
	LifecycleState BatchTaskProfileLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the batch task profile was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// The batch task profile description.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the batch task profile was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m BatchTaskProfile) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BatchTaskProfile) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBatchTaskProfileLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBatchTaskProfileLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BatchTaskProfileLifecycleStateEnum Enum with underlying type: string
type BatchTaskProfileLifecycleStateEnum string

// Set of constants representing the allowable values for BatchTaskProfileLifecycleStateEnum
const (
	BatchTaskProfileLifecycleStateActive  BatchTaskProfileLifecycleStateEnum = "ACTIVE"
	BatchTaskProfileLifecycleStateDeleted BatchTaskProfileLifecycleStateEnum = "DELETED"
)

var mappingBatchTaskProfileLifecycleStateEnum = map[string]BatchTaskProfileLifecycleStateEnum{
	"ACTIVE":  BatchTaskProfileLifecycleStateActive,
	"DELETED": BatchTaskProfileLifecycleStateDeleted,
}

var mappingBatchTaskProfileLifecycleStateEnumLowerCase = map[string]BatchTaskProfileLifecycleStateEnum{
	"active":  BatchTaskProfileLifecycleStateActive,
	"deleted": BatchTaskProfileLifecycleStateDeleted,
}

// GetBatchTaskProfileLifecycleStateEnumValues Enumerates the set of values for BatchTaskProfileLifecycleStateEnum
func GetBatchTaskProfileLifecycleStateEnumValues() []BatchTaskProfileLifecycleStateEnum {
	values := make([]BatchTaskProfileLifecycleStateEnum, 0)
	for _, v := range mappingBatchTaskProfileLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBatchTaskProfileLifecycleStateEnumStringValues Enumerates the set of values in String for BatchTaskProfileLifecycleStateEnum
func GetBatchTaskProfileLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingBatchTaskProfileLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBatchTaskProfileLifecycleStateEnum(val string) (BatchTaskProfileLifecycleStateEnum, bool) {
	enum, ok := mappingBatchTaskProfileLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
