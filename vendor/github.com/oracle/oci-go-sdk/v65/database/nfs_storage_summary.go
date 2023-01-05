// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// NfsStorageSummary NFS Storage details, including the Autonomous Container Database using the NFS storage.
type NfsStorageSummary struct {

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
	LifecycleState NfsStorageSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

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

func (m NfsStorageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NfsStorageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNfsStorageSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNfsStorageSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NfsStorageSummaryLifecycleStateEnum Enum with underlying type: string
type NfsStorageSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for NfsStorageSummaryLifecycleStateEnum
const (
	NfsStorageSummaryLifecycleStateAvailable  NfsStorageSummaryLifecycleStateEnum = "AVAILABLE"
	NfsStorageSummaryLifecycleStateAttached   NfsStorageSummaryLifecycleStateEnum = "ATTACHED"
	NfsStorageSummaryLifecycleStateAttaching  NfsStorageSummaryLifecycleStateEnum = "ATTACHING"
	NfsStorageSummaryLifecycleStateDetaching  NfsStorageSummaryLifecycleStateEnum = "DETACHING"
	NfsStorageSummaryLifecycleStateTerminated NfsStorageSummaryLifecycleStateEnum = "TERMINATED"
)

var mappingNfsStorageSummaryLifecycleStateEnum = map[string]NfsStorageSummaryLifecycleStateEnum{
	"AVAILABLE":  NfsStorageSummaryLifecycleStateAvailable,
	"ATTACHED":   NfsStorageSummaryLifecycleStateAttached,
	"ATTACHING":  NfsStorageSummaryLifecycleStateAttaching,
	"DETACHING":  NfsStorageSummaryLifecycleStateDetaching,
	"TERMINATED": NfsStorageSummaryLifecycleStateTerminated,
}

var mappingNfsStorageSummaryLifecycleStateEnumLowerCase = map[string]NfsStorageSummaryLifecycleStateEnum{
	"available":  NfsStorageSummaryLifecycleStateAvailable,
	"attached":   NfsStorageSummaryLifecycleStateAttached,
	"attaching":  NfsStorageSummaryLifecycleStateAttaching,
	"detaching":  NfsStorageSummaryLifecycleStateDetaching,
	"terminated": NfsStorageSummaryLifecycleStateTerminated,
}

// GetNfsStorageSummaryLifecycleStateEnumValues Enumerates the set of values for NfsStorageSummaryLifecycleStateEnum
func GetNfsStorageSummaryLifecycleStateEnumValues() []NfsStorageSummaryLifecycleStateEnum {
	values := make([]NfsStorageSummaryLifecycleStateEnum, 0)
	for _, v := range mappingNfsStorageSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNfsStorageSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for NfsStorageSummaryLifecycleStateEnum
func GetNfsStorageSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"ATTACHED",
		"ATTACHING",
		"DETACHING",
		"TERMINATED",
	}
}

// GetMappingNfsStorageSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNfsStorageSummaryLifecycleStateEnum(val string) (NfsStorageSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingNfsStorageSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
