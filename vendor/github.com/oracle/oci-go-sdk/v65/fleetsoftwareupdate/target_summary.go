// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TargetSummary Details of a target member of a Exadata Fleet Update Collection.
type TargetSummary struct {
	Target TargetDetails `mandatory:"false" json:"target"`

	// Current version of the target
	CurrentVersion *string `mandatory:"false" json:"currentVersion"`

	// Status of the target in the Exadata Fleet Update Collection.
	Status TargetSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Exadata Fleet Update Job OCID executing an action in the target. Null if no job is being executed.
	ExecutingFsuJobId *string `mandatory:"false" json:"executingFsuJobId"`

	// Active Exadata Fleet Update Cycle OCID. Null if no Cycle is active that has this target as member.
	ActiveFsuCycleId *string `mandatory:"false" json:"activeFsuCycleId"`

	Progress *TargetProgressSummary `mandatory:"false" json:"progress"`
}

func (m TargetSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTargetSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetTargetSummaryStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *TargetSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Target            targetdetails           `json:"target"`
		CurrentVersion    *string                 `json:"currentVersion"`
		Status            TargetSummaryStatusEnum `json:"status"`
		ExecutingFsuJobId *string                 `json:"executingFsuJobId"`
		ActiveFsuCycleId  *string                 `json:"activeFsuCycleId"`
		Progress          *TargetProgressSummary  `json:"progress"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Target.UnmarshalPolymorphicJSON(model.Target.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Target = nn.(TargetDetails)
	} else {
		m.Target = nil
	}

	m.CurrentVersion = model.CurrentVersion

	m.Status = model.Status

	m.ExecutingFsuJobId = model.ExecutingFsuJobId

	m.ActiveFsuCycleId = model.ActiveFsuCycleId

	m.Progress = model.Progress

	return
}

// TargetSummaryStatusEnum Enum with underlying type: string
type TargetSummaryStatusEnum string

// Set of constants representing the allowable values for TargetSummaryStatusEnum
const (
	TargetSummaryStatusIdle         TargetSummaryStatusEnum = "IDLE"
	TargetSummaryStatusExecutingJob TargetSummaryStatusEnum = "EXECUTING_JOB"
	TargetSummaryStatusJobFailed    TargetSummaryStatusEnum = "JOB_FAILED"
)

var mappingTargetSummaryStatusEnum = map[string]TargetSummaryStatusEnum{
	"IDLE":          TargetSummaryStatusIdle,
	"EXECUTING_JOB": TargetSummaryStatusExecutingJob,
	"JOB_FAILED":    TargetSummaryStatusJobFailed,
}

var mappingTargetSummaryStatusEnumLowerCase = map[string]TargetSummaryStatusEnum{
	"idle":          TargetSummaryStatusIdle,
	"executing_job": TargetSummaryStatusExecutingJob,
	"job_failed":    TargetSummaryStatusJobFailed,
}

// GetTargetSummaryStatusEnumValues Enumerates the set of values for TargetSummaryStatusEnum
func GetTargetSummaryStatusEnumValues() []TargetSummaryStatusEnum {
	values := make([]TargetSummaryStatusEnum, 0)
	for _, v := range mappingTargetSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetSummaryStatusEnumStringValues Enumerates the set of values in String for TargetSummaryStatusEnum
func GetTargetSummaryStatusEnumStringValues() []string {
	return []string{
		"IDLE",
		"EXECUTING_JOB",
		"JOB_FAILED",
	}
}

// GetMappingTargetSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetSummaryStatusEnum(val string) (TargetSummaryStatusEnum, bool) {
	enum, ok := mappingTargetSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
