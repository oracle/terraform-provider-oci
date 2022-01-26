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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// InstanceAgentCommand The command payload.
type InstanceAgentCommand struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the command.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment containing the command.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The target instance that the command runs on.
	Target *InstanceAgentCommandTarget `mandatory:"true" json:"target"`

	// The contents of the command.
	Content *InstanceAgentCommandContent `mandatory:"true" json:"content"`

	// A user-friendly name. Does not have to be unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time the command was created, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the command was last updated, in the format defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Whether a request was made to cancel the command. Canceling a command is a best-effort attempt.
	IsCanceled *bool `mandatory:"false" json:"isCanceled"`

	// The amount of time that Oracle Cloud Agent is given to run the command on the instance before timing
	// out. The timer starts when Oracle Cloud Agent starts the command. Zero means no timeout.
	ExecutionTimeOutInSeconds *int `mandatory:"false" json:"executionTimeOutInSeconds"`
}

func (m InstanceAgentCommand) String() string {
	return common.PointerString(m)
}
