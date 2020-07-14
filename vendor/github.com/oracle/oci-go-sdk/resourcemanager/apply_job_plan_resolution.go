// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service.
// Use this API to install, configure, and manage resources via the "infrastructure-as-code" model.
// For more information, see
// Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ApplyJobPlanResolution Deprecated. Use the property `executionPlanStrategy` in `jobOperationDetails` instead.
type ApplyJobPlanResolution struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that specifies the most recently executed plan job.
	PlanJobId *string `mandatory:"false" json:"planJobId"`

	// Specifies whether to use the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the most recently run plan job.
	// `True` if using the latest job OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be a plan job that completed successfully.
	IsUseLatestJobId *bool `mandatory:"false" json:"isUseLatestJobId"`

	// Specifies whether to use the configuration directly, without reference to a Plan job.
	// `True` if using the configuration directly. Note that it is not necessary
	// for a Plan job to have run successfully.
	IsAutoApproved *bool `mandatory:"false" json:"isAutoApproved"`
}

func (m ApplyJobPlanResolution) String() string {
	return common.PointerString(m)
}
