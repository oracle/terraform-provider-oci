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

// FsuCollectionTarget Details of a target member of a Exadata Fleet Update Collection.
type FsuCollectionTarget struct {
	Target TargetDetails `mandatory:"true" json:"target"`

	// Current version of the target.
	CurrentVersion *string `mandatory:"false" json:"currentVersion"`

	// Status of the target in the Exadata Fleet Update Collection.
	Status FsuCollectionTargetStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Exadata Fleet Update Job OCID executing an action in the target. Null if no job is being executed.
	ExecutingFsuJobId *string `mandatory:"false" json:"executingFsuJobId"`

	// Active Exadata Fleet Update Cycle OCID. Null if no Cycle is active that has this target as member.
	ActiveFsuCycleId *string `mandatory:"false" json:"activeFsuCycleId"`

	Progress *TargetProgressSummary `mandatory:"false" json:"progress"`
}

func (m FsuCollectionTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FsuCollectionTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFsuCollectionTargetStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetFsuCollectionTargetStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *FsuCollectionTarget) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CurrentVersion    *string                       `json:"currentVersion"`
		Status            FsuCollectionTargetStatusEnum `json:"status"`
		ExecutingFsuJobId *string                       `json:"executingFsuJobId"`
		ActiveFsuCycleId  *string                       `json:"activeFsuCycleId"`
		Progress          *TargetProgressSummary        `json:"progress"`
		Target            targetdetails                 `json:"target"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CurrentVersion = model.CurrentVersion

	m.Status = model.Status

	m.ExecutingFsuJobId = model.ExecutingFsuJobId

	m.ActiveFsuCycleId = model.ActiveFsuCycleId

	m.Progress = model.Progress

	nn, e = model.Target.UnmarshalPolymorphicJSON(model.Target.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Target = nn.(TargetDetails)
	} else {
		m.Target = nil
	}

	return
}

// FsuCollectionTargetStatusEnum Enum with underlying type: string
type FsuCollectionTargetStatusEnum string

// Set of constants representing the allowable values for FsuCollectionTargetStatusEnum
const (
	FsuCollectionTargetStatusIdle         FsuCollectionTargetStatusEnum = "IDLE"
	FsuCollectionTargetStatusExecutingJob FsuCollectionTargetStatusEnum = "EXECUTING_JOB"
	FsuCollectionTargetStatusJobFailed    FsuCollectionTargetStatusEnum = "JOB_FAILED"
)

var mappingFsuCollectionTargetStatusEnum = map[string]FsuCollectionTargetStatusEnum{
	"IDLE":          FsuCollectionTargetStatusIdle,
	"EXECUTING_JOB": FsuCollectionTargetStatusExecutingJob,
	"JOB_FAILED":    FsuCollectionTargetStatusJobFailed,
}

var mappingFsuCollectionTargetStatusEnumLowerCase = map[string]FsuCollectionTargetStatusEnum{
	"idle":          FsuCollectionTargetStatusIdle,
	"executing_job": FsuCollectionTargetStatusExecutingJob,
	"job_failed":    FsuCollectionTargetStatusJobFailed,
}

// GetFsuCollectionTargetStatusEnumValues Enumerates the set of values for FsuCollectionTargetStatusEnum
func GetFsuCollectionTargetStatusEnumValues() []FsuCollectionTargetStatusEnum {
	values := make([]FsuCollectionTargetStatusEnum, 0)
	for _, v := range mappingFsuCollectionTargetStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetFsuCollectionTargetStatusEnumStringValues Enumerates the set of values in String for FsuCollectionTargetStatusEnum
func GetFsuCollectionTargetStatusEnumStringValues() []string {
	return []string{
		"IDLE",
		"EXECUTING_JOB",
		"JOB_FAILED",
	}
}

// GetMappingFsuCollectionTargetStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFsuCollectionTargetStatusEnum(val string) (FsuCollectionTargetStatusEnum, bool) {
	enum, ok := mappingFsuCollectionTargetStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
