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

// SubsetDataDetails Details required to perform data subsetting on a target database using a subsetting policy
type SubsetDataDetails struct {
	TargetCredentials *Credentials `mandatory:"true" json:"targetCredentials"`

	// The OCID of the target database to be subsetted. If it's not provided, the value of the
	// targetId attribute in the SubsettingPolicy resource is used
	TargetId *string `mandatory:"false" json:"targetId"`

	// Indicates whether masking should be triggered after successful subsetting
	Masking SubsetDataDetailsMaskingEnum `mandatory:"false" json:"masking,omitempty"`

	// Indicates if the request is to rerun the previously failed subsetting job
	IsRerun *bool `mandatory:"false" json:"isRerun"`

	// Specifies the step from which subsetting needs to be rerun. This param will be used only when isRerun attribute is true.
	// If PRE_SUBSETTING_SCRIPT is passed, it will rerun the pre-subsetting script, followed by subsetting, and then the post-subsetting script.
	// If POST_SUBSETTING_SCRIPT is passed, it will rerun only the post-subsetting script.
	// If this field is not set and isRerun is set to true, then it will default to the last failed step.
	ReRunFromStep SubsetDataDetailsReRunFromStepEnum `mandatory:"false" json:"reRunFromStep,omitempty"`

	// The tablespace that should be used to create the temporary tables for data subsetting.
	// If no tablespace is provided, the DEFAULT tablespace is used.
	Tablespace *string `mandatory:"false" json:"tablespace"`

	// Indicates if redo logging is enabled during a subsetting operation. Set this attribute to true to
	// enable redo logging. If set as false, subsetting disables redo logging and flashback logging to purge any original
	// data from logs. However, in certain circumstances when you only want to test subsetting, rollback changes, and retry,
	// you could enable logging and use a flashback database to retrieve the original data after it has been subsetted.
	// If it's not provided, the value of the isRedoLoggingEnabled attribute in the SubsettingPolicy resource is used.
	IsRedoLoggingEnabled *bool `mandatory:"false" json:"isRedoLoggingEnabled"`

	// Indicates if statistics gathering is enabled. Set this attribute to false to disable statistics
	// gathering. The subsetting process gathers statistics on subset database tables after subsetting completes.
	// If it's not provided, the value of the isRefreshStatsEnabled attribute in the SubsettingPolicy resource is used.
	IsRefreshStatsEnabled *bool `mandatory:"false" json:"isRefreshStatsEnabled"`

	// Specifies options to enable parallel execution when running data subsetting. Allowed values are 'NONE' (no parallelism),
	// 'DEFAULT' (the Oracle Database computes the optimum degree of parallelism) or an integer value to be used as the degree
	// of parallelism. Parallel execution helps effectively use multiple CPUs and improve subsetting performance. Refer to the
	// Oracle Database parallel execution framework when choosing an explicit degree of parallelism.
	// https://www.oracle.com/pls/topic/lookup?ctx=dblatest&en/database/oracle/oracle-database&id=VLDBG-GUID-3E2AE088-2505-465E-A8B2-AC38813EA355
	// If it's not provided, the value of the parallelDegree attribute in the SubsettingPolicy resource is used.
	ParallelDegree *string `mandatory:"false" json:"parallelDegree"`

	// Specifies how to recompile invalid objects post data subsetting. Allowed values are 'SERIAL' (recompile in serial),
	// 'PARALLEL' (recompile in parallel), 'NONE' (do not recompile). If it's set to PARALLEL, the value of parallelDegree
	// attribute is used. Use the built-in UTL_RECOMP package to recompile any remaining invalid objects after subsetting completes.
	// If it's not provided, the value of the recompile attribute in the SubsettingPolicy resource is used.
	Recompile SubsettingPolicyRecompileEnum `mandatory:"false" json:"recompile,omitempty"`
}

func (m SubsetDataDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubsetDataDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSubsetDataDetailsMaskingEnum(string(m.Masking)); !ok && m.Masking != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Masking: %s. Supported values are: %s.", m.Masking, strings.Join(GetSubsetDataDetailsMaskingEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSubsetDataDetailsReRunFromStepEnum(string(m.ReRunFromStep)); !ok && m.ReRunFromStep != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReRunFromStep: %s. Supported values are: %s.", m.ReRunFromStep, strings.Join(GetSubsetDataDetailsReRunFromStepEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSubsettingPolicyRecompileEnum(string(m.Recompile)); !ok && m.Recompile != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Recompile: %s. Supported values are: %s.", m.Recompile, strings.Join(GetSubsettingPolicyRecompileEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SubsetDataDetailsMaskingEnum Enum with underlying type: string
type SubsetDataDetailsMaskingEnum string

// Set of constants representing the allowable values for SubsetDataDetailsMaskingEnum
const (
	SubsetDataDetailsMaskingEnabled  SubsetDataDetailsMaskingEnum = "ENABLED"
	SubsetDataDetailsMaskingDisabled SubsetDataDetailsMaskingEnum = "DISABLED"
)

var mappingSubsetDataDetailsMaskingEnum = map[string]SubsetDataDetailsMaskingEnum{
	"ENABLED":  SubsetDataDetailsMaskingEnabled,
	"DISABLED": SubsetDataDetailsMaskingDisabled,
}

var mappingSubsetDataDetailsMaskingEnumLowerCase = map[string]SubsetDataDetailsMaskingEnum{
	"enabled":  SubsetDataDetailsMaskingEnabled,
	"disabled": SubsetDataDetailsMaskingDisabled,
}

// GetSubsetDataDetailsMaskingEnumValues Enumerates the set of values for SubsetDataDetailsMaskingEnum
func GetSubsetDataDetailsMaskingEnumValues() []SubsetDataDetailsMaskingEnum {
	values := make([]SubsetDataDetailsMaskingEnum, 0)
	for _, v := range mappingSubsetDataDetailsMaskingEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsetDataDetailsMaskingEnumStringValues Enumerates the set of values in String for SubsetDataDetailsMaskingEnum
func GetSubsetDataDetailsMaskingEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingSubsetDataDetailsMaskingEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsetDataDetailsMaskingEnum(val string) (SubsetDataDetailsMaskingEnum, bool) {
	enum, ok := mappingSubsetDataDetailsMaskingEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SubsetDataDetailsReRunFromStepEnum Enum with underlying type: string
type SubsetDataDetailsReRunFromStepEnum string

// Set of constants representing the allowable values for SubsetDataDetailsReRunFromStepEnum
const (
	SubsetDataDetailsReRunFromStepPreSubsettingScript  SubsetDataDetailsReRunFromStepEnum = "PRE_SUBSETTING_SCRIPT"
	SubsetDataDetailsReRunFromStepPostSubsettingScript SubsetDataDetailsReRunFromStepEnum = "POST_SUBSETTING_SCRIPT"
)

var mappingSubsetDataDetailsReRunFromStepEnum = map[string]SubsetDataDetailsReRunFromStepEnum{
	"PRE_SUBSETTING_SCRIPT":  SubsetDataDetailsReRunFromStepPreSubsettingScript,
	"POST_SUBSETTING_SCRIPT": SubsetDataDetailsReRunFromStepPostSubsettingScript,
}

var mappingSubsetDataDetailsReRunFromStepEnumLowerCase = map[string]SubsetDataDetailsReRunFromStepEnum{
	"pre_subsetting_script":  SubsetDataDetailsReRunFromStepPreSubsettingScript,
	"post_subsetting_script": SubsetDataDetailsReRunFromStepPostSubsettingScript,
}

// GetSubsetDataDetailsReRunFromStepEnumValues Enumerates the set of values for SubsetDataDetailsReRunFromStepEnum
func GetSubsetDataDetailsReRunFromStepEnumValues() []SubsetDataDetailsReRunFromStepEnum {
	values := make([]SubsetDataDetailsReRunFromStepEnum, 0)
	for _, v := range mappingSubsetDataDetailsReRunFromStepEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsetDataDetailsReRunFromStepEnumStringValues Enumerates the set of values in String for SubsetDataDetailsReRunFromStepEnum
func GetSubsetDataDetailsReRunFromStepEnumStringValues() []string {
	return []string{
		"PRE_SUBSETTING_SCRIPT",
		"POST_SUBSETTING_SCRIPT",
	}
}

// GetMappingSubsetDataDetailsReRunFromStepEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsetDataDetailsReRunFromStepEnum(val string) (SubsetDataDetailsReRunFromStepEnum, bool) {
	enum, ok := mappingSubsetDataDetailsReRunFromStepEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
