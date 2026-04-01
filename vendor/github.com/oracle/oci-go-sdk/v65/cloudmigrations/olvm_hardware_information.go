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

// OlvmHardwareInformation Represents hardware information of host
type OlvmHardwareInformation struct {

	// Type of host???s CPU.
	Family *string `mandatory:"false" json:"family"`

	// Manufacturer of the host???s machine and hardware vendor.
	Manufacturer *string `mandatory:"false" json:"manufacturer"`

	// Host???s product name (for example RHEV Hypervisor).
	ProductName *string `mandatory:"false" json:"productName"`

	// Unique ID for host???s chassis.
	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	// Supported sources of random number generator.
	SupportedRngSources []RngSourceEnum `mandatory:"false" json:"supportedRngSources"`

	// Unique ID for each host.
	Uuid *string `mandatory:"false" json:"uuid"`

	// Unique name for each of the manufacturer.
	Version *string `mandatory:"false" json:"version"`
}

func (m OlvmHardwareInformation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmHardwareInformation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
