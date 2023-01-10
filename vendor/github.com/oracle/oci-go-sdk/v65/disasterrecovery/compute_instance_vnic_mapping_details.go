// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// ComputeInstanceVnicMappingDetails A compute instance's source and destination VNIC mapping.
type ComputeInstanceVnicMappingDetails struct {

	// The OCID of the VNIC.
	// Example: `ocid1.vnic.oc1.phx.exampleocid1`
	SourceVnicId *string `mandatory:"true" json:"sourceVnicId"`

	// The OCID of the destination (remote) subnet to which this VNIC should connect.
	// Example: `ocid1.subnet.oc1.iad.exampleocid2`
	DestinationSubnetId *string `mandatory:"true" json:"destinationSubnetId"`

	// A list of destination region's network security group (NSG) Ids which this VNIC should use.
	// Example: `[ ocid1.networksecuritygroup.oc1.iad.abcd1, ocid1.networksecuritygroup.oc1.iad.wxyz2 ]`
	DestinationNsgIdList []string `mandatory:"false" json:"destinationNsgIdList"`
}

func (m ComputeInstanceVnicMappingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeInstanceVnicMappingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
