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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateDrProtectionGroupMemberComputeInstanceMovableDetails Create properties for a Movable Compute Instance member.
type CreateDrProtectionGroupMemberComputeInstanceMovableDetails struct {

	// The OCID of the member.
	// Example: `ocid1.instance.oc1.phx.exampleocid1`
	MemberId *string `mandatory:"true" json:"memberId"`

	// A flag indicating if this compute instance should be moved to the same fault domain.
	// Compute instance launch will fail if this flag is set to true and capacity is not available in that specific fault domain in the destination region.
	// Example: `false`
	IsRetainFaultDomain *bool `mandatory:"false" json:"isRetainFaultDomain"`

	// A flag indicating if the moved compute instance should retain the same private IP addresses.
	// Compute instance launch will fail if this flag is set to true and the same private IP addresses are not available in the destination region.
	// Example: `false`
	IsRetainPrivateIPs *bool `mandatory:"false" json:"isRetainPrivateIPs"`

	// A flag indicating if the moved compute instance should retain the same hostname labels.
	// Compute instance launch will fail if this flag is set to true and the same hostname labels are not available in the destination region.
	// Example: `false`
	IsRetainHostNameLabels *bool `mandatory:"false" json:"isRetainHostNameLabels"`

	// The OCID of the capacity reservation in the destination region using which this compute instance
	// should be launched.
	// Example: `ocid1.capacityreservation.oc1..&lt;unique_id&gt;`
	DestinationCapacityReservationId *string `mandatory:"false" json:"destinationCapacityReservationId"`

	// A list of Compute Instance VNIC mappings.
	VnicMappings []ComputeInstanceMovableVnicMappingDetails `mandatory:"false" json:"vnicMappings"`

	// The OCID of the compartment for this compute instance in the destination region.
	// Example: `ocid1.compartment.oc1..&lt;unique_id&gt;`
	DestinationCompartmentId *string `mandatory:"false" json:"destinationCompartmentId"`

	// The OCID of the dedicated VM Host in the destination region where this compute instance
	// should be launched
	// Example: `ocid1.dedicatedvmhost.oc1..&lt;unique_id&gt;`
	DestinationDedicatedVmHostId *string `mandatory:"false" json:"destinationDedicatedVmHostId"`
}

//GetMemberId returns MemberId
func (m CreateDrProtectionGroupMemberComputeInstanceMovableDetails) GetMemberId() *string {
	return m.MemberId
}

func (m CreateDrProtectionGroupMemberComputeInstanceMovableDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateDrProtectionGroupMemberComputeInstanceMovableDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateDrProtectionGroupMemberComputeInstanceMovableDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDrProtectionGroupMemberComputeInstanceMovableDetails CreateDrProtectionGroupMemberComputeInstanceMovableDetails
	s := struct {
		DiscriminatorParam string `json:"memberType"`
		MarshalTypeCreateDrProtectionGroupMemberComputeInstanceMovableDetails
	}{
		"COMPUTE_INSTANCE_MOVABLE",
		(MarshalTypeCreateDrProtectionGroupMemberComputeInstanceMovableDetails)(m),
	}

	return json.Marshal(&s)
}
