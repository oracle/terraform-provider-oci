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

// DetachParentSoftwareSourceFromManagedInstanceDetails Information for detaching a software source from a managed instance
type DetachParentSoftwareSourceFromManagedInstanceDetails struct {

	// OCID for the Software Source
	SoftwareSourceId *string `mandatory:"true" json:"softwareSourceId"`
}

func (m DetachParentSoftwareSourceFromManagedInstanceDetails) String() string {
	return common.PointerString(m)
}
