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

// AddPackagesToSoftwareSourceDetails List of software package names
type AddPackagesToSoftwareSourceDetails struct {

	// the list of package names
	PackageNames []string `mandatory:"true" json:"packageNames"`
}

func (m AddPackagesToSoftwareSourceDetails) String() string {
	return common.PointerString(m)
}
