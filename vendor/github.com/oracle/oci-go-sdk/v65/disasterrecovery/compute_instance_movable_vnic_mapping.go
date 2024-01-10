// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Full Stack Disaster Recovery API
//
// Use the Full Stack Disaster Recovery (DR) API to manage disaster recovery for business applications.
// Full Stack DR is an OCI disaster recovery orchestration and management service that provides comprehensive disaster
// recovery capabilities for all layers of an application stack, including infrastructure, middleware, database,
// and application.
//

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeInstanceMovableVnicMapping Source VNIC to destination subnet mapping for a movable compute instance.
type ComputeInstanceMovableVnicMapping struct {

	// The OCID of the source VNIC.
	// Example: `ocid1.vnic.oc1..uniqueID`
	SourceVnicId *string `mandatory:"true" json:"sourceVnicId"`

	// The OCID of the destination subnet to which the source VNIC should connect.
	// Example: `ocid1.subnet.oc1..uniqueID`
	DestinationSubnetId *string `mandatory:"true" json:"destinationSubnetId"`

	// The private IP address to be assigned as the VNIC's primary IP address in the destination subnet.
	// This must be a valid IP address in the destination subnet and the IP address must be available.
	// Example: `10.0.3.3`
	DestinationPrimaryPrivateIpAddress *string `mandatory:"false" json:"destinationPrimaryPrivateIpAddress"`

	// The hostname label to be assigned in the destination subnet for the primary private IP of the source VNIC.
	// This label is the hostname portion of the private IP's fully qualified domain name (FQDN)
	// (for example, 'myhost1' in the FQDN 'myhost1.subnet123.vcn1.oraclevcn.com').
	// Example: `myhost1`
	DestinationPrimaryPrivateIpHostnameLabel *string `mandatory:"false" json:"destinationPrimaryPrivateIpHostnameLabel"`

	// A list of OCIDs of network security groups (NSG) in the destination region which should be assigned to
	// the source VNIC.
	// Example: `[ ocid1.networksecuritygroup.oc1..uniqueID, ocid1.networksecuritygroup.oc1..uniqueID ]`
	DestinationNsgIdList []string `mandatory:"false" json:"destinationNsgIdList"`
}

func (m ComputeInstanceMovableVnicMapping) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeInstanceMovableVnicMapping) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
