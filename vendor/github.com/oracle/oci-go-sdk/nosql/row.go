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

// Row The result of GetRow.
type Row struct {

	// The map of values from a row.
	Value map[string]interface{} `mandatory:"false" json:"value"`

	// The expiration time of the row. A zero value indicates that
	// the row does not expire. An RFC3339 formatted datetime
	// string.
	TimeOfExpiration *common.SDKTime `mandatory:"false" json:"timeOfExpiration"`

	Usage *RequestUsage `mandatory:"false" json:"usage"`
}

func (m Row) String() string {
	return common.PointerString(m)
}
