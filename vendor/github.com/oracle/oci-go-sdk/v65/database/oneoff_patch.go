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

// OneoffPatch One-off patches are created by specifying a database version, releaseUpdate and one-off patch number.
type OneoffPatch struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the one-off patch.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// One-off patch name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A valid Oracle Database version. For a list of supported versions, use the ListDbVersions operation.
	// This cannot be updated in parallel with any of the following: licenseModel, dbEdition, cpuCoreCount, computeCount, computeModel, adminPassword, whitelistedIps, isMTLSConnectionRequired, openMode, permissionLevel, dbWorkload, privateEndpointLabel, nsgIds, isRefreshable, dbName, scheduledOperations, dbToolsDetails, isLocalDataGuardEnabled, or isFreeTier.
	DbVersion *string `mandatory:"true" json:"dbVersion"`

	// The PSU or PBP or Release Updates. To get a list of supported versions, use the ListDbVersions operation.
	ReleaseUpdate *string `mandatory:"true" json:"releaseUpdate"`

	// The current state of the one-off patch.
	LifecycleState OneoffPatchLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time one-off patch was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// List of one-off patches for Database Homes.
	OneOffPatches []string `mandatory:"false" json:"oneOffPatches"`

	// The size of one-off patch in kilobytes.
	SizeInKBs *float32 `mandatory:"false" json:"sizeInKBs"`

	// Detailed message for the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// SHA-256 checksum of the one-off patch.
	Sha256Sum *string `mandatory:"false" json:"sha256Sum"`

	// The date and time one-off patch was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The date and time until which the one-off patch will be available for download.
	TimeOfExpiration *common.SDKTime `mandatory:"false" json:"timeOfExpiration"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m OneoffPatch) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OneoffPatch) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOneoffPatchLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOneoffPatchLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OneoffPatchLifecycleStateEnum Enum with underlying type: string
type OneoffPatchLifecycleStateEnum string

// Set of constants representing the allowable values for OneoffPatchLifecycleStateEnum
const (
	OneoffPatchLifecycleStateCreating    OneoffPatchLifecycleStateEnum = "CREATING"
	OneoffPatchLifecycleStateAvailable   OneoffPatchLifecycleStateEnum = "AVAILABLE"
	OneoffPatchLifecycleStateUpdating    OneoffPatchLifecycleStateEnum = "UPDATING"
	OneoffPatchLifecycleStateInactive    OneoffPatchLifecycleStateEnum = "INACTIVE"
	OneoffPatchLifecycleStateFailed      OneoffPatchLifecycleStateEnum = "FAILED"
	OneoffPatchLifecycleStateExpired     OneoffPatchLifecycleStateEnum = "EXPIRED"
	OneoffPatchLifecycleStateDeleting    OneoffPatchLifecycleStateEnum = "DELETING"
	OneoffPatchLifecycleStateDeleted     OneoffPatchLifecycleStateEnum = "DELETED"
	OneoffPatchLifecycleStateTerminating OneoffPatchLifecycleStateEnum = "TERMINATING"
	OneoffPatchLifecycleStateTerminated  OneoffPatchLifecycleStateEnum = "TERMINATED"
)

var mappingOneoffPatchLifecycleStateEnum = map[string]OneoffPatchLifecycleStateEnum{
	"CREATING":    OneoffPatchLifecycleStateCreating,
	"AVAILABLE":   OneoffPatchLifecycleStateAvailable,
	"UPDATING":    OneoffPatchLifecycleStateUpdating,
	"INACTIVE":    OneoffPatchLifecycleStateInactive,
	"FAILED":      OneoffPatchLifecycleStateFailed,
	"EXPIRED":     OneoffPatchLifecycleStateExpired,
	"DELETING":    OneoffPatchLifecycleStateDeleting,
	"DELETED":     OneoffPatchLifecycleStateDeleted,
	"TERMINATING": OneoffPatchLifecycleStateTerminating,
	"TERMINATED":  OneoffPatchLifecycleStateTerminated,
}

var mappingOneoffPatchLifecycleStateEnumLowerCase = map[string]OneoffPatchLifecycleStateEnum{
	"creating":    OneoffPatchLifecycleStateCreating,
	"available":   OneoffPatchLifecycleStateAvailable,
	"updating":    OneoffPatchLifecycleStateUpdating,
	"inactive":    OneoffPatchLifecycleStateInactive,
	"failed":      OneoffPatchLifecycleStateFailed,
	"expired":     OneoffPatchLifecycleStateExpired,
	"deleting":    OneoffPatchLifecycleStateDeleting,
	"deleted":     OneoffPatchLifecycleStateDeleted,
	"terminating": OneoffPatchLifecycleStateTerminating,
	"terminated":  OneoffPatchLifecycleStateTerminated,
}

// GetOneoffPatchLifecycleStateEnumValues Enumerates the set of values for OneoffPatchLifecycleStateEnum
func GetOneoffPatchLifecycleStateEnumValues() []OneoffPatchLifecycleStateEnum {
	values := make([]OneoffPatchLifecycleStateEnum, 0)
	for _, v := range mappingOneoffPatchLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOneoffPatchLifecycleStateEnumStringValues Enumerates the set of values in String for OneoffPatchLifecycleStateEnum
func GetOneoffPatchLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"AVAILABLE",
		"UPDATING",
		"INACTIVE",
		"FAILED",
		"EXPIRED",
		"DELETING",
		"DELETED",
		"TERMINATING",
		"TERMINATED",
	}
}

// GetMappingOneoffPatchLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOneoffPatchLifecycleStateEnum(val string) (OneoffPatchLifecycleStateEnum, bool) {
	enum, ok := mappingOneoffPatchLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
