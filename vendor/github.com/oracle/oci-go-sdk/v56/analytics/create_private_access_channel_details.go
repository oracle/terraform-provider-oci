// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Analytics API
//
// Analytics API.
//

package analytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreatePrivateAccessChannelDetails Input payload to create a Private Access Channel.
type CreatePrivateAccessChannelDetails struct {

	// Display Name of the Private Access Channel.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID of the customer VCN peered with private access channel.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// OCID of the customer subnet connected to private access channel.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// List of Private Source DNS zones registered with Private Access Channel,
	// where datasource hostnames from these dns zones / domains will be resolved in the peered VCN for access from Analytics Instance.
	// Min of 1 is required and Max of 30 Private Source DNS zones can be registered.
	PrivateSourceDnsZones []PrivateSourceDnsZone `mandatory:"true" json:"privateSourceDnsZones"`
}

func (m CreatePrivateAccessChannelDetails) String() string {
	return common.PointerString(m)
}
