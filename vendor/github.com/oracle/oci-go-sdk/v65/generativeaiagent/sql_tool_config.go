// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Generative AI Agents Management API
//
// OCI Generative AI Agents is a fully managed service that combines the power of large language models (LLMs) with an intelligent retrieval system to create contextually relevant answers by searching your knowledge base, making your AI applications smart and efficient.
// OCI Generative AI Agents supports several ways to onboard your data and then allows you and your customers to interact with your data using a chat interface or API.
// Use the Generative AI Agents API to create and manage agents, knowledge bases, data sources, endpoints, data ingestion jobs, and work requests.
// For creating and managing client chat sessions see the /EN/generative-ai-agents-client/latest/.
// To learn more about the service, see the Generative AI Agents documentation (https://docs.oracle.com/iaas/Content/generative-ai-agents/home.htm).
//

package generativeaiagent

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlToolConfig The configuration for SQL Tool.
type SqlToolConfig struct {
	IclExamples InputLocation `mandatory:"false" json:"iclExamples"`

	DatabaseSchema InputLocation `mandatory:"false" json:"databaseSchema"`

	// To enable/disable SQL execution.
	ShouldEnableSqlExecution *bool `mandatory:"false" json:"shouldEnableSqlExecution"`

	// To enable/disable self correction.
	ShouldEnableSelfCorrection *bool `mandatory:"false" json:"shouldEnableSelfCorrection"`

	TableAndColumnDescription InputLocation `mandatory:"false" json:"tableAndColumnDescription"`

	GenerationLlmCustomization *LlmCustomization `mandatory:"false" json:"generationLlmCustomization"`

	DatabaseConnection DatabaseConnection `mandatory:"false" json:"databaseConnection"`

	// Dialect to be used for SQL generation.
	Dialect SqlToolConfigDialectEnum `mandatory:"true" json:"dialect"`

	// Size of the model.
	ModelSize SqlToolConfigModelSizeEnum `mandatory:"false" json:"modelSize,omitempty"`
}

func (m SqlToolConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlToolConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlToolConfigDialectEnum(string(m.Dialect)); !ok && m.Dialect != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Dialect: %s. Supported values are: %s.", m.Dialect, strings.Join(GetSqlToolConfigDialectEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlToolConfigModelSizeEnum(string(m.ModelSize)); !ok && m.ModelSize != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelSize: %s. Supported values are: %s.", m.ModelSize, strings.Join(GetSqlToolConfigModelSizeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SqlToolConfig) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSqlToolConfig SqlToolConfig
	s := struct {
		DiscriminatorParam string `json:"toolConfigType"`
		MarshalTypeSqlToolConfig
	}{
		"SQL_TOOL_CONFIG",
		(MarshalTypeSqlToolConfig)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *SqlToolConfig) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		IclExamples                inputlocation              `json:"iclExamples"`
		DatabaseSchema             inputlocation              `json:"databaseSchema"`
		ShouldEnableSqlExecution   *bool                      `json:"shouldEnableSqlExecution"`
		ModelSize                  SqlToolConfigModelSizeEnum `json:"modelSize"`
		ShouldEnableSelfCorrection *bool                      `json:"shouldEnableSelfCorrection"`
		TableAndColumnDescription  inputlocation              `json:"tableAndColumnDescription"`
		GenerationLlmCustomization *LlmCustomization          `json:"generationLlmCustomization"`
		DatabaseConnection         databaseconnection         `json:"databaseConnection"`
		Dialect                    SqlToolConfigDialectEnum   `json:"dialect"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.IclExamples.UnmarshalPolymorphicJSON(model.IclExamples.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.IclExamples = nn.(InputLocation)
	} else {
		m.IclExamples = nil
	}

	nn, e = model.DatabaseSchema.UnmarshalPolymorphicJSON(model.DatabaseSchema.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatabaseSchema = nn.(InputLocation)
	} else {
		m.DatabaseSchema = nil
	}

	m.ShouldEnableSqlExecution = model.ShouldEnableSqlExecution

	m.ModelSize = model.ModelSize

	m.ShouldEnableSelfCorrection = model.ShouldEnableSelfCorrection

	nn, e = model.TableAndColumnDescription.UnmarshalPolymorphicJSON(model.TableAndColumnDescription.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TableAndColumnDescription = nn.(InputLocation)
	} else {
		m.TableAndColumnDescription = nil
	}

	m.GenerationLlmCustomization = model.GenerationLlmCustomization

	nn, e = model.DatabaseConnection.UnmarshalPolymorphicJSON(model.DatabaseConnection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatabaseConnection = nn.(DatabaseConnection)
	} else {
		m.DatabaseConnection = nil
	}

	m.Dialect = model.Dialect

	return
}

// SqlToolConfigDialectEnum Enum with underlying type: string
type SqlToolConfigDialectEnum string

// Set of constants representing the allowable values for SqlToolConfigDialectEnum
const (
	SqlToolConfigDialectSqlLite   SqlToolConfigDialectEnum = "SQL_LITE"
	SqlToolConfigDialectOracleSql SqlToolConfigDialectEnum = "ORACLE_SQL"
)

var mappingSqlToolConfigDialectEnum = map[string]SqlToolConfigDialectEnum{
	"SQL_LITE":   SqlToolConfigDialectSqlLite,
	"ORACLE_SQL": SqlToolConfigDialectOracleSql,
}

var mappingSqlToolConfigDialectEnumLowerCase = map[string]SqlToolConfigDialectEnum{
	"sql_lite":   SqlToolConfigDialectSqlLite,
	"oracle_sql": SqlToolConfigDialectOracleSql,
}

// GetSqlToolConfigDialectEnumValues Enumerates the set of values for SqlToolConfigDialectEnum
func GetSqlToolConfigDialectEnumValues() []SqlToolConfigDialectEnum {
	values := make([]SqlToolConfigDialectEnum, 0)
	for _, v := range mappingSqlToolConfigDialectEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlToolConfigDialectEnumStringValues Enumerates the set of values in String for SqlToolConfigDialectEnum
func GetSqlToolConfigDialectEnumStringValues() []string {
	return []string{
		"SQL_LITE",
		"ORACLE_SQL",
	}
}

// GetMappingSqlToolConfigDialectEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlToolConfigDialectEnum(val string) (SqlToolConfigDialectEnum, bool) {
	enum, ok := mappingSqlToolConfigDialectEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SqlToolConfigModelSizeEnum Enum with underlying type: string
type SqlToolConfigModelSizeEnum string

// Set of constants representing the allowable values for SqlToolConfigModelSizeEnum
const (
	SqlToolConfigModelSizeSmall SqlToolConfigModelSizeEnum = "SMALL"
	SqlToolConfigModelSizeLarge SqlToolConfigModelSizeEnum = "LARGE"
)

var mappingSqlToolConfigModelSizeEnum = map[string]SqlToolConfigModelSizeEnum{
	"SMALL": SqlToolConfigModelSizeSmall,
	"LARGE": SqlToolConfigModelSizeLarge,
}

var mappingSqlToolConfigModelSizeEnumLowerCase = map[string]SqlToolConfigModelSizeEnum{
	"small": SqlToolConfigModelSizeSmall,
	"large": SqlToolConfigModelSizeLarge,
}

// GetSqlToolConfigModelSizeEnumValues Enumerates the set of values for SqlToolConfigModelSizeEnum
func GetSqlToolConfigModelSizeEnumValues() []SqlToolConfigModelSizeEnum {
	values := make([]SqlToolConfigModelSizeEnum, 0)
	for _, v := range mappingSqlToolConfigModelSizeEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlToolConfigModelSizeEnumStringValues Enumerates the set of values in String for SqlToolConfigModelSizeEnum
func GetSqlToolConfigModelSizeEnumStringValues() []string {
	return []string{
		"SMALL",
		"LARGE",
	}
}

// GetMappingSqlToolConfigModelSizeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlToolConfigModelSizeEnum(val string) (SqlToolConfigModelSizeEnum, bool) {
	enum, ok := mappingSqlToolConfigModelSizeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
