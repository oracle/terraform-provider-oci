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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LoadSqlTuningSetDetails The details required to load the Sql statements into the Sql tuning set.
// It takes either credentialDetails or databaseCredential. It's recommended to provide databaseCredential
type LoadSqlTuningSetDetails struct {

	// The name of the Sql tuning set.
	Name *string `mandatory:"true" json:"name"`

	// Specifies the loading method into the Sql tuning set.
	LoadType LoadSqlTuningSetDetailsLoadTypeEnum `mandatory:"true" json:"loadType"`

	CredentialDetails SqlTuningSetAdminCredentialDetails `mandatory:"false" json:"credentialDetails"`

	DatabaseCredential DatabaseCredentialDetails `mandatory:"false" json:"databaseCredential"`

	// Flag to indicate whether to create the Sql tuning set or just display the plsql used to create Sql tuning set.
	ShowSqlOnly *int `mandatory:"false" json:"showSqlOnly"`

	// The owner of the Sql tuning set.
	Owner *string `mandatory:"false" json:"owner"`

	// Specifies the Sql predicate to filter the Sql from the Sql tuning set defined on attributes of the SQLSET_ROW.
	// User could use any combination of the following columns with appropriate values as Sql predicate
	// Refer to the documentation https://docs.oracle.com/en/database/oracle/oracle-database/18/arpls/DBMS_SQLTUNE.html#GUID-1F4AFB03-7B29-46FC-B3F2-CB01EC36326C
	BasicFilter *string `mandatory:"false" json:"basicFilter"`

	// Specifies that the filter must include recursive Sql in the Sql tuning set.
	RecursiveSql LoadSqlTuningSetDetailsRecursiveSqlEnum `mandatory:"false" json:"recursiveSql,omitempty"`

	// Specifies a filter that picks the top n% according to the supplied ranking measure.
	// Note that this parameter applies only if one ranking measure is supplied.
	ResultPercentage *float64 `mandatory:"false" json:"resultPercentage"`

	// The top limit Sql from the filtered source, ranked by the ranking measure.
	ResultLimit *int `mandatory:"false" json:"resultLimit"`

	// Specifies an ORDER BY clause on the selected Sql. User can specify upto three ranking measures.
	RankingMeasure1 RankingMeasureEnum `mandatory:"false" json:"rankingMeasure1,omitempty"`

	// Specifies an ORDER BY clause on the selected Sql. User can specify upto three ranking measures.
	RankingMeasure2 RankingMeasureEnum `mandatory:"false" json:"rankingMeasure2,omitempty"`

	// Specifies an ORDER BY clause on the selected Sql. User can specify upto three ranking measures.
	RankingMeasure3 RankingMeasureEnum `mandatory:"false" json:"rankingMeasure3,omitempty"`

	// Defines the total amount of time, in seconds, to execute.
	TotalTimeLimit *int `mandatory:"false" json:"totalTimeLimit"`

	// Defines the amount of time, in seconds, to pause between sampling.
	RepeatInterval *int `mandatory:"false" json:"repeatInterval"`

	// Specifies whether to insert new statements, update existing statements, or both.
	CaptureOption LoadSqlTuningSetDetailsCaptureOptionEnum `mandatory:"false" json:"captureOption,omitempty"`

	// Specifies the capture mode. Note that this parameter is applicable only for UPDATE and MERGE capture options.
	// Capture mode can take one of the following values
	//  - MODE_REPLACE_OLD_STATS
	//      Replaces statistics when the number of executions is greater than the number stored in the Sql tuning set
	//  - MODE_ACCUMULATE_STATS
	//      Adds new values to current values for Sql that is already stored.
	//      Note that this mode detects if a statement has been aged out, so the final value for a statistics is the sum of the statistics of all cursors that statement existed under.
	CaptureMode LoadSqlTuningSetDetailsCaptureModeEnum `mandatory:"false" json:"captureMode,omitempty"`

	// Specifies the list of Sql statement attributes to return in the result.
	// Note that this parameter cannot be made an enum since custom value can take a list of comma separated attribute names.
	// Attribute list can take one of the following values.
	//  TYPICAL - Specifies BASIC plus Sql plan (without row source statistics) and without object reference list (default).
	//  BASIC - Specifies all attributes (such as execution statistics and binds) except the plans. The execution context is always part of the result.
	//  ALL - Specifies all attributes.
	//  CUSTOM - Comma-separated list of the following attribute names.
	//           - EXECUTION_STATISTICS
	//           - BIND_LIST
	//           - OBJECT_LIST
	//           - SQL_PLAN
	//           - SQL_PLAN_STATISTICS
	// Usage examples:
	//   1. "attributeList": "TYPICAL"
	//   2. "attributeList": "ALL"
	//   3. "attributeList": "EXECUTION_STATISTICS,OBJECT_LIST,SQL_PLAN"
	AttributeList *string `mandatory:"false" json:"attributeList"`

	// Specifies which statements are loaded into the Sql tuning set.
	// The possible values are.
	//  - INSERT (default)
	//       Adds only new statements.
	//  - UPDATE
	//       Updates existing the Sql statements and ignores any new statements.
	//  - MERGE
	//       Inserts new statements and updates the information of the existing ones.
	LoadOption LoadSqlTuningSetDetailsLoadOptionEnum `mandatory:"false" json:"loadOption,omitempty"`

	// Specifies how existing Sql statements are updated.
	// This parameter is applicable only if load_option is specified with UPDATE or MERGE as an option.
	// Update option can take one of the following values.
	//    REPLACE (default) - Updates the statement using the new statistics, bind list, object list, and so on.
	//    ACCUMULATE - Combines attributes when possible (for example, statistics such as elapsed_time), otherwise replaces the existing values (for example, module and action) with the provided values.
	//    Following Sql statement attributes can be accumulated.
	//        elapsed_time
	//        buffer_gets
	//        direct_writes
	//        disk_reads
	//        row_processed
	//        fetches
	//        executions
	//        end_of_fetch_count
	//        stat_period
	//        active_stat_period
	UpdateOption LoadSqlTuningSetDetailsUpdateOptionEnum `mandatory:"false" json:"updateOption,omitempty"`

	// Specifies the list of Sql statement attributes to update during a merge or update.
	// Note that this parameter cannot be made an enum since custom value can take a list of comma separated attribute names.
	// Update attributes can take one of the following values.
	//    NULL (default) - Specifies the content of the input cursor except the execution context. On other terms, it is equivalent to ALL without execution contexts such as module and action.
	//    BASIC - Specifies statistics and binds only.
	//    TYPICAL - Specifies BASIC with Sql plans (without row source statistics) and without an object reference list.
	//    ALL - Specifies all attributes, including the execution context attributes such as module and action.
	//    CUSTOM - List of comma separated attribute names to update
	//        EXECUTION_CONTEXT
	//        EXECUTION_STATISTICS
	//        SQL_BINDS
	//        SQL_PLAN
	//        SQL_PLAN_STATISTICS (similar to SQL_PLAN with added row source statistics)
	// Usage examples:
	//   1. "updateAttributes": "TYPICAL"
	//   2. "updateAttributes": "BASIC"
	//   3. "updateAttributes": "EXECUTION_STATISTICS,SQL_PLAN_STATISTICS,SQL_PLAN"
	//   4. "updateAttributes": "EXECUTION_STATISTICS,SQL_PLAN"
	UpdateAttributes *string `mandatory:"false" json:"updateAttributes"`

	// Specifies when to perform the update.
	// The procedure only performs the update when the specified condition is satisfied.
	// The condition can refer to either the data source or destination.
	// The condition must use the following prefixes to refer to attributes from the source or the destination:
	// OLD  — Refers to statement attributes from the SQL tuning set (destination).
	// NEW  — Refers to statement attributes from the input statements (source).
	// NULL — No updates are performed.
	UpdateCondition LoadSqlTuningSetDetailsUpdateConditionEnum `mandatory:"false" json:"updateCondition,omitempty"`

	// Specifies whether to update attributes when the new value is NULL.
	// If TRUE, then the procedure does not update an attribute when the new value is NULL.
	// That is, do not override with NULL values unless intentional.
	// Possible values - true or false
	IsIgnoreNull *bool `mandatory:"false" json:"isIgnoreNull"`

	// Specifies whether to commit statements after DML.
	// If a value is provided, then the load commits after each specified number of statements is inserted.
	// If NULL is provided, then the load commits only once, at the end of the operation.
	CommitRows *int `mandatory:"false" json:"commitRows"`

	// Defines the beginning AWR snapshot (non-inclusive).
	BeginSnapshot *int64 `mandatory:"false" json:"beginSnapshot"`

	// Defines the ending AWR snapshot (inclusive).
	EndSnapshot *int64 `mandatory:"false" json:"endSnapshot"`

	// Specifies the name of the AWR baseline period.
	// When loading the sql statements from AWR, following inputs has to be provided:
	// beginSnapshot and endSnapshot
	// OR
	// baselineName
	BaselineName *string `mandatory:"false" json:"baselineName"`
}

func (m LoadSqlTuningSetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LoadSqlTuningSetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLoadSqlTuningSetDetailsLoadTypeEnum(string(m.LoadType)); !ok && m.LoadType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LoadType: %s. Supported values are: %s.", m.LoadType, strings.Join(GetLoadSqlTuningSetDetailsLoadTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingLoadSqlTuningSetDetailsRecursiveSqlEnum(string(m.RecursiveSql)); !ok && m.RecursiveSql != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecursiveSql: %s. Supported values are: %s.", m.RecursiveSql, strings.Join(GetLoadSqlTuningSetDetailsRecursiveSqlEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRankingMeasureEnum(string(m.RankingMeasure1)); !ok && m.RankingMeasure1 != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RankingMeasure1: %s. Supported values are: %s.", m.RankingMeasure1, strings.Join(GetRankingMeasureEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRankingMeasureEnum(string(m.RankingMeasure2)); !ok && m.RankingMeasure2 != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RankingMeasure2: %s. Supported values are: %s.", m.RankingMeasure2, strings.Join(GetRankingMeasureEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRankingMeasureEnum(string(m.RankingMeasure3)); !ok && m.RankingMeasure3 != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RankingMeasure3: %s. Supported values are: %s.", m.RankingMeasure3, strings.Join(GetRankingMeasureEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLoadSqlTuningSetDetailsCaptureOptionEnum(string(m.CaptureOption)); !ok && m.CaptureOption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CaptureOption: %s. Supported values are: %s.", m.CaptureOption, strings.Join(GetLoadSqlTuningSetDetailsCaptureOptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLoadSqlTuningSetDetailsCaptureModeEnum(string(m.CaptureMode)); !ok && m.CaptureMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CaptureMode: %s. Supported values are: %s.", m.CaptureMode, strings.Join(GetLoadSqlTuningSetDetailsCaptureModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLoadSqlTuningSetDetailsLoadOptionEnum(string(m.LoadOption)); !ok && m.LoadOption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LoadOption: %s. Supported values are: %s.", m.LoadOption, strings.Join(GetLoadSqlTuningSetDetailsLoadOptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLoadSqlTuningSetDetailsUpdateOptionEnum(string(m.UpdateOption)); !ok && m.UpdateOption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateOption: %s. Supported values are: %s.", m.UpdateOption, strings.Join(GetLoadSqlTuningSetDetailsUpdateOptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLoadSqlTuningSetDetailsUpdateConditionEnum(string(m.UpdateCondition)); !ok && m.UpdateCondition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateCondition: %s. Supported values are: %s.", m.UpdateCondition, strings.Join(GetLoadSqlTuningSetDetailsUpdateConditionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *LoadSqlTuningSetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CredentialDetails  sqltuningsetadmincredentialdetails         `json:"credentialDetails"`
		DatabaseCredential databasecredentialdetails                  `json:"databaseCredential"`
		ShowSqlOnly        *int                                       `json:"showSqlOnly"`
		Owner              *string                                    `json:"owner"`
		BasicFilter        *string                                    `json:"basicFilter"`
		RecursiveSql       LoadSqlTuningSetDetailsRecursiveSqlEnum    `json:"recursiveSql"`
		ResultPercentage   *float64                                   `json:"resultPercentage"`
		ResultLimit        *int                                       `json:"resultLimit"`
		RankingMeasure1    RankingMeasureEnum                         `json:"rankingMeasure1"`
		RankingMeasure2    RankingMeasureEnum                         `json:"rankingMeasure2"`
		RankingMeasure3    RankingMeasureEnum                         `json:"rankingMeasure3"`
		TotalTimeLimit     *int                                       `json:"totalTimeLimit"`
		RepeatInterval     *int                                       `json:"repeatInterval"`
		CaptureOption      LoadSqlTuningSetDetailsCaptureOptionEnum   `json:"captureOption"`
		CaptureMode        LoadSqlTuningSetDetailsCaptureModeEnum     `json:"captureMode"`
		AttributeList      *string                                    `json:"attributeList"`
		LoadOption         LoadSqlTuningSetDetailsLoadOptionEnum      `json:"loadOption"`
		UpdateOption       LoadSqlTuningSetDetailsUpdateOptionEnum    `json:"updateOption"`
		UpdateAttributes   *string                                    `json:"updateAttributes"`
		UpdateCondition    LoadSqlTuningSetDetailsUpdateConditionEnum `json:"updateCondition"`
		IsIgnoreNull       *bool                                      `json:"isIgnoreNull"`
		CommitRows         *int                                       `json:"commitRows"`
		BeginSnapshot      *int64                                     `json:"beginSnapshot"`
		EndSnapshot        *int64                                     `json:"endSnapshot"`
		BaselineName       *string                                    `json:"baselineName"`
		Name               *string                                    `json:"name"`
		LoadType           LoadSqlTuningSetDetailsLoadTypeEnum        `json:"loadType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.CredentialDetails.UnmarshalPolymorphicJSON(model.CredentialDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CredentialDetails = nn.(SqlTuningSetAdminCredentialDetails)
	} else {
		m.CredentialDetails = nil
	}

	nn, e = model.DatabaseCredential.UnmarshalPolymorphicJSON(model.DatabaseCredential.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatabaseCredential = nn.(DatabaseCredentialDetails)
	} else {
		m.DatabaseCredential = nil
	}

	m.ShowSqlOnly = model.ShowSqlOnly

	m.Owner = model.Owner

	m.BasicFilter = model.BasicFilter

	m.RecursiveSql = model.RecursiveSql

	m.ResultPercentage = model.ResultPercentage

	m.ResultLimit = model.ResultLimit

	m.RankingMeasure1 = model.RankingMeasure1

	m.RankingMeasure2 = model.RankingMeasure2

	m.RankingMeasure3 = model.RankingMeasure3

	m.TotalTimeLimit = model.TotalTimeLimit

	m.RepeatInterval = model.RepeatInterval

	m.CaptureOption = model.CaptureOption

	m.CaptureMode = model.CaptureMode

	m.AttributeList = model.AttributeList

	m.LoadOption = model.LoadOption

	m.UpdateOption = model.UpdateOption

	m.UpdateAttributes = model.UpdateAttributes

	m.UpdateCondition = model.UpdateCondition

	m.IsIgnoreNull = model.IsIgnoreNull

	m.CommitRows = model.CommitRows

	m.BeginSnapshot = model.BeginSnapshot

	m.EndSnapshot = model.EndSnapshot

	m.BaselineName = model.BaselineName

	m.Name = model.Name

	m.LoadType = model.LoadType

	return
}

// LoadSqlTuningSetDetailsLoadTypeEnum Enum with underlying type: string
type LoadSqlTuningSetDetailsLoadTypeEnum string

// Set of constants representing the allowable values for LoadSqlTuningSetDetailsLoadTypeEnum
const (
	LoadSqlTuningSetDetailsLoadTypeIncrementalCursorCache LoadSqlTuningSetDetailsLoadTypeEnum = "INCREMENTAL_CURSOR_CACHE"
	LoadSqlTuningSetDetailsLoadTypeCurrentCursorCache     LoadSqlTuningSetDetailsLoadTypeEnum = "CURRENT_CURSOR_CACHE"
	LoadSqlTuningSetDetailsLoadTypeAwr                    LoadSqlTuningSetDetailsLoadTypeEnum = "AWR"
)

var mappingLoadSqlTuningSetDetailsLoadTypeEnum = map[string]LoadSqlTuningSetDetailsLoadTypeEnum{
	"INCREMENTAL_CURSOR_CACHE": LoadSqlTuningSetDetailsLoadTypeIncrementalCursorCache,
	"CURRENT_CURSOR_CACHE":     LoadSqlTuningSetDetailsLoadTypeCurrentCursorCache,
	"AWR":                      LoadSqlTuningSetDetailsLoadTypeAwr,
}

var mappingLoadSqlTuningSetDetailsLoadTypeEnumLowerCase = map[string]LoadSqlTuningSetDetailsLoadTypeEnum{
	"incremental_cursor_cache": LoadSqlTuningSetDetailsLoadTypeIncrementalCursorCache,
	"current_cursor_cache":     LoadSqlTuningSetDetailsLoadTypeCurrentCursorCache,
	"awr":                      LoadSqlTuningSetDetailsLoadTypeAwr,
}

// GetLoadSqlTuningSetDetailsLoadTypeEnumValues Enumerates the set of values for LoadSqlTuningSetDetailsLoadTypeEnum
func GetLoadSqlTuningSetDetailsLoadTypeEnumValues() []LoadSqlTuningSetDetailsLoadTypeEnum {
	values := make([]LoadSqlTuningSetDetailsLoadTypeEnum, 0)
	for _, v := range mappingLoadSqlTuningSetDetailsLoadTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLoadSqlTuningSetDetailsLoadTypeEnumStringValues Enumerates the set of values in String for LoadSqlTuningSetDetailsLoadTypeEnum
func GetLoadSqlTuningSetDetailsLoadTypeEnumStringValues() []string {
	return []string{
		"INCREMENTAL_CURSOR_CACHE",
		"CURRENT_CURSOR_CACHE",
		"AWR",
	}
}

// GetMappingLoadSqlTuningSetDetailsLoadTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLoadSqlTuningSetDetailsLoadTypeEnum(val string) (LoadSqlTuningSetDetailsLoadTypeEnum, bool) {
	enum, ok := mappingLoadSqlTuningSetDetailsLoadTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LoadSqlTuningSetDetailsRecursiveSqlEnum Enum with underlying type: string
type LoadSqlTuningSetDetailsRecursiveSqlEnum string

// Set of constants representing the allowable values for LoadSqlTuningSetDetailsRecursiveSqlEnum
const (
	LoadSqlTuningSetDetailsRecursiveSqlHasRecursiveSql LoadSqlTuningSetDetailsRecursiveSqlEnum = "HAS_RECURSIVE_SQL"
	LoadSqlTuningSetDetailsRecursiveSqlNoRecursiveSql  LoadSqlTuningSetDetailsRecursiveSqlEnum = "NO_RECURSIVE_SQL"
)

var mappingLoadSqlTuningSetDetailsRecursiveSqlEnum = map[string]LoadSqlTuningSetDetailsRecursiveSqlEnum{
	"HAS_RECURSIVE_SQL": LoadSqlTuningSetDetailsRecursiveSqlHasRecursiveSql,
	"NO_RECURSIVE_SQL":  LoadSqlTuningSetDetailsRecursiveSqlNoRecursiveSql,
}

var mappingLoadSqlTuningSetDetailsRecursiveSqlEnumLowerCase = map[string]LoadSqlTuningSetDetailsRecursiveSqlEnum{
	"has_recursive_sql": LoadSqlTuningSetDetailsRecursiveSqlHasRecursiveSql,
	"no_recursive_sql":  LoadSqlTuningSetDetailsRecursiveSqlNoRecursiveSql,
}

// GetLoadSqlTuningSetDetailsRecursiveSqlEnumValues Enumerates the set of values for LoadSqlTuningSetDetailsRecursiveSqlEnum
func GetLoadSqlTuningSetDetailsRecursiveSqlEnumValues() []LoadSqlTuningSetDetailsRecursiveSqlEnum {
	values := make([]LoadSqlTuningSetDetailsRecursiveSqlEnum, 0)
	for _, v := range mappingLoadSqlTuningSetDetailsRecursiveSqlEnum {
		values = append(values, v)
	}
	return values
}

// GetLoadSqlTuningSetDetailsRecursiveSqlEnumStringValues Enumerates the set of values in String for LoadSqlTuningSetDetailsRecursiveSqlEnum
func GetLoadSqlTuningSetDetailsRecursiveSqlEnumStringValues() []string {
	return []string{
		"HAS_RECURSIVE_SQL",
		"NO_RECURSIVE_SQL",
	}
}

// GetMappingLoadSqlTuningSetDetailsRecursiveSqlEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLoadSqlTuningSetDetailsRecursiveSqlEnum(val string) (LoadSqlTuningSetDetailsRecursiveSqlEnum, bool) {
	enum, ok := mappingLoadSqlTuningSetDetailsRecursiveSqlEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LoadSqlTuningSetDetailsCaptureOptionEnum Enum with underlying type: string
type LoadSqlTuningSetDetailsCaptureOptionEnum string

// Set of constants representing the allowable values for LoadSqlTuningSetDetailsCaptureOptionEnum
const (
	LoadSqlTuningSetDetailsCaptureOptionInsert LoadSqlTuningSetDetailsCaptureOptionEnum = "INSERT"
	LoadSqlTuningSetDetailsCaptureOptionUpdate LoadSqlTuningSetDetailsCaptureOptionEnum = "UPDATE"
	LoadSqlTuningSetDetailsCaptureOptionMerge  LoadSqlTuningSetDetailsCaptureOptionEnum = "MERGE"
)

var mappingLoadSqlTuningSetDetailsCaptureOptionEnum = map[string]LoadSqlTuningSetDetailsCaptureOptionEnum{
	"INSERT": LoadSqlTuningSetDetailsCaptureOptionInsert,
	"UPDATE": LoadSqlTuningSetDetailsCaptureOptionUpdate,
	"MERGE":  LoadSqlTuningSetDetailsCaptureOptionMerge,
}

var mappingLoadSqlTuningSetDetailsCaptureOptionEnumLowerCase = map[string]LoadSqlTuningSetDetailsCaptureOptionEnum{
	"insert": LoadSqlTuningSetDetailsCaptureOptionInsert,
	"update": LoadSqlTuningSetDetailsCaptureOptionUpdate,
	"merge":  LoadSqlTuningSetDetailsCaptureOptionMerge,
}

// GetLoadSqlTuningSetDetailsCaptureOptionEnumValues Enumerates the set of values for LoadSqlTuningSetDetailsCaptureOptionEnum
func GetLoadSqlTuningSetDetailsCaptureOptionEnumValues() []LoadSqlTuningSetDetailsCaptureOptionEnum {
	values := make([]LoadSqlTuningSetDetailsCaptureOptionEnum, 0)
	for _, v := range mappingLoadSqlTuningSetDetailsCaptureOptionEnum {
		values = append(values, v)
	}
	return values
}

// GetLoadSqlTuningSetDetailsCaptureOptionEnumStringValues Enumerates the set of values in String for LoadSqlTuningSetDetailsCaptureOptionEnum
func GetLoadSqlTuningSetDetailsCaptureOptionEnumStringValues() []string {
	return []string{
		"INSERT",
		"UPDATE",
		"MERGE",
	}
}

// GetMappingLoadSqlTuningSetDetailsCaptureOptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLoadSqlTuningSetDetailsCaptureOptionEnum(val string) (LoadSqlTuningSetDetailsCaptureOptionEnum, bool) {
	enum, ok := mappingLoadSqlTuningSetDetailsCaptureOptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LoadSqlTuningSetDetailsCaptureModeEnum Enum with underlying type: string
type LoadSqlTuningSetDetailsCaptureModeEnum string

// Set of constants representing the allowable values for LoadSqlTuningSetDetailsCaptureModeEnum
const (
	LoadSqlTuningSetDetailsCaptureModeReplaceOldStats LoadSqlTuningSetDetailsCaptureModeEnum = "MODE_REPLACE_OLD_STATS"
	LoadSqlTuningSetDetailsCaptureModeAccumulateStats LoadSqlTuningSetDetailsCaptureModeEnum = "MODE_ACCUMULATE_STATS"
)

var mappingLoadSqlTuningSetDetailsCaptureModeEnum = map[string]LoadSqlTuningSetDetailsCaptureModeEnum{
	"MODE_REPLACE_OLD_STATS": LoadSqlTuningSetDetailsCaptureModeReplaceOldStats,
	"MODE_ACCUMULATE_STATS":  LoadSqlTuningSetDetailsCaptureModeAccumulateStats,
}

var mappingLoadSqlTuningSetDetailsCaptureModeEnumLowerCase = map[string]LoadSqlTuningSetDetailsCaptureModeEnum{
	"mode_replace_old_stats": LoadSqlTuningSetDetailsCaptureModeReplaceOldStats,
	"mode_accumulate_stats":  LoadSqlTuningSetDetailsCaptureModeAccumulateStats,
}

// GetLoadSqlTuningSetDetailsCaptureModeEnumValues Enumerates the set of values for LoadSqlTuningSetDetailsCaptureModeEnum
func GetLoadSqlTuningSetDetailsCaptureModeEnumValues() []LoadSqlTuningSetDetailsCaptureModeEnum {
	values := make([]LoadSqlTuningSetDetailsCaptureModeEnum, 0)
	for _, v := range mappingLoadSqlTuningSetDetailsCaptureModeEnum {
		values = append(values, v)
	}
	return values
}

// GetLoadSqlTuningSetDetailsCaptureModeEnumStringValues Enumerates the set of values in String for LoadSqlTuningSetDetailsCaptureModeEnum
func GetLoadSqlTuningSetDetailsCaptureModeEnumStringValues() []string {
	return []string{
		"MODE_REPLACE_OLD_STATS",
		"MODE_ACCUMULATE_STATS",
	}
}

// GetMappingLoadSqlTuningSetDetailsCaptureModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLoadSqlTuningSetDetailsCaptureModeEnum(val string) (LoadSqlTuningSetDetailsCaptureModeEnum, bool) {
	enum, ok := mappingLoadSqlTuningSetDetailsCaptureModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LoadSqlTuningSetDetailsLoadOptionEnum Enum with underlying type: string
type LoadSqlTuningSetDetailsLoadOptionEnum string

// Set of constants representing the allowable values for LoadSqlTuningSetDetailsLoadOptionEnum
const (
	LoadSqlTuningSetDetailsLoadOptionInsert LoadSqlTuningSetDetailsLoadOptionEnum = "INSERT"
	LoadSqlTuningSetDetailsLoadOptionUpdate LoadSqlTuningSetDetailsLoadOptionEnum = "UPDATE"
	LoadSqlTuningSetDetailsLoadOptionMerge  LoadSqlTuningSetDetailsLoadOptionEnum = "MERGE"
)

var mappingLoadSqlTuningSetDetailsLoadOptionEnum = map[string]LoadSqlTuningSetDetailsLoadOptionEnum{
	"INSERT": LoadSqlTuningSetDetailsLoadOptionInsert,
	"UPDATE": LoadSqlTuningSetDetailsLoadOptionUpdate,
	"MERGE":  LoadSqlTuningSetDetailsLoadOptionMerge,
}

var mappingLoadSqlTuningSetDetailsLoadOptionEnumLowerCase = map[string]LoadSqlTuningSetDetailsLoadOptionEnum{
	"insert": LoadSqlTuningSetDetailsLoadOptionInsert,
	"update": LoadSqlTuningSetDetailsLoadOptionUpdate,
	"merge":  LoadSqlTuningSetDetailsLoadOptionMerge,
}

// GetLoadSqlTuningSetDetailsLoadOptionEnumValues Enumerates the set of values for LoadSqlTuningSetDetailsLoadOptionEnum
func GetLoadSqlTuningSetDetailsLoadOptionEnumValues() []LoadSqlTuningSetDetailsLoadOptionEnum {
	values := make([]LoadSqlTuningSetDetailsLoadOptionEnum, 0)
	for _, v := range mappingLoadSqlTuningSetDetailsLoadOptionEnum {
		values = append(values, v)
	}
	return values
}

// GetLoadSqlTuningSetDetailsLoadOptionEnumStringValues Enumerates the set of values in String for LoadSqlTuningSetDetailsLoadOptionEnum
func GetLoadSqlTuningSetDetailsLoadOptionEnumStringValues() []string {
	return []string{
		"INSERT",
		"UPDATE",
		"MERGE",
	}
}

// GetMappingLoadSqlTuningSetDetailsLoadOptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLoadSqlTuningSetDetailsLoadOptionEnum(val string) (LoadSqlTuningSetDetailsLoadOptionEnum, bool) {
	enum, ok := mappingLoadSqlTuningSetDetailsLoadOptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LoadSqlTuningSetDetailsUpdateOptionEnum Enum with underlying type: string
type LoadSqlTuningSetDetailsUpdateOptionEnum string

// Set of constants representing the allowable values for LoadSqlTuningSetDetailsUpdateOptionEnum
const (
	LoadSqlTuningSetDetailsUpdateOptionReplace    LoadSqlTuningSetDetailsUpdateOptionEnum = "REPLACE"
	LoadSqlTuningSetDetailsUpdateOptionAccumulate LoadSqlTuningSetDetailsUpdateOptionEnum = "ACCUMULATE"
)

var mappingLoadSqlTuningSetDetailsUpdateOptionEnum = map[string]LoadSqlTuningSetDetailsUpdateOptionEnum{
	"REPLACE":    LoadSqlTuningSetDetailsUpdateOptionReplace,
	"ACCUMULATE": LoadSqlTuningSetDetailsUpdateOptionAccumulate,
}

var mappingLoadSqlTuningSetDetailsUpdateOptionEnumLowerCase = map[string]LoadSqlTuningSetDetailsUpdateOptionEnum{
	"replace":    LoadSqlTuningSetDetailsUpdateOptionReplace,
	"accumulate": LoadSqlTuningSetDetailsUpdateOptionAccumulate,
}

// GetLoadSqlTuningSetDetailsUpdateOptionEnumValues Enumerates the set of values for LoadSqlTuningSetDetailsUpdateOptionEnum
func GetLoadSqlTuningSetDetailsUpdateOptionEnumValues() []LoadSqlTuningSetDetailsUpdateOptionEnum {
	values := make([]LoadSqlTuningSetDetailsUpdateOptionEnum, 0)
	for _, v := range mappingLoadSqlTuningSetDetailsUpdateOptionEnum {
		values = append(values, v)
	}
	return values
}

// GetLoadSqlTuningSetDetailsUpdateOptionEnumStringValues Enumerates the set of values in String for LoadSqlTuningSetDetailsUpdateOptionEnum
func GetLoadSqlTuningSetDetailsUpdateOptionEnumStringValues() []string {
	return []string{
		"REPLACE",
		"ACCUMULATE",
	}
}

// GetMappingLoadSqlTuningSetDetailsUpdateOptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLoadSqlTuningSetDetailsUpdateOptionEnum(val string) (LoadSqlTuningSetDetailsUpdateOptionEnum, bool) {
	enum, ok := mappingLoadSqlTuningSetDetailsUpdateOptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// LoadSqlTuningSetDetailsUpdateConditionEnum Enum with underlying type: string
type LoadSqlTuningSetDetailsUpdateConditionEnum string

// Set of constants representing the allowable values for LoadSqlTuningSetDetailsUpdateConditionEnum
const (
	LoadSqlTuningSetDetailsUpdateConditionOld LoadSqlTuningSetDetailsUpdateConditionEnum = "OLD"
	LoadSqlTuningSetDetailsUpdateConditionNew LoadSqlTuningSetDetailsUpdateConditionEnum = "NEW"
)

var mappingLoadSqlTuningSetDetailsUpdateConditionEnum = map[string]LoadSqlTuningSetDetailsUpdateConditionEnum{
	"OLD": LoadSqlTuningSetDetailsUpdateConditionOld,
	"NEW": LoadSqlTuningSetDetailsUpdateConditionNew,
}

var mappingLoadSqlTuningSetDetailsUpdateConditionEnumLowerCase = map[string]LoadSqlTuningSetDetailsUpdateConditionEnum{
	"old": LoadSqlTuningSetDetailsUpdateConditionOld,
	"new": LoadSqlTuningSetDetailsUpdateConditionNew,
}

// GetLoadSqlTuningSetDetailsUpdateConditionEnumValues Enumerates the set of values for LoadSqlTuningSetDetailsUpdateConditionEnum
func GetLoadSqlTuningSetDetailsUpdateConditionEnumValues() []LoadSqlTuningSetDetailsUpdateConditionEnum {
	values := make([]LoadSqlTuningSetDetailsUpdateConditionEnum, 0)
	for _, v := range mappingLoadSqlTuningSetDetailsUpdateConditionEnum {
		values = append(values, v)
	}
	return values
}

// GetLoadSqlTuningSetDetailsUpdateConditionEnumStringValues Enumerates the set of values in String for LoadSqlTuningSetDetailsUpdateConditionEnum
func GetLoadSqlTuningSetDetailsUpdateConditionEnumStringValues() []string {
	return []string{
		"OLD",
		"NEW",
	}
}

// GetMappingLoadSqlTuningSetDetailsUpdateConditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLoadSqlTuningSetDetailsUpdateConditionEnum(val string) (LoadSqlTuningSetDetailsUpdateConditionEnum, bool) {
	enum, ok := mappingLoadSqlTuningSetDetailsUpdateConditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
