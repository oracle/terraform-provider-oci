// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Content and Experience API
//
// Oracle Content and Experience is a cloud-based content hub to drive omni-channel content management and accelerate experience delivery
//

package oce

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeOceInstanceCompartmentDetails The information about compartment details.
type ChangeOceInstanceCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment
	// into which the OceInstance should be moved.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeOceInstanceCompartmentDetails) String() string {
	return common.PointerString(m)
}
