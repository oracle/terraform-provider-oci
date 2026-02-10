// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// BulkDeletePrivateIpsDetails Bulk Secondary IPv4 deletion object.
type BulkDeletePrivateIpsDetails struct {

	// Secondary IPv4 addresses to deleted
	BulkDeletePrivateIpItem []BulkDeletePrivateIpItem `mandatory:"true" json:"bulkDeletePrivateIpItem"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC of which
	// private IPs should be deleted. The VNIC and private IPs must be in the same subnet.
	VnicId *string `mandatory:"false" json:"vnicId"`

	// Use this attribute only with the Oracle Cloud VMware Solution.
	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN from which the private IP is to be deleted.
	VlanId *string `mandatory:"false" json:"vlanId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet from which the private IPs is to be deleted.
	SubnetId *string `mandatory:"false" json:"subnetId"`
}

func (m BulkDeletePrivateIpsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkDeletePrivateIpsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
