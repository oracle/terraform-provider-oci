// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlPlanBaseline The details of a SQL plan baseline.
type SqlPlanBaseline struct {

	// The unique plan identifier.
	PlanName *string `mandatory:"true" json:"planName"`

	// The unique SQL identifier.
	SqlHandle *string `mandatory:"true" json:"sqlHandle"`

	// The SQL text.
	SqlText *string `mandatory:"true" json:"sqlText"`

	// The date and time when the plan baseline was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The execution plan for the SQL statement.
	ExecutionPlan *string `mandatory:"true" json:"executionPlan"`

	// The origin of the SQL plan baseline.
	Origin SqlPlanBaselineOriginEnum `mandatory:"false" json:"origin,omitempty"`

	// The date and time when the plan baseline was last modified.
	TimeLastModified *common.SDKTime `mandatory:"false" json:"timeLastModified"`

	// The date and time when the plan baseline was last executed.
	// **Note:** For performance reasons, database does not update this value
	// immediately after each execution of the plan baseline. Therefore, the plan
	// baseline may have been executed more recently than this value indicates.
	TimeLastExecuted *common.SDKTime `mandatory:"false" json:"timeLastExecuted"`

	// Indicates whether the plan baseline is enabled (`YES`) or disabled (`NO`).
	Enabled SqlPlanBaselineEnabledEnum `mandatory:"false" json:"enabled,omitempty"`

	// Indicates whether the plan baseline is accepted (`YES`) or not (`NO`).
	Accepted SqlPlanBaselineAcceptedEnum `mandatory:"false" json:"accepted,omitempty"`

	// Indicates whether the plan baseline is fixed (`YES`) or not (`NO`).
	Fixed SqlPlanBaselineFixedEnum `mandatory:"false" json:"fixed,omitempty"`

	// Indicates whether the optimizer was able to reproduce the plan (`YES`) or not (`NO`).
	// The value is set to `YES` when a plan is initially added to the plan baseline.
	Reproduced SqlPlanBaselineReproducedEnum `mandatory:"false" json:"reproduced,omitempty"`

	// Indicates whether the plan baseline is auto-purged (`YES`) or not (`NO`).
	AutoPurge SqlPlanBaselineAutoPurgeEnum `mandatory:"false" json:"autoPurge,omitempty"`

	// Indicates whether a plan that is automatically captured by SQL plan management is marked adaptive or not.
	// When a new adaptive plan is found for a SQL statement that has an existing SQL plan baseline, that new plan
	// will be added to the SQL plan baseline as an unaccepted plan, and the `ADAPTIVE` property will be marked `YES`.
	// When this new plan is verified (either manually or via the auto evolve task), the plan will be test executed
	// and the final plan determined at execution will become an accepted plan if its performance is better than
	// the existing plan baseline. At this point, the value of the `ADAPTIVE` property is set to `NO` since the plan
	// is no longer adaptive, but resolved.
	Adaptive SqlPlanBaselineAdaptiveEnum `mandatory:"false" json:"adaptive,omitempty"`

	// The application module name.
	Module *string `mandatory:"false" json:"module"`

	// The application action.
	Action *string `mandatory:"false" json:"action"`
}

func (m SqlPlanBaseline) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlPlanBaseline) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSqlPlanBaselineOriginEnum(string(m.Origin)); !ok && m.Origin != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Origin: %s. Supported values are: %s.", m.Origin, strings.Join(GetSqlPlanBaselineOriginEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlPlanBaselineEnabledEnum(string(m.Enabled)); !ok && m.Enabled != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Enabled: %s. Supported values are: %s.", m.Enabled, strings.Join(GetSqlPlanBaselineEnabledEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlPlanBaselineAcceptedEnum(string(m.Accepted)); !ok && m.Accepted != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Accepted: %s. Supported values are: %s.", m.Accepted, strings.Join(GetSqlPlanBaselineAcceptedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlPlanBaselineFixedEnum(string(m.Fixed)); !ok && m.Fixed != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fixed: %s. Supported values are: %s.", m.Fixed, strings.Join(GetSqlPlanBaselineFixedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlPlanBaselineReproducedEnum(string(m.Reproduced)); !ok && m.Reproduced != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Reproduced: %s. Supported values are: %s.", m.Reproduced, strings.Join(GetSqlPlanBaselineReproducedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlPlanBaselineAutoPurgeEnum(string(m.AutoPurge)); !ok && m.AutoPurge != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutoPurge: %s. Supported values are: %s.", m.AutoPurge, strings.Join(GetSqlPlanBaselineAutoPurgeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlPlanBaselineAdaptiveEnum(string(m.Adaptive)); !ok && m.Adaptive != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Adaptive: %s. Supported values are: %s.", m.Adaptive, strings.Join(GetSqlPlanBaselineAdaptiveEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlPlanBaselineEnabledEnum Enum with underlying type: string
type SqlPlanBaselineEnabledEnum string

// Set of constants representing the allowable values for SqlPlanBaselineEnabledEnum
const (
	SqlPlanBaselineEnabledYes SqlPlanBaselineEnabledEnum = "YES"
	SqlPlanBaselineEnabledNo  SqlPlanBaselineEnabledEnum = "NO"
)

var mappingSqlPlanBaselineEnabledEnum = map[string]SqlPlanBaselineEnabledEnum{
	"YES": SqlPlanBaselineEnabledYes,
	"NO":  SqlPlanBaselineEnabledNo,
}

var mappingSqlPlanBaselineEnabledEnumLowerCase = map[string]SqlPlanBaselineEnabledEnum{
	"yes": SqlPlanBaselineEnabledYes,
	"no":  SqlPlanBaselineEnabledNo,
}

// GetSqlPlanBaselineEnabledEnumValues Enumerates the set of values for SqlPlanBaselineEnabledEnum
func GetSqlPlanBaselineEnabledEnumValues() []SqlPlanBaselineEnabledEnum {
	values := make([]SqlPlanBaselineEnabledEnum, 0)
	for _, v := range mappingSqlPlanBaselineEnabledEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlPlanBaselineEnabledEnumStringValues Enumerates the set of values in String for SqlPlanBaselineEnabledEnum
func GetSqlPlanBaselineEnabledEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingSqlPlanBaselineEnabledEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlPlanBaselineEnabledEnum(val string) (SqlPlanBaselineEnabledEnum, bool) {
	enum, ok := mappingSqlPlanBaselineEnabledEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlPlanBaselineAcceptedEnum Enum with underlying type: string
type SqlPlanBaselineAcceptedEnum string

// Set of constants representing the allowable values for SqlPlanBaselineAcceptedEnum
const (
	SqlPlanBaselineAcceptedYes SqlPlanBaselineAcceptedEnum = "YES"
	SqlPlanBaselineAcceptedNo  SqlPlanBaselineAcceptedEnum = "NO"
)

var mappingSqlPlanBaselineAcceptedEnum = map[string]SqlPlanBaselineAcceptedEnum{
	"YES": SqlPlanBaselineAcceptedYes,
	"NO":  SqlPlanBaselineAcceptedNo,
}

var mappingSqlPlanBaselineAcceptedEnumLowerCase = map[string]SqlPlanBaselineAcceptedEnum{
	"yes": SqlPlanBaselineAcceptedYes,
	"no":  SqlPlanBaselineAcceptedNo,
}

// GetSqlPlanBaselineAcceptedEnumValues Enumerates the set of values for SqlPlanBaselineAcceptedEnum
func GetSqlPlanBaselineAcceptedEnumValues() []SqlPlanBaselineAcceptedEnum {
	values := make([]SqlPlanBaselineAcceptedEnum, 0)
	for _, v := range mappingSqlPlanBaselineAcceptedEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlPlanBaselineAcceptedEnumStringValues Enumerates the set of values in String for SqlPlanBaselineAcceptedEnum
func GetSqlPlanBaselineAcceptedEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingSqlPlanBaselineAcceptedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlPlanBaselineAcceptedEnum(val string) (SqlPlanBaselineAcceptedEnum, bool) {
	enum, ok := mappingSqlPlanBaselineAcceptedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlPlanBaselineFixedEnum Enum with underlying type: string
type SqlPlanBaselineFixedEnum string

// Set of constants representing the allowable values for SqlPlanBaselineFixedEnum
const (
	SqlPlanBaselineFixedYes SqlPlanBaselineFixedEnum = "YES"
	SqlPlanBaselineFixedNo  SqlPlanBaselineFixedEnum = "NO"
)

var mappingSqlPlanBaselineFixedEnum = map[string]SqlPlanBaselineFixedEnum{
	"YES": SqlPlanBaselineFixedYes,
	"NO":  SqlPlanBaselineFixedNo,
}

var mappingSqlPlanBaselineFixedEnumLowerCase = map[string]SqlPlanBaselineFixedEnum{
	"yes": SqlPlanBaselineFixedYes,
	"no":  SqlPlanBaselineFixedNo,
}

// GetSqlPlanBaselineFixedEnumValues Enumerates the set of values for SqlPlanBaselineFixedEnum
func GetSqlPlanBaselineFixedEnumValues() []SqlPlanBaselineFixedEnum {
	values := make([]SqlPlanBaselineFixedEnum, 0)
	for _, v := range mappingSqlPlanBaselineFixedEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlPlanBaselineFixedEnumStringValues Enumerates the set of values in String for SqlPlanBaselineFixedEnum
func GetSqlPlanBaselineFixedEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingSqlPlanBaselineFixedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlPlanBaselineFixedEnum(val string) (SqlPlanBaselineFixedEnum, bool) {
	enum, ok := mappingSqlPlanBaselineFixedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlPlanBaselineReproducedEnum Enum with underlying type: string
type SqlPlanBaselineReproducedEnum string

// Set of constants representing the allowable values for SqlPlanBaselineReproducedEnum
const (
	SqlPlanBaselineReproducedYes SqlPlanBaselineReproducedEnum = "YES"
	SqlPlanBaselineReproducedNo  SqlPlanBaselineReproducedEnum = "NO"
)

var mappingSqlPlanBaselineReproducedEnum = map[string]SqlPlanBaselineReproducedEnum{
	"YES": SqlPlanBaselineReproducedYes,
	"NO":  SqlPlanBaselineReproducedNo,
}

var mappingSqlPlanBaselineReproducedEnumLowerCase = map[string]SqlPlanBaselineReproducedEnum{
	"yes": SqlPlanBaselineReproducedYes,
	"no":  SqlPlanBaselineReproducedNo,
}

// GetSqlPlanBaselineReproducedEnumValues Enumerates the set of values for SqlPlanBaselineReproducedEnum
func GetSqlPlanBaselineReproducedEnumValues() []SqlPlanBaselineReproducedEnum {
	values := make([]SqlPlanBaselineReproducedEnum, 0)
	for _, v := range mappingSqlPlanBaselineReproducedEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlPlanBaselineReproducedEnumStringValues Enumerates the set of values in String for SqlPlanBaselineReproducedEnum
func GetSqlPlanBaselineReproducedEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingSqlPlanBaselineReproducedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlPlanBaselineReproducedEnum(val string) (SqlPlanBaselineReproducedEnum, bool) {
	enum, ok := mappingSqlPlanBaselineReproducedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlPlanBaselineAutoPurgeEnum Enum with underlying type: string
type SqlPlanBaselineAutoPurgeEnum string

// Set of constants representing the allowable values for SqlPlanBaselineAutoPurgeEnum
const (
	SqlPlanBaselineAutoPurgeYes SqlPlanBaselineAutoPurgeEnum = "YES"
	SqlPlanBaselineAutoPurgeNo  SqlPlanBaselineAutoPurgeEnum = "NO"
)

var mappingSqlPlanBaselineAutoPurgeEnum = map[string]SqlPlanBaselineAutoPurgeEnum{
	"YES": SqlPlanBaselineAutoPurgeYes,
	"NO":  SqlPlanBaselineAutoPurgeNo,
}

var mappingSqlPlanBaselineAutoPurgeEnumLowerCase = map[string]SqlPlanBaselineAutoPurgeEnum{
	"yes": SqlPlanBaselineAutoPurgeYes,
	"no":  SqlPlanBaselineAutoPurgeNo,
}

// GetSqlPlanBaselineAutoPurgeEnumValues Enumerates the set of values for SqlPlanBaselineAutoPurgeEnum
func GetSqlPlanBaselineAutoPurgeEnumValues() []SqlPlanBaselineAutoPurgeEnum {
	values := make([]SqlPlanBaselineAutoPurgeEnum, 0)
	for _, v := range mappingSqlPlanBaselineAutoPurgeEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlPlanBaselineAutoPurgeEnumStringValues Enumerates the set of values in String for SqlPlanBaselineAutoPurgeEnum
func GetSqlPlanBaselineAutoPurgeEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingSqlPlanBaselineAutoPurgeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlPlanBaselineAutoPurgeEnum(val string) (SqlPlanBaselineAutoPurgeEnum, bool) {
	enum, ok := mappingSqlPlanBaselineAutoPurgeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlPlanBaselineAdaptiveEnum Enum with underlying type: string
type SqlPlanBaselineAdaptiveEnum string

// Set of constants representing the allowable values for SqlPlanBaselineAdaptiveEnum
const (
	SqlPlanBaselineAdaptiveYes SqlPlanBaselineAdaptiveEnum = "YES"
	SqlPlanBaselineAdaptiveNo  SqlPlanBaselineAdaptiveEnum = "NO"
)

var mappingSqlPlanBaselineAdaptiveEnum = map[string]SqlPlanBaselineAdaptiveEnum{
	"YES": SqlPlanBaselineAdaptiveYes,
	"NO":  SqlPlanBaselineAdaptiveNo,
}

var mappingSqlPlanBaselineAdaptiveEnumLowerCase = map[string]SqlPlanBaselineAdaptiveEnum{
	"yes": SqlPlanBaselineAdaptiveYes,
	"no":  SqlPlanBaselineAdaptiveNo,
}

// GetSqlPlanBaselineAdaptiveEnumValues Enumerates the set of values for SqlPlanBaselineAdaptiveEnum
func GetSqlPlanBaselineAdaptiveEnumValues() []SqlPlanBaselineAdaptiveEnum {
	values := make([]SqlPlanBaselineAdaptiveEnum, 0)
	for _, v := range mappingSqlPlanBaselineAdaptiveEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlPlanBaselineAdaptiveEnumStringValues Enumerates the set of values in String for SqlPlanBaselineAdaptiveEnum
func GetSqlPlanBaselineAdaptiveEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingSqlPlanBaselineAdaptiveEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlPlanBaselineAdaptiveEnum(val string) (SqlPlanBaselineAdaptiveEnum, bool) {
	enum, ok := mappingSqlPlanBaselineAdaptiveEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
