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

// CryptoAssessmentSummary The summary of a crypto assessment.
type CryptoAssessmentSummary struct {

	// The OCID of the crypto assessment.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the crypto assessment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the crypto assessment was created, in RFC3339 format.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current lifecycle state of the crypto assessment.
	LifecycleState CryptoAssessmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Indicates whether this summary row is for a target database or a target database group.
	TargetType CryptoAssessmentTargetTypeEnum `mandatory:"true" json:"targetType"`

	// The display name of the crypto assessment.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the crypto assessment.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the crypto assessment was last updated, in RFC3339 format.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The date and time the crypto posture was last assessed, in RFC3339 format.
	TimeLastAssessed *common.SDKTime `mandatory:"false" json:"timeLastAssessed"`

	// Details about the current lifecycle state of the crypto assessment.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The type of this crypto assessment.
	Type CryptoAssessmentTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The actor that created the assessment.
	TriggeredBy CryptoAssessmentTriggeredByEnum `mandatory:"false" json:"triggeredBy,omitempty"`

	// The OCID of the target database. This is returned when `targetType` is `TARGET_DATABASE`.
	TargetId *string `mandatory:"false" json:"targetId"`

	// The OCID of the target database group. This is returned when `targetType` is `TARGET_DATABASE_GROUP`.
	TargetDatabaseGroupId *string `mandatory:"false" json:"targetDatabaseGroupId"`

	// The version of the assessed target database.
	DatabaseVersion *string `mandatory:"false" json:"databaseVersion"`

	// Overall posture category for the crypto assessment.
	PostureCategory CryptoPostureCategoryEnum `mandatory:"false" json:"postureCategory,omitempty"`

	// The schedule used to run the crypto assessment periodically. The schedule uses the format:
	// <version-string>;<version-specific-schedule>
	// For v1, the version-specific schedule format is:
	// <ss> <mm> <hh> <day-of-week> <day-of-month>
	// Specify either day-of-week for weekly schedules or day-of-month for monthly schedules. Do not specify both.
	// For monthly schedules, day-of-month must be between 1 and 28. If the service generates a default monthly schedule for an assessment created on day 29, 30, or 31 of a month, it uses day 28.
	Schedule *string `mandatory:"false" json:"schedule"`

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

func (m CryptoAssessmentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoAssessmentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCryptoAssessmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCryptoAssessmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoAssessmentTargetTypeEnum(string(m.TargetType)); !ok && m.TargetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetType: %s. Supported values are: %s.", m.TargetType, strings.Join(GetCryptoAssessmentTargetTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCryptoAssessmentTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetCryptoAssessmentTypeEnumStringValues(), ",")))
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
