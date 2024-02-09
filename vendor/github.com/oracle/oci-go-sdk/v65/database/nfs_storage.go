// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NfsStorage NFS Storage details.
type NfsStorage struct {

	// The user-provided name of the NFS storage.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Specifies the directory on NFS server that should be mounted.
	NfsServerExport *string `mandatory:"true" json:"nfsServerExport"`

	// IP addresses and hostnames for NFS Storage.
	NfsServers []string `mandatory:"true" json:"nfsServers"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the NFS Storage.
	Id *string `mandatory:"false" json:"id"`

	// ocid of the ACD attached to the NFS Storage.
	AssociatedAutonomousContainerDatabase *string `mandatory:"false" json:"associatedAutonomousContainerDatabase"`

	// Name of the ACD attached to the NFS Storage.
	AssociatedAutonomousContainerDatabaseName *string `mandatory:"false" json:"associatedAutonomousContainerDatabaseName"`

	// The current lifecycle state of the NFS Storage.
	LifecycleState NfsStorageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A descriptive text associated with the lifecycleState.
	// Typically contains additional displayable text
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time that the NFS Storage was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time that the NFS Storage was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m NfsStorage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NfsStorage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNfsStorageLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNfsStorageLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NfsStorageLifecycleStateEnum Enum with underlying type: string
type NfsStorageLifecycleStateEnum string

// Set of constants representing the allowable values for NfsStorageLifecycleStateEnum
const (
	NfsStorageLifecycleStateAvailable  NfsStorageLifecycleStateEnum = "AVAILABLE"
	NfsStorageLifecycleStateAttached   NfsStorageLifecycleStateEnum = "ATTACHED"
	NfsStorageLifecycleStateAttaching  NfsStorageLifecycleStateEnum = "ATTACHING"
	NfsStorageLifecycleStateDetaching  NfsStorageLifecycleStateEnum = "DETACHING"
	NfsStorageLifecycleStateTerminated NfsStorageLifecycleStateEnum = "TERMINATED"
)

var mappingNfsStorageLifecycleStateEnum = map[string]NfsStorageLifecycleStateEnum{
	"AVAILABLE":  NfsStorageLifecycleStateAvailable,
	"ATTACHED":   NfsStorageLifecycleStateAttached,
	"ATTACHING":  NfsStorageLifecycleStateAttaching,
	"DETACHING":  NfsStorageLifecycleStateDetaching,
	"TERMINATED": NfsStorageLifecycleStateTerminated,
}

var mappingNfsStorageLifecycleStateEnumLowerCase = map[string]NfsStorageLifecycleStateEnum{
	"available":  NfsStorageLifecycleStateAvailable,
	"attached":   NfsStorageLifecycleStateAttached,
	"attaching":  NfsStorageLifecycleStateAttaching,
	"detaching":  NfsStorageLifecycleStateDetaching,
	"terminated": NfsStorageLifecycleStateTerminated,
}

// GetNfsStorageLifecycleStateEnumValues Enumerates the set of values for NfsStorageLifecycleStateEnum
func GetNfsStorageLifecycleStateEnumValues() []NfsStorageLifecycleStateEnum {
	values := make([]NfsStorageLifecycleStateEnum, 0)
	for _, v := range mappingNfsStorageLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNfsStorageLifecycleStateEnumStringValues Enumerates the set of values in String for NfsStorageLifecycleStateEnum
func GetNfsStorageLifecycleStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"ATTACHED",
		"ATTACHING",
		"DETACHING",
		"TERMINATED",
	}
}

// GetMappingNfsStorageLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNfsStorageLifecycleStateEnum(val string) (NfsStorageLifecycleStateEnum, bool) {
	enum, ok := mappingNfsStorageLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
