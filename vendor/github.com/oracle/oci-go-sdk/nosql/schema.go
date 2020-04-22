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

// Schema The table schema information as a JSON object.
type Schema struct {

	// The columns of a table.
	Columns []Column `mandatory:"true" json:"columns"`

	// A list of column names that make up a key.
	PrimaryKey []string `mandatory:"true" json:"primaryKey"`

	// A list of column names that make up a key.
	ShardKey []string `mandatory:"true" json:"shardKey"`

	// The default Time-to-Live for the table, in days.
	Ttl *int `mandatory:"true" json:"ttl"`
}

func (m Schema) String() string {
	return common.PointerString(m)
}
