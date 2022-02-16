// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// PrivateEndpointDetails OCI Private Endpoint configuration details.
type PrivateEndpointDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment to contain the
	// private endpoint.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VCN where the Private Endpoint will be bound to.
	VcnId *string `mandatory:"false" json:"vcnId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the customer's
	// subnet where the private endpoint VNIC will reside.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of a previously created Private Endpoint.
	Id *string `mandatory:"false" json:"id"`
}

func (m PrivateEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
