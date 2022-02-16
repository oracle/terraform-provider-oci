// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// This describes the network load balancer API.
//

package networkloadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

	// IP version associated with this IP address.
	IpVersion IpVersionEnum `mandatory:"false" json:"ipVersion,omitempty"`

	ReservedIp *ReservedIp `mandatory:"false" json:"reservedIp"`
}

func (m IpAddress) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IpAddress) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingIpVersionEnum(string(m.IpVersion)); !ok && m.IpVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IpVersion: %s. Supported values are: %s.", m.IpVersion, strings.Join(GetIpVersionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
