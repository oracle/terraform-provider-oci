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

// ChangeTableCompartmentDetails Specification of both from and to compartments.
type ChangeTableCompartmentDetails struct {

	// The OCID of the table's new compartment.
	ToCompartmentId *string `mandatory:"true" json:"toCompartmentId"`

	// The OCID of the table's current compartment.  Required
	// if the tableNameOrId path parameter is a table name.
	// Optional if tableNameOrId is an OCID.  If tableNameOrId
	// is an OCID, and fromCompartmentId is supplied, the latter
	// must match the identified table's current compartmentId.
	FromCompartmentId *string `mandatory:"false" json:"fromCompartmentId"`
}

func (m ChangeTableCompartmentDetails) String() string {
	return common.PointerString(m)
}
