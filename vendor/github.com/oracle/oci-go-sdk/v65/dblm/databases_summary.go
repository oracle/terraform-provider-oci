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

// DatabasesSummary Summary of a databases.
type DatabasesSummary struct {

	// Database ocid.
	DatabaseId *string `mandatory:"true" json:"databaseId"`

	ImageDetails *ImageDetails `mandatory:"true" json:"imageDetails"`

	PatchComplianceDetails *PatchComplianceDetails `mandatory:"true" json:"patchComplianceDetails"`

	PatchActivityDetails *PatchActivityDetails `mandatory:"true" json:"patchActivityDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Database name.
	DatabaseName *string `mandatory:"false" json:"databaseName"`

	// Database type.
	DatabaseType DatabasesSummaryDatabaseTypeEnum `mandatory:"false" json:"databaseType,omitempty"`

	// Database release.
	Release *string `mandatory:"false" json:"release"`

	// Database release full version.
	ReleaseFullVersion *string `mandatory:"false" json:"releaseFullVersion"`

	// Path to the Oracle home.
	OracleHomePath *string `mandatory:"false" json:"oracleHomePath"`

	// This is the hashcode representing the list of patches applied.
	CurrentPatchWatermark *string `mandatory:"false" json:"currentPatchWatermark"`

	// For SI, hosted on host and for RAC, host on cluster.
	HostOrCluster *string `mandatory:"false" json:"hostOrCluster"`

	// Intermediate user to be used for patching, created and maintained by customers. This user requires sudo access to switch as Oracle home owner and root user
	PatchUser *string `mandatory:"false" json:"patchUser"`

	// Path to sudo binary (executable) file
	SudoFilePath *string `mandatory:"false" json:"sudoFilePath"`

	// List of additional patches on database.
	AdditionalPatches []AdditionalPatches `mandatory:"false" json:"additionalPatches"`

	// Summary of vulnerabilities found in registered resources grouped by severity.
	VulnerabilitiesSummary *interface{} `mandatory:"false" json:"vulnerabilitiesSummary"`

	// The current state of the database.
	LifecycleState DatabasesSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DatabasesSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabasesSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDatabasesSummaryDatabaseTypeEnum(string(m.DatabaseType)); !ok && m.DatabaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", m.DatabaseType, strings.Join(GetDatabasesSummaryDatabaseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDatabasesSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDatabasesSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabasesSummaryDatabaseTypeEnum Enum with underlying type: string
type DatabasesSummaryDatabaseTypeEnum string

// Set of constants representing the allowable values for DatabasesSummaryDatabaseTypeEnum
const (
	DatabasesSummaryDatabaseTypeSi  DatabasesSummaryDatabaseTypeEnum = "SI"
	DatabasesSummaryDatabaseTypeRac DatabasesSummaryDatabaseTypeEnum = "RAC"
)

var mappingDatabasesSummaryDatabaseTypeEnum = map[string]DatabasesSummaryDatabaseTypeEnum{
	"SI":  DatabasesSummaryDatabaseTypeSi,
	"RAC": DatabasesSummaryDatabaseTypeRac,
}

var mappingDatabasesSummaryDatabaseTypeEnumLowerCase = map[string]DatabasesSummaryDatabaseTypeEnum{
	"si":  DatabasesSummaryDatabaseTypeSi,
	"rac": DatabasesSummaryDatabaseTypeRac,
}

// GetDatabasesSummaryDatabaseTypeEnumValues Enumerates the set of values for DatabasesSummaryDatabaseTypeEnum
func GetDatabasesSummaryDatabaseTypeEnumValues() []DatabasesSummaryDatabaseTypeEnum {
	values := make([]DatabasesSummaryDatabaseTypeEnum, 0)
	for _, v := range mappingDatabasesSummaryDatabaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabasesSummaryDatabaseTypeEnumStringValues Enumerates the set of values in String for DatabasesSummaryDatabaseTypeEnum
func GetDatabasesSummaryDatabaseTypeEnumStringValues() []string {
	return []string{
		"SI",
		"RAC",
	}
}

// GetMappingDatabasesSummaryDatabaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabasesSummaryDatabaseTypeEnum(val string) (DatabasesSummaryDatabaseTypeEnum, bool) {
	enum, ok := mappingDatabasesSummaryDatabaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DatabasesSummaryLifecycleStateEnum Enum with underlying type: string
type DatabasesSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for DatabasesSummaryLifecycleStateEnum
const (
	DatabasesSummaryLifecycleStateCreating DatabasesSummaryLifecycleStateEnum = "CREATING"
	DatabasesSummaryLifecycleStateUpdating DatabasesSummaryLifecycleStateEnum = "UPDATING"
	DatabasesSummaryLifecycleStateActive   DatabasesSummaryLifecycleStateEnum = "ACTIVE"
	DatabasesSummaryLifecycleStateDeleting DatabasesSummaryLifecycleStateEnum = "DELETING"
	DatabasesSummaryLifecycleStateDeleted  DatabasesSummaryLifecycleStateEnum = "DELETED"
	DatabasesSummaryLifecycleStateFailed   DatabasesSummaryLifecycleStateEnum = "FAILED"
)

var mappingDatabasesSummaryLifecycleStateEnum = map[string]DatabasesSummaryLifecycleStateEnum{
	"CREATING": DatabasesSummaryLifecycleStateCreating,
	"UPDATING": DatabasesSummaryLifecycleStateUpdating,
	"ACTIVE":   DatabasesSummaryLifecycleStateActive,
	"DELETING": DatabasesSummaryLifecycleStateDeleting,
	"DELETED":  DatabasesSummaryLifecycleStateDeleted,
	"FAILED":   DatabasesSummaryLifecycleStateFailed,
}

var mappingDatabasesSummaryLifecycleStateEnumLowerCase = map[string]DatabasesSummaryLifecycleStateEnum{
	"creating": DatabasesSummaryLifecycleStateCreating,
	"updating": DatabasesSummaryLifecycleStateUpdating,
	"active":   DatabasesSummaryLifecycleStateActive,
	"deleting": DatabasesSummaryLifecycleStateDeleting,
	"deleted":  DatabasesSummaryLifecycleStateDeleted,
	"failed":   DatabasesSummaryLifecycleStateFailed,
}

// GetDatabasesSummaryLifecycleStateEnumValues Enumerates the set of values for DatabasesSummaryLifecycleStateEnum
func GetDatabasesSummaryLifecycleStateEnumValues() []DatabasesSummaryLifecycleStateEnum {
	values := make([]DatabasesSummaryLifecycleStateEnum, 0)
	for _, v := range mappingDatabasesSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabasesSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for DatabasesSummaryLifecycleStateEnum
func GetDatabasesSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingDatabasesSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabasesSummaryLifecycleStateEnum(val string) (DatabasesSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingDatabasesSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
