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

// InstanceAgentCommandSourceViaObjectStorageTupleDetails The source of the command when imported from an Object Storage bucket.
type InstanceAgentCommandSourceViaObjectStorageTupleDetails struct {

	// The Object Storage bucket for the command.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The Object Storage namespace for the command.
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	// The Object Storage object name for the command.
	ObjectName *string `mandatory:"true" json:"objectName"`
}

func (m InstanceAgentCommandSourceViaObjectStorageTupleDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstanceAgentCommandSourceViaObjectStorageTupleDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceAgentCommandSourceViaObjectStorageTupleDetails InstanceAgentCommandSourceViaObjectStorageTupleDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeInstanceAgentCommandSourceViaObjectStorageTupleDetails
	}{
		"OBJECT_STORAGE_TUPLE",
		(MarshalTypeInstanceAgentCommandSourceViaObjectStorageTupleDetails)(m),
	}

	return json.Marshal(&s)
}
