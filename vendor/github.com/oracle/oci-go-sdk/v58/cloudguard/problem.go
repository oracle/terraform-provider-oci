// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Problem Problem Definition.
type Problem struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier where the resource is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Identifier of the rule
	DetectorRuleId *string `mandatory:"false" json:"detectorRuleId"`

	// DEPRECATED
	Region *string `mandatory:"false" json:"region"`

	// Regions where the problem is found
	Regions []string `mandatory:"false" json:"regions"`

	// The Risk Level
	RiskLevel RiskLevelEnum `mandatory:"false" json:"riskLevel,omitempty"`

	// Identifier of the Resource
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// DisplayName of the Resource
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// Type of the Resource
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// user defined labels on the problem
	Labels []string `mandatory:"false" json:"labels"`

	// The date and time the problem was last detected. Format defined by RFC3339.
	TimeLastDetected *common.SDKTime `mandatory:"false" json:"timeLastDetected"`

	// The date and time the problem was first detected. Format defined by RFC3339.
	TimeFirstDetected *common.SDKTime `mandatory:"false" json:"timeFirstDetected"`

	// The current state of the Problem.
	LifecycleState ProblemLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The lifecycleDetail will give more detail on the substate of the lifecycleState.
	LifecycleDetail ProblemLifecycleDetailEnum `mandatory:"false" json:"lifecycleDetail,omitempty"`

	// Id of the detector associated with the Problem.
	DetectorId DetectorEnumEnum `mandatory:"false" json:"detectorId,omitempty"`

	// targetId of the problem
	TargetId *string `mandatory:"false" json:"targetId"`

	// The additional details of the Problem
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`

	// Description of the problem
	Description *string `mandatory:"false" json:"description"`

	// Recommendation for the problem
	Recommendation *string `mandatory:"false" json:"recommendation"`

	// User Comments
	Comment *string `mandatory:"false" json:"comment"`
}

func (m Problem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Problem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRiskLevelEnum(string(m.RiskLevel)); !ok && m.RiskLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RiskLevel: %s. Supported values are: %s.", m.RiskLevel, strings.Join(GetRiskLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingProblemLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetProblemLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingProblemLifecycleDetailEnum(string(m.LifecycleDetail)); !ok && m.LifecycleDetail != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetail: %s. Supported values are: %s.", m.LifecycleDetail, strings.Join(GetProblemLifecycleDetailEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDetectorEnumEnum(string(m.DetectorId)); !ok && m.DetectorId != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DetectorId: %s. Supported values are: %s.", m.DetectorId, strings.Join(GetDetectorEnumEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
