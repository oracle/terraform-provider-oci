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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// InstanceAgentCommandExecutionOutputViaObjectStorageUriDetails The execution output from a command when saved to an Object Storage URL.
type InstanceAgentCommandExecutionOutputViaObjectStorageUriDetails struct {

	// The exit code for the command. Exit code `0` indicates success.
	ExitCode *int `mandatory:"true" json:"exitCode"`

	// The Object Storage URL or pre-authenticated request (PAR) for the command output.
	OutputUri *string `mandatory:"true" json:"outputUri"`

	// An optional status message that Oracle Cloud Agent can populate for additional troubleshooting.
	Message *string `mandatory:"false" json:"message"`
}

//GetExitCode returns ExitCode
func (m InstanceAgentCommandExecutionOutputViaObjectStorageUriDetails) GetExitCode() *int {
	return m.ExitCode
}

//GetMessage returns Message
func (m InstanceAgentCommandExecutionOutputViaObjectStorageUriDetails) GetMessage() *string {
	return m.Message
}

func (m InstanceAgentCommandExecutionOutputViaObjectStorageUriDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceAgentCommandExecutionOutputViaObjectStorageUriDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InstanceAgentCommandExecutionOutputViaObjectStorageUriDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceAgentCommandExecutionOutputViaObjectStorageUriDetails InstanceAgentCommandExecutionOutputViaObjectStorageUriDetails
	s := struct {
		DiscriminatorParam string `json:"outputType"`
		MarshalTypeInstanceAgentCommandExecutionOutputViaObjectStorageUriDetails
	}{
		"OBJECT_STORAGE_URI",
		(MarshalTypeInstanceAgentCommandExecutionOutputViaObjectStorageUriDetails)(m),
	}

	return json.Marshal(&s)
}
