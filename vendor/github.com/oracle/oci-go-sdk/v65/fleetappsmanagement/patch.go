// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Patch Patch metadata for Custom and Oracle patches.
type Patch struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Should be unique within the tenancy, and cannot be changed after creation.
	// Avoid entering confidential information.
	Name *string `mandatory:"true" json:"name"`

	PatchType *PatchType `mandatory:"true" json:"patchType"`

	// Patch Severity.
	Severity PatchSeverityEnum `mandatory:"true" json:"severity"`

	// Date when the patch was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	ArtifactDetails ArtifactDetails `mandatory:"true" json:"artifactDetails"`

	Product *PatchProduct `mandatory:"true" json:"product"`

	// The current state of the Patch.
	LifecycleState PatchLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// OCID of the compartment to which the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Provide information on who defined the patch.
	// Example: For Custom Patches the value will be USER_DEFINED
	// For Oracle Defined Patches the value will be ORACLE_DEFINED
	Type PatchTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Dependent Patches for this patch.
	DependentPatches []DependentPatchDetails `mandatory:"false" json:"dependentPatches"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Associated region
	ResourceRegion *string `mandatory:"false" json:"resourceRegion"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Patch) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Patch) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetPatchSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPatchLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPatchTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetPatchTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Patch) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description      *string                           `json:"description"`
		Type             PatchTypeEnum                     `json:"type"`
		DependentPatches []DependentPatchDetails           `json:"dependentPatches"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		ResourceRegion   *string                           `json:"resourceRegion"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		SystemTags       map[string]map[string]interface{} `json:"systemTags"`
		Id               *string                           `json:"id"`
		Name             *string                           `json:"name"`
		PatchType        *PatchType                        `json:"patchType"`
		Severity         PatchSeverityEnum                 `json:"severity"`
		TimeReleased     *common.SDKTime                   `json:"timeReleased"`
		ArtifactDetails  artifactdetails                   `json:"artifactDetails"`
		Product          *PatchProduct                     `json:"product"`
		LifecycleState   PatchLifecycleStateEnum           `json:"lifecycleState"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
		CompartmentId    *string                           `json:"compartmentId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Type = model.Type

	m.DependentPatches = make([]DependentPatchDetails, len(model.DependentPatches))
	copy(m.DependentPatches, model.DependentPatches)
	m.LifecycleDetails = model.LifecycleDetails

	m.ResourceRegion = model.ResourceRegion

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.Name = model.Name

	m.PatchType = model.PatchType

	m.Severity = model.Severity

	m.TimeReleased = model.TimeReleased

	nn, e = model.ArtifactDetails.UnmarshalPolymorphicJSON(model.ArtifactDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ArtifactDetails = nn.(ArtifactDetails)
	} else {
		m.ArtifactDetails = nil
	}

	m.Product = model.Product

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.CompartmentId = model.CompartmentId

	return
}

// PatchTypeEnum Enum with underlying type: string
type PatchTypeEnum string

// Set of constants representing the allowable values for PatchTypeEnum
const (
	PatchTypeUserDefined   PatchTypeEnum = "USER_DEFINED"
	PatchTypeOracleDefined PatchTypeEnum = "ORACLE_DEFINED"
)

var mappingPatchTypeEnum = map[string]PatchTypeEnum{
	"USER_DEFINED":   PatchTypeUserDefined,
	"ORACLE_DEFINED": PatchTypeOracleDefined,
}

var mappingPatchTypeEnumLowerCase = map[string]PatchTypeEnum{
	"user_defined":   PatchTypeUserDefined,
	"oracle_defined": PatchTypeOracleDefined,
}

// GetPatchTypeEnumValues Enumerates the set of values for PatchTypeEnum
func GetPatchTypeEnumValues() []PatchTypeEnum {
	values := make([]PatchTypeEnum, 0)
	for _, v := range mappingPatchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchTypeEnumStringValues Enumerates the set of values in String for PatchTypeEnum
func GetPatchTypeEnumStringValues() []string {
	return []string{
		"USER_DEFINED",
		"ORACLE_DEFINED",
	}
}

// GetMappingPatchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchTypeEnum(val string) (PatchTypeEnum, bool) {
	enum, ok := mappingPatchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchLifecycleStateEnum Enum with underlying type: string
type PatchLifecycleStateEnum string

// Set of constants representing the allowable values for PatchLifecycleStateEnum
const (
	PatchLifecycleStateActive   PatchLifecycleStateEnum = "ACTIVE"
	PatchLifecycleStateInactive PatchLifecycleStateEnum = "INACTIVE"
	PatchLifecycleStateDeleted  PatchLifecycleStateEnum = "DELETED"
	PatchLifecycleStateDeleting PatchLifecycleStateEnum = "DELETING"
	PatchLifecycleStateFailed   PatchLifecycleStateEnum = "FAILED"
	PatchLifecycleStateUpdating PatchLifecycleStateEnum = "UPDATING"
)

var mappingPatchLifecycleStateEnum = map[string]PatchLifecycleStateEnum{
	"ACTIVE":   PatchLifecycleStateActive,
	"INACTIVE": PatchLifecycleStateInactive,
	"DELETED":  PatchLifecycleStateDeleted,
	"DELETING": PatchLifecycleStateDeleting,
	"FAILED":   PatchLifecycleStateFailed,
	"UPDATING": PatchLifecycleStateUpdating,
}

var mappingPatchLifecycleStateEnumLowerCase = map[string]PatchLifecycleStateEnum{
	"active":   PatchLifecycleStateActive,
	"inactive": PatchLifecycleStateInactive,
	"deleted":  PatchLifecycleStateDeleted,
	"deleting": PatchLifecycleStateDeleting,
	"failed":   PatchLifecycleStateFailed,
	"updating": PatchLifecycleStateUpdating,
}

// GetPatchLifecycleStateEnumValues Enumerates the set of values for PatchLifecycleStateEnum
func GetPatchLifecycleStateEnumValues() []PatchLifecycleStateEnum {
	values := make([]PatchLifecycleStateEnum, 0)
	for _, v := range mappingPatchLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchLifecycleStateEnumStringValues Enumerates the set of values in String for PatchLifecycleStateEnum
func GetPatchLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingPatchLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchLifecycleStateEnum(val string) (PatchLifecycleStateEnum, bool) {
	enum, ok := mappingPatchLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
