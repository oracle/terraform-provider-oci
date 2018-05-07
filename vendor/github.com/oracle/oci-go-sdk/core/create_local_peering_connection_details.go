// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateLocalPeeringConnectionDetails The representation of CreateLocalPeeringConnectionDetails
type CreateLocalPeeringConnectionDetails struct {

	// The OCID of the compartment containing the local peering connection.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the VCN the local peering connection belongs to.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m CreateLocalPeeringConnectionDetails) String() string {
	return common.PointerString(m)
}
