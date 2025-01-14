// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// GovernanceProblemSummary Summary information for a governance problem.
type GovernanceProblemSummary struct {

	// Unique identifier that can't be changed after creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment OCID of subject tenancy where the resource is created
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the governance problem was created. Format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current lifecycle state of the governance problem
	LifecycleState ProblemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Unique identifier of the detector rule
	DetectorRuleId *string `mandatory:"false" json:"detectorRuleId"`

	// The risk level of the governance problem
	RiskLevel RiskLevelEnum `mandatory:"false" json:"riskLevel,omitempty"`

	// The risk score for the governance problem
	RiskScore *float64 `mandatory:"false" json:"riskScore"`

	// Unique identifier of the resource that's impacted by the governance problem
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Display name of the resource impacted by the governance problem
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// Type of the resource impacted by the governance problem
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// User-defined labels on the governance problem
	Labels []string `mandatory:"false" json:"labels"`

	// The date and time the governance problem was first detected. Format defined by RFC3339.
	TimeFirstDetected *common.SDKTime `mandatory:"false" json:"timeFirstDetected"`

	// The date and time the governance problem was last detected. Format defined by RFC3339.
	TimeLastDetected *common.SDKTime `mandatory:"false" json:"timeLastDetected"`

	// Additional details on the substate of the lifecycle state
	LifecycleDetail ProblemLifecycleDetailEnum `mandatory:"false" json:"lifecycleDetail,omitempty"`

	// Unique identifier of the detector associated with the governance problem
	DetectorId DetectorEnumEnum `mandatory:"false" json:"detectorId,omitempty"`

	// List of regions where the governance problem is found
	Regions []string `mandatory:"false" json:"regions"`

	// Unique identifier of the subject tenancy associated with the governance problem
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// Unique identifier of the subject target associated with the governance problem
	TargetId *string `mandatory:"false" json:"targetId"`

	// Target OCID of the governance target
	GovernanceTargetId *string `mandatory:"false" json:"governanceTargetId"`

	// Compartment OCID where the governance target is created
	GovernanceTargetCompartmentId *string `mandatory:"false" json:"governanceTargetCompartmentId"`

	// Indicates if problem was created by current tenancy or governing tenancy.
	ManageType ManageTypeEnum `mandatory:"false" json:"manageType,omitempty"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	// Avoid entering confidential information.
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// System tags can be viewed by users, but can only be created by the system.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m GovernanceProblemSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GovernanceProblemSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingProblemLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetProblemLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingRiskLevelEnum(string(m.RiskLevel)); !ok && m.RiskLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RiskLevel: %s. Supported values are: %s.", m.RiskLevel, strings.Join(GetRiskLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingProblemLifecycleDetailEnum(string(m.LifecycleDetail)); !ok && m.LifecycleDetail != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetail: %s. Supported values are: %s.", m.LifecycleDetail, strings.Join(GetProblemLifecycleDetailEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDetectorEnumEnum(string(m.DetectorId)); !ok && m.DetectorId != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DetectorId: %s. Supported values are: %s.", m.DetectorId, strings.Join(GetDetectorEnumEnumStringValues(), ",")))
	}
	if _, ok := GetMappingManageTypeEnum(string(m.ManageType)); !ok && m.ManageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManageType: %s. Supported values are: %s.", m.ManageType, strings.Join(GetManageTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
