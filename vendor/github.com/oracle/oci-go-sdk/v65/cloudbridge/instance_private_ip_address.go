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

// InstancePrivateIpAddress Describes a private IPv4 address.
type InstancePrivateIpAddress struct {
	Association *InstanceNetworkInterfaceAssociation `mandatory:"false" json:"association"`

	// Indicates whether this IPv4 address is the primary private IP address of the network interface.
	IsPrimary *bool `mandatory:"false" json:"isPrimary"`

	// The private IPv4 DNS name.
	PrivateDnsName *string `mandatory:"false" json:"privateDnsName"`

	// The private IPv4 address of the network interface.
	PrivateIpAddress *string `mandatory:"false" json:"privateIpAddress"`
}

func (m InstancePrivateIpAddress) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstancePrivateIpAddress) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
