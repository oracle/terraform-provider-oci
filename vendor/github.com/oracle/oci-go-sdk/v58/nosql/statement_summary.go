// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NoSQL Database API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// StatementSummary Information derived from parsing a NoSQL SQL statement.
type StatementSummary struct {

	// The operation represented in the statement, e.g. CREATE_TABLE.
	Operation StatementSummaryOperationEnum `mandatory:"false" json:"operation,omitempty"`

	// The table name from the SQL statement.
	TableName *string `mandatory:"false" json:"tableName"`

	// The index name from the SQL statement, if present.
	IndexName *string `mandatory:"false" json:"indexName"`

	// True if the statement includes "IF EXISTS."
	IsIfExists *bool `mandatory:"false" json:"isIfExists"`

	// True if the statement includes "IF NOT EXISTS."
	IsIfNotExists *bool `mandatory:"false" json:"isIfNotExists"`

	// If present, indicates a syntax error in the statement.
	SyntaxError *string `mandatory:"false" json:"syntaxError"`
}

func (m StatementSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StatementSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStatementSummaryOperationEnum(string(m.Operation)); !ok && m.Operation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operation: %s. Supported values are: %s.", m.Operation, strings.Join(GetStatementSummaryOperationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StatementSummaryOperationEnum Enum with underlying type: string
type StatementSummaryOperationEnum string

// Set of constants representing the allowable values for StatementSummaryOperationEnum
const (
	StatementSummaryOperationCreateTable StatementSummaryOperationEnum = "CREATE_TABLE"
	StatementSummaryOperationAlterTable  StatementSummaryOperationEnum = "ALTER_TABLE"
	StatementSummaryOperationDropTable   StatementSummaryOperationEnum = "DROP_TABLE"
	StatementSummaryOperationCreateIndex StatementSummaryOperationEnum = "CREATE_INDEX"
	StatementSummaryOperationDropIndex   StatementSummaryOperationEnum = "DROP_INDEX"
	StatementSummaryOperationSelect      StatementSummaryOperationEnum = "SELECT"
	StatementSummaryOperationUpdate      StatementSummaryOperationEnum = "UPDATE"
	StatementSummaryOperationInsert      StatementSummaryOperationEnum = "INSERT"
	StatementSummaryOperationDelete      StatementSummaryOperationEnum = "DELETE"
)

var mappingStatementSummaryOperationEnum = map[string]StatementSummaryOperationEnum{
	"CREATE_TABLE": StatementSummaryOperationCreateTable,
	"ALTER_TABLE":  StatementSummaryOperationAlterTable,
	"DROP_TABLE":   StatementSummaryOperationDropTable,
	"CREATE_INDEX": StatementSummaryOperationCreateIndex,
	"DROP_INDEX":   StatementSummaryOperationDropIndex,
	"SELECT":       StatementSummaryOperationSelect,
	"UPDATE":       StatementSummaryOperationUpdate,
	"INSERT":       StatementSummaryOperationInsert,
	"DELETE":       StatementSummaryOperationDelete,
}

// GetStatementSummaryOperationEnumValues Enumerates the set of values for StatementSummaryOperationEnum
func GetStatementSummaryOperationEnumValues() []StatementSummaryOperationEnum {
	values := make([]StatementSummaryOperationEnum, 0)
	for _, v := range mappingStatementSummaryOperationEnum {
		values = append(values, v)
	}
	return values
}

// GetStatementSummaryOperationEnumStringValues Enumerates the set of values in String for StatementSummaryOperationEnum
func GetStatementSummaryOperationEnumStringValues() []string {
	return []string{
		"CREATE_TABLE",
		"ALTER_TABLE",
		"DROP_TABLE",
		"CREATE_INDEX",
		"DROP_INDEX",
		"SELECT",
		"UPDATE",
		"INSERT",
		"DELETE",
	}
}

// GetMappingStatementSummaryOperationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStatementSummaryOperationEnum(val string) (StatementSummaryOperationEnum, bool) {
	mappingStatementSummaryOperationEnumIgnoreCase := make(map[string]StatementSummaryOperationEnum)
	for k, v := range mappingStatementSummaryOperationEnum {
		mappingStatementSummaryOperationEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingStatementSummaryOperationEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
