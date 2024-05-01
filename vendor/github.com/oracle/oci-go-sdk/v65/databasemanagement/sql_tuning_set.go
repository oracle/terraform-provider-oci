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

// SqlTuningSet Details of the Sql tuning set.
type SqlTuningSet struct {

	// The owner of the Sql tuning set.
	Owner *string `mandatory:"true" json:"owner"`

	// The name of the Sql tuning set.
	Name *string `mandatory:"true" json:"name"`

	// The unique Sql tuning set identifier.
	Id *int `mandatory:"false" json:"id"`

	// Number of statements in the Sql tuning set
	StatementCount *int `mandatory:"false" json:"statementCount"`

	// The created time of the Sql tuning set.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The description of the Sql tuning set.
	Description *string `mandatory:"false" json:"description"`

	// Last modified time of the Sql tuning set.
	TimeLastModified *common.SDKTime `mandatory:"false" json:"timeLastModified"`

	// Current status of the Sql tuning set.
	Status SqlTuningSetStatusTypesEnum `mandatory:"false" json:"status,omitempty"`

	// Name of the Sql tuning set scheduler job.
	ScheduledJobName *string `mandatory:"false" json:"scheduledJobName"`

	// Latest execution error of the plsql that was submitted as a scheduler job.
	ErrorMessage *string `mandatory:"false" json:"errorMessage"`

	// In OCI database management, there is a limit to fetch only 2000 rows.
	// This flag indicates whether all Sql statements of this Sql tuning set matching the filter criteria are fetched or not.
	// Possible values are 'Yes' or 'No'
	//   - Yes - All Sql statements matching the filter criteria are fetched.
	//   - No  - There are more Sql statements matching the fitler criteria.
	//           User should fine tune the filter criteria to narrow down the result set.
	AllSqlStatementsFetched SqlTuningSetAllSqlStatementsFetchedEnum `mandatory:"false" json:"allSqlStatementsFetched,omitempty"`

	// A list of the Sqls associated with the Sql tuning set.
	SqlList []SqlInSqlTuningSet `mandatory:"false" json:"sqlList"`
}

func (m SqlTuningSet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlTuningSet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSqlTuningSetStatusTypesEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetSqlTuningSetStatusTypesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSqlTuningSetAllSqlStatementsFetchedEnum(string(m.AllSqlStatementsFetched)); !ok && m.AllSqlStatementsFetched != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AllSqlStatementsFetched: %s. Supported values are: %s.", m.AllSqlStatementsFetched, strings.Join(GetSqlTuningSetAllSqlStatementsFetchedEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SqlTuningSetAllSqlStatementsFetchedEnum Enum with underlying type: string
type SqlTuningSetAllSqlStatementsFetchedEnum string

// Set of constants representing the allowable values for SqlTuningSetAllSqlStatementsFetchedEnum
const (
	SqlTuningSetAllSqlStatementsFetchedYes SqlTuningSetAllSqlStatementsFetchedEnum = "YES"
	SqlTuningSetAllSqlStatementsFetchedNo  SqlTuningSetAllSqlStatementsFetchedEnum = "NO"
)

var mappingSqlTuningSetAllSqlStatementsFetchedEnum = map[string]SqlTuningSetAllSqlStatementsFetchedEnum{
	"YES": SqlTuningSetAllSqlStatementsFetchedYes,
	"NO":  SqlTuningSetAllSqlStatementsFetchedNo,
}

var mappingSqlTuningSetAllSqlStatementsFetchedEnumLowerCase = map[string]SqlTuningSetAllSqlStatementsFetchedEnum{
	"yes": SqlTuningSetAllSqlStatementsFetchedYes,
	"no":  SqlTuningSetAllSqlStatementsFetchedNo,
}

// GetSqlTuningSetAllSqlStatementsFetchedEnumValues Enumerates the set of values for SqlTuningSetAllSqlStatementsFetchedEnum
func GetSqlTuningSetAllSqlStatementsFetchedEnumValues() []SqlTuningSetAllSqlStatementsFetchedEnum {
	values := make([]SqlTuningSetAllSqlStatementsFetchedEnum, 0)
	for _, v := range mappingSqlTuningSetAllSqlStatementsFetchedEnum {
		values = append(values, v)
	}
	return values
}

// GetSqlTuningSetAllSqlStatementsFetchedEnumStringValues Enumerates the set of values in String for SqlTuningSetAllSqlStatementsFetchedEnum
func GetSqlTuningSetAllSqlStatementsFetchedEnumStringValues() []string {
	return []string{
		"YES",
		"NO",
	}
}

// GetMappingSqlTuningSetAllSqlStatementsFetchedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSqlTuningSetAllSqlStatementsFetchedEnum(val string) (SqlTuningSetAllSqlStatementsFetchedEnum, bool) {
	enum, ok := mappingSqlTuningSetAllSqlStatementsFetchedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
