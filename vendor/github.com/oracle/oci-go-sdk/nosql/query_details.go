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

// QueryDetails All the information surrounding a query, including the query statement,
// limits, consistency, and so forth.
type QueryDetails struct {

	// Compartment OCID, to provide context for a table name in
	// the given statement.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A NoSQL SQL query statement; or a Base64-encoded prepared statement.
	Statement *string `mandatory:"true" json:"statement"`

	// If true, the statement is a prepared statement.
	IsPrepared *bool `mandatory:"false" json:"isPrepared"`

	// Consistency requirement for a read operation.
	Consistency QueryDetailsConsistencyEnum `mandatory:"false" json:"consistency,omitempty"`

	// A limit on the total amount of data read during this operation, in KB.
	MaxReadInKBs *int `mandatory:"false" json:"maxReadInKBs"`

	// A map of prepared statement variables to values.
	Variables map[string]interface{} `mandatory:"false" json:"variables"`

	// Timeout setting for the query.
	TimeoutInMs *int `mandatory:"false" json:"timeoutInMs"`
}

func (m QueryDetails) String() string {
	return common.PointerString(m)
}

// QueryDetailsConsistencyEnum Enum with underlying type: string
type QueryDetailsConsistencyEnum string

// Set of constants representing the allowable values for QueryDetailsConsistencyEnum
const (
	QueryDetailsConsistencyEventual QueryDetailsConsistencyEnum = "EVENTUAL"
	QueryDetailsConsistencyAbsolute QueryDetailsConsistencyEnum = "ABSOLUTE"
)

var mappingQueryDetailsConsistency = map[string]QueryDetailsConsistencyEnum{
	"EVENTUAL": QueryDetailsConsistencyEventual,
	"ABSOLUTE": QueryDetailsConsistencyAbsolute,
}

// GetQueryDetailsConsistencyEnumValues Enumerates the set of values for QueryDetailsConsistencyEnum
func GetQueryDetailsConsistencyEnumValues() []QueryDetailsConsistencyEnum {
	values := make([]QueryDetailsConsistencyEnum, 0)
	for _, v := range mappingQueryDetailsConsistency {
		values = append(values, v)
	}
	return values
}
