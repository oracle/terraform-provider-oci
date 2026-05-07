// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RegistrationPolicy A registration policy.
// This object contains detailed information about a registration policy, including its ID, compartment ID, display name, and features.
type RegistrationPolicy struct {

	// The OCID for the compartment containing the registration policy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the registration policy.
	Id *string `mandatory:"true" json:"id"`

	// The display name of the registration policy.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the resource used in the registration policy.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// The Data Safe features granted to the databases registering under the registration policy.
	Features []RegistrationPolicyFeaturesEnum `mandatory:"true" json:"features"`

	// The lifecycle state of the registration policy.
	// - CREATING - The registration policy is getting created.
	// - ACTIVE  - The registration policy has been successfully created.
	// - UPDATING - The registration policy is getting updated.
	// - NEEDS_ATTENTION - The registration policy needs attention.
	// - FAILED - The registration policy creation or update has failed.
	// - DELETING - The registration policy is getting deleted.
	LifecycleState RegistrationPolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the registration policy was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time when the registration policy was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// A description of the registration policy.
	Description *string `mandatory:"false" json:"description"`

	// The resource type which has been opted in for the registration policy.
	// - CLOUD_VMCLUSTER  - The Registration policy is applied at the ExaCS VM Cluster level, enabling opt-in for this resource
	// - VMCLUSTER  - The registration policy will be opted in at Exadata Cloud@Customer instances.
	// - EXADB_VMCLUSTER - The registration policy will be opted in at Exadata VM cluster on Exascale Infrastructure
	// - DATABASE - The registration policy will be opted in at the Container Database level.
	ResourceType RegistrationPolicyResourceTypeEnum `mandatory:"false" json:"resourceType,omitempty"`

	// Indicates whether features will be overridden.
	CanOverrideFeatures *bool `mandatory:"false" json:"canOverrideFeatures"`

	// Details about the lifecycle state of the registration policy
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	ConnectionOption *PolicyConnectionOption `mandatory:"false" json:"connectionOption"`
}

func (m RegistrationPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RegistrationPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.Features {
		if _, ok := GetMappingRegistrationPolicyFeaturesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Features: %s. Supported values are: %s.", val, strings.Join(GetRegistrationPolicyFeaturesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingRegistrationPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRegistrationPolicyLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingRegistrationPolicyResourceTypeEnum(string(m.ResourceType)); !ok && m.ResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceType: %s. Supported values are: %s.", m.ResourceType, strings.Join(GetRegistrationPolicyResourceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RegistrationPolicyResourceTypeEnum Enum with underlying type: string
type RegistrationPolicyResourceTypeEnum string

// Set of constants representing the allowable values for RegistrationPolicyResourceTypeEnum
const (
	RegistrationPolicyResourceTypeCloudVmCluster RegistrationPolicyResourceTypeEnum = "CLOUD_VM_CLUSTER"
	RegistrationPolicyResourceTypeVmCluster      RegistrationPolicyResourceTypeEnum = "VM_CLUSTER"
	RegistrationPolicyResourceTypeExadbVmCluster RegistrationPolicyResourceTypeEnum = "EXADB_VM_CLUSTER"
	RegistrationPolicyResourceTypeDatabase       RegistrationPolicyResourceTypeEnum = "DATABASE"
)

var mappingRegistrationPolicyResourceTypeEnum = map[string]RegistrationPolicyResourceTypeEnum{
	"CLOUD_VM_CLUSTER": RegistrationPolicyResourceTypeCloudVmCluster,
	"VM_CLUSTER":       RegistrationPolicyResourceTypeVmCluster,
	"EXADB_VM_CLUSTER": RegistrationPolicyResourceTypeExadbVmCluster,
	"DATABASE":         RegistrationPolicyResourceTypeDatabase,
}

var mappingRegistrationPolicyResourceTypeEnumLowerCase = map[string]RegistrationPolicyResourceTypeEnum{
	"cloud_vm_cluster": RegistrationPolicyResourceTypeCloudVmCluster,
	"vm_cluster":       RegistrationPolicyResourceTypeVmCluster,
	"exadb_vm_cluster": RegistrationPolicyResourceTypeExadbVmCluster,
	"database":         RegistrationPolicyResourceTypeDatabase,
}

// GetRegistrationPolicyResourceTypeEnumValues Enumerates the set of values for RegistrationPolicyResourceTypeEnum
func GetRegistrationPolicyResourceTypeEnumValues() []RegistrationPolicyResourceTypeEnum {
	values := make([]RegistrationPolicyResourceTypeEnum, 0)
	for _, v := range mappingRegistrationPolicyResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRegistrationPolicyResourceTypeEnumStringValues Enumerates the set of values in String for RegistrationPolicyResourceTypeEnum
func GetRegistrationPolicyResourceTypeEnumStringValues() []string {
	return []string{
		"CLOUD_VM_CLUSTER",
		"VM_CLUSTER",
		"EXADB_VM_CLUSTER",
		"DATABASE",
	}
}

// GetMappingRegistrationPolicyResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRegistrationPolicyResourceTypeEnum(val string) (RegistrationPolicyResourceTypeEnum, bool) {
	enum, ok := mappingRegistrationPolicyResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RegistrationPolicyLifecycleStateEnum Enum with underlying type: string
type RegistrationPolicyLifecycleStateEnum string

// Set of constants representing the allowable values for RegistrationPolicyLifecycleStateEnum
const (
	RegistrationPolicyLifecycleStateCreating       RegistrationPolicyLifecycleStateEnum = "CREATING"
	RegistrationPolicyLifecycleStateActive         RegistrationPolicyLifecycleStateEnum = "ACTIVE"
	RegistrationPolicyLifecycleStateUpdating       RegistrationPolicyLifecycleStateEnum = "UPDATING"
	RegistrationPolicyLifecycleStateNeedsAttention RegistrationPolicyLifecycleStateEnum = "NEEDS_ATTENTION"
	RegistrationPolicyLifecycleStateFailed         RegistrationPolicyLifecycleStateEnum = "FAILED"
	RegistrationPolicyLifecycleStateDeleting       RegistrationPolicyLifecycleStateEnum = "DELETING"
)

var mappingRegistrationPolicyLifecycleStateEnum = map[string]RegistrationPolicyLifecycleStateEnum{
	"CREATING":        RegistrationPolicyLifecycleStateCreating,
	"ACTIVE":          RegistrationPolicyLifecycleStateActive,
	"UPDATING":        RegistrationPolicyLifecycleStateUpdating,
	"NEEDS_ATTENTION": RegistrationPolicyLifecycleStateNeedsAttention,
	"FAILED":          RegistrationPolicyLifecycleStateFailed,
	"DELETING":        RegistrationPolicyLifecycleStateDeleting,
}

var mappingRegistrationPolicyLifecycleStateEnumLowerCase = map[string]RegistrationPolicyLifecycleStateEnum{
	"creating":        RegistrationPolicyLifecycleStateCreating,
	"active":          RegistrationPolicyLifecycleStateActive,
	"updating":        RegistrationPolicyLifecycleStateUpdating,
	"needs_attention": RegistrationPolicyLifecycleStateNeedsAttention,
	"failed":          RegistrationPolicyLifecycleStateFailed,
	"deleting":        RegistrationPolicyLifecycleStateDeleting,
}

// GetRegistrationPolicyLifecycleStateEnumValues Enumerates the set of values for RegistrationPolicyLifecycleStateEnum
func GetRegistrationPolicyLifecycleStateEnumValues() []RegistrationPolicyLifecycleStateEnum {
	values := make([]RegistrationPolicyLifecycleStateEnum, 0)
	for _, v := range mappingRegistrationPolicyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetRegistrationPolicyLifecycleStateEnumStringValues Enumerates the set of values in String for RegistrationPolicyLifecycleStateEnum
func GetRegistrationPolicyLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"NEEDS_ATTENTION",
		"FAILED",
		"DELETING",
	}
}

// GetMappingRegistrationPolicyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRegistrationPolicyLifecycleStateEnum(val string) (RegistrationPolicyLifecycleStateEnum, bool) {
	enum, ok := mappingRegistrationPolicyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
