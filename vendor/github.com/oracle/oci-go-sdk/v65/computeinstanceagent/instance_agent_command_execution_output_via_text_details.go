// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstanceAgentCommandExecutionOutputViaTextDetails The execution output from a command when returned in plain text.
type InstanceAgentCommandExecutionOutputViaTextDetails struct {

	// The exit code for the command. Exit code `0` indicates success.
	ExitCode *int `mandatory:"true" json:"exitCode"`

	// An optional status message that Oracle Cloud Agent can populate for additional troubleshooting.
	Message *string `mandatory:"false" json:"message"`

	// The command output.
	Text *string `mandatory:"false" json:"text"`

	// SHA-256 checksum value of the text content.
	TextSha256 *string `mandatory:"false" json:"textSha256"`
}

// GetExitCode returns ExitCode
func (m InstanceAgentCommandExecutionOutputViaTextDetails) GetExitCode() *int {
	return m.ExitCode
}

// GetMessage returns Message
func (m InstanceAgentCommandExecutionOutputViaTextDetails) GetMessage() *string {
	return m.Message
}

func (m InstanceAgentCommandExecutionOutputViaTextDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceAgentCommandExecutionOutputViaTextDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m InstanceAgentCommandExecutionOutputViaTextDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceAgentCommandExecutionOutputViaTextDetails InstanceAgentCommandExecutionOutputViaTextDetails
	s := struct {
		DiscriminatorParam string `json:"outputType"`
		MarshalTypeInstanceAgentCommandExecutionOutputViaTextDetails
	}{
		"TEXT",
		(MarshalTypeInstanceAgentCommandExecutionOutputViaTextDetails)(m),
	}

	return json.Marshal(&s)
}
