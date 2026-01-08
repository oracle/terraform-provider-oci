// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// FsuReadinessCheck Exadata Fleet Update Readiness Check resource details.
type FsuReadinessCheck interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Readiness Check.
	GetId() *string

	// The user-friendly name for the Exadata Fleet Update Readiness Check resource.
	GetDisplayName() *string

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Compartment.
	GetCompartmentId() *string

	// Number of issues found during the Exadata Fleet Update Readiness Check run.
	GetIssueCount() *int

	// The date and time the Exadata Fleet Update Readiness Check was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	GetTimeCreated() *common.SDKTime

	// Possible lifecycle states for the Exadata Fleet Update Readiness Check resource.
	GetLifecycleState() FsuReadinessCheckLifecycleStateEnum

	// Issues found during the Exadata Fleet Update Readiness Check run.
	GetIssues() []PatchingIssueEntry

	// The date and time the Exadata Fleet Update Readiness Check was updated,
	// as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	GetTimeUpdated() *common.SDKTime

	// The date and time the Exadata Fleet Update Readiness Check was finished,
	// as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	GetTimeFinished() *common.SDKTime

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in Failed state.
	GetLifecycleDetails() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type fsureadinesscheck struct {
	JsonData         []byte
	Issues           []PatchingIssueEntry                `mandatory:"false" json:"issues"`
	TimeUpdated      *common.SDKTime                     `mandatory:"false" json:"timeUpdated"`
	TimeFinished     *common.SDKTime                     `mandatory:"false" json:"timeFinished"`
	LifecycleDetails *string                             `mandatory:"false" json:"lifecycleDetails"`
	FreeformTags     map[string]string                   `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{}   `mandatory:"false" json:"definedTags"`
	SystemTags       map[string]map[string]interface{}   `mandatory:"false" json:"systemTags"`
	Id               *string                             `mandatory:"true" json:"id"`
	DisplayName      *string                             `mandatory:"true" json:"displayName"`
	CompartmentId    *string                             `mandatory:"true" json:"compartmentId"`
	IssueCount       *int                                `mandatory:"true" json:"issueCount"`
	TimeCreated      *common.SDKTime                     `mandatory:"true" json:"timeCreated"`
	LifecycleState   FsuReadinessCheckLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	Type             string                              `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *fsureadinesscheck) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerfsureadinesscheck fsureadinesscheck
	s := struct {
		Model Unmarshalerfsureadinesscheck
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.DisplayName = s.Model.DisplayName
	m.CompartmentId = s.Model.CompartmentId
	m.IssueCount = s.Model.IssueCount
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.Issues = s.Model.Issues
	m.TimeUpdated = s.Model.TimeUpdated
	m.TimeFinished = s.Model.TimeFinished
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *fsureadinesscheck) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "TARGET":
		mm := TargetFsuReadinessCheck{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for FsuReadinessCheck: %s.", m.Type)
		return *m, nil
	}
}

// GetIssues returns Issues
func (m fsureadinesscheck) GetIssues() []PatchingIssueEntry {
	return m.Issues
}

// GetTimeUpdated returns TimeUpdated
func (m fsureadinesscheck) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetTimeFinished returns TimeFinished
func (m fsureadinesscheck) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

// GetLifecycleDetails returns LifecycleDetails
func (m fsureadinesscheck) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m fsureadinesscheck) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m fsureadinesscheck) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m fsureadinesscheck) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m fsureadinesscheck) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m fsureadinesscheck) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m fsureadinesscheck) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetIssueCount returns IssueCount
func (m fsureadinesscheck) GetIssueCount() *int {
	return m.IssueCount
}

// GetTimeCreated returns TimeCreated
func (m fsureadinesscheck) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m fsureadinesscheck) GetLifecycleState() FsuReadinessCheckLifecycleStateEnum {
	return m.LifecycleState
}

func (m fsureadinesscheck) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m fsureadinesscheck) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFsuReadinessCheckLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFsuReadinessCheckLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FsuReadinessCheckLifecycleStateEnum Enum with underlying type: string
type FsuReadinessCheckLifecycleStateEnum string

// Set of constants representing the allowable values for FsuReadinessCheckLifecycleStateEnum
const (
	FsuReadinessCheckLifecycleStateAccepted       FsuReadinessCheckLifecycleStateEnum = "ACCEPTED"
	FsuReadinessCheckLifecycleStateInProgress     FsuReadinessCheckLifecycleStateEnum = "IN_PROGRESS"
	FsuReadinessCheckLifecycleStateFailed         FsuReadinessCheckLifecycleStateEnum = "FAILED"
	FsuReadinessCheckLifecycleStateNeedsAttention FsuReadinessCheckLifecycleStateEnum = "NEEDS_ATTENTION"
	FsuReadinessCheckLifecycleStateSucceeded      FsuReadinessCheckLifecycleStateEnum = "SUCCEEDED"
	FsuReadinessCheckLifecycleStateWaiting        FsuReadinessCheckLifecycleStateEnum = "WAITING"
	FsuReadinessCheckLifecycleStateCanceling      FsuReadinessCheckLifecycleStateEnum = "CANCELING"
	FsuReadinessCheckLifecycleStateCanceled       FsuReadinessCheckLifecycleStateEnum = "CANCELED"
	FsuReadinessCheckLifecycleStateDeleting       FsuReadinessCheckLifecycleStateEnum = "DELETING"
	FsuReadinessCheckLifecycleStateDeleted        FsuReadinessCheckLifecycleStateEnum = "DELETED"
)

var mappingFsuReadinessCheckLifecycleStateEnum = map[string]FsuReadinessCheckLifecycleStateEnum{
	"ACCEPTED":        FsuReadinessCheckLifecycleStateAccepted,
	"IN_PROGRESS":     FsuReadinessCheckLifecycleStateInProgress,
	"FAILED":          FsuReadinessCheckLifecycleStateFailed,
	"NEEDS_ATTENTION": FsuReadinessCheckLifecycleStateNeedsAttention,
	"SUCCEEDED":       FsuReadinessCheckLifecycleStateSucceeded,
	"WAITING":         FsuReadinessCheckLifecycleStateWaiting,
	"CANCELING":       FsuReadinessCheckLifecycleStateCanceling,
	"CANCELED":        FsuReadinessCheckLifecycleStateCanceled,
	"DELETING":        FsuReadinessCheckLifecycleStateDeleting,
	"DELETED":         FsuReadinessCheckLifecycleStateDeleted,
}

var mappingFsuReadinessCheckLifecycleStateEnumLowerCase = map[string]FsuReadinessCheckLifecycleStateEnum{
	"accepted":        FsuReadinessCheckLifecycleStateAccepted,
	"in_progress":     FsuReadinessCheckLifecycleStateInProgress,
	"failed":          FsuReadinessCheckLifecycleStateFailed,
	"needs_attention": FsuReadinessCheckLifecycleStateNeedsAttention,
	"succeeded":       FsuReadinessCheckLifecycleStateSucceeded,
	"waiting":         FsuReadinessCheckLifecycleStateWaiting,
	"canceling":       FsuReadinessCheckLifecycleStateCanceling,
	"canceled":        FsuReadinessCheckLifecycleStateCanceled,
	"deleting":        FsuReadinessCheckLifecycleStateDeleting,
	"deleted":         FsuReadinessCheckLifecycleStateDeleted,
}

// GetFsuReadinessCheckLifecycleStateEnumValues Enumerates the set of values for FsuReadinessCheckLifecycleStateEnum
func GetFsuReadinessCheckLifecycleStateEnumValues() []FsuReadinessCheckLifecycleStateEnum {
	values := make([]FsuReadinessCheckLifecycleStateEnum, 0)
	for _, v := range mappingFsuReadinessCheckLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFsuReadinessCheckLifecycleStateEnumStringValues Enumerates the set of values in String for FsuReadinessCheckLifecycleStateEnum
func GetFsuReadinessCheckLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"NEEDS_ATTENTION",
		"SUCCEEDED",
		"WAITING",
		"CANCELING",
		"CANCELED",
		"DELETING",
		"DELETED",
	}
}

// GetMappingFsuReadinessCheckLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFsuReadinessCheckLifecycleStateEnum(val string) (FsuReadinessCheckLifecycleStateEnum, bool) {
	enum, ok := mappingFsuReadinessCheckLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FsuReadinessCheckTypeEnum Enum with underlying type: string
type FsuReadinessCheckTypeEnum string

// Set of constants representing the allowable values for FsuReadinessCheckTypeEnum
const (
	FsuReadinessCheckTypeTarget FsuReadinessCheckTypeEnum = "TARGET"
)

var mappingFsuReadinessCheckTypeEnum = map[string]FsuReadinessCheckTypeEnum{
	"TARGET": FsuReadinessCheckTypeTarget,
}

var mappingFsuReadinessCheckTypeEnumLowerCase = map[string]FsuReadinessCheckTypeEnum{
	"target": FsuReadinessCheckTypeTarget,
}

// GetFsuReadinessCheckTypeEnumValues Enumerates the set of values for FsuReadinessCheckTypeEnum
func GetFsuReadinessCheckTypeEnumValues() []FsuReadinessCheckTypeEnum {
	values := make([]FsuReadinessCheckTypeEnum, 0)
	for _, v := range mappingFsuReadinessCheckTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFsuReadinessCheckTypeEnumStringValues Enumerates the set of values in String for FsuReadinessCheckTypeEnum
func GetFsuReadinessCheckTypeEnumStringValues() []string {
	return []string{
		"TARGET",
	}
}

// GetMappingFsuReadinessCheckTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFsuReadinessCheckTypeEnum(val string) (FsuReadinessCheckTypeEnum, bool) {
	enum, ok := mappingFsuReadinessCheckTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
