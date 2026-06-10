// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Ipv6AddressIpv6SubnetCidrPairDetails Details to assign an IPv6 subnet prefix or IPv6 address to a resource.
type Ipv6AddressIpv6SubnetCidrPairDetails struct {

	// The IPv6 prefix allocated to the subnet.
	Ipv6SubnetCidr *string `mandatory:"false" json:"ipv6SubnetCidr"`

	// An IPv6 address of your choice. Must be an available IPv6 address within the subnet's prefix.
	Ipv6Address *string `mandatory:"false" json:"ipv6Address"`
}

func (m Ipv6AddressIpv6SubnetCidrPairDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Ipv6AddressIpv6SubnetCidrPairDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
