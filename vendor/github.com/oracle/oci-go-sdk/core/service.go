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

// Service Information of a particular Service that can be exposed through Service
// Gateway.
type Service struct {

	// This value will be used as Destination CidrBlock while creating a route rule with service gateway as target.
	CidrBlock *string `mandatory:"true" json:"cidrBlock"`

	// Description of this particular Service, provided by the Service owner.
	Description *string `mandatory:"true" json:"description"`

	// The Service's Oracle ID ([OCID])(/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"true" json:"id"`

	// Name of the Service.
	Name *string `mandatory:"true" json:"name"`
}

func (m Service) String() string {
	return common.PointerString(m)
}
