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

// CreateInstanceAgentCommandDetails Creation details for an Oracle Cloud Agent command.
type CreateInstanceAgentCommandDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment to create the command in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The amount of time that Oracle Cloud Agent is given to run the command on the instance before timing
	// out. The timer starts when Oracle Cloud Agent starts the command. Zero means no timeout.
	ExecutionTimeOutInSeconds *int `mandatory:"true" json:"executionTimeOutInSeconds"`

	// The target instance to run the command on.
	Target *InstanceAgentCommandTarget `mandatory:"true" json:"target"`

	// The contents of the command.
	Content *InstanceAgentCommandContent `mandatory:"true" json:"content"`

	// A user-friendly name for the command. It does not have to be unique.
	// Avoid entering confidential information.
	// Example: `Database Backup Script`
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m CreateInstanceAgentCommandDetails) String() string {
	return common.PointerString(m)
}
