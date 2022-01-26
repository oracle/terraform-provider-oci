// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// UpdateResourceActionDetails The request object for updating the resource action details.
type UpdateResourceActionDetails struct {

	// The status of the resource action.
	Status StatusEnum `mandatory:"true" json:"status"`

	// The date and time the current status will change. The format is defined by RFC3339.
	// For example, "The current `postponed` status of the resource action will end and change to `pending` on this
	// date and time."
	TimeStatusEnd *common.SDKTime `mandatory:"false" json:"timeStatusEnd"`
}

func (m UpdateResourceActionDetails) String() string {
	return common.PointerString(m)
}
