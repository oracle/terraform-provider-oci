// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secure Desktops API
//
// Create and manage cloud-hosted desktops which can be accessed from a web browser or installed client.
//

package desktops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DesktopPoolPrivateAccessDetails The details of the desktop's private access network connectivity that were used to create the pool.
type DesktopPoolPrivateAccessDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet in the customer VCN where the
	// connectivity will be established.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the customer VCN.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// A list of network security groups for the private access.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The IPv4 address from the provided OCI subnet which needs to be assigned to the VNIC. If not provided, it will
	// be auto-assigned with an available IPv4 address from the subnet.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// The three-label FQDN to use for the private endpoint. The customer VCN's DNS records are
	// updated with this FQDN. This enables the customer to use the FQDN instead of the private endpoint's
	// private IP address to access the service (for example, xyz.oraclecloud.com).
	EndpointFqdn *string `mandatory:"false" json:"endpointFqdn"`
}

func (m DesktopPoolPrivateAccessDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DesktopPoolPrivateAccessDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
