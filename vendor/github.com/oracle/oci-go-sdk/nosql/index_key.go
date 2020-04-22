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

// IndexKey Specifies a single key in a secondary index.
type IndexKey struct {

	// The name of a column to be included as an index key.
	ColumnName *string `mandatory:"true" json:"columnName"`

	// If the specified column is of type JSON, jsonPath contains
	// a dotted path indicating the field within the JSON object
	// that will be the index key.
	JsonPath *string `mandatory:"false" json:"jsonPath"`

	// If the specified column is of type JSON, jsonFieldType contains
	// the type of the field indicated by jsonPath.
	JsonFieldType *string `mandatory:"false" json:"jsonFieldType"`
}

func (m IndexKey) String() string {
	return common.PointerString(m)
}
