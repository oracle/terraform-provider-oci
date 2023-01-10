// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Mesh The mesh resource is the top-level container that represents the logical boundary of application traffic between the services and deployments that reside within it. A mesh also provides a unit of access control.
type Mesh struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. The name does not have to be unique and can be changed after creation.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time when this resource was created in an RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when this resource was updated in an RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the Resource.
	LifecycleState MeshLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Description of the resource. It can be changed after creation.
	// Avoid entering confidential information.
	// Example: `This is my new resource`
	Description *string `mandatory:"false" json:"description"`

	// A list of certificate authority resources to use for creating leaf certificates for mTLS authentication.
	// Currently we only support one certificate authority, but this may expand in future releases. Request with
	// more than one certificate authority will be rejected.
	CertificateAuthorities []CertificateAuthority `mandatory:"false" json:"certificateAuthorities"`

	Mtls *MeshMutualTransportLayerSecurity `mandatory:"false" json:"mtls"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Mesh) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Mesh) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMeshLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMeshLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MeshLifecycleStateEnum Enum with underlying type: string
type MeshLifecycleStateEnum string

// Set of constants representing the allowable values for MeshLifecycleStateEnum
const (
	MeshLifecycleStateCreating MeshLifecycleStateEnum = "CREATING"
	MeshLifecycleStateUpdating MeshLifecycleStateEnum = "UPDATING"
	MeshLifecycleStateActive   MeshLifecycleStateEnum = "ACTIVE"
	MeshLifecycleStateDeleting MeshLifecycleStateEnum = "DELETING"
	MeshLifecycleStateDeleted  MeshLifecycleStateEnum = "DELETED"
	MeshLifecycleStateFailed   MeshLifecycleStateEnum = "FAILED"
)

var mappingMeshLifecycleStateEnum = map[string]MeshLifecycleStateEnum{
	"CREATING": MeshLifecycleStateCreating,
	"UPDATING": MeshLifecycleStateUpdating,
	"ACTIVE":   MeshLifecycleStateActive,
	"DELETING": MeshLifecycleStateDeleting,
	"DELETED":  MeshLifecycleStateDeleted,
	"FAILED":   MeshLifecycleStateFailed,
}

var mappingMeshLifecycleStateEnumLowerCase = map[string]MeshLifecycleStateEnum{
	"creating": MeshLifecycleStateCreating,
	"updating": MeshLifecycleStateUpdating,
	"active":   MeshLifecycleStateActive,
	"deleting": MeshLifecycleStateDeleting,
	"deleted":  MeshLifecycleStateDeleted,
	"failed":   MeshLifecycleStateFailed,
}

// GetMeshLifecycleStateEnumValues Enumerates the set of values for MeshLifecycleStateEnum
func GetMeshLifecycleStateEnumValues() []MeshLifecycleStateEnum {
	values := make([]MeshLifecycleStateEnum, 0)
	for _, v := range mappingMeshLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMeshLifecycleStateEnumStringValues Enumerates the set of values in String for MeshLifecycleStateEnum
func GetMeshLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingMeshLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMeshLifecycleStateEnum(val string) (MeshLifecycleStateEnum, bool) {
	enum, ok := mappingMeshLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
