// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// IpMaxConnections An object that species the maximum number of connections the listed IPs can make to a listener.
type IpMaxConnections struct {

	// Each element in the list should be valid IPv4 or IPv6 CIDR Block address.
	// Example: '["129.213.176.0/24", "150.136.187.0/24", "2002::1234:abcd:ffff:c0a8:101/64"]'
	IpAddresses []string `mandatory:"true" json:"ipAddresses"`

	// The maximum number of simultaneous connections that the specified IPs can make to the
	// Listener. IPs without a maxConnections setting can make either defaultMaxConnections
	// simultaneous connections to a listener or, if no defaultMaxConnections is specified, an
	// unlimited number of simultaneous connections to a listener.
	MaxConnections *int `mandatory:"true" json:"maxConnections"`
}

func (m IpMaxConnections) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IpMaxConnections) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
