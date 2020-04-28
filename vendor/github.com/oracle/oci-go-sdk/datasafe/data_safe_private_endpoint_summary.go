// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/common"
)

// DataSafePrivateEndpointSummary Summary of a Data Safe private endpoint.
type DataSafePrivateEndpointSummary struct {

	// The OCID of the Data Safe private endpoint.
	Id *string `mandatory:"true" json:"id"`

	// The display name of the private endpoint.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the VCN.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The OCID of the subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The OCID of the private endpoint.
	PrivateEndpointId *string `mandatory:"true" json:"privateEndpointId"`

	// The description of the private endpoint.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the private endpoint was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The current state of the private endpoint.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m DataSafePrivateEndpointSummary) String() string {
	return common.PointerString(m)
}
