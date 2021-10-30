// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// ChangeDomainCompartmentDetails The representation of ChangeDomainCompartmentDetails
type ChangeDomainCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the destination compartment
	// into which to move the domain.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeDomainCompartmentDetails) String() string {
	return common.PointerString(m)
}
