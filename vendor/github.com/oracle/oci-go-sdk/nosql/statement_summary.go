// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ndcs-control-plane API
//
// The control plane API for NoSQL Database Cloud Service HTTPS
// provides endpoints to perform NDCS operations, including creation
// and deletion of tables and indexes; population and access of data
// in tables; and access of table usage metrics.
//

package nosql

import (
	"github.com/oracle/oci-go-sdk/common"
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

var mappingStatementSummaryOperation = map[string]StatementSummaryOperationEnum{
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
	for _, v := range mappingStatementSummaryOperation {
		values = append(values, v)
	}
	return values
}
