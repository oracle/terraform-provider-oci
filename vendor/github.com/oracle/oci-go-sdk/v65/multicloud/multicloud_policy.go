// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see Oracle Multicloud Hub (https://docs.oracle.com/iaas/Content/multicloud-hub/home.htm).
//

package multicloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MulticloudPolicy A missing IAM policy required for multicloud operation.
type MulticloudPolicy struct {

	// Compartment The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) where the policy is configured.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of the missing policy.
	Name *string `mandatory:"true" json:"name"`

	// IAM policy statements required.
	Statements []string `mandatory:"true" json:"statements"`

	// Description of the compartment e.g. Base Compartment, Root Compartment
	CompartmentName *string `mandatory:"false" json:"compartmentName"`

	// Description of the policy purpose.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the Multicloud Policy.
	LifecycleState MulticloudPolicyLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m MulticloudPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MulticloudPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMulticloudPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMulticloudPolicyLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MulticloudPolicyLifecycleStateEnum Enum with underlying type: string
type MulticloudPolicyLifecycleStateEnum string

// Set of constants representing the allowable values for MulticloudPolicyLifecycleStateEnum
const (
	MulticloudPolicyLifecycleStateCreating MulticloudPolicyLifecycleStateEnum = "CREATING"
	MulticloudPolicyLifecycleStateUpdating MulticloudPolicyLifecycleStateEnum = "UPDATING"
	MulticloudPolicyLifecycleStateActive   MulticloudPolicyLifecycleStateEnum = "ACTIVE"
	MulticloudPolicyLifecycleStateDeleting MulticloudPolicyLifecycleStateEnum = "DELETING"
	MulticloudPolicyLifecycleStateDeleted  MulticloudPolicyLifecycleStateEnum = "DELETED"
	MulticloudPolicyLifecycleStateFailed   MulticloudPolicyLifecycleStateEnum = "FAILED"
)

var mappingMulticloudPolicyLifecycleStateEnum = map[string]MulticloudPolicyLifecycleStateEnum{
	"CREATING": MulticloudPolicyLifecycleStateCreating,
	"UPDATING": MulticloudPolicyLifecycleStateUpdating,
	"ACTIVE":   MulticloudPolicyLifecycleStateActive,
	"DELETING": MulticloudPolicyLifecycleStateDeleting,
	"DELETED":  MulticloudPolicyLifecycleStateDeleted,
	"FAILED":   MulticloudPolicyLifecycleStateFailed,
}

var mappingMulticloudPolicyLifecycleStateEnumLowerCase = map[string]MulticloudPolicyLifecycleStateEnum{
	"creating": MulticloudPolicyLifecycleStateCreating,
	"updating": MulticloudPolicyLifecycleStateUpdating,
	"active":   MulticloudPolicyLifecycleStateActive,
	"deleting": MulticloudPolicyLifecycleStateDeleting,
	"deleted":  MulticloudPolicyLifecycleStateDeleted,
	"failed":   MulticloudPolicyLifecycleStateFailed,
}

// GetMulticloudPolicyLifecycleStateEnumValues Enumerates the set of values for MulticloudPolicyLifecycleStateEnum
func GetMulticloudPolicyLifecycleStateEnumValues() []MulticloudPolicyLifecycleStateEnum {
	values := make([]MulticloudPolicyLifecycleStateEnum, 0)
	for _, v := range mappingMulticloudPolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMulticloudPolicyLifecycleStateEnumStringValues Enumerates the set of values in String for MulticloudPolicyLifecycleStateEnum
func GetMulticloudPolicyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingMulticloudPolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMulticloudPolicyLifecycleStateEnum(val string) (MulticloudPolicyLifecycleStateEnum, bool) {
	enum, ok := mappingMulticloudPolicyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
