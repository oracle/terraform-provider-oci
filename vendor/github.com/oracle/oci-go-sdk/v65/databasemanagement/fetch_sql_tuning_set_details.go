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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FetchSqlTuningSetDetails The details required to fetch the Sql tuning set details.
// It takes either credentialDetails or databaseCredential. It's recommended to provide databaseCredential
type FetchSqlTuningSetDetails struct {

	// The owner of the Sql tuning set.
	Owner *string `mandatory:"true" json:"owner"`

	// The name of the Sql tuning set.
	Name *string `mandatory:"true" json:"name"`

	CredentialDetails SqlTuningSetAdminCredentialDetails `mandatory:"false" json:"credentialDetails"`

	DatabaseCredential DatabaseCredentialDetails `mandatory:"false" json:"databaseCredential"`

	// Specifies the Sql predicate to filter the Sql from the Sql tuning set defined on attributes of the SQLSET_ROW.
	// User could use any combination of the following columns with appropriate values as Sql predicate
	// Refer to the documentation https://docs.oracle.com/en/database/oracle/oracle-database/18/arpls/DBMS_SQLTUNE.html#GUID-1F4AFB03-7B29-46FC-B3F2-CB01EC36326C
	BasicFilter *string `mandatory:"false" json:"basicFilter"`

	// Specifies that the filter must include recursive Sql in the Sql tuning set.
	RecursiveSql FetchSqlTuningSetDetailsRecursiveSqlEnum `mandatory:"false" json:"recursiveSql,omitempty"`

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
}

func (m FetchSqlTuningSetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FetchSqlTuningSetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFetchSqlTuningSetDetailsRecursiveSqlEnum(string(m.RecursiveSql)); !ok && m.RecursiveSql != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RecursiveSql: %s. Supported values are: %s.", m.RecursiveSql, strings.Join(GetFetchSqlTuningSetDetailsRecursiveSqlEnumStringValues(), ",")))
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
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *FetchSqlTuningSetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CredentialDetails  sqltuningsetadmincredentialdetails       `json:"credentialDetails"`
		DatabaseCredential databasecredentialdetails                `json:"databaseCredential"`
		BasicFilter        *string                                  `json:"basicFilter"`
		RecursiveSql       FetchSqlTuningSetDetailsRecursiveSqlEnum `json:"recursiveSql"`
		ResultPercentage   *float64                                 `json:"resultPercentage"`
		ResultLimit        *int                                     `json:"resultLimit"`
		RankingMeasure1    RankingMeasureEnum                       `json:"rankingMeasure1"`
		RankingMeasure2    RankingMeasureEnum                       `json:"rankingMeasure2"`
		RankingMeasure3    RankingMeasureEnum                       `json:"rankingMeasure3"`
		Owner              *string                                  `json:"owner"`
		Name               *string                                  `json:"name"`
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

	m.BasicFilter = model.BasicFilter

	m.RecursiveSql = model.RecursiveSql

	m.ResultPercentage = model.ResultPercentage

	m.ResultLimit = model.ResultLimit

	m.RankingMeasure1 = model.RankingMeasure1

	m.RankingMeasure2 = model.RankingMeasure2

	m.RankingMeasure3 = model.RankingMeasure3

	m.Owner = model.Owner

	m.Name = model.Name

	return
}

// FetchSqlTuningSetDetailsRecursiveSqlEnum Enum with underlying type: string
type FetchSqlTuningSetDetailsRecursiveSqlEnum string

// Set of constants representing the allowable values for FetchSqlTuningSetDetailsRecursiveSqlEnum
const (
	FetchSqlTuningSetDetailsRecursiveSqlHasRecursiveSql FetchSqlTuningSetDetailsRecursiveSqlEnum = "HAS_RECURSIVE_SQL"
	FetchSqlTuningSetDetailsRecursiveSqlNoRecursiveSql  FetchSqlTuningSetDetailsRecursiveSqlEnum = "NO_RECURSIVE_SQL"
)

var mappingFetchSqlTuningSetDetailsRecursiveSqlEnum = map[string]FetchSqlTuningSetDetailsRecursiveSqlEnum{
	"HAS_RECURSIVE_SQL": FetchSqlTuningSetDetailsRecursiveSqlHasRecursiveSql,
	"NO_RECURSIVE_SQL":  FetchSqlTuningSetDetailsRecursiveSqlNoRecursiveSql,
}

var mappingFetchSqlTuningSetDetailsRecursiveSqlEnumLowerCase = map[string]FetchSqlTuningSetDetailsRecursiveSqlEnum{
	"has_recursive_sql": FetchSqlTuningSetDetailsRecursiveSqlHasRecursiveSql,
	"no_recursive_sql":  FetchSqlTuningSetDetailsRecursiveSqlNoRecursiveSql,
}

// GetFetchSqlTuningSetDetailsRecursiveSqlEnumValues Enumerates the set of values for FetchSqlTuningSetDetailsRecursiveSqlEnum
func GetFetchSqlTuningSetDetailsRecursiveSqlEnumValues() []FetchSqlTuningSetDetailsRecursiveSqlEnum {
	values := make([]FetchSqlTuningSetDetailsRecursiveSqlEnum, 0)
	for _, v := range mappingFetchSqlTuningSetDetailsRecursiveSqlEnum {
		values = append(values, v)
	}
	return values
}

// GetFetchSqlTuningSetDetailsRecursiveSqlEnumStringValues Enumerates the set of values in String for FetchSqlTuningSetDetailsRecursiveSqlEnum
func GetFetchSqlTuningSetDetailsRecursiveSqlEnumStringValues() []string {
	return []string{
		"HAS_RECURSIVE_SQL",
		"NO_RECURSIVE_SQL",
	}
}

// GetMappingFetchSqlTuningSetDetailsRecursiveSqlEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFetchSqlTuningSetDetailsRecursiveSqlEnum(val string) (FetchSqlTuningSetDetailsRecursiveSqlEnum, bool) {
	enum, ok := mappingFetchSqlTuningSetDetailsRecursiveSqlEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
