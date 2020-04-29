// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
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
