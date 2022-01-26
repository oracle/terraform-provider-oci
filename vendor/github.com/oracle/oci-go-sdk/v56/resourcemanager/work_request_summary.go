// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

	// The date and time when the work request was created.
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// The date and time when the work request transitioned from ACCEPTED to IN_PROGRESS.
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time when the work request reached a terminal state (FAILED or SUCCEEDED).
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequestSummary) String() string {
	return common.PointerString(m)
}
