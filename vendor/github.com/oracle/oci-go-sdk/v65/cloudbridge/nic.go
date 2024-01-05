// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Nic The VNIC configuration.
type Nic struct {

	// Provides a label and summary information for the device.
	Label *string `mandatory:"false" json:"label"`

	// Switch name.
	SwitchName *string `mandatory:"false" json:"switchName"`

	// Mac address of the VM.
	MacAddress *string `mandatory:"false" json:"macAddress"`

	// Mac address type.
	MacAddressType *string `mandatory:"false" json:"macAddressType"`

	// Network name.
	NetworkName *string `mandatory:"false" json:"networkName"`

	// List of IP addresses.
	IpAddresses []string `mandatory:"false" json:"ipAddresses"`
}

func (m Nic) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Nic) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
