// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ChangeOnPremConnectorCompartmentDetails The details used to change the compartment of a on-premises connector.
type ChangeOnPremConnectorCompartmentDetails struct {

	// The OCID of the new compartment where you want to move the on-premises connector.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeOnPremConnectorCompartmentDetails) String() string {
	return common.PointerString(m)
}
