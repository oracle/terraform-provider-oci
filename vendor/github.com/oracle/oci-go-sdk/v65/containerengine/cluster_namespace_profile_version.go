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

// ClusterNamespaceProfileVersion A version of a Cluster Namespace Profile.
type ClusterNamespaceProfileVersion struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// A name for the Cluster Namespace Profile Version. Names are unique across versions in a Cluster Namespace Profile Profiles.
	Name *string `mandatory:"true" json:"name"`

	// OCID of compartment owning the Cluster Namespace Profile Version.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the ClusterNamespaceProfile
	ClusterNamespaceProfileId *string `mandatory:"true" json:"clusterNamespaceProfileId"`

	// Name of the ClusterRole to bind to the admin account in the resulting namespace.
	AdminClusterRoleName *string `mandatory:"true" json:"adminClusterRoleName"`

	// The time when this resource was created in an RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when this resource was updated in an RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the resource.
	LifecycleState ClusterNamespaceProfileVersionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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
	FixedNamespaceLabels []NamespaceLabel `mandatory:"false" json:"fixedNamespaceLabels"`

	// List of Kubernetes annotations to apply to the resulting namespace.
	FixedNamespaceAnnotations []NamespaceAnnotation `mandatory:"false" json:"fixedNamespaceAnnotations"`

	// List of Kubernetes labels that may be specified via Cluster Namespaces.
	AllowedNamespaceLabels []AllowedNamespaceLabel `mandatory:"false" json:"allowedNamespaceLabels"`

	// List of Kubernetes annotations that may be specified via Cluster Namespaces.
	AllowedNamespaceAnnotations []AllowedNamespaceAnnotation `mandatory:"false" json:"allowedNamespaceAnnotations"`

	// List of Kubernetes labels that must be specified via Cluster Namespaces.
	RequiredNamespaceLabels []RequiredNamespaceLabel `mandatory:"false" json:"requiredNamespaceLabels"`

	// List of Kubernetes annotations that must be specified via Cluster Namespaces.
	RequiredNamespaceAnnotations []RequiredNamespaceAnnotation `mandatory:"false" json:"requiredNamespaceAnnotations"`

	// If set to true, the Cluster Namespace Profile Version is not consumable by new Cluster Namespace configurations.
	IsDeprecated *bool `mandatory:"false" json:"isDeprecated"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m ClusterNamespaceProfileVersion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ClusterNamespaceProfileVersion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingClusterNamespaceProfileVersionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetClusterNamespaceProfileVersionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ClusterNamespaceProfileVersionLifecycleStateEnum Enum with underlying type: string
type ClusterNamespaceProfileVersionLifecycleStateEnum string

// Set of constants representing the allowable values for ClusterNamespaceProfileVersionLifecycleStateEnum
const (
	ClusterNamespaceProfileVersionLifecycleStateCreating ClusterNamespaceProfileVersionLifecycleStateEnum = "CREATING"
	ClusterNamespaceProfileVersionLifecycleStateUpdating ClusterNamespaceProfileVersionLifecycleStateEnum = "UPDATING"
	ClusterNamespaceProfileVersionLifecycleStateActive   ClusterNamespaceProfileVersionLifecycleStateEnum = "ACTIVE"
	ClusterNamespaceProfileVersionLifecycleStateDeleting ClusterNamespaceProfileVersionLifecycleStateEnum = "DELETING"
	ClusterNamespaceProfileVersionLifecycleStateDeleted  ClusterNamespaceProfileVersionLifecycleStateEnum = "DELETED"
	ClusterNamespaceProfileVersionLifecycleStateFailed   ClusterNamespaceProfileVersionLifecycleStateEnum = "FAILED"
)

var mappingClusterNamespaceProfileVersionLifecycleStateEnum = map[string]ClusterNamespaceProfileVersionLifecycleStateEnum{
	"CREATING": ClusterNamespaceProfileVersionLifecycleStateCreating,
	"UPDATING": ClusterNamespaceProfileVersionLifecycleStateUpdating,
	"ACTIVE":   ClusterNamespaceProfileVersionLifecycleStateActive,
	"DELETING": ClusterNamespaceProfileVersionLifecycleStateDeleting,
	"DELETED":  ClusterNamespaceProfileVersionLifecycleStateDeleted,
	"FAILED":   ClusterNamespaceProfileVersionLifecycleStateFailed,
}

var mappingClusterNamespaceProfileVersionLifecycleStateEnumLowerCase = map[string]ClusterNamespaceProfileVersionLifecycleStateEnum{
	"creating": ClusterNamespaceProfileVersionLifecycleStateCreating,
	"updating": ClusterNamespaceProfileVersionLifecycleStateUpdating,
	"active":   ClusterNamespaceProfileVersionLifecycleStateActive,
	"deleting": ClusterNamespaceProfileVersionLifecycleStateDeleting,
	"deleted":  ClusterNamespaceProfileVersionLifecycleStateDeleted,
	"failed":   ClusterNamespaceProfileVersionLifecycleStateFailed,
}

// GetClusterNamespaceProfileVersionLifecycleStateEnumValues Enumerates the set of values for ClusterNamespaceProfileVersionLifecycleStateEnum
func GetClusterNamespaceProfileVersionLifecycleStateEnumValues() []ClusterNamespaceProfileVersionLifecycleStateEnum {
	values := make([]ClusterNamespaceProfileVersionLifecycleStateEnum, 0)
	for _, v := range mappingClusterNamespaceProfileVersionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetClusterNamespaceProfileVersionLifecycleStateEnumStringValues Enumerates the set of values in String for ClusterNamespaceProfileVersionLifecycleStateEnum
func GetClusterNamespaceProfileVersionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingClusterNamespaceProfileVersionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClusterNamespaceProfileVersionLifecycleStateEnum(val string) (ClusterNamespaceProfileVersionLifecycleStateEnum, bool) {
	enum, ok := mappingClusterNamespaceProfileVersionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
