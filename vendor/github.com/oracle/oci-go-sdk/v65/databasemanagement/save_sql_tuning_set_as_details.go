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

// SaveSqlTuningSetAsDetails Save current list of Sql statements into another Sql tuning set.
type SaveSqlTuningSetAsDetails struct {
	CredentialDetails SqlTuningSetAdminCredentialDetails `mandatory:"true" json:"credentialDetails"`

	// The name of the Sql tuning set.
	Name *string `mandatory:"true" json:"name"`

	// The name of the destination Sql tuning set.
	DestinationSqlTuningSetName *string `mandatory:"true" json:"destinationSqlTuningSetName"`

	// Specifies whether to create a new Sql tuning set or not.
	// Possible values
	// 1 - Create a new Sql tuning set
	// 0 - Do not create a new Sql tuning set
	CreateNew *int `mandatory:"true" json:"createNew"`

	// Flag to indicate whether to save the Sql tuning set or just display the plsql used to save Sql tuning set.
	ShowSqlOnly *int `mandatory:"false" json:"showSqlOnly"`

	// The owner of the Sql tuning set.
	Owner *string `mandatory:"false" json:"owner"`

	// The description for the destination Sql tuning set.
	DestinationSqlTuningSetDescription *string `mandatory:"false" json:"destinationSqlTuningSetDescription"`

	// Owner of the destination Sql tuning set.
	DestinationSqlTuningSetOwner *string `mandatory:"false" json:"destinationSqlTuningSetOwner"`

	// Specifies the Sql predicate to filter the Sql from the Sql tuning set defined on attributes of the SQLSET_ROW.
	// User could use any combination of the following columns with appropriate values as Sql predicate
	// Refer to the documentation https://docs.oracle.com/en/database/oracle/oracle-database/18/arpls/DBMS_SQLTUNE.html#GUID-1F4AFB03-7B29-46FC-B3F2-CB01EC36326C
	BasicFilter *string `mandatory:"false" json:"basicFilter"`

	// Specifies the plan filter.
	// This parameter enables you to select a single plan when a statement has multiple plans.
	// Refer to the documentation https://docs.oracle.com/en/database/oracle/oracle-database/19/arpls/DBMS_SQLSET.html#GUID-9D995019-91AB-4B1E-9EAF-031050789B21
	PlanFilter SaveSqlTuningSetAsDetailsPlanFilterEnum `mandatory:"false" json:"planFilter,omitempty"`

	// Specifies that the filter must include recursive Sql in the Sql tuning set.
	RecursiveSql SaveSqlTuningSetAsDetailsRecursiveSqlEnum `mandatory:"false" json:"recursiveSql,omitempty"`

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
	//   - INSERT (default)
	//     Adds only new statements.
	//   - UPDATE
	//     Updates existing the Sql statements and ignores any new statements.
	//   - MERGE
	//     Inserts new statements and updates the information of the existing ones.
	LoadOption SaveSqlTuningSetAsDetailsLoadOptionEnum `mandatory:"false" json:"loadOption,omitempty"`

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
	UpdateOption SaveSqlTuningSetAsDetailsUpdateOptionEnum `mandatory:"false" json:"updateOption,omitempty"`

	// Specifies when to perform the update.
	// The procedure only performs the update when the specified condition is satisfied.
	// The condition can refer to either the data source or destination.
	// The condition must use the following prefixes to refer to attributes from the source or the destination:
	//   OLD  — Refers to statement attributes from the SQL tuning set (destination).
	//   NEW  — Refers to statement attributes from the input statements (source).
	//   NULL — No updates are performed.
	UpdateCondition SaveSqlTuningSetAsDetailsUpdateConditionEnum `mandatory:"false" json:"updateCondition,omitempty"`

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

	// Specifies whether to update attributes when the new value is NULL.
	// If TRUE, then the procedure does not update an attribute when the new value is NULL.
	// That is, do not override with NULL values unless intentional.
	// Possible values - true or false
	IsIgnoreNull *bool `mandatory:"false" json:"isIgnoreNull"`

	// Specifies whether to commit statements after DML.
	// If a value is provided, then the load commits after each specified number of statements is inserted.
	// If NULL is provided, then the load commits only once, at the end of the operation.
	CommitRows *int `mandatory:"false" json:"commitRows"`
}

func (m SaveSqlTuningSetAsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SaveSqlTuningSetAsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSaveSqlTuningSetAsDetailsPlanFilterEnum(string(m.PlanFilter)); !ok && m.PlanFilter != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlanFilter: %s. Supported values are: %s.", m.PlanFilter, strings.Join(GetSaveSqlTuningSetAsDetailsPlanFilterEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSaveSqlTuningSetAsDetailsRecursiveSqlEnum(string(m.RecursiveSql)); !ok && m.RecursiveSql != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecursiveSql: %s. Supported values are: %s.", m.RecursiveSql, strings.Join(GetSaveSqlTuningSetAsDetailsRecursiveSqlEnumStringValues(), ",")))
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
	if _, ok := GetMappingSaveSqlTuningSetAsDetailsLoadOptionEnum(string(m.LoadOption)); !ok && m.LoadOption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LoadOption: %s. Supported values are: %s.", m.LoadOption, strings.Join(GetSaveSqlTuningSetAsDetailsLoadOptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSaveSqlTuningSetAsDetailsUpdateOptionEnum(string(m.UpdateOption)); !ok && m.UpdateOption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateOption: %s. Supported values are: %s.", m.UpdateOption, strings.Join(GetSaveSqlTuningSetAsDetailsUpdateOptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSaveSqlTuningSetAsDetailsUpdateConditionEnum(string(m.UpdateCondition)); !ok && m.UpdateCondition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateCondition: %s. Supported values are: %s.", m.UpdateCondition, strings.Join(GetSaveSqlTuningSetAsDetailsUpdateConditionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SaveSqlTuningSetAsDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ShowSqlOnly                        *int                                         `json:"showSqlOnly"`
		Owner                              *string                                      `json:"owner"`
		DestinationSqlTuningSetDescription *string                                      `json:"destinationSqlTuningSetDescription"`
		DestinationSqlTuningSetOwner       *string                                      `json:"destinationSqlTuningSetOwner"`
		BasicFilter                        *string                                      `json:"basicFilter"`
		PlanFilter                         SaveSqlTuningSetAsDetailsPlanFilterEnum      `json:"planFilter"`
		RecursiveSql                       SaveSqlTuningSetAsDetailsRecursiveSqlEnum    `json:"recursiveSql"`
		ResultPercentage                   *float64                                     `json:"resultPercentage"`
		ResultLimit                        *int                                         `json:"resultLimit"`
		RankingMeasure1                    RankingMeasureEnum                           `json:"rankingMeasure1"`
		RankingMeasure2                    RankingMeasureEnum                           `json:"rankingMeasure2"`
		RankingMeasure3                    RankingMeasureEnum                           `json:"rankingMeasure3"`
		AttributeList                      *string                                      `json:"attributeList"`
		LoadOption                         SaveSqlTuningSetAsDetailsLoadOptionEnum      `json:"loadOption"`
		UpdateOption                       SaveSqlTuningSetAsDetailsUpdateOptionEnum    `json:"updateOption"`
		UpdateCondition                    SaveSqlTuningSetAsDetailsUpdateConditionEnum `json:"updateCondition"`
		UpdateAttributes                   *string                                      `json:"updateAttributes"`
		IsIgnoreNull                       *bool                                        `json:"isIgnoreNull"`
		CommitRows                         *int                                         `json:"commitRows"`
		CredentialDetails                  sqltuningsetadmincredentialdetails           `json:"credentialDetails"`
		Name                               *string                                      `json:"name"`
		DestinationSqlTuningSetName        *string                                      `json:"destinationSqlTuningSetName"`
		CreateNew                          *int                                         `json:"createNew"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ShowSqlOnly = model.ShowSqlOnly

	m.Owner = model.Owner

	m.DestinationSqlTuningSetDescription = model.DestinationSqlTuningSetDescription

	m.DestinationSqlTuningSetOwner = model.DestinationSqlTuningSetOwner

	m.BasicFilter = model.BasicFilter

	m.PlanFilter = model.PlanFilter

	m.RecursiveSql = model.RecursiveSql

	m.ResultPercentage = model.ResultPercentage

	m.ResultLimit = model.ResultLimit

	m.RankingMeasure1 = model.RankingMeasure1

	m.RankingMeasure2 = model.RankingMeasure2

	m.RankingMeasure3 = model.RankingMeasure3

	m.AttributeList = model.AttributeList

	m.LoadOption = model.LoadOption

	m.UpdateOption = model.UpdateOption

	m.UpdateCondition = model.UpdateCondition

	m.UpdateAttributes = model.UpdateAttributes

	m.IsIgnoreNull = model.IsIgnoreNull

	m.CommitRows = model.CommitRows

	nn, e = model.CredentialDetails.UnmarshalPolymorphicJSON(model.CredentialDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.CredentialDetails = nn.(SqlTuningSetAdminCredentialDetails)
	} else {
		m.CredentialDetails = nil
	}

	m.Name = model.Name

	m.DestinationSqlTuningSetName = model.DestinationSqlTuningSetName

	m.CreateNew = model.CreateNew

	return
}

// SaveSqlTuningSetAsDetailsPlanFilterEnum Enum with underlying type: string
type SaveSqlTuningSetAsDetailsPlanFilterEnum string

// Set of constants representing the allowable values for SaveSqlTuningSetAsDetailsPlanFilterEnum
const (
	SaveSqlTuningSetAsDetailsPlanFilterLastGenerated    SaveSqlTuningSetAsDetailsPlanFilterEnum = "LAST_GENERATED"
	SaveSqlTuningSetAsDetailsPlanFilterFirstGenerated   SaveSqlTuningSetAsDetailsPlanFilterEnum = "FIRST_GENERATED"
	SaveSqlTuningSetAsDetailsPlanFilterLastLoaded       SaveSqlTuningSetAsDetailsPlanFilterEnum = "LAST_LOADED"
	SaveSqlTuningSetAsDetailsPlanFilterFirstLoaded      SaveSqlTuningSetAsDetailsPlanFilterEnum = "FIRST_LOADED"
	SaveSqlTuningSetAsDetailsPlanFilterMaxElapsedTime   SaveSqlTuningSetAsDetailsPlanFilterEnum = "MAX_ELAPSED_TIME"
	SaveSqlTuningSetAsDetailsPlanFilterMaxBufferGets    SaveSqlTuningSetAsDetailsPlanFilterEnum = "MAX_BUFFER_GETS"
	SaveSqlTuningSetAsDetailsPlanFilterMaxDiskReads     SaveSqlTuningSetAsDetailsPlanFilterEnum = "MAX_DISK_READS"
	SaveSqlTuningSetAsDetailsPlanFilterMaxDirectWrites  SaveSqlTuningSetAsDetailsPlanFilterEnum = "MAX_DIRECT_WRITES"
	SaveSqlTuningSetAsDetailsPlanFilterMaxOptimizerCost SaveSqlTuningSetAsDetailsPlanFilterEnum = "MAX_OPTIMIZER_COST"
)

var mappingSaveSqlTuningSetAsDetailsPlanFilterEnum = map[string]SaveSqlTuningSetAsDetailsPlanFilterEnum{
	"LAST_GENERATED":     SaveSqlTuningSetAsDetailsPlanFilterLastGenerated,
	"FIRST_GENERATED":    SaveSqlTuningSetAsDetailsPlanFilterFirstGenerated,
	"LAST_LOADED":        SaveSqlTuningSetAsDetailsPlanFilterLastLoaded,
	"FIRST_LOADED":       SaveSqlTuningSetAsDetailsPlanFilterFirstLoaded,
	"MAX_ELAPSED_TIME":   SaveSqlTuningSetAsDetailsPlanFilterMaxElapsedTime,
	"MAX_BUFFER_GETS":    SaveSqlTuningSetAsDetailsPlanFilterMaxBufferGets,
	"MAX_DISK_READS":     SaveSqlTuningSetAsDetailsPlanFilterMaxDiskReads,
	"MAX_DIRECT_WRITES":  SaveSqlTuningSetAsDetailsPlanFilterMaxDirectWrites,
	"MAX_OPTIMIZER_COST": SaveSqlTuningSetAsDetailsPlanFilterMaxOptimizerCost,
}

var mappingSaveSqlTuningSetAsDetailsPlanFilterEnumLowerCase = map[string]SaveSqlTuningSetAsDetailsPlanFilterEnum{
	"last_generated":     SaveSqlTuningSetAsDetailsPlanFilterLastGenerated,
	"first_generated":    SaveSqlTuningSetAsDetailsPlanFilterFirstGenerated,
	"last_loaded":        SaveSqlTuningSetAsDetailsPlanFilterLastLoaded,
	"first_loaded":       SaveSqlTuningSetAsDetailsPlanFilterFirstLoaded,
	"max_elapsed_time":   SaveSqlTuningSetAsDetailsPlanFilterMaxElapsedTime,
	"max_buffer_gets":    SaveSqlTuningSetAsDetailsPlanFilterMaxBufferGets,
	"max_disk_reads":     SaveSqlTuningSetAsDetailsPlanFilterMaxDiskReads,
	"max_direct_writes":  SaveSqlTuningSetAsDetailsPlanFilterMaxDirectWrites,
	"max_optimizer_cost": SaveSqlTuningSetAsDetailsPlanFilterMaxOptimizerCost,
}

// GetSaveSqlTuningSetAsDetailsPlanFilterEnumValues Enumerates the set of values for SaveSqlTuningSetAsDetailsPlanFilterEnum
func GetSaveSqlTuningSetAsDetailsPlanFilterEnumValues() []SaveSqlTuningSetAsDetailsPlanFilterEnum {
	values := make([]SaveSqlTuningSetAsDetailsPlanFilterEnum, 0)
	for _, v := range mappingSaveSqlTuningSetAsDetailsPlanFilterEnum {
		values = append(values, v)
	}
	return values
}

// GetSaveSqlTuningSetAsDetailsPlanFilterEnumStringValues Enumerates the set of values in String for SaveSqlTuningSetAsDetailsPlanFilterEnum
func GetSaveSqlTuningSetAsDetailsPlanFilterEnumStringValues() []string {
	return []string{
		"LAST_GENERATED",
		"FIRST_GENERATED",
		"LAST_LOADED",
		"FIRST_LOADED",
		"MAX_ELAPSED_TIME",
		"MAX_BUFFER_GETS",
		"MAX_DISK_READS",
		"MAX_DIRECT_WRITES",
		"MAX_OPTIMIZER_COST",
	}
}

// GetMappingSaveSqlTuningSetAsDetailsPlanFilterEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSaveSqlTuningSetAsDetailsPlanFilterEnum(val string) (SaveSqlTuningSetAsDetailsPlanFilterEnum, bool) {
	enum, ok := mappingSaveSqlTuningSetAsDetailsPlanFilterEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SaveSqlTuningSetAsDetailsRecursiveSqlEnum Enum with underlying type: string
type SaveSqlTuningSetAsDetailsRecursiveSqlEnum string

// Set of constants representing the allowable values for SaveSqlTuningSetAsDetailsRecursiveSqlEnum
const (
	SaveSqlTuningSetAsDetailsRecursiveSqlHasRecursiveSql SaveSqlTuningSetAsDetailsRecursiveSqlEnum = "HAS_RECURSIVE_SQL"
	SaveSqlTuningSetAsDetailsRecursiveSqlNoRecursiveSql  SaveSqlTuningSetAsDetailsRecursiveSqlEnum = "NO_RECURSIVE_SQL"
)

var mappingSaveSqlTuningSetAsDetailsRecursiveSqlEnum = map[string]SaveSqlTuningSetAsDetailsRecursiveSqlEnum{
	"HAS_RECURSIVE_SQL": SaveSqlTuningSetAsDetailsRecursiveSqlHasRecursiveSql,
	"NO_RECURSIVE_SQL":  SaveSqlTuningSetAsDetailsRecursiveSqlNoRecursiveSql,
}

var mappingSaveSqlTuningSetAsDetailsRecursiveSqlEnumLowerCase = map[string]SaveSqlTuningSetAsDetailsRecursiveSqlEnum{
	"has_recursive_sql": SaveSqlTuningSetAsDetailsRecursiveSqlHasRecursiveSql,
	"no_recursive_sql":  SaveSqlTuningSetAsDetailsRecursiveSqlNoRecursiveSql,
}

// GetSaveSqlTuningSetAsDetailsRecursiveSqlEnumValues Enumerates the set of values for SaveSqlTuningSetAsDetailsRecursiveSqlEnum
func GetSaveSqlTuningSetAsDetailsRecursiveSqlEnumValues() []SaveSqlTuningSetAsDetailsRecursiveSqlEnum {
	values := make([]SaveSqlTuningSetAsDetailsRecursiveSqlEnum, 0)
	for _, v := range mappingSaveSqlTuningSetAsDetailsRecursiveSqlEnum {
		values = append(values, v)
	}
	return values
}

// GetSaveSqlTuningSetAsDetailsRecursiveSqlEnumStringValues Enumerates the set of values in String for SaveSqlTuningSetAsDetailsRecursiveSqlEnum
func GetSaveSqlTuningSetAsDetailsRecursiveSqlEnumStringValues() []string {
	return []string{
		"HAS_RECURSIVE_SQL",
		"NO_RECURSIVE_SQL",
	}
}

// GetMappingSaveSqlTuningSetAsDetailsRecursiveSqlEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSaveSqlTuningSetAsDetailsRecursiveSqlEnum(val string) (SaveSqlTuningSetAsDetailsRecursiveSqlEnum, bool) {
	enum, ok := mappingSaveSqlTuningSetAsDetailsRecursiveSqlEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SaveSqlTuningSetAsDetailsLoadOptionEnum Enum with underlying type: string
type SaveSqlTuningSetAsDetailsLoadOptionEnum string

// Set of constants representing the allowable values for SaveSqlTuningSetAsDetailsLoadOptionEnum
const (
	SaveSqlTuningSetAsDetailsLoadOptionInsert SaveSqlTuningSetAsDetailsLoadOptionEnum = "INSERT"
	SaveSqlTuningSetAsDetailsLoadOptionUpdate SaveSqlTuningSetAsDetailsLoadOptionEnum = "UPDATE"
	SaveSqlTuningSetAsDetailsLoadOptionMerge  SaveSqlTuningSetAsDetailsLoadOptionEnum = "MERGE"
)

var mappingSaveSqlTuningSetAsDetailsLoadOptionEnum = map[string]SaveSqlTuningSetAsDetailsLoadOptionEnum{
	"INSERT": SaveSqlTuningSetAsDetailsLoadOptionInsert,
	"UPDATE": SaveSqlTuningSetAsDetailsLoadOptionUpdate,
	"MERGE":  SaveSqlTuningSetAsDetailsLoadOptionMerge,
}

var mappingSaveSqlTuningSetAsDetailsLoadOptionEnumLowerCase = map[string]SaveSqlTuningSetAsDetailsLoadOptionEnum{
	"insert": SaveSqlTuningSetAsDetailsLoadOptionInsert,
	"update": SaveSqlTuningSetAsDetailsLoadOptionUpdate,
	"merge":  SaveSqlTuningSetAsDetailsLoadOptionMerge,
}

// GetSaveSqlTuningSetAsDetailsLoadOptionEnumValues Enumerates the set of values for SaveSqlTuningSetAsDetailsLoadOptionEnum
func GetSaveSqlTuningSetAsDetailsLoadOptionEnumValues() []SaveSqlTuningSetAsDetailsLoadOptionEnum {
	values := make([]SaveSqlTuningSetAsDetailsLoadOptionEnum, 0)
	for _, v := range mappingSaveSqlTuningSetAsDetailsLoadOptionEnum {
		values = append(values, v)
	}
	return values
}

// GetSaveSqlTuningSetAsDetailsLoadOptionEnumStringValues Enumerates the set of values in String for SaveSqlTuningSetAsDetailsLoadOptionEnum
func GetSaveSqlTuningSetAsDetailsLoadOptionEnumStringValues() []string {
	return []string{
		"INSERT",
		"UPDATE",
		"MERGE",
	}
}

// GetMappingSaveSqlTuningSetAsDetailsLoadOptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSaveSqlTuningSetAsDetailsLoadOptionEnum(val string) (SaveSqlTuningSetAsDetailsLoadOptionEnum, bool) {
	enum, ok := mappingSaveSqlTuningSetAsDetailsLoadOptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SaveSqlTuningSetAsDetailsUpdateOptionEnum Enum with underlying type: string
type SaveSqlTuningSetAsDetailsUpdateOptionEnum string

// Set of constants representing the allowable values for SaveSqlTuningSetAsDetailsUpdateOptionEnum
const (
	SaveSqlTuningSetAsDetailsUpdateOptionReplace    SaveSqlTuningSetAsDetailsUpdateOptionEnum = "REPLACE"
	SaveSqlTuningSetAsDetailsUpdateOptionAccumulate SaveSqlTuningSetAsDetailsUpdateOptionEnum = "ACCUMULATE"
)

var mappingSaveSqlTuningSetAsDetailsUpdateOptionEnum = map[string]SaveSqlTuningSetAsDetailsUpdateOptionEnum{
	"REPLACE":    SaveSqlTuningSetAsDetailsUpdateOptionReplace,
	"ACCUMULATE": SaveSqlTuningSetAsDetailsUpdateOptionAccumulate,
}

var mappingSaveSqlTuningSetAsDetailsUpdateOptionEnumLowerCase = map[string]SaveSqlTuningSetAsDetailsUpdateOptionEnum{
	"replace":    SaveSqlTuningSetAsDetailsUpdateOptionReplace,
	"accumulate": SaveSqlTuningSetAsDetailsUpdateOptionAccumulate,
}

// GetSaveSqlTuningSetAsDetailsUpdateOptionEnumValues Enumerates the set of values for SaveSqlTuningSetAsDetailsUpdateOptionEnum
func GetSaveSqlTuningSetAsDetailsUpdateOptionEnumValues() []SaveSqlTuningSetAsDetailsUpdateOptionEnum {
	values := make([]SaveSqlTuningSetAsDetailsUpdateOptionEnum, 0)
	for _, v := range mappingSaveSqlTuningSetAsDetailsUpdateOptionEnum {
		values = append(values, v)
	}
	return values
}

// GetSaveSqlTuningSetAsDetailsUpdateOptionEnumStringValues Enumerates the set of values in String for SaveSqlTuningSetAsDetailsUpdateOptionEnum
func GetSaveSqlTuningSetAsDetailsUpdateOptionEnumStringValues() []string {
	return []string{
		"REPLACE",
		"ACCUMULATE",
	}
}

// GetMappingSaveSqlTuningSetAsDetailsUpdateOptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSaveSqlTuningSetAsDetailsUpdateOptionEnum(val string) (SaveSqlTuningSetAsDetailsUpdateOptionEnum, bool) {
	enum, ok := mappingSaveSqlTuningSetAsDetailsUpdateOptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SaveSqlTuningSetAsDetailsUpdateConditionEnum Enum with underlying type: string
type SaveSqlTuningSetAsDetailsUpdateConditionEnum string

// Set of constants representing the allowable values for SaveSqlTuningSetAsDetailsUpdateConditionEnum
const (
	SaveSqlTuningSetAsDetailsUpdateConditionOld SaveSqlTuningSetAsDetailsUpdateConditionEnum = "OLD"
	SaveSqlTuningSetAsDetailsUpdateConditionNew SaveSqlTuningSetAsDetailsUpdateConditionEnum = "NEW"
)

var mappingSaveSqlTuningSetAsDetailsUpdateConditionEnum = map[string]SaveSqlTuningSetAsDetailsUpdateConditionEnum{
	"OLD": SaveSqlTuningSetAsDetailsUpdateConditionOld,
	"NEW": SaveSqlTuningSetAsDetailsUpdateConditionNew,
}

var mappingSaveSqlTuningSetAsDetailsUpdateConditionEnumLowerCase = map[string]SaveSqlTuningSetAsDetailsUpdateConditionEnum{
	"old": SaveSqlTuningSetAsDetailsUpdateConditionOld,
	"new": SaveSqlTuningSetAsDetailsUpdateConditionNew,
}

// GetSaveSqlTuningSetAsDetailsUpdateConditionEnumValues Enumerates the set of values for SaveSqlTuningSetAsDetailsUpdateConditionEnum
func GetSaveSqlTuningSetAsDetailsUpdateConditionEnumValues() []SaveSqlTuningSetAsDetailsUpdateConditionEnum {
	values := make([]SaveSqlTuningSetAsDetailsUpdateConditionEnum, 0)
	for _, v := range mappingSaveSqlTuningSetAsDetailsUpdateConditionEnum {
		values = append(values, v)
	}
	return values
}

// GetSaveSqlTuningSetAsDetailsUpdateConditionEnumStringValues Enumerates the set of values in String for SaveSqlTuningSetAsDetailsUpdateConditionEnum
func GetSaveSqlTuningSetAsDetailsUpdateConditionEnumStringValues() []string {
	return []string{
		"OLD",
		"NEW",
	}
}

// GetMappingSaveSqlTuningSetAsDetailsUpdateConditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSaveSqlTuningSetAsDetailsUpdateConditionEnum(val string) (SaveSqlTuningSetAsDetailsUpdateConditionEnum, bool) {
	enum, ok := mappingSaveSqlTuningSetAsDetailsUpdateConditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
