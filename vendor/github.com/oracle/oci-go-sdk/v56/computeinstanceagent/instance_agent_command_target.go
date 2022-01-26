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

// InstanceAgentCommandTarget The target instance that the command runs on.
type InstanceAgentCommandTarget struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the target instance.
	InstanceId *string `mandatory:"false" json:"instanceId"`
}

func (m InstanceAgentCommandTarget) String() string {
	return common.PointerString(m)
}
