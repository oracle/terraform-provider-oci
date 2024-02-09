// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.cloud.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.cloud.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RecommendationSummary Recommendation definition.
type RecommendationSummary struct {

	// Unique identifier for the recommendation
	Id *string `mandatory:"true" json:"id"`

	// Compartment OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Target ID associated with the problem
	TargetId *string `mandatory:"true" json:"targetId"`

	// Recommendation details
	Details map[string]string `mandatory:"true" json:"details"`

	// Count number of the problem
	ProblemCount *int64 `mandatory:"true" json:"problemCount"`

	// The current lifecycle state of the recommendation
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Additional details on the substate of the lifecycle state
	LifecycleDetail RecommendationLifecycleDetailEnum `mandatory:"true" json:"lifecycleDetail"`

	// Recommendation string that appears in the UI for the problem
	Name *string `mandatory:"true" json:"name"`

	// Description of the recommendation
	Description *string `mandatory:"true" json:"description"`

	// Recommendation type
	Type RecommendationTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Tenant identifier
	TenantId *string `mandatory:"false" json:"tenantId"`

	// The risk level of the problem
	RiskLevel RiskLevelEnum `mandatory:"false" json:"riskLevel,omitempty"`

	// The date and time the problem was first created
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the problem was last updated
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m RecommendationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecommendationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRecommendationLifecycleDetailEnum(string(m.LifecycleDetail)); !ok && m.LifecycleDetail != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetail: %s. Supported values are: %s.", m.LifecycleDetail, strings.Join(GetRecommendationLifecycleDetailEnumStringValues(), ",")))
	}

	if _, ok := GetMappingRecommendationTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetRecommendationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRiskLevelEnum(string(m.RiskLevel)); !ok && m.RiskLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RiskLevel: %s. Supported values are: %s.", m.RiskLevel, strings.Join(GetRiskLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
