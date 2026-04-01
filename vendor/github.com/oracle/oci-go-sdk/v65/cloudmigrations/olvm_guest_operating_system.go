// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmGuestOperatingSystem Represents an operating system installed on the virtual machine.
type OlvmGuestOperatingSystem struct {

	// The architecture of the operating system, such as x86_64.
	Architecture *string `mandatory:"false" json:"architecture"`

	// Code name of the operating system, such as Maipo.
	Codename *string `mandatory:"false" json:"codename"`

	// Full name of operating system distribution.
	Distribution *string `mandatory:"false" json:"distribution"`

	// Family of operating system, such as Linux.
	Family *string `mandatory:"false" json:"family"`

	Kernel *OlvmKernel `mandatory:"false" json:"kernel"`

	Version *OlvmVersion `mandatory:"false" json:"version"`
}

func (m OlvmGuestOperatingSystem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmGuestOperatingSystem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
