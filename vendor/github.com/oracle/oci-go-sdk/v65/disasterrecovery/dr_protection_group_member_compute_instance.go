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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrProtectionGroupMemberComputeInstance Properties for a Compute Instance member of a DR Protection Group.
type DrProtectionGroupMemberComputeInstance struct {

	// The OCID of the member.
	// Example: `ocid1.instance.oc1.phx.exampleocid1`
	MemberId *string `mandatory:"true" json:"memberId"`

	// A flag indicating if this compute instance should be moved during DR operations.
	// Example: `false`
	IsMovable *bool `mandatory:"false" json:"isMovable"`

	// A list of compute instance VNIC mappings.
	VnicMapping []ComputeInstanceVnicMapping `mandatory:"false" json:"vnicMapping"`

	// The OCID of the compartment for this compute instance in the destination region.
	// Example: `ocid1.compartment.oc1..exampleocid`
	DestinationCompartmentId *string `mandatory:"false" json:"destinationCompartmentId"`

	// The OCID of the dedicated VM Host for this compute instance in the destination region.
	// Example: `ocid1.dedicatedvmhost.oc1.iad.exampleocid`
	DestinationDedicatedVmHostId *string `mandatory:"false" json:"destinationDedicatedVmHostId"`
}

//GetMemberId returns MemberId
func (m DrProtectionGroupMemberComputeInstance) GetMemberId() *string {
	return m.MemberId
}

func (m DrProtectionGroupMemberComputeInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrProtectionGroupMemberComputeInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DrProtectionGroupMemberComputeInstance) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDrProtectionGroupMemberComputeInstance DrProtectionGroupMemberComputeInstance
	s := struct {
		DiscriminatorParam string `json:"memberType"`
		MarshalTypeDrProtectionGroupMemberComputeInstance
	}{
		"COMPUTE_INSTANCE",
		(MarshalTypeDrProtectionGroupMemberComputeInstance)(m),
	}

	return json.Marshal(&s)
}
