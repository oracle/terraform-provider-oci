// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use this API to manage your Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"github.com/oracle/oci-go-sdk/v37/common"
)

// SupportedVmwareSoftwareVersionSummary A specific version of bundled VMware software supported by the Oracle Cloud
// VMware Solution.
type SupportedVmwareSoftwareVersionSummary struct {

	// A short, unique string that identifies the version of bundled software.
	Version *string `mandatory:"true" json:"version"`

	// A description of the software in the bundle.
	Description *string `mandatory:"true" json:"description"`
}

func (m SupportedVmwareSoftwareVersionSummary) String() string {
	return common.PointerString(m)
}
