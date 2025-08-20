// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CloudAsmInstance The details of a cloud ASM instance.
type CloudAsmInstance struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud ASM instance.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the ASM instance. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The name of the cloud ASM instance.
	ComponentName *string `mandatory:"true" json:"componentName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud ASM that the ASM instance belongs to.
	CloudAsmId *string `mandatory:"true" json:"cloudAsmId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system that the ASM instance is a part of.
	CloudDbSystemId *string `mandatory:"true" json:"cloudDbSystemId"`

	// The current lifecycle state of the cloud ASM instance.
	LifecycleState CloudAsmInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) in DBaas service.
	DbaasId *string `mandatory:"false" json:"dbaasId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB node on which the ASM instance is running.
	CloudDbNodeId *string `mandatory:"false" json:"cloudDbNodeId"`

	// The Automatic Diagnostic Repository (ADR) home directory for the ASM instance.
	AdrHomeDirectory *string `mandatory:"false" json:"adrHomeDirectory"`

	// The name of the host on which the ASM instance is running.
	HostName *string `mandatory:"false" json:"hostName"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the cloud ASM instance was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the cloud ASM instance was last updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CloudAsmInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudAsmInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCloudAsmInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCloudAsmInstanceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CloudAsmInstanceLifecycleStateEnum Enum with underlying type: string
type CloudAsmInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for CloudAsmInstanceLifecycleStateEnum
const (
	CloudAsmInstanceLifecycleStateCreating CloudAsmInstanceLifecycleStateEnum = "CREATING"
	CloudAsmInstanceLifecycleStateActive   CloudAsmInstanceLifecycleStateEnum = "ACTIVE"
	CloudAsmInstanceLifecycleStateInactive CloudAsmInstanceLifecycleStateEnum = "INACTIVE"
	CloudAsmInstanceLifecycleStateUpdating CloudAsmInstanceLifecycleStateEnum = "UPDATING"
	CloudAsmInstanceLifecycleStateDeleting CloudAsmInstanceLifecycleStateEnum = "DELETING"
	CloudAsmInstanceLifecycleStateDeleted  CloudAsmInstanceLifecycleStateEnum = "DELETED"
	CloudAsmInstanceLifecycleStateFailed   CloudAsmInstanceLifecycleStateEnum = "FAILED"
)

var mappingCloudAsmInstanceLifecycleStateEnum = map[string]CloudAsmInstanceLifecycleStateEnum{
	"CREATING": CloudAsmInstanceLifecycleStateCreating,
	"ACTIVE":   CloudAsmInstanceLifecycleStateActive,
	"INACTIVE": CloudAsmInstanceLifecycleStateInactive,
	"UPDATING": CloudAsmInstanceLifecycleStateUpdating,
	"DELETING": CloudAsmInstanceLifecycleStateDeleting,
	"DELETED":  CloudAsmInstanceLifecycleStateDeleted,
	"FAILED":   CloudAsmInstanceLifecycleStateFailed,
}

var mappingCloudAsmInstanceLifecycleStateEnumLowerCase = map[string]CloudAsmInstanceLifecycleStateEnum{
	"creating": CloudAsmInstanceLifecycleStateCreating,
	"active":   CloudAsmInstanceLifecycleStateActive,
	"inactive": CloudAsmInstanceLifecycleStateInactive,
	"updating": CloudAsmInstanceLifecycleStateUpdating,
	"deleting": CloudAsmInstanceLifecycleStateDeleting,
	"deleted":  CloudAsmInstanceLifecycleStateDeleted,
	"failed":   CloudAsmInstanceLifecycleStateFailed,
}

// GetCloudAsmInstanceLifecycleStateEnumValues Enumerates the set of values for CloudAsmInstanceLifecycleStateEnum
func GetCloudAsmInstanceLifecycleStateEnumValues() []CloudAsmInstanceLifecycleStateEnum {
	values := make([]CloudAsmInstanceLifecycleStateEnum, 0)
	for _, v := range mappingCloudAsmInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCloudAsmInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for CloudAsmInstanceLifecycleStateEnum
func GetCloudAsmInstanceLifecycleStateEnumStringValues() []string {
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

// GetMappingCloudAsmInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCloudAsmInstanceLifecycleStateEnum(val string) (CloudAsmInstanceLifecycleStateEnum, bool) {
	enum, ok := mappingCloudAsmInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
