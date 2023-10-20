// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// RemediationRunStage A remediation run stage is one step of an remediation run. Each stage provides output logs and has a specific type.
// The stages are: DETECT, RECOMMEND, VERIFY, and APPLY.
type RemediationRunStage interface {

	// The current status of a remediation run stage.
	GetStatus() RemediationRunStageStatusEnum

	// The creation date and time of the remediation run stage (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	GetTimeCreated() *common.SDKTime

	// The Oracle Cloud identifier (OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)) of the remediation run.
	GetRemediationRunId() *string

	// The date and time of the start of the remediation run stage (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	GetTimeStarted() *common.SDKTime

	// The date and time of the finish of the remediation run stage (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	GetTimeFinished() *common.SDKTime

	// Information about the current step within the stage.
	GetSummary() *string

	// The previous type of stage in the remediation run.
	GetPreviousStageType() RemediationRunStageTypeEnum

	// The next type of stage in the remediation run.
	GetNextStageType() RemediationRunStageTypeEnum
}

type remediationrunstage struct {
	JsonData          []byte
	TimeStarted       *common.SDKTime               `mandatory:"false" json:"timeStarted"`
	TimeFinished      *common.SDKTime               `mandatory:"false" json:"timeFinished"`
	Summary           *string                       `mandatory:"false" json:"summary"`
	PreviousStageType RemediationRunStageTypeEnum   `mandatory:"false" json:"previousStageType,omitempty"`
	NextStageType     RemediationRunStageTypeEnum   `mandatory:"false" json:"nextStageType,omitempty"`
	Status            RemediationRunStageStatusEnum `mandatory:"true" json:"status"`
	TimeCreated       *common.SDKTime               `mandatory:"true" json:"timeCreated"`
	RemediationRunId  *string                       `mandatory:"true" json:"remediationRunId"`
	Type              string                        `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *remediationrunstage) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerremediationrunstage remediationrunstage
	s := struct {
		Model Unmarshalerremediationrunstage
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Status = s.Model.Status
	m.TimeCreated = s.Model.TimeCreated
	m.RemediationRunId = s.Model.RemediationRunId
	m.TimeStarted = s.Model.TimeStarted
	m.TimeFinished = s.Model.TimeFinished
	m.Summary = s.Model.Summary
	m.PreviousStageType = s.Model.PreviousStageType
	m.NextStageType = s.Model.NextStageType
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *remediationrunstage) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "RECOMMEND":
		mm := RecommendStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "VERIFY":
		mm := VerifyStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "APPLY":
		mm := ApplyStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DETECT":
		mm := DetectStage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for RemediationRunStage: %s.", m.Type)
		return *m, nil
	}
}

// GetTimeStarted returns TimeStarted
func (m remediationrunstage) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

// GetTimeFinished returns TimeFinished
func (m remediationrunstage) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

// GetSummary returns Summary
func (m remediationrunstage) GetSummary() *string {
	return m.Summary
}

// GetPreviousStageType returns PreviousStageType
func (m remediationrunstage) GetPreviousStageType() RemediationRunStageTypeEnum {
	return m.PreviousStageType
}

// GetNextStageType returns NextStageType
func (m remediationrunstage) GetNextStageType() RemediationRunStageTypeEnum {
	return m.NextStageType
}

// GetStatus returns Status
func (m remediationrunstage) GetStatus() RemediationRunStageStatusEnum {
	return m.Status
}

// GetTimeCreated returns TimeCreated
func (m remediationrunstage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetRemediationRunId returns RemediationRunId
func (m remediationrunstage) GetRemediationRunId() *string {
	return m.RemediationRunId
}

func (m remediationrunstage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m remediationrunstage) ValidateEnumValue() (bool, error) {
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

// RemediationRunStageStatusEnum Enum with underlying type: string
type RemediationRunStageStatusEnum string

// Set of constants representing the allowable values for RemediationRunStageStatusEnum
const (
	RemediationRunStageStatusCreated    RemediationRunStageStatusEnum = "CREATED"
	RemediationRunStageStatusInProgress RemediationRunStageStatusEnum = "IN_PROGRESS"
	RemediationRunStageStatusSucceeded  RemediationRunStageStatusEnum = "SUCCEEDED"
	RemediationRunStageStatusFailed     RemediationRunStageStatusEnum = "FAILED"
	RemediationRunStageStatusCanceling  RemediationRunStageStatusEnum = "CANCELING"
	RemediationRunStageStatusCanceled   RemediationRunStageStatusEnum = "CANCELED"
)

var mappingRemediationRunStageStatusEnum = map[string]RemediationRunStageStatusEnum{
	"CREATED":     RemediationRunStageStatusCreated,
	"IN_PROGRESS": RemediationRunStageStatusInProgress,
	"SUCCEEDED":   RemediationRunStageStatusSucceeded,
	"FAILED":      RemediationRunStageStatusFailed,
	"CANCELING":   RemediationRunStageStatusCanceling,
	"CANCELED":    RemediationRunStageStatusCanceled,
}

var mappingRemediationRunStageStatusEnumLowerCase = map[string]RemediationRunStageStatusEnum{
	"created":     RemediationRunStageStatusCreated,
	"in_progress": RemediationRunStageStatusInProgress,
	"succeeded":   RemediationRunStageStatusSucceeded,
	"failed":      RemediationRunStageStatusFailed,
	"canceling":   RemediationRunStageStatusCanceling,
	"canceled":    RemediationRunStageStatusCanceled,
}

// GetRemediationRunStageStatusEnumValues Enumerates the set of values for RemediationRunStageStatusEnum
func GetRemediationRunStageStatusEnumValues() []RemediationRunStageStatusEnum {
	values := make([]RemediationRunStageStatusEnum, 0)
	for _, v := range mappingRemediationRunStageStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetRemediationRunStageStatusEnumStringValues Enumerates the set of values in String for RemediationRunStageStatusEnum
func GetRemediationRunStageStatusEnumStringValues() []string {
	return []string{
		"CREATED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingRemediationRunStageStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRemediationRunStageStatusEnum(val string) (RemediationRunStageStatusEnum, bool) {
	enum, ok := mappingRemediationRunStageStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
