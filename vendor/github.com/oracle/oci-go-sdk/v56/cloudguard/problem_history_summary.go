// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ProblemHistorySummary Problem History Definition.
type ProblemHistorySummary struct {

	// Unique identifier for the history record
	Id *string `mandatory:"true" json:"id"`

	// problemId for which history is associated to.
	ProblemId *string `mandatory:"true" json:"problemId"`

	// Actor type who performed the operation
	ActorType ActorTypeEnum `mandatory:"true" json:"actorType"`

	// Resource Name who performed activity
	ActorName *string `mandatory:"true" json:"actorName"`

	// Activity explanation details
	Explanation *string `mandatory:"true" json:"explanation"`

	// Problem Lifecycle Detail Status
	LifecycleDetail ProblemLifecycleDetailEnum `mandatory:"true" json:"lifecycleDetail"`

	// Type of the Entity
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Impacted Resource Names in a comma-separated string.
	Delta *string `mandatory:"true" json:"delta"`

	// Event status
	EventStatus EventStatusEnum `mandatory:"false" json:"eventStatus,omitempty"`

	// User Defined Comments
	Comment *string `mandatory:"false" json:"comment"`
}

func (m ProblemHistorySummary) String() string {
	return common.PointerString(m)
}
