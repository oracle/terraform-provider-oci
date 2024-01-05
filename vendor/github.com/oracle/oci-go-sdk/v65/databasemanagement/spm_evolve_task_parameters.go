// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SpmEvolveTaskParameters The set of parameters used in an SPM evolve task.
type SpmEvolveTaskParameters struct {

	// Determines which sources to search for additional plans.
	AlternatePlanSources []SpmEvolveTaskParametersAlternatePlanSourcesEnum `mandatory:"false" json:"alternatePlanSources,omitempty"`

	// Determines which alternative plans should be loaded.
	AlternatePlanBaselines []SpmEvolveTaskParametersAlternatePlanBaselinesEnum `mandatory:"false" json:"alternatePlanBaselines,omitempty"`

	// Specifies the maximum number of plans to load in total (that is, not
	// the limit for each SQL statement). A value of zero indicates `UNLIMITED`
	// number of plans.
	AlternatePlanLimit *int `mandatory:"false" json:"alternatePlanLimit"`

	// Specifies whether to accept recommended plans automatically.
	ArePlansAutoAccepted *bool `mandatory:"false" json:"arePlansAutoAccepted"`

	// The global time limit in seconds. This is the total time allowed for the task.
	AllowedTimeLimit *int `mandatory:"false" json:"allowedTimeLimit"`
}

func (m SpmEvolveTaskParameters) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SpmEvolveTaskParameters) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.AlternatePlanSources {
		if _, ok := GetMappingSpmEvolveTaskParametersAlternatePlanSourcesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AlternatePlanSources: %s. Supported values are: %s.", val, strings.Join(GetSpmEvolveTaskParametersAlternatePlanSourcesEnumStringValues(), ",")))
		}
	}

	for _, val := range m.AlternatePlanBaselines {
		if _, ok := GetMappingSpmEvolveTaskParametersAlternatePlanBaselinesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AlternatePlanBaselines: %s. Supported values are: %s.", val, strings.Join(GetSpmEvolveTaskParametersAlternatePlanBaselinesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SpmEvolveTaskParametersAlternatePlanSourcesEnum Enum with underlying type: string
type SpmEvolveTaskParametersAlternatePlanSourcesEnum string

// Set of constants representing the allowable values for SpmEvolveTaskParametersAlternatePlanSourcesEnum
const (
	SpmEvolveTaskParametersAlternatePlanSourcesAuto                        SpmEvolveTaskParametersAlternatePlanSourcesEnum = "AUTO"
	SpmEvolveTaskParametersAlternatePlanSourcesAutomaticWorkloadRepository SpmEvolveTaskParametersAlternatePlanSourcesEnum = "AUTOMATIC_WORKLOAD_REPOSITORY"
	SpmEvolveTaskParametersAlternatePlanSourcesCursorCache                 SpmEvolveTaskParametersAlternatePlanSourcesEnum = "CURSOR_CACHE"
	SpmEvolveTaskParametersAlternatePlanSourcesSqlTuningSet                SpmEvolveTaskParametersAlternatePlanSourcesEnum = "SQL_TUNING_SET"
)

var mappingSpmEvolveTaskParametersAlternatePlanSourcesEnum = map[string]SpmEvolveTaskParametersAlternatePlanSourcesEnum{
	"AUTO":                          SpmEvolveTaskParametersAlternatePlanSourcesAuto,
	"AUTOMATIC_WORKLOAD_REPOSITORY": SpmEvolveTaskParametersAlternatePlanSourcesAutomaticWorkloadRepository,
	"CURSOR_CACHE":                  SpmEvolveTaskParametersAlternatePlanSourcesCursorCache,
	"SQL_TUNING_SET":                SpmEvolveTaskParametersAlternatePlanSourcesSqlTuningSet,
}

var mappingSpmEvolveTaskParametersAlternatePlanSourcesEnumLowerCase = map[string]SpmEvolveTaskParametersAlternatePlanSourcesEnum{
	"auto":                          SpmEvolveTaskParametersAlternatePlanSourcesAuto,
	"automatic_workload_repository": SpmEvolveTaskParametersAlternatePlanSourcesAutomaticWorkloadRepository,
	"cursor_cache":                  SpmEvolveTaskParametersAlternatePlanSourcesCursorCache,
	"sql_tuning_set":                SpmEvolveTaskParametersAlternatePlanSourcesSqlTuningSet,
}

// GetSpmEvolveTaskParametersAlternatePlanSourcesEnumValues Enumerates the set of values for SpmEvolveTaskParametersAlternatePlanSourcesEnum
func GetSpmEvolveTaskParametersAlternatePlanSourcesEnumValues() []SpmEvolveTaskParametersAlternatePlanSourcesEnum {
	values := make([]SpmEvolveTaskParametersAlternatePlanSourcesEnum, 0)
	for _, v := range mappingSpmEvolveTaskParametersAlternatePlanSourcesEnum {
		values = append(values, v)
	}
	return values
}

// GetSpmEvolveTaskParametersAlternatePlanSourcesEnumStringValues Enumerates the set of values in String for SpmEvolveTaskParametersAlternatePlanSourcesEnum
func GetSpmEvolveTaskParametersAlternatePlanSourcesEnumStringValues() []string {
	return []string{
		"AUTO",
		"AUTOMATIC_WORKLOAD_REPOSITORY",
		"CURSOR_CACHE",
		"SQL_TUNING_SET",
	}
}

// GetMappingSpmEvolveTaskParametersAlternatePlanSourcesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSpmEvolveTaskParametersAlternatePlanSourcesEnum(val string) (SpmEvolveTaskParametersAlternatePlanSourcesEnum, bool) {
	enum, ok := mappingSpmEvolveTaskParametersAlternatePlanSourcesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SpmEvolveTaskParametersAlternatePlanBaselinesEnum Enum with underlying type: string
type SpmEvolveTaskParametersAlternatePlanBaselinesEnum string

// Set of constants representing the allowable values for SpmEvolveTaskParametersAlternatePlanBaselinesEnum
const (
	SpmEvolveTaskParametersAlternatePlanBaselinesAuto     SpmEvolveTaskParametersAlternatePlanBaselinesEnum = "AUTO"
	SpmEvolveTaskParametersAlternatePlanBaselinesExisting SpmEvolveTaskParametersAlternatePlanBaselinesEnum = "EXISTING"
	SpmEvolveTaskParametersAlternatePlanBaselinesNew      SpmEvolveTaskParametersAlternatePlanBaselinesEnum = "NEW"
)

var mappingSpmEvolveTaskParametersAlternatePlanBaselinesEnum = map[string]SpmEvolveTaskParametersAlternatePlanBaselinesEnum{
	"AUTO":     SpmEvolveTaskParametersAlternatePlanBaselinesAuto,
	"EXISTING": SpmEvolveTaskParametersAlternatePlanBaselinesExisting,
	"NEW":      SpmEvolveTaskParametersAlternatePlanBaselinesNew,
}

var mappingSpmEvolveTaskParametersAlternatePlanBaselinesEnumLowerCase = map[string]SpmEvolveTaskParametersAlternatePlanBaselinesEnum{
	"auto":     SpmEvolveTaskParametersAlternatePlanBaselinesAuto,
	"existing": SpmEvolveTaskParametersAlternatePlanBaselinesExisting,
	"new":      SpmEvolveTaskParametersAlternatePlanBaselinesNew,
}

// GetSpmEvolveTaskParametersAlternatePlanBaselinesEnumValues Enumerates the set of values for SpmEvolveTaskParametersAlternatePlanBaselinesEnum
func GetSpmEvolveTaskParametersAlternatePlanBaselinesEnumValues() []SpmEvolveTaskParametersAlternatePlanBaselinesEnum {
	values := make([]SpmEvolveTaskParametersAlternatePlanBaselinesEnum, 0)
	for _, v := range mappingSpmEvolveTaskParametersAlternatePlanBaselinesEnum {
		values = append(values, v)
	}
	return values
}

// GetSpmEvolveTaskParametersAlternatePlanBaselinesEnumStringValues Enumerates the set of values in String for SpmEvolveTaskParametersAlternatePlanBaselinesEnum
func GetSpmEvolveTaskParametersAlternatePlanBaselinesEnumStringValues() []string {
	return []string{
		"AUTO",
		"EXISTING",
		"NEW",
	}
}

// GetMappingSpmEvolveTaskParametersAlternatePlanBaselinesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSpmEvolveTaskParametersAlternatePlanBaselinesEnum(val string) (SpmEvolveTaskParametersAlternatePlanBaselinesEnum, bool) {
	enum, ok := mappingSpmEvolveTaskParametersAlternatePlanBaselinesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
