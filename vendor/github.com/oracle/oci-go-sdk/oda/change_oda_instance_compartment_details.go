// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Digital Assistant Control Plane API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeOdaInstanceCompartmentDetails Properties required to move a Digital Assistant instance from one compartment to another.
type ChangeOdaInstanceCompartmentDetails struct {

	// Identifier of the compartment into which the Digital Assistant instance should be moved.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeOdaInstanceCompartmentDetails) String() string {
	return common.PointerString(m)
}
