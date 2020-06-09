// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
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
