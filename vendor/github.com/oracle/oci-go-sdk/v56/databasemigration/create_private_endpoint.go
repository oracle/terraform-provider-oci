// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreatePrivateEndpoint OCI Private Endpoint configuration details.
// Not required for source container database connections, it will default to the specified Source Database Connection Private Endpoint.
type CreatePrivateEndpoint struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment to contain the
	// private endpoint.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VCN where the Private Endpoint will be bound to.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the customer's subnet where the private endpoint VNIC
	// will reside.
	SubnetId *string `mandatory:"true" json:"subnetId"`
}

func (m CreatePrivateEndpoint) String() string {
	return common.PointerString(m)
}
