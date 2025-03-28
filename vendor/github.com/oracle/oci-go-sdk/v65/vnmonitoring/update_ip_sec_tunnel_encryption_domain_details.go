// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateIpSecTunnelEncryptionDomainDetails Request to update a multi-encryption domain policy on the IPSec tunnel.
// There can't be more than 50 security associations in use at one time. See Encryption domain for policy-based
// tunnels (https://docs.oracle.com/iaas/Content/Network/Tasks/ipsecencryptiondomains.htm#spi_policy_based_tunnel) for more.
type UpdateIpSecTunnelEncryptionDomainDetails struct {

	// Lists IPv4 or IPv6-enabled subnets in your Oracle tenancy.
	OracleTrafficSelector []string `mandatory:"false" json:"oracleTrafficSelector"`

	// Lists IPv4 or IPv6-enabled subnets in your on-premises network.
	CpeTrafficSelector []string `mandatory:"false" json:"cpeTrafficSelector"`
}

func (m UpdateIpSecTunnelEncryptionDomainDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateIpSecTunnelEncryptionDomainDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
