// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExternalAsmInstance The details of an external ASM instance.
type ExternalAsmInstance struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external ASM instance.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the ASM instance. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the external ASM instance.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external ASM that the ASM instance belongs to.
	ExternalAsmId *string `mandatory:"true" json:"externalAsmId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB system that the ASM instance is a part of.
	ExternalDbSystemId *string `mandatory:"true" json:"externalDbSystemId"`

	// The current lifecycle state of the external ASM instance.
	LifecycleState ExternalAsmInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the external DB node on which the ASM instance is running.
	ExternalDbNodeId *string `mandatory:"false" json:"externalDbNodeId"`

	// The Automatic Diagnostic Repository (ADR) home directory for the ASM instance.
	AdrHomeDirectory *string `mandatory:"false" json:"adrHomeDirectory"`

	// The name of the host on which the ASM instance is running.
	HostName *string `mandatory:"false" json:"hostName"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the external ASM instance was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the external ASM instance was last updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ExternalAsmInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExternalAsmInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExternalAsmInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExternalAsmInstanceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExternalAsmInstanceLifecycleStateEnum Enum with underlying type: string
type ExternalAsmInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for ExternalAsmInstanceLifecycleStateEnum
const (
	ExternalAsmInstanceLifecycleStateCreating ExternalAsmInstanceLifecycleStateEnum = "CREATING"
	ExternalAsmInstanceLifecycleStateActive   ExternalAsmInstanceLifecycleStateEnum = "ACTIVE"
	ExternalAsmInstanceLifecycleStateInactive ExternalAsmInstanceLifecycleStateEnum = "INACTIVE"
	ExternalAsmInstanceLifecycleStateUpdating ExternalAsmInstanceLifecycleStateEnum = "UPDATING"
	ExternalAsmInstanceLifecycleStateDeleting ExternalAsmInstanceLifecycleStateEnum = "DELETING"
	ExternalAsmInstanceLifecycleStateDeleted  ExternalAsmInstanceLifecycleStateEnum = "DELETED"
	ExternalAsmInstanceLifecycleStateFailed   ExternalAsmInstanceLifecycleStateEnum = "FAILED"
)

var mappingExternalAsmInstanceLifecycleStateEnum = map[string]ExternalAsmInstanceLifecycleStateEnum{
	"CREATING": ExternalAsmInstanceLifecycleStateCreating,
	"ACTIVE":   ExternalAsmInstanceLifecycleStateActive,
	"INACTIVE": ExternalAsmInstanceLifecycleStateInactive,
	"UPDATING": ExternalAsmInstanceLifecycleStateUpdating,
	"DELETING": ExternalAsmInstanceLifecycleStateDeleting,
	"DELETED":  ExternalAsmInstanceLifecycleStateDeleted,
	"FAILED":   ExternalAsmInstanceLifecycleStateFailed,
}

var mappingExternalAsmInstanceLifecycleStateEnumLowerCase = map[string]ExternalAsmInstanceLifecycleStateEnum{
	"creating": ExternalAsmInstanceLifecycleStateCreating,
	"active":   ExternalAsmInstanceLifecycleStateActive,
	"inactive": ExternalAsmInstanceLifecycleStateInactive,
	"updating": ExternalAsmInstanceLifecycleStateUpdating,
	"deleting": ExternalAsmInstanceLifecycleStateDeleting,
	"deleted":  ExternalAsmInstanceLifecycleStateDeleted,
	"failed":   ExternalAsmInstanceLifecycleStateFailed,
}

// GetExternalAsmInstanceLifecycleStateEnumValues Enumerates the set of values for ExternalAsmInstanceLifecycleStateEnum
func GetExternalAsmInstanceLifecycleStateEnumValues() []ExternalAsmInstanceLifecycleStateEnum {
	values := make([]ExternalAsmInstanceLifecycleStateEnum, 0)
	for _, v := range mappingExternalAsmInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExternalAsmInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for ExternalAsmInstanceLifecycleStateEnum
func GetExternalAsmInstanceLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingExternalAsmInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExternalAsmInstanceLifecycleStateEnum(val string) (ExternalAsmInstanceLifecycleStateEnum, bool) {
	enum, ok := mappingExternalAsmInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
