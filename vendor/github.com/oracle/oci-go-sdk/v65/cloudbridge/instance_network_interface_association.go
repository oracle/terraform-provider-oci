// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstanceNetworkInterfaceAssociation Describes association information for an Elastic IP address (IPv4).
type InstanceNetworkInterfaceAssociation struct {

	// The carrier IP address associated with the network interface.
	CarrierIp *string `mandatory:"false" json:"carrierIp"`

	// The customer-owned IP address associated with the network interface.
	CustomerOwnedIp *string `mandatory:"false" json:"customerOwnedIp"`

	// The ID of the owner of the Elastic IP address.
	IpOwnerKey *string `mandatory:"false" json:"ipOwnerKey"`

	// The public DNS name.
	PublicDnsName *string `mandatory:"false" json:"publicDnsName"`

	// The public IP address or Elastic IP address bound to the network interface.
	PublicIp *string `mandatory:"false" json:"publicIp"`
}

func (m InstanceNetworkInterfaceAssociation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceNetworkInterfaceAssociation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
