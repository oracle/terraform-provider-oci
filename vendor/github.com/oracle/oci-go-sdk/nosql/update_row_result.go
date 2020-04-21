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

// UpdateRowResult The result of an UpdateRow operation.
type UpdateRowResult struct {

	// An opaque version string associated with the row.
	Version *string `mandatory:"false" json:"version"`

	// The version string associated with the existing row.
	// Returned if the put fails due to options setting in the
	// request.
	ExistingVersion *string `mandatory:"false" json:"existingVersion"`

	// The map of values from a row.
	ExistingValue map[string]interface{} `mandatory:"false" json:"existingValue"`

	// The value generated if the operation created a new value for
	// an identity column. If the table has no identity column, this value
	// is null. If it has an identity column, and a value was generated for
	// that column, it is non-null.
	GeneratedValue *string `mandatory:"false" json:"generatedValue"`

	Usage *RequestUsage `mandatory:"false" json:"usage"`
}

func (m UpdateRowResult) String() string {
	return common.PointerString(m)
}
