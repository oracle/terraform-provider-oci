// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FppServer Description of FPP Server.
type FppServer struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// FPP Server Identifier, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the the FPP server was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the FPP server.
	LifecycleState FppServerLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Unique Management Agent identifier
	ManagementAgentId *string `mandatory:"false" json:"managementAgentId"`

	// Unique Management Agent Compartment identifier
	MacAgentCompartmentId *string `mandatory:"false" json:"macAgentCompartmentId"`

	// REST endpoint URL of FPP Server
	FppRestUrl *string `mandatory:"false" json:"fppRestUrl"`

	// Absolute path of JKS Trust Store
	TrustStoreLocation *string `mandatory:"false" json:"trustStoreLocation"`

	// JKS Trust Store Password
	TrustStorePassword *string `mandatory:"false" json:"trustStorePassword"`

	// FPP REST User Name
	RestUserName *string `mandatory:"false" json:"restUserName"`

	// FPP REST User Password
	RestUserPassword *string `mandatory:"false" json:"restUserPassword"`

	// The time the FPP server was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m FppServer) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FppServer) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFppServerLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFppServerLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FppServerLifecycleStateEnum Enum with underlying type: string
type FppServerLifecycleStateEnum string

// Set of constants representing the allowable values for FppServerLifecycleStateEnum
const (
	FppServerLifecycleStateCreating FppServerLifecycleStateEnum = "CREATING"
	FppServerLifecycleStateUpdating FppServerLifecycleStateEnum = "UPDATING"
	FppServerLifecycleStateActive   FppServerLifecycleStateEnum = "ACTIVE"
	FppServerLifecycleStateDeleting FppServerLifecycleStateEnum = "DELETING"
	FppServerLifecycleStateDeleted  FppServerLifecycleStateEnum = "DELETED"
	FppServerLifecycleStateFailed   FppServerLifecycleStateEnum = "FAILED"
)

var mappingFppServerLifecycleStateEnum = map[string]FppServerLifecycleStateEnum{
	"CREATING": FppServerLifecycleStateCreating,
	"UPDATING": FppServerLifecycleStateUpdating,
	"ACTIVE":   FppServerLifecycleStateActive,
	"DELETING": FppServerLifecycleStateDeleting,
	"DELETED":  FppServerLifecycleStateDeleted,
	"FAILED":   FppServerLifecycleStateFailed,
}

var mappingFppServerLifecycleStateEnumLowerCase = map[string]FppServerLifecycleStateEnum{
	"creating": FppServerLifecycleStateCreating,
	"updating": FppServerLifecycleStateUpdating,
	"active":   FppServerLifecycleStateActive,
	"deleting": FppServerLifecycleStateDeleting,
	"deleted":  FppServerLifecycleStateDeleted,
	"failed":   FppServerLifecycleStateFailed,
}

// GetFppServerLifecycleStateEnumValues Enumerates the set of values for FppServerLifecycleStateEnum
func GetFppServerLifecycleStateEnumValues() []FppServerLifecycleStateEnum {
	values := make([]FppServerLifecycleStateEnum, 0)
	for _, v := range mappingFppServerLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFppServerLifecycleStateEnumStringValues Enumerates the set of values in String for FppServerLifecycleStateEnum
func GetFppServerLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingFppServerLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFppServerLifecycleStateEnum(val string) (FppServerLifecycleStateEnum, bool) {
	enum, ok := mappingFppServerLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
