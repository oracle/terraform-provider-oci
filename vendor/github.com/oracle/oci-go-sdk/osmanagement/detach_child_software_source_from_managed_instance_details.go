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

// DetachChildSoftwareSourceFromManagedInstanceDetails Information for detaching a software source from a managed instance
type DetachChildSoftwareSourceFromManagedInstanceDetails struct {

	// OCID for the Software Source
	SoftwareSourceId *string `mandatory:"true" json:"softwareSourceId"`
}

func (m DetachChildSoftwareSourceFromManagedInstanceDetails) String() string {
	return common.PointerString(m)
}
