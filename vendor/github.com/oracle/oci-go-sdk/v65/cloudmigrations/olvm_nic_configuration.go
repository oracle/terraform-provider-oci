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

// OlvmNicConfiguration The type describes the configuration of a virtual network interface.
type OlvmNicConfiguration struct {
	BootProtocol *OlvmBootProtocol `mandatory:"false" json:"bootProtocol"`

	Ip *OlvmIp `mandatory:"false" json:"ip"`

	Ipv6 *OlvmIp `mandatory:"false" json:"ipv6"`

	Ipv6BootProtocol *OlvmBootProtocol `mandatory:"false" json:"ipv6BootProtocol"`

	// Network interface name.
	Name *string `mandatory:"false" json:"name"`

	// Specifies whether the network interface should be activated on the virtual machine guest operating system boot.
	IsOnBoot *bool `mandatory:"false" json:"isOnBoot"`
}

func (m OlvmNicConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmNicConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
