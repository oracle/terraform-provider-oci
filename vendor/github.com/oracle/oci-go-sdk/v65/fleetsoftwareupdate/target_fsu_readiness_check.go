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

// TargetFsuReadinessCheck Patch Exadata Fleet Update Readiness Check resource details.
type TargetFsuReadinessCheck struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Fleet Update Readiness Check.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name for the Exadata Fleet Update Readiness Check resource.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Number of issues found during the Exadata Fleet Update Readiness Check run.
	IssueCount *int `mandatory:"true" json:"issueCount"`

	// The date and time the Exadata Fleet Update Readiness Check was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// List of targets that will run the Exadata Fleet Update Readiness Check.
	// The targets have to be of the same entity type.
	Targets []ReadinessCheckTargetEntry `mandatory:"true" json:"targets"`

	// Issues found during the Exadata Fleet Update Readiness Check run.
	Issues []PatchingIssueEntry `mandatory:"false" json:"issues"`

	// The date and time the Exadata Fleet Update Readiness Check was updated,
	// as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The date and time the Exadata Fleet Update Readiness Check was finished,
	// as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Possible lifecycle states for the Exadata Fleet Update Readiness Check resource.
	LifecycleState FsuReadinessCheckLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m TargetFsuReadinessCheck) GetId() *string {
	return m.Id
}

// GetDisplayName returns DisplayName
func (m TargetFsuReadinessCheck) GetDisplayName() *string {
	return m.DisplayName
}

// GetCompartmentId returns CompartmentId
func (m TargetFsuReadinessCheck) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetIssueCount returns IssueCount
func (m TargetFsuReadinessCheck) GetIssueCount() *int {
	return m.IssueCount
}

// GetIssues returns Issues
func (m TargetFsuReadinessCheck) GetIssues() []PatchingIssueEntry {
	return m.Issues
}

// GetTimeCreated returns TimeCreated
func (m TargetFsuReadinessCheck) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m TargetFsuReadinessCheck) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetTimeFinished returns TimeFinished
func (m TargetFsuReadinessCheck) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

// GetLifecycleState returns LifecycleState
func (m TargetFsuReadinessCheck) GetLifecycleState() FsuReadinessCheckLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m TargetFsuReadinessCheck) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetFreeformTags returns FreeformTags
func (m TargetFsuReadinessCheck) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m TargetFsuReadinessCheck) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m TargetFsuReadinessCheck) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m TargetFsuReadinessCheck) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetFsuReadinessCheck) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFsuReadinessCheckLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFsuReadinessCheckLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TargetFsuReadinessCheck) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTargetFsuReadinessCheck TargetFsuReadinessCheck
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTargetFsuReadinessCheck
	}{
		"TARGET",
		(MarshalTypeTargetFsuReadinessCheck)(m),
	}

	return json.Marshal(&s)
}
