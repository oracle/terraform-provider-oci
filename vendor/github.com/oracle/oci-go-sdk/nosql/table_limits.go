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

// TableLimits Throughput and storage limits configuration of a table.
type TableLimits struct {

	// Maximum sustained read throughput limit for the table.
	MaxReadUnits *int `mandatory:"true" json:"maxReadUnits"`

	// Maximum sustained write throughput limit for the table.
	MaxWriteUnits *int `mandatory:"true" json:"maxWriteUnits"`

	// Maximum size of storage used by the table.
	MaxStorageInGBs *int `mandatory:"true" json:"maxStorageInGBs"`
}

func (m TableLimits) String() string {
	return common.PointerString(m)
}
