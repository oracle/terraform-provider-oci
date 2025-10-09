// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Kubernetes Engine API
//
// API for the Kubernetes Engine service (also known as the Container Engine for Kubernetes service). Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Kubernetes Engine (https://docs.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"encoding/json"
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
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"true" json:"systemTags"`

	// Description of the resource. It can be changed after creation.
	Description *string `mandatory:"false" json:"description"`

	// Suffix to append to the end of the namespaces generated from this Profile
	NamespaceSuffix *string `mandatory:"false" json:"namespaceSuffix"`

	ClusterNamespaceSchedulingPolicy ClusterNamespaceSchedulingPolicy `mandatory:"false" json:"clusterNamespaceSchedulingPolicy"`

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
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ClusterNamespaceProfile) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                      *string                                   `json:"description"`
		NamespaceSuffix                  *string                                   `json:"namespaceSuffix"`
		ClusterNamespaceSchedulingPolicy clusternamespaceschedulingpolicy          `json:"clusterNamespaceSchedulingPolicy"`
		LifecycleDetails                 *string                                   `json:"lifecycleDetails"`
		Id                               *string                                   `json:"id"`
		DisplayName                      *string                                   `json:"displayName"`
		CompartmentId                    *string                                   `json:"compartmentId"`
		KubernetesVersion                *string                                   `json:"kubernetesVersion"`
		TimeCreated                      *common.SDKTime                           `json:"timeCreated"`
		TimeUpdated                      *common.SDKTime                           `json:"timeUpdated"`
		LifecycleState                   ClusterNamespaceProfileLifecycleStateEnum `json:"lifecycleState"`
		FreeformTags                     map[string]string                         `json:"freeformTags"`
		DefinedTags                      map[string]map[string]interface{}         `json:"definedTags"`
		SystemTags                       map[string]map[string]interface{}         `json:"systemTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.NamespaceSuffix = model.NamespaceSuffix

	nn, e = model.ClusterNamespaceSchedulingPolicy.UnmarshalPolymorphicJSON(model.ClusterNamespaceSchedulingPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ClusterNamespaceSchedulingPolicy = nn.(ClusterNamespaceSchedulingPolicy)
	} else {
		m.ClusterNamespaceSchedulingPolicy = nil
	}

	m.LifecycleDetails = model.LifecycleDetails

	m.Id = model.Id

	m.DisplayName = model.DisplayName

	m.CompartmentId = model.CompartmentId

	m.KubernetesVersion = model.KubernetesVersion

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	return
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
