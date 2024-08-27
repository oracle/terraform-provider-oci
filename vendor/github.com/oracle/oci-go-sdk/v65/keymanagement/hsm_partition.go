// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Key Management API
//
// Use the Key Management API to manage vaults and keys. For more information, see Managing Vaults (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingvaults.htm) and Managing Keys (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingkeys.htm).
//

package keymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HsmPartition Dedicated KMS-HSM Partition Management
type HsmPartition struct {

	// The OCID of the HSM resource. Each HSM resource has a unique OCID as an identifier.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains a particular HSM resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Details of a single portInformation item include the PortNumber (an integer used as an identifier) and the PortType (refers to either an enum value of Managementutility,Clientutility, or null)
	PortInformation []PortInformation `mandatory:"true" json:"portInformation"`

	// The date and time a HSMPartition was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time a HSMPartition was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2018-04-03T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// A HSMCluster resource's current lifecycle state.
	// Example: `ACTIVE`
	LifecycleState HsmPartitionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

func (m HsmPartition) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HsmPartition) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHsmPartitionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetHsmPartitionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HsmPartitionLifecycleStateEnum Enum with underlying type: string
type HsmPartitionLifecycleStateEnum string

// Set of constants representing the allowable values for HsmPartitionLifecycleStateEnum
const (
	HsmPartitionLifecycleStateActive             HsmPartitionLifecycleStateEnum = "ACTIVE"
	HsmPartitionLifecycleStateInactive           HsmPartitionLifecycleStateEnum = "INACTIVE"
	HsmPartitionLifecycleStateActivating         HsmPartitionLifecycleStateEnum = "ACTIVATING"
	HsmPartitionLifecycleStateActivationRequired HsmPartitionLifecycleStateEnum = "ACTIVATION_REQUIRED"
	HsmPartitionLifecycleStateSchedulingDeletion HsmPartitionLifecycleStateEnum = "SCHEDULING_DELETION"
	HsmPartitionLifecycleStatePendingDeletion    HsmPartitionLifecycleStateEnum = "PENDING_DELETION"
	HsmPartitionLifecycleStateDeleting           HsmPartitionLifecycleStateEnum = "DELETING"
	HsmPartitionLifecycleStateDeleted            HsmPartitionLifecycleStateEnum = "DELETED"
)

var mappingHsmPartitionLifecycleStateEnum = map[string]HsmPartitionLifecycleStateEnum{
	"ACTIVE":              HsmPartitionLifecycleStateActive,
	"INACTIVE":            HsmPartitionLifecycleStateInactive,
	"ACTIVATING":          HsmPartitionLifecycleStateActivating,
	"ACTIVATION_REQUIRED": HsmPartitionLifecycleStateActivationRequired,
	"SCHEDULING_DELETION": HsmPartitionLifecycleStateSchedulingDeletion,
	"PENDING_DELETION":    HsmPartitionLifecycleStatePendingDeletion,
	"DELETING":            HsmPartitionLifecycleStateDeleting,
	"DELETED":             HsmPartitionLifecycleStateDeleted,
}

var mappingHsmPartitionLifecycleStateEnumLowerCase = map[string]HsmPartitionLifecycleStateEnum{
	"active":              HsmPartitionLifecycleStateActive,
	"inactive":            HsmPartitionLifecycleStateInactive,
	"activating":          HsmPartitionLifecycleStateActivating,
	"activation_required": HsmPartitionLifecycleStateActivationRequired,
	"scheduling_deletion": HsmPartitionLifecycleStateSchedulingDeletion,
	"pending_deletion":    HsmPartitionLifecycleStatePendingDeletion,
	"deleting":            HsmPartitionLifecycleStateDeleting,
	"deleted":             HsmPartitionLifecycleStateDeleted,
}

// GetHsmPartitionLifecycleStateEnumValues Enumerates the set of values for HsmPartitionLifecycleStateEnum
func GetHsmPartitionLifecycleStateEnumValues() []HsmPartitionLifecycleStateEnum {
	values := make([]HsmPartitionLifecycleStateEnum, 0)
	for _, v := range mappingHsmPartitionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetHsmPartitionLifecycleStateEnumStringValues Enumerates the set of values in String for HsmPartitionLifecycleStateEnum
func GetHsmPartitionLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"ACTIVATING",
		"ACTIVATION_REQUIRED",
		"SCHEDULING_DELETION",
		"PENDING_DELETION",
		"DELETING",
		"DELETED",
	}
}

// GetMappingHsmPartitionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHsmPartitionLifecycleStateEnum(val string) (HsmPartitionLifecycleStateEnum, bool) {
	enum, ok := mappingHsmPartitionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
