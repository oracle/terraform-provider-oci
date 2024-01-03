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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrProtectionGroupMemberComputeInstanceNonMovable Properties for a non-movable compute instance member of a DR protection group.
type DrProtectionGroupMemberComputeInstanceNonMovable struct {

	// The OCID of the member.
	// Example: `ocid1.instance.oc1..uniqueID`
	MemberId *string `mandatory:"true" json:"memberId"`

	// A flag indicating whether the non-movable compute instance needs to be started and stopped during DR operations.
	IsStartStopEnabled *bool `mandatory:"false" json:"isStartStopEnabled"`

	// Operations performed on a list of file systems used on the non-movable compute instance.
	FileSystemOperations []ComputeInstanceNonMovableFileSystemOperation `mandatory:"false" json:"fileSystemOperations"`

	// Operations performed on a list of block volumes used on the non-movable compute instance.
	BlockVolumeOperations []ComputeInstanceNonMovableBlockVolumeOperation `mandatory:"false" json:"blockVolumeOperations"`
}

// GetMemberId returns MemberId
func (m DrProtectionGroupMemberComputeInstanceNonMovable) GetMemberId() *string {
	return m.MemberId
}

func (m DrProtectionGroupMemberComputeInstanceNonMovable) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DrProtectionGroupMemberComputeInstanceNonMovable) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DrProtectionGroupMemberComputeInstanceNonMovable) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDrProtectionGroupMemberComputeInstanceNonMovable DrProtectionGroupMemberComputeInstanceNonMovable
	s := struct {
		DiscriminatorParam string `json:"memberType"`
		MarshalTypeDrProtectionGroupMemberComputeInstanceNonMovable
	}{
		"COMPUTE_INSTANCE_NON_MOVABLE",
		(MarshalTypeDrProtectionGroupMemberComputeInstanceNonMovable)(m),
	}

	return json.Marshal(&s)
}
