// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InternalServiceVnicAdConfig Contains ad config for service vnics
type InternalServiceVnicAdConfig struct {

	// Name of the ad
	// Example: `ad1`
	AdName *string `mandatory:"true" json:"adName"`

	// This ip address (anycast ip) is used as destination address for routing packets from vnics attached to caviums, and elpaso.
	IpAddressForClient *string `mandatory:"true" json:"ipAddressForClient"`

	// This ip address (anycast ip) is used as destination address for routing packets to NLB service vnics.
	IpAddressForNlbService *string `mandatory:"true" json:"ipAddressForNlbService"`

	// List of site group ids mapped to this shard
	SiteGroupIds []string `mandatory:"false" json:"siteGroupIds"`

	// This ip address (anycast ip) is used as destination address for routing packets from PA DP.
	IpAddressForService *string `mandatory:"false" json:"ipAddressForService"`

	// This ip address (anycast ip) is used as destination address for routing packets to micro-vnics.
	IpAddressForMicroVnic *string `mandatory:"false" json:"ipAddressForMicroVnic"`
}

func (m InternalServiceVnicAdConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalServiceVnicAdConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
