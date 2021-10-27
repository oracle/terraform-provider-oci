// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// A description of the network load balancer API
//

package networkloadbalancer

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// IpAddress A load balancer IP address.
type IpAddress struct {

	// An IP address.
	// Example: `192.168.0.3`
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// Whether the IP address is public or private.
	// If "true", then the IP address is public and accessible from the internet.
	// If "false", then the IP address is private and accessible only from within the associated virtual cloud network.
	IsPublic *bool `mandatory:"false" json:"isPublic"`

	ReservedIp *ReservedIp `mandatory:"false" json:"reservedIp"`
}

func (m IpAddress) String() string {
	return common.PointerString(m)
}
