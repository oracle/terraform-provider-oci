// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Whitelist An array of IP addresses that bypass the Web Application Firewall. Supports both single IP addresses or subnet masks (CIDR notation).
type Whitelist struct {

	// The unique name of the whitelist.
	Name *string `mandatory:"true" json:"name"`

	// A set of IP addresses or CIDR notations to include in the whitelist.
	Addresses []string `mandatory:"false" json:"addresses"`

	// A list of OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of IP address lists to include in the whitelist.
	AddressLists []string `mandatory:"false" json:"addressLists"`
}

func (m Whitelist) String() string {
	return common.PointerString(m)
}
