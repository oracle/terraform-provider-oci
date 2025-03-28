// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrProtectionGroupMemberComputeInstance Deprecated. Properties for a compute instance member of a DR protection group.
type DrProtectionGroupMemberComputeInstance struct {

	// The OCID of the member.
	// Example: `ocid1.instance.oc1..uniqueID`
	MemberId *string `mandatory:"true" json:"memberId"`

	// A flag indicating if the compute instance should be moved during DR operations.
	// Example: `false`
	IsMovable *bool `mandatory:"false" json:"isMovable"`

	// A list of compute instance VNIC mappings.
	VnicMapping []ComputeInstanceVnicMapping `mandatory:"false" json:"vnicMapping"`

	// The OCID of a compartment in the destination region in which the compute instance
	// should be launched.
	// Example: `ocid1.compartment.oc1..uniqueID`
	DestinationCompartmentId *string `mandatory:"false" json:"destinationCompartmentId"`

	// The OCID of a dedicated VM host in the destination region where the compute instance
	// should be launched.
	// Example: `ocid1.dedicatedvmhost.oc1..uniqueID`
	DestinationDedicatedVmHostId *string `mandatory:"false" json:"destinationDedicatedVmHostId"`
}

// GetMemberId returns MemberId
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
