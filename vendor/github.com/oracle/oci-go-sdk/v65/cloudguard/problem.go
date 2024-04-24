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

// Problem Problems are at the core of Cloud Guard’s functionality. A Problem resource is created whenever an action or a configuration on a resource triggers a rule in a detector that’s attached to the target containing the compartment where the resource is located. Each Problem resource contains all the details for a single problem. This is the information for the problem that appears on the Cloud Guard Problems page.
type Problem struct {

	// Unique identifier that can't be changed after creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment OCID where the resource is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Unique identifier of the detector rule that triggered the problem
	DetectorRuleId *string `mandatory:"false" json:"detectorRuleId"`

	// DEPRECATED
	Region *string `mandatory:"false" json:"region"`

	// Regions where the problem is found
	Regions []string `mandatory:"false" json:"regions"`

	// The risk level for the problem
	RiskLevel RiskLevelEnum `mandatory:"false" json:"riskLevel,omitempty"`

	// The risk score for the problem
	RiskScore *float64 `mandatory:"false" json:"riskScore"`

	// The date and time for the peak risk score that is observed for the problem. Format defined by RFC3339.
	PeakRiskScoreDate *string `mandatory:"false" json:"peakRiskScoreDate"`

	// Peak risk score for the problem
	PeakRiskScore *float64 `mandatory:"false" json:"peakRiskScore"`

	// The date and time when the problem will be auto resolved. Format defined by RFC3339.
	AutoResolveDate *string `mandatory:"false" json:"autoResolveDate"`

	// Number of days for which peak score is calculated for the problem
	PeakRiskScoreLookupPeriodInDays *int `mandatory:"false" json:"peakRiskScoreLookupPeriodInDays"`

	// Unique identifier of the resource affected by the problem
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Display name of the affected resource
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// Type of the affected resource
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// User-defined labels on the problem
	Labels []string `mandatory:"false" json:"labels"`

	// The date and time the problem was last detected. Format defined by RFC3339.
	TimeLastDetected *common.SDKTime `mandatory:"false" json:"timeLastDetected"`

	// The date and time the problem was first detected. Format defined by RFC3339.
	TimeFirstDetected *common.SDKTime `mandatory:"false" json:"timeFirstDetected"`

	// The current lifecycle state of the problem
	LifecycleState ProblemLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Additional details on the substate of the lifecycle state
	LifecycleDetail ProblemLifecycleDetailEnum `mandatory:"false" json:"lifecycleDetail,omitempty"`

	// Unique identifier of the detector rule that triggered the problem
	DetectorId DetectorEnumEnum `mandatory:"false" json:"detectorId,omitempty"`

	// Unique identifier of the target associated with the problem
	TargetId *string `mandatory:"false" json:"targetId"`

	// The additional details of the problem
	AdditionalDetails map[string]string `mandatory:"false" json:"additionalDetails"`

	// Description of the problem
	Description *string `mandatory:"false" json:"description"`

	// Recommendation for the problem
	Recommendation *string `mandatory:"false" json:"recommendation"`

	// User comments on the problem
	Comment *string `mandatory:"false" json:"comment"`

	// Unique identifier of the resource impacted by the problem
	ImpactedResourceId *string `mandatory:"false" json:"impactedResourceId"`

	// Display name of the impacted resource
	ImpactedResourceName *string `mandatory:"false" json:"impactedResourceName"`

	// Type of the impacted resource
	ImpactedResourceType *string `mandatory:"false" json:"impactedResourceType"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`
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
