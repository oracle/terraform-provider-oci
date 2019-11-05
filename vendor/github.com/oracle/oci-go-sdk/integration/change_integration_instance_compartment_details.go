// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeIntegrationInstanceCompartmentDetails The information to be updated.
type ChangeIntegrationInstanceCompartmentDetails struct {

	// Compartment Identifier.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m ChangeIntegrationInstanceCompartmentDetails) String() string {
	return common.PointerString(m)
}
