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

// InstanceAgentCommandExecutionOutputViaObjectStorageTupleDetails The execution output from a command when saved to an Object Storage bucket.
type InstanceAgentCommandExecutionOutputViaObjectStorageTupleDetails struct {

	// The exit code for the command. Exit code `0` indicates success.
	ExitCode *int `mandatory:"true" json:"exitCode"`

	// The Object Storage bucket for the command output.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The Object Storage namespace for the command output.
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	// The Object Storage object name for the command output.
	ObjectName *string `mandatory:"true" json:"objectName"`

	// An optional status message that Oracle Cloud Agent can populate for additional troubleshooting.
	Message *string `mandatory:"false" json:"message"`
}

//GetExitCode returns ExitCode
func (m InstanceAgentCommandExecutionOutputViaObjectStorageTupleDetails) GetExitCode() *int {
	return m.ExitCode
}

//GetMessage returns Message
func (m InstanceAgentCommandExecutionOutputViaObjectStorageTupleDetails) GetMessage() *string {
	return m.Message
}

func (m InstanceAgentCommandExecutionOutputViaObjectStorageTupleDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstanceAgentCommandExecutionOutputViaObjectStorageTupleDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceAgentCommandExecutionOutputViaObjectStorageTupleDetails InstanceAgentCommandExecutionOutputViaObjectStorageTupleDetails
	s := struct {
		DiscriminatorParam string `json:"outputType"`
		MarshalTypeInstanceAgentCommandExecutionOutputViaObjectStorageTupleDetails
	}{
		"OBJECT_STORAGE_TUPLE",
		(MarshalTypeInstanceAgentCommandExecutionOutputViaObjectStorageTupleDetails)(m),
	}

	return json.Marshal(&s)
}
