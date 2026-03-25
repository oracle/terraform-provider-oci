// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PrivateAccessChannel Analytics instance private access channel model.
type PrivateAccessChannel struct {

	// Private access channel unique identifier key.
	Key *string `mandatory:"true" json:"key"`

	// Display name of the private access channel.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID of the customer VCN peered with the private access channel.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// OCID of the customer subnet connected to the private access channel.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// IP address of the private access channel.
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// List of IP addresses from the customer subnet connected to the private access channel, used as a source IP by the private access channel
	// for network traffic from the Analytics instance to the private sources.
	EgressSourceIpAddresses []string `mandatory:"true" json:"egressSourceIpAddresses"`

	// List of private source DNS zones registered with the private access channel. The
	//  datasource hostnames from these DNS zones / domains will be resolved in the peered VCN for access from  the Analytics instance.
	// Minimum 1 private source is required. Maximum 30 private source DNS zones can be registered.
	PrivateSourceDnsZones []PrivateSourceDnsZone `mandatory:"false" json:"privateSourceDnsZones"`

	// List of private source database SCAN hosts registered with the private access channel for access from the Analytics instance.
	PrivateSourceScanHosts []PrivateSourceScanHost `mandatory:"false" json:"privateSourceScanHosts"`

	// Network Security Group OCIDs for the Analytics instance.
	NetworkSecurityGroupIds []string `mandatory:"false" json:"networkSecurityGroupIds"`
}

func (m PrivateAccessChannel) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateAccessChannel) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
