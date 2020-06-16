// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// APIs for dynamically scaling Compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.cloud.oracle.com/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Overview of the Compute Service (https://docs.cloud.oracle.com/Content/Compute/Concepts/computeoverview.htm).
// **Note:** Autoscaling is not available in US Government Cloud tenancies. For more information, see
// Oracle Cloud Infrastructure US Government Cloud (https://docs.cloud.oracle.com/Content/General/Concepts/govoverview.htm).
//

package autoscaling

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ChangeAutoScalingCompartmentDetails The configuration details for the move operation.
type ChangeAutoScalingCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to move the autoscaling configuration to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeAutoScalingCompartmentDetails) String() string {
	return common.PointerString(m)
}
