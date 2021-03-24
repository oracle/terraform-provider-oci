// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v37/common"
)

// EnrollmentStatusSummary The metadata associated with the enrollment status summary.
type EnrollmentStatusSummary struct {

	// The OCID of the enrollment status.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The enrollment status' current state.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The current Cloud Advisor enrollment status.
	Status OptimizerEnrollmentStatusEnum `mandatory:"true" json:"status"`

	// The reason for the enrollment status of the tenancy.
	StatusReason *string `mandatory:"false" json:"statusReason"`

	// The date and time the enrollment status was created, in the format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the enrollment status was last updated, in the format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m EnrollmentStatusSummary) String() string {
	return common.PointerString(m)
}
