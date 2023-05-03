// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// ClusterNamespaceProfile Description of ClusterNamespaceProfile.
type ClusterNamespaceProfile struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Name of the cluster namespace.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID of compartment owning the cluster namespace.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Minimum Kubernetes version supported by the Cluster
	// Namespace Profile. Effectively the minimum version of
	// Kubernetes clusters attached to the Profile.
	KubernetesVersion *string `mandatory:"true" json:"kubernetesVersion"`

	// The time when this resource was created in an RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when this resource was updated in an RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the ClusterNamespaceProfile.
	LifecycleState ClusterNamespaceProfileLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// Description of the resource. It can be changed after creation.
	Description *string `mandatory:"false" json:"description"`

	// Suffix to append to the end of the namespaces generated from this Profile
	NamespaceSuffix *string `mandatory:"false" json:"namespaceSuffix"`

	// A name for the Cluster Namespace Profile. Names (when not null) are unique across Cluster Namespace Profile versions.
	VersionName *string `mandatory:"false" json:"versionName"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m ClusterNamespaceProfile) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ClusterNamespaceProfile) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingClusterNamespaceProfileLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetClusterNamespaceProfileLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ClusterNamespaceProfileLifecycleStateEnum Enum with underlying type: string
type ClusterNamespaceProfileLifecycleStateEnum string

// Set of constants representing the allowable values for ClusterNamespaceProfileLifecycleStateEnum
const (
	ClusterNamespaceProfileLifecycleStateCreating ClusterNamespaceProfileLifecycleStateEnum = "CREATING"
	ClusterNamespaceProfileLifecycleStateUpdating ClusterNamespaceProfileLifecycleStateEnum = "UPDATING"
	ClusterNamespaceProfileLifecycleStateActive   ClusterNamespaceProfileLifecycleStateEnum = "ACTIVE"
	ClusterNamespaceProfileLifecycleStateDeleting ClusterNamespaceProfileLifecycleStateEnum = "DELETING"
	ClusterNamespaceProfileLifecycleStateDeleted  ClusterNamespaceProfileLifecycleStateEnum = "DELETED"
	ClusterNamespaceProfileLifecycleStateFailed   ClusterNamespaceProfileLifecycleStateEnum = "FAILED"
)

var mappingClusterNamespaceProfileLifecycleStateEnum = map[string]ClusterNamespaceProfileLifecycleStateEnum{
	"CREATING": ClusterNamespaceProfileLifecycleStateCreating,
	"UPDATING": ClusterNamespaceProfileLifecycleStateUpdating,
	"ACTIVE":   ClusterNamespaceProfileLifecycleStateActive,
	"DELETING": ClusterNamespaceProfileLifecycleStateDeleting,
	"DELETED":  ClusterNamespaceProfileLifecycleStateDeleted,
	"FAILED":   ClusterNamespaceProfileLifecycleStateFailed,
}

var mappingClusterNamespaceProfileLifecycleStateEnumLowerCase = map[string]ClusterNamespaceProfileLifecycleStateEnum{
	"creating": ClusterNamespaceProfileLifecycleStateCreating,
	"updating": ClusterNamespaceProfileLifecycleStateUpdating,
	"active":   ClusterNamespaceProfileLifecycleStateActive,
	"deleting": ClusterNamespaceProfileLifecycleStateDeleting,
	"deleted":  ClusterNamespaceProfileLifecycleStateDeleted,
	"failed":   ClusterNamespaceProfileLifecycleStateFailed,
}

// GetClusterNamespaceProfileLifecycleStateEnumValues Enumerates the set of values for ClusterNamespaceProfileLifecycleStateEnum
func GetClusterNamespaceProfileLifecycleStateEnumValues() []ClusterNamespaceProfileLifecycleStateEnum {
	values := make([]ClusterNamespaceProfileLifecycleStateEnum, 0)
	for _, v := range mappingClusterNamespaceProfileLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetClusterNamespaceProfileLifecycleStateEnumStringValues Enumerates the set of values in String for ClusterNamespaceProfileLifecycleStateEnum
func GetClusterNamespaceProfileLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingClusterNamespaceProfileLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClusterNamespaceProfileLifecycleStateEnum(val string) (ClusterNamespaceProfileLifecycleStateEnum, bool) {
	enum, ok := mappingClusterNamespaceProfileLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
