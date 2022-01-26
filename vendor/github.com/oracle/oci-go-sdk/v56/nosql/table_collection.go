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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// TableCollection Results of ListTables.
type TableCollection struct {

	// A page of TableSummary objects.
	Items []TableSummary `mandatory:"false" json:"items"`

	// The maximum number of reclaimable tables allowed in the tenancy.
	MaxAutoReclaimableTables *int `mandatory:"false" json:"maxAutoReclaimableTables"`

	// The current number of reclaimable tables in the tenancy.
	AutoReclaimableTables *int `mandatory:"false" json:"autoReclaimableTables"`
}

func (m TableCollection) String() string {
	return common.PointerString(m)
}
