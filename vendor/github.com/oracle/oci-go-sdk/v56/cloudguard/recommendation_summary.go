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

// RecommendationSummary Recommendation Definition.
type RecommendationSummary struct {

	// Unique identifier for Recommendation
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// targetId associated with the problem
	TargetId *string `mandatory:"true" json:"targetId"`

	// Recommendation details
	Details map[string]string `mandatory:"true" json:"details"`

	// Count number of the problem
	ProblemCount *int64 `mandatory:"true" json:"problemCount"`

	// The current state of the Recommendation.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The lifecycleDetail will give more detail on the substate of the lifecycleState.
	LifecycleDetail RecommendationLifecycleDetailEnum `mandatory:"true" json:"lifecycleDetail"`

	// recommendation string showing on UX
	Name *string `mandatory:"true" json:"name"`

	// description of the recommendation
	Description *string `mandatory:"true" json:"description"`

	// Recommendation type
	Type RecommendationTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Tenant Identifier
	TenantId *string `mandatory:"false" json:"tenantId"`

	// The Risk Level
	RiskLevel RiskLevelEnum `mandatory:"false" json:"riskLevel,omitempty"`

	// problem creating time
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// problem updating time
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m RecommendationSummary) String() string {
	return common.PointerString(m)
}
