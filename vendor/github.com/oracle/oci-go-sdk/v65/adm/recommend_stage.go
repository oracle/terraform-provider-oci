// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Dependency Management API
//
// Use the Application Dependency Management API to create knowledge bases and vulnerability audits.  For more information, see ADM (https://docs.cloud.oracle.com/Content/application-dependency-management/home.htm).
//

package adm

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RecommendStage A recommend stage determines the recommended third-party dependencies that the app should use and (if any) opens a pull request to propose these changes.
type RecommendStage struct {

	// The creation date and time of the remediation run stage (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The Oracle Cloud identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the remediation run.
	RemediationRunId *string `mandatory:"true" json:"remediationRunId"`

	// The date and time of the start of the remediation run stage (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time of the finish of the remediation run stage (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// Information about the current step within the stage.
	Summary *string `mandatory:"false" json:"summary"`

	// Count of recommended application dependencies to update.
	RecommendedUpdatesCount *int `mandatory:"false" json:"recommendedUpdatesCount"`

	PullRequestProperties *PullRequestProperties `mandatory:"false" json:"pullRequestProperties"`

	// The current status of a remediation run stage.
	Status RemediationRunStageStatusEnum `mandatory:"true" json:"status"`

	// The previous type of stage in the remediation run.
	PreviousStageType RemediationRunStageTypeEnum `mandatory:"false" json:"previousStageType,omitempty"`

	// The next type of stage in the remediation run.
	NextStageType RemediationRunStageTypeEnum `mandatory:"false" json:"nextStageType,omitempty"`
}

// GetStatus returns Status
func (m RecommendStage) GetStatus() RemediationRunStageStatusEnum {
	return m.Status
}

// GetTimeCreated returns TimeCreated
func (m RecommendStage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeStarted returns TimeStarted
func (m RecommendStage) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

// GetTimeFinished returns TimeFinished
func (m RecommendStage) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

// GetSummary returns Summary
func (m RecommendStage) GetSummary() *string {
	return m.Summary
}

// GetRemediationRunId returns RemediationRunId
func (m RecommendStage) GetRemediationRunId() *string {
	return m.RemediationRunId
}

// GetPreviousStageType returns PreviousStageType
func (m RecommendStage) GetPreviousStageType() RemediationRunStageTypeEnum {
	return m.PreviousStageType
}

// GetNextStageType returns NextStageType
func (m RecommendStage) GetNextStageType() RemediationRunStageTypeEnum {
	return m.NextStageType
}

func (m RecommendStage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecommendStage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRemediationRunStageStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetRemediationRunStageStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRemediationRunStageTypeEnum(string(m.PreviousStageType)); !ok && m.PreviousStageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PreviousStageType: %s. Supported values are: %s.", m.PreviousStageType, strings.Join(GetRemediationRunStageTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRemediationRunStageTypeEnum(string(m.NextStageType)); !ok && m.NextStageType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NextStageType: %s. Supported values are: %s.", m.NextStageType, strings.Join(GetRemediationRunStageTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RecommendStage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRecommendStage RecommendStage
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeRecommendStage
	}{
		"RECOMMEND",
		(MarshalTypeRecommendStage)(m),
	}

	return json.Marshal(&s)
}
