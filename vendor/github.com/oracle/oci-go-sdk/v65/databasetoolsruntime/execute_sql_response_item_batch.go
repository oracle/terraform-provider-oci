// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExecuteSqlResponseItemBatch A statement execution response for batched type requests.
type ExecuteSqlResponseItemBatch struct {

	// The sequence number of the statement. Sequence number of the first statement is 1.
	StatementId *int `mandatory:"false" json:"statementId"`

	// The Statement type.
	StatementType StatementTypeEnum `mandatory:"false" json:"statementType,omitempty"`

	StatementPos *ExecuteSqlResponseItemStatementPos `mandatory:"false" json:"statementPos"`

	// DML statements to execute in jdbc batch mode
	BatchStatementTexts []string `mandatory:"false" json:"batchStatementTexts"`

	Error *ExecuteSqlResponseItemError `mandatory:"false" json:"error"`

	// Output from DBMS_OUTPUT package. Server output must be enabled (e.g., SET SERVEROUTPUT ON).
	DbmsOutput *string `mandatory:"false" json:"dbmsOutput"`

	Properties *ExecuteSqlResponseItemProperties `mandatory:"false" json:"properties"`

	// Responses generated when executing the statements.
	Responses []string `mandatory:"false" json:"responses"`

	// Results generated when executing the statements.
	Results []int `mandatory:"false" json:"results"`
}

func (m ExecuteSqlResponseItemBatch) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteSqlResponseItemBatch) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStatementTypeEnum(string(m.StatementType)); !ok && m.StatementType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StatementType: %s. Supported values are: %s.", m.StatementType, strings.Join(GetStatementTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
