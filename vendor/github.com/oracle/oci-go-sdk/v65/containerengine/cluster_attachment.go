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

// ClusterAttachment Description of ClusterAttachment.
type ClusterAttachment struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// Name of the Cluster Namespace.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID of compartment owning the Cluster Namespace.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCID of the Cluster
	ClusterId *string `mandatory:"true" json:"clusterId"`

	// OCID of the Cluster Namespace Profile
	ClusterNamespaceProfileId *string `mandatory:"true" json:"clusterNamespaceProfileId"`

	// The time when this resource was created in an RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the ClusterAttachment.
	LifecycleState ClusterAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

	// The time when this resource was updated in an RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m ClusterAttachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ClusterAttachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingClusterAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetClusterAttachmentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ClusterAttachmentLifecycleStateEnum Enum with underlying type: string
type ClusterAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for ClusterAttachmentLifecycleStateEnum
const (
	ClusterAttachmentLifecycleStateCreating       ClusterAttachmentLifecycleStateEnum = "CREATING"
	ClusterAttachmentLifecycleStateUpdating       ClusterAttachmentLifecycleStateEnum = "UPDATING"
	ClusterAttachmentLifecycleStateNeedsAttention ClusterAttachmentLifecycleStateEnum = "NEEDS_ATTENTION"
	ClusterAttachmentLifecycleStateActive         ClusterAttachmentLifecycleStateEnum = "ACTIVE"
	ClusterAttachmentLifecycleStateDeleting       ClusterAttachmentLifecycleStateEnum = "DELETING"
	ClusterAttachmentLifecycleStateDeleted        ClusterAttachmentLifecycleStateEnum = "DELETED"
	ClusterAttachmentLifecycleStateFailed         ClusterAttachmentLifecycleStateEnum = "FAILED"
)

var mappingClusterAttachmentLifecycleStateEnum = map[string]ClusterAttachmentLifecycleStateEnum{
	"CREATING":        ClusterAttachmentLifecycleStateCreating,
	"UPDATING":        ClusterAttachmentLifecycleStateUpdating,
	"NEEDS_ATTENTION": ClusterAttachmentLifecycleStateNeedsAttention,
	"ACTIVE":          ClusterAttachmentLifecycleStateActive,
	"DELETING":        ClusterAttachmentLifecycleStateDeleting,
	"DELETED":         ClusterAttachmentLifecycleStateDeleted,
	"FAILED":          ClusterAttachmentLifecycleStateFailed,
}

var mappingClusterAttachmentLifecycleStateEnumLowerCase = map[string]ClusterAttachmentLifecycleStateEnum{
	"creating":        ClusterAttachmentLifecycleStateCreating,
	"updating":        ClusterAttachmentLifecycleStateUpdating,
	"needs_attention": ClusterAttachmentLifecycleStateNeedsAttention,
	"active":          ClusterAttachmentLifecycleStateActive,
	"deleting":        ClusterAttachmentLifecycleStateDeleting,
	"deleted":         ClusterAttachmentLifecycleStateDeleted,
	"failed":          ClusterAttachmentLifecycleStateFailed,
}

// GetClusterAttachmentLifecycleStateEnumValues Enumerates the set of values for ClusterAttachmentLifecycleStateEnum
func GetClusterAttachmentLifecycleStateEnumValues() []ClusterAttachmentLifecycleStateEnum {
	values := make([]ClusterAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingClusterAttachmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetClusterAttachmentLifecycleStateEnumStringValues Enumerates the set of values in String for ClusterAttachmentLifecycleStateEnum
func GetClusterAttachmentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"NEEDS_ATTENTION",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingClusterAttachmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClusterAttachmentLifecycleStateEnum(val string) (ClusterAttachmentLifecycleStateEnum, bool) {
	enum, ok := mappingClusterAttachmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
