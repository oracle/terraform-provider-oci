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

// CreateIndexDetails Specifications for the new index.
type CreateIndexDetails struct {

	// Index name.
	Name *string `mandatory:"true" json:"name"`

	// A set of keys for a secondary index.
	Keys []IndexKey `mandatory:"true" json:"keys"`

	// The OCID of the table's compartment.  Required
	// if the tableNameOrId path parameter is a table name.
	// Optional if tableNameOrId is an OCID.  If tableNameOrId
	// is an OCID, and compartmentId is supplied, the latter
	// must match the identified table's compartmentId.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// If true, the operation completes successfully even when the
	// index exists.  Otherwise, an attempt to create an index
	// that already exists will return an error.
	IsIfNotExists *bool `mandatory:"false" json:"isIfNotExists"`
}

func (m CreateIndexDetails) String() string {
	return common.PointerString(m)
}
