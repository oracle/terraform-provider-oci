// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
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

// QueryResultCollection The result of a query.
type QueryResultCollection struct {

	// Array of objects representing query results.
	Items []map[string]interface{} `mandatory:"false" json:"items"`

	Usage *RequestUsage `mandatory:"false" json:"usage"`
}

func (m QueryResultCollection) String() string {
	return common.PointerString(m)
}
