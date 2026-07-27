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

// CreateReadEndpointDetails Details required to create a Read Endpoint.
type CreateReadEndpointDetails struct {

	// Specifies if the read endpoint is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The IP address the read endpoint is configured to listen on.
	// A private IP address of your choice to assign to the read endpoint.
	// Must be an available IP address within the subnet's CIDR. If you don't specify a value,
	// Oracle automatically assigns a private IP address from the subnet. This should be a
	// "dotted-quad" style IPv4 address.
	ReadEndpointIpAddress *string `mandatory:"false" json:"readEndpointIpAddress"`

	ReadEndpointIpv6AddressIpv6SubnetCidrPairDetails *Ipv6AddressIpv6SubnetCidrPairDetails `mandatory:"false" json:"readEndpointIpv6AddressIpv6SubnetCidrPairDetails"`

	// The hostname for the read endpoint. Used for DNS.
	// The value is the hostname portion of the read endpoint's fully qualified domain name (FQDN).
	// (for example, "reader" in FQDN "reader.mysql.us-ashburn-1.oraclecloud.com").
	// Must be unique across all VNICs in the subnet and comply with RFC 952 and RFC 1123.
	ReadEndpointHostnameLabel *string `mandatory:"false" json:"readEndpointHostnameLabel"`

	// A list of IP addresses for read replicas (in a DB System) or DB Nodes (in shared-storage DB cluster) that are excluded from serving read requests.
	ExcludeIps []string `mandatory:"false" json:"excludeIps"`
}

func (m CreateReadEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateReadEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
