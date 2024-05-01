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

// SqlPlanBaselineSummary The summary of a SQL plan baseline.
type SqlPlanBaselineSummary struct {

	// The unique plan identifier.
	PlanName *string `mandatory:"true" json:"planName"`

	// The unique SQL identifier.
	SqlHandle *string `mandatory:"true" json:"sqlHandle"`

	// The SQL text (truncated to the first 50 characters).
	SqlText *string `mandatory:"true" json:"sqlText"`

	// The date and time when the plan baseline was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

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
	Enabled SqlPlanBaselineSummaryEnabledEnum `mandatory:"false" json:"enabled,omitempty"`

	// Indicates whether the plan baseline is accepted (`YES`) or not (`NO`).
	Accepted SqlPlanBaselineSummaryAcceptedEnum `mandatory:"false" json:"accepted,omitempty"`

	// Indicates whether the plan baseline is fixed (`YES`) or not (`NO`).
	Fixed SqlPlanBaselineSummaryFixedEnum `mandatory:"false" json:"fixed,omitempty"`

	// Indicates whether the optimizer was able to reproduce the plan (`YES`) or not (`NO`).
	// The value is set to `YES` when a plan is initially added to the plan baseline.
	Reproduced SqlPlanBaselineSummaryReproducedEnum `mandatory:"false" json:"reproduced,omitempty"`

	// Indicates whether the plan baseline is auto-purged (`YES`) or not (`NO`).
	AutoPurge SqlPlanBaselineSummaryAutoPurgeEnum `mandatory:"false" json:"autoPurge,omitempty"`

	// Indicates whether a plan that is automatically captured by SQL plan management is marked adaptive or not.
	// When a new adaptive plan is found for a SQL statement that has an existing SQL plan baseline, that new plan
	// will be added to the SQL plan baseline as an unaccepted plan, and the `ADAPTIVE` property will be marked `YES`.
	// When this new plan is verified (either manually or via the auto evolve task), the plan will be test executed
	// and the final plan determined at execution will become an accepted plan if its performance is better than
	// the existing plan baseline. At this point, the value of the `ADAPTIVE` property is set to `NO` since the plan
	// is no longer adaptive, but resolved.
	Adaptive SqlPlanBaselineSummaryAdaptiveEnum `mandatory:"false" json:"adaptive,omitempty"`
}

func (m SqlPlanBaselineSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlPlanBaselineSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSqlPlanBaselineOriginEnum(string(m.Origin)); !ok && m.Origin != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Origin: %s. Supported values are: %s.", m.Origin, strings.Join(GetSqlPlanBaselineOriginEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlPlanBaselineSummaryEnabledEnum(string(m.Enabled)); !ok && m.Enabled != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Enabled: %s. Supported values are: %s.", m.Enabled, strings.Join(GetSqlPlanBaselineSummaryEnabledEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlPlanBaselineSummaryAcceptedEnum(string(m.Accepted)); !ok && m.Accepted != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Accepted: %s. Supported values are: %s.", m.Accepted, strings.Join(GetSqlPlanBaselineSummaryAcceptedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlPlanBaselineSummaryFixedEnum(string(m.Fixed)); !ok && m.Fixed != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fixed: %s. Supported values are: %s.", m.Fixed, strings.Join(GetSqlPlanBaselineSummaryFixedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlPlanBaselineSummaryReproducedEnum(string(m.Reproduced)); !ok && m.Reproduced != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Reproduced: %s. Supported values are: %s.", m.Reproduced, strings.Join(GetSqlPlanBaselineSummaryReproducedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlPlanBaselineSummaryAutoPurgeEnum(string(m.AutoPurge)); !ok && m.AutoPurge != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutoPurge: %s. Supported values are: %s.", m.AutoPurge, strings.Join(GetSqlPlanBaselineSummaryAutoPurgeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlPlanBaselineSummaryAdaptiveEnum(string(m.Adaptive)); !ok && m.Adaptive != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Adaptive: %s. Supported values are: %s.", m.Adaptive, strings.Join(GetSqlPlanBaselineSummaryAdaptiveEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlPlanBaselineSummaryEnabledEnum Enum with underlying type: string
type SqlPlanBaselineSummaryEnabledEnum string

// Set of constants representing the allowable values for SqlPlanBaselineSummaryEnabledEnum
const (
	SqlPlanBaselineSummaryEnabledYes SqlPlanBaselineSummaryEnabledEnum = "YES"
	SqlPlanBaselineSummaryEnabledNo  SqlPlanBaselineSummaryEnabledEnum = "NO"
)

var mappingSqlPlanBaselineSummaryEnabledEnum = map[string]SqlPlanBaselineSummaryEnabledEnum{
	"YES": SqlPlanBaselineSummaryEnabledYes,
	"NO":  SqlPlanBaselineSummaryEnabledNo,
}

var mappingSqlPlanBaselineSummaryEnabledEnumLowerCase = map[string]SqlPlanBaselineSummaryEnabledEnum{
	"yes": SqlPlanBaselineSummaryEnabledYes,
	"no":  SqlPlanBaselineSummaryEnabledNo,
}

// GetSqlPlanBaselineSummaryEnabledEnumValues Enumerates the set of values for SqlPlanBaselineSummaryEnabledEnum
func GetSqlPlanBaselineSummaryEnabledEnumValues() []SqlPlanBaselineSummaryEnabledEnum {
	values := make([]SqlPlanBaselineSummaryEnabledEnum, 0)
	for _, v := range mappingSqlPlanBaselineSummaryEnabledEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlPlanBaselineSummaryEnabledEnumStringValues Enumerates the set of values in String for SqlPlanBaselineSummaryEnabledEnum
func GetSqlPlanBaselineSummaryEnabledEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingSqlPlanBaselineSummaryEnabledEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlPlanBaselineSummaryEnabledEnum(val string) (SqlPlanBaselineSummaryEnabledEnum, bool) {
	enum, ok := mappingSqlPlanBaselineSummaryEnabledEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlPlanBaselineSummaryAcceptedEnum Enum with underlying type: string
type SqlPlanBaselineSummaryAcceptedEnum string

// Set of constants representing the allowable values for SqlPlanBaselineSummaryAcceptedEnum
const (
	SqlPlanBaselineSummaryAcceptedYes SqlPlanBaselineSummaryAcceptedEnum = "YES"
	SqlPlanBaselineSummaryAcceptedNo  SqlPlanBaselineSummaryAcceptedEnum = "NO"
)

var mappingSqlPlanBaselineSummaryAcceptedEnum = map[string]SqlPlanBaselineSummaryAcceptedEnum{
	"YES": SqlPlanBaselineSummaryAcceptedYes,
	"NO":  SqlPlanBaselineSummaryAcceptedNo,
}

var mappingSqlPlanBaselineSummaryAcceptedEnumLowerCase = map[string]SqlPlanBaselineSummaryAcceptedEnum{
	"yes": SqlPlanBaselineSummaryAcceptedYes,
	"no":  SqlPlanBaselineSummaryAcceptedNo,
}

// GetSqlPlanBaselineSummaryAcceptedEnumValues Enumerates the set of values for SqlPlanBaselineSummaryAcceptedEnum
func GetSqlPlanBaselineSummaryAcceptedEnumValues() []SqlPlanBaselineSummaryAcceptedEnum {
	values := make([]SqlPlanBaselineSummaryAcceptedEnum, 0)
	for _, v := range mappingSqlPlanBaselineSummaryAcceptedEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlPlanBaselineSummaryAcceptedEnumStringValues Enumerates the set of values in String for SqlPlanBaselineSummaryAcceptedEnum
func GetSqlPlanBaselineSummaryAcceptedEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingSqlPlanBaselineSummaryAcceptedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlPlanBaselineSummaryAcceptedEnum(val string) (SqlPlanBaselineSummaryAcceptedEnum, bool) {
	enum, ok := mappingSqlPlanBaselineSummaryAcceptedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlPlanBaselineSummaryFixedEnum Enum with underlying type: string
type SqlPlanBaselineSummaryFixedEnum string

// Set of constants representing the allowable values for SqlPlanBaselineSummaryFixedEnum
const (
	SqlPlanBaselineSummaryFixedYes SqlPlanBaselineSummaryFixedEnum = "YES"
	SqlPlanBaselineSummaryFixedNo  SqlPlanBaselineSummaryFixedEnum = "NO"
)

var mappingSqlPlanBaselineSummaryFixedEnum = map[string]SqlPlanBaselineSummaryFixedEnum{
	"YES": SqlPlanBaselineSummaryFixedYes,
	"NO":  SqlPlanBaselineSummaryFixedNo,
}

var mappingSqlPlanBaselineSummaryFixedEnumLowerCase = map[string]SqlPlanBaselineSummaryFixedEnum{
	"yes": SqlPlanBaselineSummaryFixedYes,
	"no":  SqlPlanBaselineSummaryFixedNo,
}

// GetSqlPlanBaselineSummaryFixedEnumValues Enumerates the set of values for SqlPlanBaselineSummaryFixedEnum
func GetSqlPlanBaselineSummaryFixedEnumValues() []SqlPlanBaselineSummaryFixedEnum {
	values := make([]SqlPlanBaselineSummaryFixedEnum, 0)
	for _, v := range mappingSqlPlanBaselineSummaryFixedEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlPlanBaselineSummaryFixedEnumStringValues Enumerates the set of values in String for SqlPlanBaselineSummaryFixedEnum
func GetSqlPlanBaselineSummaryFixedEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingSqlPlanBaselineSummaryFixedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlPlanBaselineSummaryFixedEnum(val string) (SqlPlanBaselineSummaryFixedEnum, bool) {
	enum, ok := mappingSqlPlanBaselineSummaryFixedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlPlanBaselineSummaryReproducedEnum Enum with underlying type: string
type SqlPlanBaselineSummaryReproducedEnum string

// Set of constants representing the allowable values for SqlPlanBaselineSummaryReproducedEnum
const (
	SqlPlanBaselineSummaryReproducedYes SqlPlanBaselineSummaryReproducedEnum = "YES"
	SqlPlanBaselineSummaryReproducedNo  SqlPlanBaselineSummaryReproducedEnum = "NO"
)

var mappingSqlPlanBaselineSummaryReproducedEnum = map[string]SqlPlanBaselineSummaryReproducedEnum{
	"YES": SqlPlanBaselineSummaryReproducedYes,
	"NO":  SqlPlanBaselineSummaryReproducedNo,
}

var mappingSqlPlanBaselineSummaryReproducedEnumLowerCase = map[string]SqlPlanBaselineSummaryReproducedEnum{
	"yes": SqlPlanBaselineSummaryReproducedYes,
	"no":  SqlPlanBaselineSummaryReproducedNo,
}

// GetSqlPlanBaselineSummaryReproducedEnumValues Enumerates the set of values for SqlPlanBaselineSummaryReproducedEnum
func GetSqlPlanBaselineSummaryReproducedEnumValues() []SqlPlanBaselineSummaryReproducedEnum {
	values := make([]SqlPlanBaselineSummaryReproducedEnum, 0)
	for _, v := range mappingSqlPlanBaselineSummaryReproducedEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlPlanBaselineSummaryReproducedEnumStringValues Enumerates the set of values in String for SqlPlanBaselineSummaryReproducedEnum
func GetSqlPlanBaselineSummaryReproducedEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingSqlPlanBaselineSummaryReproducedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlPlanBaselineSummaryReproducedEnum(val string) (SqlPlanBaselineSummaryReproducedEnum, bool) {
	enum, ok := mappingSqlPlanBaselineSummaryReproducedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlPlanBaselineSummaryAutoPurgeEnum Enum with underlying type: string
type SqlPlanBaselineSummaryAutoPurgeEnum string

// Set of constants representing the allowable values for SqlPlanBaselineSummaryAutoPurgeEnum
const (
	SqlPlanBaselineSummaryAutoPurgeYes SqlPlanBaselineSummaryAutoPurgeEnum = "YES"
	SqlPlanBaselineSummaryAutoPurgeNo  SqlPlanBaselineSummaryAutoPurgeEnum = "NO"
)

var mappingSqlPlanBaselineSummaryAutoPurgeEnum = map[string]SqlPlanBaselineSummaryAutoPurgeEnum{
	"YES": SqlPlanBaselineSummaryAutoPurgeYes,
	"NO":  SqlPlanBaselineSummaryAutoPurgeNo,
}

var mappingSqlPlanBaselineSummaryAutoPurgeEnumLowerCase = map[string]SqlPlanBaselineSummaryAutoPurgeEnum{
	"yes": SqlPlanBaselineSummaryAutoPurgeYes,
	"no":  SqlPlanBaselineSummaryAutoPurgeNo,
}

// GetSqlPlanBaselineSummaryAutoPurgeEnumValues Enumerates the set of values for SqlPlanBaselineSummaryAutoPurgeEnum
func GetSqlPlanBaselineSummaryAutoPurgeEnumValues() []SqlPlanBaselineSummaryAutoPurgeEnum {
	values := make([]SqlPlanBaselineSummaryAutoPurgeEnum, 0)
	for _, v := range mappingSqlPlanBaselineSummaryAutoPurgeEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlPlanBaselineSummaryAutoPurgeEnumStringValues Enumerates the set of values in String for SqlPlanBaselineSummaryAutoPurgeEnum
func GetSqlPlanBaselineSummaryAutoPurgeEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingSqlPlanBaselineSummaryAutoPurgeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlPlanBaselineSummaryAutoPurgeEnum(val string) (SqlPlanBaselineSummaryAutoPurgeEnum, bool) {
	enum, ok := mappingSqlPlanBaselineSummaryAutoPurgeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlPlanBaselineSummaryAdaptiveEnum Enum with underlying type: string
type SqlPlanBaselineSummaryAdaptiveEnum string

// Set of constants representing the allowable values for SqlPlanBaselineSummaryAdaptiveEnum
const (
	SqlPlanBaselineSummaryAdaptiveYes SqlPlanBaselineSummaryAdaptiveEnum = "YES"
	SqlPlanBaselineSummaryAdaptiveNo  SqlPlanBaselineSummaryAdaptiveEnum = "NO"
)

var mappingSqlPlanBaselineSummaryAdaptiveEnum = map[string]SqlPlanBaselineSummaryAdaptiveEnum{
	"YES": SqlPlanBaselineSummaryAdaptiveYes,
	"NO":  SqlPlanBaselineSummaryAdaptiveNo,
}

var mappingSqlPlanBaselineSummaryAdaptiveEnumLowerCase = map[string]SqlPlanBaselineSummaryAdaptiveEnum{
	"yes": SqlPlanBaselineSummaryAdaptiveYes,
	"no":  SqlPlanBaselineSummaryAdaptiveNo,
}

// GetSqlPlanBaselineSummaryAdaptiveEnumValues Enumerates the set of values for SqlPlanBaselineSummaryAdaptiveEnum
func GetSqlPlanBaselineSummaryAdaptiveEnumValues() []SqlPlanBaselineSummaryAdaptiveEnum {
	values := make([]SqlPlanBaselineSummaryAdaptiveEnum, 0)
	for _, v := range mappingSqlPlanBaselineSummaryAdaptiveEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlPlanBaselineSummaryAdaptiveEnumStringValues Enumerates the set of values in String for SqlPlanBaselineSummaryAdaptiveEnum
func GetSqlPlanBaselineSummaryAdaptiveEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingSqlPlanBaselineSummaryAdaptiveEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlPlanBaselineSummaryAdaptiveEnum(val string) (SqlPlanBaselineSummaryAdaptiveEnum, bool) {
	enum, ok := mappingSqlPlanBaselineSummaryAdaptiveEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
