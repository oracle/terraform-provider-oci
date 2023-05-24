// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (FSDR) API to manage disaster recovery for business applications.
// FSDR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster recovery
// capabilities for all layers of an application stack, including infrastructure, middleware, database, and application.
//

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeInstanceMovableVnicMappingDetails A movable compute instance's source and destination VNIC mapping.
type ComputeInstanceMovableVnicMappingDetails struct {

	// The OCID of the VNIC.
	// Example: `ocid1.vnic.oc1..&lt;unique_id&gt;`
	SourceVnicId *string `mandatory:"true" json:"sourceVnicId"`

	// The OCID of the destination (remote) subnet to which this VNIC should connect.
	// Example: `ocid1.subnet.oc1..&lt;unique_id&gt;`
	DestinationSubnetId *string `mandatory:"true" json:"destinationSubnetId"`

	// The primary private IP address to assign. This address must belong to the destination subnet.
	// Example: `10.0.3.3`
	DestinationPrimaryPrivateIpAddress *string `mandatory:"false" json:"destinationPrimaryPrivateIpAddress"`

	// The hostname to assign for this primary private IP.
	// The value is the hostname portion of the private IP's fully qualified domain name (FQDN)
	// (for example, bminstance1 in FQDN bminstance1.subnet123.vcn1.oraclevcn.com).
	// Example: `bminstance1`
	DestinationPrimaryPrivateIpHostnameLabel *string `mandatory:"false" json:"destinationPrimaryPrivateIpHostnameLabel"`

	// A list of network security group (NSG) IDs in the destination region which this VNIC should use.
	// Example: `[ ocid1.networksecuritygroup.oc1..&lt;unique_id&gt;, ocid1.networksecuritygroup.oc1..&lt;unique_id&gt; ]`
	DestinationNsgIdList []string `mandatory:"false" json:"destinationNsgIdList"`
}

func (m ComputeInstanceMovableVnicMappingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeInstanceMovableVnicMappingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
