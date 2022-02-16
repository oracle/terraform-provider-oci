// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// PrivateAccessChannel Analytics Instance Private Access Channel model.
type PrivateAccessChannel struct {

	// Private Access Channel unique identifier key.
	Key *string `mandatory:"true" json:"key"`

	// Display Name of the Private Access Channel.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID of the customer VCN peered with private access channel.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// OCID of the customer subnet connected to private access channel.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// IP Address of the Private Access channel.
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// The list of IP addresses from the customer subnet connected to private access channel, used as a source Ip by Private Access Channel
	// for network traffic from the AnalyticsInstance to Private Sources.
	EgressSourceIpAddresses []string `mandatory:"true" json:"egressSourceIpAddresses"`

	// List of Private Source DNS zones registered with Private Access Channel,
	// where datasource hostnames from these dns zones / domains will be resolved in the peered VCN for access from Analytics Instance.
	// Min of 1 is required and Max of 30 Private Source DNS zones can be registered.
	PrivateSourceDnsZones []PrivateSourceDnsZone `mandatory:"false" json:"privateSourceDnsZones"`
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
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
