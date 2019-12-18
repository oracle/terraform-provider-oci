// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeSoftwareSourceCompartmentDetails Compartment id for a software source
type ChangeSoftwareSourceCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the
	// compartment into which the resource should be moved.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m ChangeSoftwareSourceCompartmentDetails) String() string {
	return common.PointerString(m)
}
