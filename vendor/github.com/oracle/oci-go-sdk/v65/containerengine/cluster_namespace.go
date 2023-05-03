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

// ClusterNamespace Description of ClusterNamespace.
type ClusterNamespace struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Name of the cluster namespace.
	Name *string `mandatory:"true" json:"name"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// List of IAM Group OCIDs to allow admin access
	// within the cluster namespace.
	AdminGroupIds []string `mandatory:"true" json:"adminGroupIds"`

	// OCID of Cluster Namespace Profile Version to use.
	ClusterNamespaceProfileVersionId *string `mandatory:"true" json:"clusterNamespaceProfileVersionId"`

	// List of OKE Cluster OCIDs the Cluster Namespace is provisioned upon
	ClusterIds []string `mandatory:"true" json:"clusterIds"`

	// Name of the resulting Kubernetes namespace
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	// The time when this resource was created in an RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when this resource was updated in an RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the ClusterNamespace.
	LifecycleState ClusterNamespaceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

	// List of Kubernetes labels to apply to the resulting namespace.
	NamespaceLabels []NamespaceLabel `mandatory:"false" json:"namespaceLabels"`

	// List of Kubernetes annotations to apply to the resulting namespace.
	NamespaceAnnotations []NamespaceAnnotation `mandatory:"false" json:"namespaceAnnotations"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m ClusterNamespace) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ClusterNamespace) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingClusterNamespaceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetClusterNamespaceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ClusterNamespaceLifecycleStateEnum Enum with underlying type: string
type ClusterNamespaceLifecycleStateEnum string

// Set of constants representing the allowable values for ClusterNamespaceLifecycleStateEnum
const (
	ClusterNamespaceLifecycleStateCreating ClusterNamespaceLifecycleStateEnum = "CREATING"
	ClusterNamespaceLifecycleStateUpdating ClusterNamespaceLifecycleStateEnum = "UPDATING"
	ClusterNamespaceLifecycleStateActive   ClusterNamespaceLifecycleStateEnum = "ACTIVE"
	ClusterNamespaceLifecycleStateDeleting ClusterNamespaceLifecycleStateEnum = "DELETING"
	ClusterNamespaceLifecycleStateDeleted  ClusterNamespaceLifecycleStateEnum = "DELETED"
	ClusterNamespaceLifecycleStateFailed   ClusterNamespaceLifecycleStateEnum = "FAILED"
)

var mappingClusterNamespaceLifecycleStateEnum = map[string]ClusterNamespaceLifecycleStateEnum{
	"CREATING": ClusterNamespaceLifecycleStateCreating,
	"UPDATING": ClusterNamespaceLifecycleStateUpdating,
	"ACTIVE":   ClusterNamespaceLifecycleStateActive,
	"DELETING": ClusterNamespaceLifecycleStateDeleting,
	"DELETED":  ClusterNamespaceLifecycleStateDeleted,
	"FAILED":   ClusterNamespaceLifecycleStateFailed,
}

var mappingClusterNamespaceLifecycleStateEnumLowerCase = map[string]ClusterNamespaceLifecycleStateEnum{
	"creating": ClusterNamespaceLifecycleStateCreating,
	"updating": ClusterNamespaceLifecycleStateUpdating,
	"active":   ClusterNamespaceLifecycleStateActive,
	"deleting": ClusterNamespaceLifecycleStateDeleting,
	"deleted":  ClusterNamespaceLifecycleStateDeleted,
	"failed":   ClusterNamespaceLifecycleStateFailed,
}

// GetClusterNamespaceLifecycleStateEnumValues Enumerates the set of values for ClusterNamespaceLifecycleStateEnum
func GetClusterNamespaceLifecycleStateEnumValues() []ClusterNamespaceLifecycleStateEnum {
	values := make([]ClusterNamespaceLifecycleStateEnum, 0)
	for _, v := range mappingClusterNamespaceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetClusterNamespaceLifecycleStateEnumStringValues Enumerates the set of values in String for ClusterNamespaceLifecycleStateEnum
func GetClusterNamespaceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingClusterNamespaceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClusterNamespaceLifecycleStateEnum(val string) (ClusterNamespaceLifecycleStateEnum, bool) {
	enum, ok := mappingClusterNamespaceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
