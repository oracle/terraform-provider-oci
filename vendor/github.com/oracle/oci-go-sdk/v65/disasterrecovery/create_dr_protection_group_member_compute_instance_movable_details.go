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

// CreateDrProtectionGroupMemberComputeInstanceMovableDetails Create properties for a movable compute instance member.
type CreateDrProtectionGroupMemberComputeInstanceMovableDetails struct {

	// The OCID of the member.
	// Example: `ocid1.instance.oc1..uniqueID`
	MemberId *string `mandatory:"true" json:"memberId"`

	// A flag indicating if the compute instance should be moved to the same fault domain in the destination region.
	// The compute instance launch will fail if this flag is set to true and capacity is not available in the
	// specified fault domain in the destination region.
	// Example: `false`
	IsRetainFaultDomain *bool `mandatory:"false" json:"isRetainFaultDomain"`

	// The OCID of a capacity reservation in the destination region which will be used to launch
	// the compute instance.
	// Example: `ocid1.capacityreservation.oc1..uniqueID`
	DestinationCapacityReservationId *string `mandatory:"false" json:"destinationCapacityReservationId"`

	// A list of compute instance VNIC mappings.
	VnicMappings []ComputeInstanceMovableVnicMappingDetails `mandatory:"false" json:"vnicMappings"`

	// The OCID of a compartment in the destination region in which the compute instance
	// should be launched.
	// Example: `ocid1.compartment.oc1..uniqueID`
	DestinationCompartmentId *string `mandatory:"false" json:"destinationCompartmentId"`

	// The OCID of a dedicated VM host in the destination region where the compute instance
	// should be launched.
	// Example: `ocid1.dedicatedvmhost.oc1..uniqueID`
	DestinationDedicatedVmHostId *string `mandatory:"false" json:"destinationDedicatedVmHostId"`

	// A list of operations performed on file systems used by the compute instance.
	FileSystemOperations []CreateComputeInstanceMovableFileSystemOperationDetails `mandatory:"false" json:"fileSystemOperations"`
}

// GetMemberId returns MemberId
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
