// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service. Use this API to install, configure, and manage resources via the "infrastructure-as-code" model. For more information, see Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestSummary A summary of the status of a work request.
type WorkRequestSummary struct {

	// The asynchronous operation tracked by this work request.
	OperationType WorkRequestOperationTypeEnum `mandatory:"true" json:"operationType"`

	// The status of the specified work request.
	Status WorkRequestStatusEnum `mandatory:"true" json:"status"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) identifying this work request.
	Id *string `mandatory:"true" json:"id"`

	// Unique identifier (OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the compartment that contains the work request.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The resources affected by this work request.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// Percentage of the work request completed.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// When the work request was created.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// When the work request transitioned from ACCEPTED to IN_PROGRESS.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// When the work request reached a terminal state (FAILED or SUCCEEDED).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequestSummary) String() string {
	return common.PointerString(m)
}
