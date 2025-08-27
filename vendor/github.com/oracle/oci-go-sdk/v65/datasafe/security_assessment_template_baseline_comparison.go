// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// SecurityAssessmentTemplateBaselineComparison Provides a list of the differences in a comparison of the security assessment with the template baseline value.
type SecurityAssessmentTemplateBaselineComparison struct {

	// The OCID of the security assessment that is being compared with a template baseline security assessment.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the security assessment that is set as a template baseline.
	TemplateBaselineId *string `mandatory:"true" json:"templateBaselineId"`

	// The current state of the security assessment comparison.
	LifecycleState SecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the security assessment comparison was created. Conforms to the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The display name of the security assessment that is set as a template baseline.
	TemplateBaselineName *string `mandatory:"false" json:"templateBaselineName"`

	// A comparison between findings belonging to Auditing category.
	Auditing []TemplateBaselineDiffs `mandatory:"false" json:"auditing"`

	// A comparison between findings belonging to Authorization Control category.
	AuthorizationControl []TemplateBaselineDiffs `mandatory:"false" json:"authorizationControl"`

	// Comparison between findings belonging to Data Encryption category.
	DataEncryption []TemplateBaselineDiffs `mandatory:"false" json:"dataEncryption"`

	// Comparison between findings belonging to Database Configuration category.
	DbConfiguration []TemplateBaselineDiffs `mandatory:"false" json:"dbConfiguration"`

	// Comparison between findings belonging to Fine-Grained Access Control category.
	FineGrainedAccessControl []TemplateBaselineDiffs `mandatory:"false" json:"fineGrainedAccessControl"`

	// Comparison between findings belonging to Privileges and Roles category.
	PrivilegesAndRoles []TemplateBaselineDiffs `mandatory:"false" json:"privilegesAndRoles"`

	// Comparison between findings belonging to User Accounts category.
	UserAccounts []TemplateBaselineDiffs `mandatory:"false" json:"userAccounts"`
}

func (m SecurityAssessmentTemplateBaselineComparison) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityAssessmentTemplateBaselineComparison) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSecurityAssessmentTemplateBaselineComparisonLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum Enum with underlying type: string
type SecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum string

// Set of constants representing the allowable values for SecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum
const (
	SecurityAssessmentTemplateBaselineComparisonLifecycleStateInProgress SecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum = "IN_PROGRESS"
	SecurityAssessmentTemplateBaselineComparisonLifecycleStateSucceeded  SecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum = "SUCCEEDED"
	SecurityAssessmentTemplateBaselineComparisonLifecycleStateFailed     SecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum = "FAILED"
	SecurityAssessmentTemplateBaselineComparisonLifecycleStateDeleted    SecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum = "DELETED"
	SecurityAssessmentTemplateBaselineComparisonLifecycleStateDeleting   SecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum = "DELETING"
)

var mappingSecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum = map[string]SecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum{
	"IN_PROGRESS": SecurityAssessmentTemplateBaselineComparisonLifecycleStateInProgress,
	"SUCCEEDED":   SecurityAssessmentTemplateBaselineComparisonLifecycleStateSucceeded,
	"FAILED":      SecurityAssessmentTemplateBaselineComparisonLifecycleStateFailed,
	"DELETED":     SecurityAssessmentTemplateBaselineComparisonLifecycleStateDeleted,
	"DELETING":    SecurityAssessmentTemplateBaselineComparisonLifecycleStateDeleting,
}

var mappingSecurityAssessmentTemplateBaselineComparisonLifecycleStateEnumLowerCase = map[string]SecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum{
	"in_progress": SecurityAssessmentTemplateBaselineComparisonLifecycleStateInProgress,
	"succeeded":   SecurityAssessmentTemplateBaselineComparisonLifecycleStateSucceeded,
	"failed":      SecurityAssessmentTemplateBaselineComparisonLifecycleStateFailed,
	"deleted":     SecurityAssessmentTemplateBaselineComparisonLifecycleStateDeleted,
	"deleting":    SecurityAssessmentTemplateBaselineComparisonLifecycleStateDeleting,
}

// GetSecurityAssessmentTemplateBaselineComparisonLifecycleStateEnumValues Enumerates the set of values for SecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum
func GetSecurityAssessmentTemplateBaselineComparisonLifecycleStateEnumValues() []SecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum {
	values := make([]SecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum, 0)
	for _, v := range mappingSecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityAssessmentTemplateBaselineComparisonLifecycleStateEnumStringValues Enumerates the set of values in String for SecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum
func GetSecurityAssessmentTemplateBaselineComparisonLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"DELETED",
		"DELETING",
	}
}

// GetMappingSecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum(val string) (SecurityAssessmentTemplateBaselineComparisonLifecycleStateEnum, bool) {
	enum, ok := mappingSecurityAssessmentTemplateBaselineComparisonLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
