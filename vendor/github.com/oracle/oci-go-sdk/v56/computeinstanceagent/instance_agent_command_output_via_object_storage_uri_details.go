// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Agent API
//
// API for the Oracle Cloud Agent software running on compute instances. Oracle Cloud Agent
// is a lightweight process that monitors and manages compute instances.
//

package computeinstanceagent

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// InstanceAgentCommandOutputViaObjectStorageUriDetails The command output destination when saved to an Object Storage URL.
type InstanceAgentCommandOutputViaObjectStorageUriDetails struct {

	// The Object Storage URL or pre-authenticated request (PAR) for the command output.
	OutputUri *string `mandatory:"true" json:"outputUri"`
}

func (m InstanceAgentCommandOutputViaObjectStorageUriDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstanceAgentCommandOutputViaObjectStorageUriDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceAgentCommandOutputViaObjectStorageUriDetails InstanceAgentCommandOutputViaObjectStorageUriDetails
	s := struct {
		DiscriminatorParam string `json:"outputType"`
		MarshalTypeInstanceAgentCommandOutputViaObjectStorageUriDetails
	}{
		"OBJECT_STORAGE_URI",
		(MarshalTypeInstanceAgentCommandOutputViaObjectStorageUriDetails)(m),
	}

	return json.Marshal(&s)
}
