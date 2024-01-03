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

// OneoffPatchSummary An Oracle one-off patch for a specified database version.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type OneoffPatchSummary struct {

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
	LifecycleState OneoffPatchSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

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

func (m OneoffPatchSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OneoffPatchSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOneoffPatchSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOneoffPatchSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OneoffPatchSummaryLifecycleStateEnum Enum with underlying type: string
type OneoffPatchSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for OneoffPatchSummaryLifecycleStateEnum
const (
	OneoffPatchSummaryLifecycleStateCreating    OneoffPatchSummaryLifecycleStateEnum = "CREATING"
	OneoffPatchSummaryLifecycleStateAvailable   OneoffPatchSummaryLifecycleStateEnum = "AVAILABLE"
	OneoffPatchSummaryLifecycleStateUpdating    OneoffPatchSummaryLifecycleStateEnum = "UPDATING"
	OneoffPatchSummaryLifecycleStateInactive    OneoffPatchSummaryLifecycleStateEnum = "INACTIVE"
	OneoffPatchSummaryLifecycleStateFailed      OneoffPatchSummaryLifecycleStateEnum = "FAILED"
	OneoffPatchSummaryLifecycleStateExpired     OneoffPatchSummaryLifecycleStateEnum = "EXPIRED"
	OneoffPatchSummaryLifecycleStateDeleting    OneoffPatchSummaryLifecycleStateEnum = "DELETING"
	OneoffPatchSummaryLifecycleStateDeleted     OneoffPatchSummaryLifecycleStateEnum = "DELETED"
	OneoffPatchSummaryLifecycleStateTerminating OneoffPatchSummaryLifecycleStateEnum = "TERMINATING"
	OneoffPatchSummaryLifecycleStateTerminated  OneoffPatchSummaryLifecycleStateEnum = "TERMINATED"
)

var mappingOneoffPatchSummaryLifecycleStateEnum = map[string]OneoffPatchSummaryLifecycleStateEnum{
	"CREATING":    OneoffPatchSummaryLifecycleStateCreating,
	"AVAILABLE":   OneoffPatchSummaryLifecycleStateAvailable,
	"UPDATING":    OneoffPatchSummaryLifecycleStateUpdating,
	"INACTIVE":    OneoffPatchSummaryLifecycleStateInactive,
	"FAILED":      OneoffPatchSummaryLifecycleStateFailed,
	"EXPIRED":     OneoffPatchSummaryLifecycleStateExpired,
	"DELETING":    OneoffPatchSummaryLifecycleStateDeleting,
	"DELETED":     OneoffPatchSummaryLifecycleStateDeleted,
	"TERMINATING": OneoffPatchSummaryLifecycleStateTerminating,
	"TERMINATED":  OneoffPatchSummaryLifecycleStateTerminated,
}

var mappingOneoffPatchSummaryLifecycleStateEnumLowerCase = map[string]OneoffPatchSummaryLifecycleStateEnum{
	"creating":    OneoffPatchSummaryLifecycleStateCreating,
	"available":   OneoffPatchSummaryLifecycleStateAvailable,
	"updating":    OneoffPatchSummaryLifecycleStateUpdating,
	"inactive":    OneoffPatchSummaryLifecycleStateInactive,
	"failed":      OneoffPatchSummaryLifecycleStateFailed,
	"expired":     OneoffPatchSummaryLifecycleStateExpired,
	"deleting":    OneoffPatchSummaryLifecycleStateDeleting,
	"deleted":     OneoffPatchSummaryLifecycleStateDeleted,
	"terminating": OneoffPatchSummaryLifecycleStateTerminating,
	"terminated":  OneoffPatchSummaryLifecycleStateTerminated,
}

// GetOneoffPatchSummaryLifecycleStateEnumValues Enumerates the set of values for OneoffPatchSummaryLifecycleStateEnum
func GetOneoffPatchSummaryLifecycleStateEnumValues() []OneoffPatchSummaryLifecycleStateEnum {
	values := make([]OneoffPatchSummaryLifecycleStateEnum, 0)
	for _, v := range mappingOneoffPatchSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOneoffPatchSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for OneoffPatchSummaryLifecycleStateEnum
func GetOneoffPatchSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingOneoffPatchSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOneoffPatchSummaryLifecycleStateEnum(val string) (OneoffPatchSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingOneoffPatchSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
