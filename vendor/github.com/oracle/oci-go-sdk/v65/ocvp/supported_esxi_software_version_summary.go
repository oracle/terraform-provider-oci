// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use the Oracle Cloud VMware API to create SDDCs and manage ESXi hosts and software.
// For more information, see Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SupportedEsxiSoftwareVersionSummary A specific version of bundled ESXi software supported by the Oracle Cloud
// VMware Solution.
type SupportedEsxiSoftwareVersionSummary struct {

	// A short, unique string that identifies the version of bundled software.
	Version *string `mandatory:"true" json:"version"`

	// A description of the software in the bundle.
	Description *string `mandatory:"true" json:"description"`

	// A list of ESXi host shapes supported by the version of bundled software.
	SupportedHostShapeNames []string `mandatory:"true" json:"supportedHostShapeNames"`
}

func (m SupportedEsxiSoftwareVersionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SupportedEsxiSoftwareVersionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
