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

// LoadSqlPlanBaselinesFromCursorCacheDetails The details of SQL statements and plans to be loaded from cursor cache. You can specify
// the plans to load using SQL ID, plan identifier, or filterName and filterValue pair.
// You can also control the SQL plan baseline into which the plans are loaded using either
// SQL text or SQL handle.
type LoadSqlPlanBaselinesFromCursorCacheDetails struct {

	// The name of the database job used for loading SQL plan baselines.
	JobName *string `mandatory:"true" json:"jobName"`

	Credentials ManagedDatabaseCredential `mandatory:"true" json:"credentials"`

	// The description of the job.
	JobDescription *string `mandatory:"false" json:"jobDescription"`

	// The SQL statement identifier. Identifies a SQL statement in the cursor cache.
	SqlId *string `mandatory:"false" json:"sqlId"`

	// The plan identifier. By default, all plans present in the cursor cache
	// for the SQL statement identified by `sqlId` are captured.
	PlanHash *float32 `mandatory:"false" json:"planHash"`

	// The SQL text to use in identifying the SQL plan baseline into which the plans
	// are loaded. If the SQL plan baseline does not exist, it is created.
	SqlText *string `mandatory:"false" json:"sqlText"`

	// The SQL handle to use in identifying the SQL plan baseline into which
	// the plans are loaded.
	SqlHandle *string `mandatory:"false" json:"sqlHandle"`

	// The name of the filter.
	// - SQL_TEXT: Search pattern to apply to SQL text.
	// - PARSING_SCHEMA_NAME: Name of the parsing schema.
	// - MODULE: Name of the module.
	// - ACTION: Name of the action.
	FilterName LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum `mandatory:"false" json:"filterName,omitempty"`

	// The filter value. It is upper-cased except when it is enclosed in
	// double quotes or filter name is `SQL_TEXT`.
	FilterValue *string `mandatory:"false" json:"filterValue"`

	// Indicates whether the plans are loaded as fixed plans (`true`) or non-fixed plans (`false`).
	// By default, they are loaded as non-fixed plans.
	IsFixed *bool `mandatory:"false" json:"isFixed"`

	// Indicates whether the loaded plans are enabled (`true`) or not (`false`).
	// By default, they are enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`
}

func (m LoadSqlPlanBaselinesFromCursorCacheDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LoadSqlPlanBaselinesFromCursorCacheDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum(string(m.FilterName)); !ok && m.FilterName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FilterName: %s. Supported values are: %s.", m.FilterName, strings.Join(GetLoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *LoadSqlPlanBaselinesFromCursorCacheDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		JobDescription *string                                                  `json:"jobDescription"`
		SqlId          *string                                                  `json:"sqlId"`
		PlanHash       *float32                                                 `json:"planHash"`
		SqlText        *string                                                  `json:"sqlText"`
		SqlHandle      *string                                                  `json:"sqlHandle"`
		FilterName     LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum `json:"filterName"`
		FilterValue    *string                                                  `json:"filterValue"`
		IsFixed        *bool                                                    `json:"isFixed"`
		IsEnabled      *bool                                                    `json:"isEnabled"`
		JobName        *string                                                  `json:"jobName"`
		Credentials    manageddatabasecredential                                `json:"credentials"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.JobDescription = model.JobDescription

	m.SqlId = model.SqlId

	m.PlanHash = model.PlanHash

	m.SqlText = model.SqlText

	m.SqlHandle = model.SqlHandle

	m.FilterName = model.FilterName

	m.FilterValue = model.FilterValue

	m.IsFixed = model.IsFixed

	m.IsEnabled = model.IsEnabled

	m.JobName = model.JobName

	nn, e = model.Credentials.UnmarshalPolymorphicJSON(model.Credentials.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Credentials = nn.(ManagedDatabaseCredential)
	} else {
		m.Credentials = nil
	}

	return
}

// LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum Enum with underlying type: string
type LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum string

// Set of constants representing the allowable values for LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum
const (
	LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameSqlText           LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum = "SQL_TEXT"
	LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameParsingSchemaName LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum = "PARSING_SCHEMA_NAME"
	LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameModule            LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum = "MODULE"
	LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameAction            LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum = "ACTION"
)

var mappingLoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum = map[string]LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum{
	"SQL_TEXT":            LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameSqlText,
	"PARSING_SCHEMA_NAME": LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameParsingSchemaName,
	"MODULE":              LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameModule,
	"ACTION":              LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameAction,
}

var mappingLoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnumLowerCase = map[string]LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum{
	"sql_text":            LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameSqlText,
	"parsing_schema_name": LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameParsingSchemaName,
	"module":              LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameModule,
	"action":              LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameAction,
}

// GetLoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnumValues Enumerates the set of values for LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum
func GetLoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnumValues() []LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum {
	values := make([]LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum, 0)
	for _, v := range mappingLoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum {
		values = append(values, v)
	}
	return values
}

// GetLoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnumStringValues Enumerates the set of values in String for LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum
func GetLoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnumStringValues() []string {
	return []string{
		"SQL_TEXT",
		"PARSING_SCHEMA_NAME",
		"MODULE",
		"ACTION",
	}
}

// GetMappingLoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum(val string) (LoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnum, bool) {
	enum, ok := mappingLoadSqlPlanBaselinesFromCursorCacheDetailsFilterNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
