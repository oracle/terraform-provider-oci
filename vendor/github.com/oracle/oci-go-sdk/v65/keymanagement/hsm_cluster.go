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

// HsmCluster Dedicated KMS-HSM Cluster Management
type HsmCluster struct {

	// The OCID of the HSMCluster resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains this HSMCluster resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name for the HSMCluster resource. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time this HSM resource was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2023-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time this HSM resource was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2023-04-03T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The HSMCluster's current state.
	// Example: `ACTIVE`
	LifecycleState HsmClusterLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// DNS name for the Hsm Cluster.
	DnsName *string `mandatory:"true" json:"dnsName"`

	// An optional property indicating when to delete the key, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m HsmCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HsmCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHsmClusterLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetHsmClusterLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HsmClusterLifecycleStateEnum Enum with underlying type: string
type HsmClusterLifecycleStateEnum string

// Set of constants representing the allowable values for HsmClusterLifecycleStateEnum
const (
	HsmClusterLifecycleStateCreating               HsmClusterLifecycleStateEnum = "CREATING"
	HsmClusterLifecycleStateInitializationRequired HsmClusterLifecycleStateEnum = "INITIALIZATION_REQUIRED"
	HsmClusterLifecycleStateInitializing           HsmClusterLifecycleStateEnum = "INITIALIZING"
	HsmClusterLifecycleStateActivationRequired     HsmClusterLifecycleStateEnum = "ACTIVATION_REQUIRED"
	HsmClusterLifecycleStateActivating             HsmClusterLifecycleStateEnum = "ACTIVATING"
	HsmClusterLifecycleStateActive                 HsmClusterLifecycleStateEnum = "ACTIVE"
	HsmClusterLifecycleStateDeleting               HsmClusterLifecycleStateEnum = "DELETING"
	HsmClusterLifecycleStateDeleted                HsmClusterLifecycleStateEnum = "DELETED"
	HsmClusterLifecycleStatePendingDeletion        HsmClusterLifecycleStateEnum = "PENDING_DELETION"
	HsmClusterLifecycleStateSchedulingDeletion     HsmClusterLifecycleStateEnum = "SCHEDULING_DELETION"
	HsmClusterLifecycleStateCancellingDeletion     HsmClusterLifecycleStateEnum = "CANCELLING_DELETION"
)

var mappingHsmClusterLifecycleStateEnum = map[string]HsmClusterLifecycleStateEnum{
	"CREATING":                HsmClusterLifecycleStateCreating,
	"INITIALIZATION_REQUIRED": HsmClusterLifecycleStateInitializationRequired,
	"INITIALIZING":            HsmClusterLifecycleStateInitializing,
	"ACTIVATION_REQUIRED":     HsmClusterLifecycleStateActivationRequired,
	"ACTIVATING":              HsmClusterLifecycleStateActivating,
	"ACTIVE":                  HsmClusterLifecycleStateActive,
	"DELETING":                HsmClusterLifecycleStateDeleting,
	"DELETED":                 HsmClusterLifecycleStateDeleted,
	"PENDING_DELETION":        HsmClusterLifecycleStatePendingDeletion,
	"SCHEDULING_DELETION":     HsmClusterLifecycleStateSchedulingDeletion,
	"CANCELLING_DELETION":     HsmClusterLifecycleStateCancellingDeletion,
}

var mappingHsmClusterLifecycleStateEnumLowerCase = map[string]HsmClusterLifecycleStateEnum{
	"creating":                HsmClusterLifecycleStateCreating,
	"initialization_required": HsmClusterLifecycleStateInitializationRequired,
	"initializing":            HsmClusterLifecycleStateInitializing,
	"activation_required":     HsmClusterLifecycleStateActivationRequired,
	"activating":              HsmClusterLifecycleStateActivating,
	"active":                  HsmClusterLifecycleStateActive,
	"deleting":                HsmClusterLifecycleStateDeleting,
	"deleted":                 HsmClusterLifecycleStateDeleted,
	"pending_deletion":        HsmClusterLifecycleStatePendingDeletion,
	"scheduling_deletion":     HsmClusterLifecycleStateSchedulingDeletion,
	"cancelling_deletion":     HsmClusterLifecycleStateCancellingDeletion,
}

// GetHsmClusterLifecycleStateEnumValues Enumerates the set of values for HsmClusterLifecycleStateEnum
func GetHsmClusterLifecycleStateEnumValues() []HsmClusterLifecycleStateEnum {
	values := make([]HsmClusterLifecycleStateEnum, 0)
	for _, v := range mappingHsmClusterLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetHsmClusterLifecycleStateEnumStringValues Enumerates the set of values in String for HsmClusterLifecycleStateEnum
func GetHsmClusterLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"INITIALIZATION_REQUIRED",
		"INITIALIZING",
		"ACTIVATION_REQUIRED",
		"ACTIVATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"PENDING_DELETION",
		"SCHEDULING_DELETION",
		"CANCELLING_DELETION",
	}
}

// GetMappingHsmClusterLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHsmClusterLifecycleStateEnum(val string) (HsmClusterLifecycleStateEnum, bool) {
	enum, ok := mappingHsmClusterLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
