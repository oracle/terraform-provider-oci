// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use this API to manage the Oracle Cloud VMware Solution.
//

package ocvp

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SupportedVmwareSoftwareVersionCollection A list of the supported versions of bundled VMware software.
type SupportedVmwareSoftwareVersionCollection struct {

	// A list of the supported versions of bundled VMware software.
	Items []SupportedVmwareSoftwareVersionSummary `mandatory:"true" json:"items"`
}

func (m SupportedVmwareSoftwareVersionCollection) String() string {
	return common.PointerString(m)
}
