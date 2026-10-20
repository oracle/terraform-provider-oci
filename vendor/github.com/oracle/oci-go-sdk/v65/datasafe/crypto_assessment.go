// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CryptoAssessment A crypto assessment that provides insight into database cryptographic posture.
// The assessment evaluates data and network encryption, certificates, key management, backups, and quantum-readiness settings for a target database.
type CryptoAssessment struct {

	// The OCID of the crypto assessment.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the crypto assessment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the crypto assessment was created, in RFC3339 format.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the crypto assessment was last updated, in RFC3339 format.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current lifecycle state of the crypto assessment.
	LifecycleState CryptoAssessmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The type of this crypto assessment.
	Type CryptoAssessmentTypeEnum `mandatory:"true" json:"type"`

	// The target type of the crypto assessment.
	TargetType CryptoAssessmentTargetTypeEnum `mandatory:"true" json:"targetType"`

	// The display name of the crypto assessment.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the crypto assessment.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the crypto posture was last assessed, in RFC3339 format.
	TimeLastAssessed *common.SDKTime `mandatory:"false" json:"timeLastAssessed"`

	// Details about the current lifecycle state of the crypto assessment.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The actor that created the assessment.
	TriggeredBy CryptoAssessmentTriggeredByEnum `mandatory:"false" json:"triggeredBy,omitempty"`

	// The OCID of the target database.
	TargetId *string `mandatory:"false" json:"targetId"`

	// The OCID of the target database group. This is returned when `targetType` is `TARGET_DATABASE_GROUP`.
	TargetDatabaseGroupId *string `mandatory:"false" json:"targetDatabaseGroupId"`

	// The version of the assessed target database.
	DatabaseVersion *string `mandatory:"false" json:"databaseVersion"`

	// The name of the assessed target database.
	DatabaseName *string `mandatory:"false" json:"databaseName"`

	// The architecture of the assessed target database.
	DatabaseArchitecture *string `mandatory:"false" json:"databaseArchitecture"`

	// Cryptographic provider and version information observed on the target.
	CryptoProvider *string `mandatory:"false" json:"cryptoProvider"`

	// Overall posture category for the crypto assessment.
	PostureCategory CryptoPostureCategoryEnum `mandatory:"false" json:"postureCategory,omitempty"`

	// The schedule used to run the crypto assessment periodically. The schedule uses the format:
	// <version-string>;<version-specific-schedule>
	// For v1, the version-specific schedule format is:
	// <ss> <mm> <hh> <day-of-week> <day-of-month>
	// Specify either day-of-week for weekly schedules or day-of-month for monthly schedules. Do not specify both.
	// For monthly schedules, day-of-month must be between 1 and 28. If the service generates a default monthly schedule for an assessment created on day 29, 30, or 31 of a month, it uses day 28.
	Schedule *string `mandatory:"false" json:"schedule"`

	CryptoPosture *CryptoPosture `mandatory:"false" json:"cryptoPosture"`

	// Number of crypto issues detected in this assessment.
	IssueCount *int `mandatory:"false" json:"issueCount"`

	// Number of assessed targets with one or more crypto issues. For a target database assessment, this value is 1 when the target has issues and 0 otherwise. For a target database group assessment, this value is the number of targets in the group that have issues.
	TargetsWithIssuesCount *int `mandatory:"false" json:"targetsWithIssuesCount"`

	// Indicates whether scheduled execution is active for this crypto assessment. When false, the schedule value is retained but scheduled execution is paused.
	IsAssessmentScheduled *bool `mandatory:"false" json:"isAssessmentScheduled"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CryptoAssessment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoAssessment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCryptoAssessmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCryptoAssessmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoAssessmentTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCryptoAssessmentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoAssessmentTargetTypeEnum(string(m.TargetType)); !ok && m.TargetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetType: %s. Supported values are: %s.", m.TargetType, strings.Join(GetCryptoAssessmentTargetTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCryptoAssessmentTriggeredByEnum(string(m.TriggeredBy)); !ok && m.TriggeredBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TriggeredBy: %s. Supported values are: %s.", m.TriggeredBy, strings.Join(GetCryptoAssessmentTriggeredByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoPostureCategoryEnum(string(m.PostureCategory)); !ok && m.PostureCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PostureCategory: %s. Supported values are: %s.", m.PostureCategory, strings.Join(GetCryptoPostureCategoryEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CryptoAssessmentLifecycleStateEnum Enum with underlying type: string
type CryptoAssessmentLifecycleStateEnum string

// Set of constants representing the allowable values for CryptoAssessmentLifecycleStateEnum
const (
	CryptoAssessmentLifecycleStateCreating CryptoAssessmentLifecycleStateEnum = "CREATING"
	CryptoAssessmentLifecycleStateActive   CryptoAssessmentLifecycleStateEnum = "ACTIVE"
	CryptoAssessmentLifecycleStateUpdating CryptoAssessmentLifecycleStateEnum = "UPDATING"
	CryptoAssessmentLifecycleStateDeleting CryptoAssessmentLifecycleStateEnum = "DELETING"
	CryptoAssessmentLifecycleStateDeleted  CryptoAssessmentLifecycleStateEnum = "DELETED"
	CryptoAssessmentLifecycleStateFailed   CryptoAssessmentLifecycleStateEnum = "FAILED"
)

var mappingCryptoAssessmentLifecycleStateEnum = map[string]CryptoAssessmentLifecycleStateEnum{
	"CREATING": CryptoAssessmentLifecycleStateCreating,
	"ACTIVE":   CryptoAssessmentLifecycleStateActive,
	"UPDATING": CryptoAssessmentLifecycleStateUpdating,
	"DELETING": CryptoAssessmentLifecycleStateDeleting,
	"DELETED":  CryptoAssessmentLifecycleStateDeleted,
	"FAILED":   CryptoAssessmentLifecycleStateFailed,
}

var mappingCryptoAssessmentLifecycleStateEnumLowerCase = map[string]CryptoAssessmentLifecycleStateEnum{
	"creating": CryptoAssessmentLifecycleStateCreating,
	"active":   CryptoAssessmentLifecycleStateActive,
	"updating": CryptoAssessmentLifecycleStateUpdating,
	"deleting": CryptoAssessmentLifecycleStateDeleting,
	"deleted":  CryptoAssessmentLifecycleStateDeleted,
	"failed":   CryptoAssessmentLifecycleStateFailed,
}

// GetCryptoAssessmentLifecycleStateEnumValues Enumerates the set of values for CryptoAssessmentLifecycleStateEnum
func GetCryptoAssessmentLifecycleStateEnumValues() []CryptoAssessmentLifecycleStateEnum {
	values := make([]CryptoAssessmentLifecycleStateEnum, 0)
	for _, v := range mappingCryptoAssessmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoAssessmentLifecycleStateEnumStringValues Enumerates the set of values in String for CryptoAssessmentLifecycleStateEnum
func GetCryptoAssessmentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingCryptoAssessmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoAssessmentLifecycleStateEnum(val string) (CryptoAssessmentLifecycleStateEnum, bool) {
	enum, ok := mappingCryptoAssessmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CryptoAssessmentTypeEnum Enum with underlying type: string
type CryptoAssessmentTypeEnum string

// Set of constants representing the allowable values for CryptoAssessmentTypeEnum
const (
	CryptoAssessmentTypeLatest CryptoAssessmentTypeEnum = "LATEST"
	CryptoAssessmentTypeSaved  CryptoAssessmentTypeEnum = "SAVED"
)

var mappingCryptoAssessmentTypeEnum = map[string]CryptoAssessmentTypeEnum{
	"LATEST": CryptoAssessmentTypeLatest,
	"SAVED":  CryptoAssessmentTypeSaved,
}

var mappingCryptoAssessmentTypeEnumLowerCase = map[string]CryptoAssessmentTypeEnum{
	"latest": CryptoAssessmentTypeLatest,
	"saved":  CryptoAssessmentTypeSaved,
}

// GetCryptoAssessmentTypeEnumValues Enumerates the set of values for CryptoAssessmentTypeEnum
func GetCryptoAssessmentTypeEnumValues() []CryptoAssessmentTypeEnum {
	values := make([]CryptoAssessmentTypeEnum, 0)
	for _, v := range mappingCryptoAssessmentTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoAssessmentTypeEnumStringValues Enumerates the set of values in String for CryptoAssessmentTypeEnum
func GetCryptoAssessmentTypeEnumStringValues() []string {
	return []string{
		"LATEST",
		"SAVED",
	}
}

// GetMappingCryptoAssessmentTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoAssessmentTypeEnum(val string) (CryptoAssessmentTypeEnum, bool) {
	enum, ok := mappingCryptoAssessmentTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CryptoAssessmentTriggeredByEnum Enum with underlying type: string
type CryptoAssessmentTriggeredByEnum string

// Set of constants representing the allowable values for CryptoAssessmentTriggeredByEnum
const (
	CryptoAssessmentTriggeredByUser   CryptoAssessmentTriggeredByEnum = "USER"
	CryptoAssessmentTriggeredBySystem CryptoAssessmentTriggeredByEnum = "SYSTEM"
)

var mappingCryptoAssessmentTriggeredByEnum = map[string]CryptoAssessmentTriggeredByEnum{
	"USER":   CryptoAssessmentTriggeredByUser,
	"SYSTEM": CryptoAssessmentTriggeredBySystem,
}

var mappingCryptoAssessmentTriggeredByEnumLowerCase = map[string]CryptoAssessmentTriggeredByEnum{
	"user":   CryptoAssessmentTriggeredByUser,
	"system": CryptoAssessmentTriggeredBySystem,
}

// GetCryptoAssessmentTriggeredByEnumValues Enumerates the set of values for CryptoAssessmentTriggeredByEnum
func GetCryptoAssessmentTriggeredByEnumValues() []CryptoAssessmentTriggeredByEnum {
	values := make([]CryptoAssessmentTriggeredByEnum, 0)
	for _, v := range mappingCryptoAssessmentTriggeredByEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoAssessmentTriggeredByEnumStringValues Enumerates the set of values in String for CryptoAssessmentTriggeredByEnum
func GetCryptoAssessmentTriggeredByEnumStringValues() []string {
	return []string{
		"USER",
		"SYSTEM",
	}
}

// GetMappingCryptoAssessmentTriggeredByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoAssessmentTriggeredByEnum(val string) (CryptoAssessmentTriggeredByEnum, bool) {
	enum, ok := mappingCryptoAssessmentTriggeredByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CryptoAssessmentTargetTypeEnum Enum with underlying type: string
type CryptoAssessmentTargetTypeEnum string

// Set of constants representing the allowable values for CryptoAssessmentTargetTypeEnum
const (
	CryptoAssessmentTargetTypeDatabase      CryptoAssessmentTargetTypeEnum = "TARGET_DATABASE"
	CryptoAssessmentTargetTypeDatabaseGroup CryptoAssessmentTargetTypeEnum = "TARGET_DATABASE_GROUP"
)

var mappingCryptoAssessmentTargetTypeEnum = map[string]CryptoAssessmentTargetTypeEnum{
	"TARGET_DATABASE":       CryptoAssessmentTargetTypeDatabase,
	"TARGET_DATABASE_GROUP": CryptoAssessmentTargetTypeDatabaseGroup,
}

var mappingCryptoAssessmentTargetTypeEnumLowerCase = map[string]CryptoAssessmentTargetTypeEnum{
	"target_database":       CryptoAssessmentTargetTypeDatabase,
	"target_database_group": CryptoAssessmentTargetTypeDatabaseGroup,
}

// GetCryptoAssessmentTargetTypeEnumValues Enumerates the set of values for CryptoAssessmentTargetTypeEnum
func GetCryptoAssessmentTargetTypeEnumValues() []CryptoAssessmentTargetTypeEnum {
	values := make([]CryptoAssessmentTargetTypeEnum, 0)
	for _, v := range mappingCryptoAssessmentTargetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoAssessmentTargetTypeEnumStringValues Enumerates the set of values in String for CryptoAssessmentTargetTypeEnum
func GetCryptoAssessmentTargetTypeEnumStringValues() []string {
	return []string{
		"TARGET_DATABASE",
		"TARGET_DATABASE_GROUP",
	}
}

// GetMappingCryptoAssessmentTargetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoAssessmentTargetTypeEnum(val string) (CryptoAssessmentTargetTypeEnum, bool) {
	enum, ok := mappingCryptoAssessmentTargetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
